---
author: "Aaron Schlesinger"
date: 2021-10-29T21:58:32Z
title: 'Synchronizing the KEDA HTTP Addon Request Routing Table Across Hundreds of Interceptor Pods'
slug: "keda-http-addon-routing-table"
tags: ['distributed-systems', 'event-loops', 'concurrency']
---

The [KEDA HTTP Addon project](https://github.com/kedacore/http-add-on) contains three major components: the [operator](https://github.com/kedacore/http-add-on/tree/main/operator), [scaler](https://github.com/kedacore/http-add-on/tree/main/scaler) and [interceptor](https://github.com/kedacore/http-add-on/tree/main/interceptor).

Of these, the interceptor is the only component that sits in the critical path of all incoming HTTP requests. We also run them in a _fleet_ that is horizontally scaled by software.

We're going to focus on how we ensure that any interceptor replica can route an incoming request to the correct backing application at any time.

## Implications of Multi-Tenancy

The interceptor component is designed to run in a `Deployment` that KEDA automatically scales. This high-level design has a few implications:

1. There must be a centralized, durable copy of a lookup table -- called a _routing table_ -- that maps any incoming request to the correct backing `Service` and port.
1. All interceptor pods must reliably stay up to date with the central routing table
1. All interceptor pods must be able to handle any valid incoming request, regardless of application it was intended for
1. All interceptor pods must be able to quickly execute lookups to the routing table

## Keeping up to Date with the Routing Table

Since the interceptor needs to do a lookup to the routing table before forwarding any request (or returning an error code), lookups need to be as fast as possible. That means storing the routing table in memory and keeping each interceptor's in-memory copy up to date with the central copy.

We do this wth a relatively simple event loop, outlined in the below ([Go](https://golang.org)-like) pseudocode:

```go
table = fetch_table_from_kubernetes()
report_alive()
ticker = start_ticker(every_1_second)
for {
    select {
        case new_table <- kubernetes_events_chan:
            table = new_table
        case <-ticker:
            table = fetch_table_from_kubernetes()
    }
}
```

A few important points to note about this event loop:

1. On startup, we fetch a complete copy of the routing table from Kubernetes before we report to the [Kubernetes liveness probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/). This means that Kubernetes doesn't consider any interceptor pod "ready" until it has a complete initial copy of the routing table.
2. The `kubernetes_events_chan` receives notifications about changes to the central routing table in near-real time. When we get a change notification, we immediately update the in-memory copy.
3. The `ticker` fires a signal every second, at which time we do a full refresh of the routing table. This mechanism ensures that all interceptor replicas receive changes to the routing table within 1 second of the change being made to the central copy.

## The Event Loop Pattern in Action

The KEDA HTTP Addon runs a fleet of interceptors and enlists KEDA to actively scale the fleet. This means that as we send more HTTP traffic to the cluster, we expect the interceptors to automatically scale up. This behavior is one of the most important features of the HTTP Addon.

We've built this event loop into the interceptors to ensure that there can be thousands [1] of them running at once, and they all stay up to date with the central routing table -- data that they need to do their job.

---

[1] As we see in the pseudocode above, each interceptor issues requests to the Kubernetes API, so as we scale them out, we generate more consistent traffic to the cluster API. As you scale further than the low thousands of replicas, you would need to add an intermediate layer of caching between the interceptors and the API to ensure that you don't crash the cluster control plane. 
