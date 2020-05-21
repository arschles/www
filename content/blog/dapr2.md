---
author: "Aaron Schlesinger"
date: 2020-05-04T16:40:09-07:00
title: 'What is Dapr? (Dapr Series #2)'
tags: ["dapr"]

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Welcome back to the dapr series!

[Last time](https://dev.to/arschles/how-to-make-your-app-harder-to-write-k35), we talked about the challenges of getting software right. If you need a quick refresher:

- The easiest thing about software is writing it
- The next hardest thing is getting it right
- The hardest thing is not crashing and burning in production

You might ask this very reasonable question:

>But writing code is pretty hard already. Does it really only get harder from there?

And I'll hit you with one-worder:

>yes

### Production?

Before we move on, let me take a second to explain what "production" is:

- It's a [distributed system](https://en.wikipedia.org/wiki/Distributed_computing). Wikipedia says it best: `A distributed system is a system whose components are located on different networked computers`
- Your code getting relentlessly pounded by the forces of nature

(If you have an app with a DB, you're running a distributed system)

In other words, "production" is a harsh, brutal hellscape that'll chew your code up and spit it out like it was nothing 😱.

![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/x7svgfpd5rdqokswg2rf.jpg)

## Uh, thanks a lot, sad man!

Yea, that was a little bit of hyperbole but to be honest, I'm hitting you with all these scary things to ... scare you. But only a bit.

Why? To get my point across, of course! And to recommend some tools to make your life easier now and down the line.

Obviously I'm gonna talk about Dapr here, but here are other good tools you can and should consider. Here are some I can personally recommend from experience:

- [PaaS systems](https://en.wikipedia.org/wiki/Platform_as_a_service) take your code, run it on some internal/proprietary system, and give you a few APIs you can call to do useful things.
    - Azure App Services, Google App Engine, Heroku
- [FaaS platforms](https://en.wikipedia.org/wiki/Function_as_a_service) let you upload a function in the language of your choice to the system, and tell it when to run the function (e.g. an HTTP request to `GET /dostuff`)
    - Azure Functions, Google Cloud Functions
- [Orchestrators](https://en.wikipedia.org/wiki/Orchestration_%28computing%29) take your code and make sure it's always running exactly as you tell it to
    - [Kubernetes](https://kubernetes.io/), [Nomad](https://www.nomadproject.io/)
- [Service meshes](https://en.wikipedia.org/wiki/Service_mesh) take care of making talking over any network more reliable. You get [mTLS](https://www.codeproject.com/articles/326574/an-introduction-to-mutual-ssl-authentication) for free too 🎉
    - [LinkerD](https://linkerd.io/)

## Dapr

Dapr takes some of the above things and builds on some others. The idea is that you run this `daprd` process (and maybe some others) "next to" your code (on `localhost`). You talk to this thing via a standard API (HTTP or [gRPC](https://grpc.io)) to get lots of things done:

- [Service discovery](https://github.com/dapr/docs/blob/master/concepts/service-invocation/README.md)
    - Getting the IP for the REST API that you need, when all you know is the name. Kind of like DNS for your app, but designed for apps in production instead of billions of connected devices on the internet. Design tradeoffs, amirite!?
- [Databases](https://github.com/dapr/docs/blob/master/concepts/state-management/README.md)
    - Storing your data somewhere, of course :)
    - Dapr calls these "State Stores"
- [Publish/subscribe](https://github.com/dapr/docs/blob/master/concepts/publish-subscribe-messaging/README.md)
    - One part of your app broadcasts "I did X" into the world. Another app hears it and does things
- [Bindings](https://github.com/dapr/docs/blob/master/concepts/bindings/README.md)
    - Automatically calling your app when something in the outside world happens
    - Or automatically calling something in the outside world when your app does something
    - These have some overlap with publish/subscribe 
- [Actors](https://github.com/dapr/docs/blob/master/concepts/actors/README.md) 
    - Write pieces of your app like it's running on their own, and let dapr run each piece in the right place at the right time (or keep them running forever!)
    - These are great if you need to break your app up into tons of little pieces
- [Observability](https://github.com/dapr/docs/blob/master/concepts/observability/README.md)
    - Collecting logs and various metrics so you can figure out how to fix your app when it spectacularly fails :)
    - and so you can figure out how it's doing before it fails!
- [Secrets management](https://github.com/dapr/docs/blob/master/concepts/secrets/README.md)
    - Safely storing your database password, API keys, and other stuff you'd rather not check into your GitHub repo 
    - _And_, getting them back when your app starts up

Each of these things is a standard API backed by pluggable implementations. Dapr takes care of the dirty details behind the scenes. Database implementations take care of getting the right SDK, hooking you to the right DB in the right place, helping you make sure you don't accidentally overwrite or corrupt your data at runtime, and more.

As I write this, there are [15 implementations](https://github.com/dapr/components-contrib/tree/master/state) you can use for Databases, from Aerospike to Zookeeper.

Possibly most importantly, you don't have to pick up any SDKs to use any of this stuff. It's compiled into Dapr itself.

## Next up

If you took nothing else away, just remember that you should pick up "distributed systems tools" to help you build your app. Even if you don't think that your app is a "hardcore distributed system". The best ones will make your life easier _while you're building it_, too.

I think Dapr is one of those tools. Right when I start building, I can just use the simple(ish) REST API for my database instead of picking up some new SDK and reading all the DB docs up front. Same goes for some of the other features, they just come later. I recommend giving [the README](https://github.com/dapr/dapr) a skim and seeing if it fits your app.
