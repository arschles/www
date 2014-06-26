---
layout: post
title: Concurrency is the New Memory Management
summary: Distributed systems demand concurrency. What are we doing about it?
---

These days, we all need to build "cloud apps." We've heard "cloud" so many
times at this point that it's a buzz word, but underneath the hype lives a real
fact for us developers: everything we build needs to be a distributed system.

For our apps to run, we must have servers and they need to respond to requests
from our apps, all the time. So we need to build systems that have many
computers to provide all the fault tolerance, throughput, etc... we need.
You've probably heard this all before.

# Distributed Systems are Hard

This is a fact that we all know by now. Computers die, networks get congested,
etc...

The bottom line is that our programs are deployed onto multiple nodes as a
process on each, and each process communicates with all the others.

The process model is still powerful but it has changed on the server side. Now
every process communicates with N other identical ones, making a *cluster*.

Cluster is just a fancy word for distributed systems, so they come with the
same pitfalls. Most notably we're forced into a world of hard concurrency and
networking problems when we build a cluster.

We need to address them or the thing will crash. Most likely late at night
on a weekend, 5 months from now after you've moved on to the next project.

# Yesterday's State of the Art

I'm defining the "languages of yesterday" as those that the vast majority of
developers were using 5 years ago to build their servers. I'm thinking C++, C,
Java, PHP, Ruby, Python, etc...

Those languages have a very rich set of libraries and strong communities for
server developers to this day. That's for a reason - they're well suited
in some way to build *servers*. Some made it easy to write servers, some made
it possible to write efficient servers, and others somewhere in between.

Regardless, none are well suited to build *clusters*. Of course that problem
has been addressed with frameworks, new languages or new runtimes. By no means
are these languages ignored now.

# Building Clusters

Languages like [Go](http://golang.org), [Akka](http://akka.io) and
[Erlang](http://erlang.org) have come up now because they help solve the
hard networking and concurrency problems that we need to build clusters.

They keep lots of the syntax and semantics from the "languages of yesterday".
For example:

- automatic memory management
- polymorphism
- syntax

But they all build on top with built-in concurrency features:

1. non-blocking I/O
2. user-space concurrency
3. easy messaging
4. advanced scheduling

These 4 concurrency features make clusters **much** easier to build because
they take a large amount of cognitive load off of the developer,
leaving it to the runtime.

When we can stop thinking about single process
concurrency as much, we can start thinking about clustering issues more. That
advancement, I believe, is the key to enabling more powerful and functional
clustering technologies, and better apps built on top of them

# A New Era

The concurrency features don't come for free, however. The concurrency
trade-offs are very similar to those with memory in garbage collected languages.

In both cases, the runtime will take responsibility and control from the
developer to manage a resource (memory and CPU, respectively).

When we're building a cluster, we're often interested in correctness,
latency and throughput of the *system*. This concept is best explained by
an example: If by porting our system from Go to C we could squeeze 5% more
CPU utilization or remove 10% memory utilization, we won't. **We're thinking
first about clusters now, not processes.**

In a cluster, we're not optimizing first for single-process performance.
Instead, we're optimizing to get our systems right.

We're building concurrency into our languages and runtimes because we need
it for the systems we're building today.
