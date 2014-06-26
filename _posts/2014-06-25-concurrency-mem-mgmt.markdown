---
layout: post
title: Concurrency is the New Memory Management
summary: Distributed systems demand concurrency. What are we doing about it?
---

These days, we all need to build "cloud apps." We've heard "cloud" so many times at this point that it's a buzz word, but underneath the hype lives a real fact for us developers: everything we build needs to be a distributed system.

For our apps to run, we must have servers and they need to respond to requests from our apps, so we need to build systems that have many computers to provide all the fault tolerance, throughput, etc... we need. You've probably heard this all before.

# Distributed Systems are Hard

This is a fact that we all know by now. Computers die, networks get congested, etc...

The bottom line is that our programs are deployed onto multiple nodes as a process on each, and each process communicates with all the others.

The process model is still powerful but it has changed on the server side. Each process communicates with N other identical ones now, making a *cluster*. And when we build a cluster, we're forced into a world of hard concurrency and networking problems.

# Yesterday's State of the Art

I'm defining the "languages of yesterday" as those that the vast majority of developers were using 5 years ago to build their programs. I'm thinking C++, C, Java, PHP, Ruby, Python, etc...

Those languages have a very rich set of libraries and strong communities for server devleoper to this day. And that's for a reason - they were well suited in some way to build *servers*. Some made it easy to write servers, some made it possible to write efficient servers, and others somewhere in between.

Regardless, none made it easy to build *clusters*.

# Building Clusters

Languages like [Go](http://golang.org), [Akka](http://akka.io) and [Erlang](http://erlang.org) have come up now because they help solve the hard networking and concurrency problems that we need to build clusters.

They keep lots of the syntax and semantics from the "languages of yesterday". For example:

- automatic memory management
- polymorphism
- syntax

But they all build on top with built-in concurrency features:

1. non-blocking I/O
2. user-space concurrency
3. easy messaging
4. advanced scheduling

These 4 concurrency features make clusters **much** easier to build because they take a large amount of cognitive load off of the developer, leaving it to the runtime.

# A New Era

The concurrency features don't come for free, however. The concurrency trade-offs are very similar to those with memory in garbage collected languages.

In both cases, the runtime will take responsibility and control from the developer to manage a resource (memory and CPU, respectively).

When we're building a cluster, we're often interested in correctness, latency and throughput of the system. If by porting our system from Go to C we could squeeze 5% more CPU utilization or remove 10% memory utilization, we won't.

In a cluster, we're not optimizing first for single-process performance. Instead, we're optimizing to get our systems right.

We're building concurrency into our languages and runtimes because we need it for the systems we're building today.