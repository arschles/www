---
author: "Aaron Schlesinger"
date: 2019-11-27T17:09:38-08:00
title: 'Athens on Azure Kubernetes Service'
description: 'The hosted athens.azurefd.net, now with more Kubernetes!'
tags: ["Athens", "Athens Hosting", "Kubernetes", "DevOps"]

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
meta_img: "/images/athens-gopher.png"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

I've been working on Athens for well over a year now, and I've mentioned before that I have been running the `athens.azurefd.net` server on my own for almost all that time. The infrastructure that powered that server looked like this:

```console
you ---> Azure Front Door
                |
                |
                |
        ACI Running Athens
                |
                |
                |
        Azure Cosmos DB
```

If you're not familiar with the acronyms, here are some short explanations for you:

- [ACI](https://azure.microsoft.com/en-us/services/container-instances/) stands for "Azure Container Instances". It's a handy way to run a group of Docker containers together in Azure, and get a public IP & hostname for them
- [Front Door](https://azure.microsoft.com/en-us/services/frontdoor/) is Azure's reverse proxy. It does a bunch of stuff, but the things that I care about are TLS termination and global caching. I've written and spoken about front door a lot in the past, and I'm a big believer in it. I think it's one of Azure's most handy services
- [Cosmos DB](https://azure.microsoft.com/en-us/services/cosmos-db/) is Azure's global "NoSQL" database. It supports the MongoDB API and does replication across Azure cloud regions very nicely, with no intervention from the operator (AKA me!)

>It's important to note that Athens can run anywhere. I work for Microsoft so I chose to run it on Azure infrastructure and with CosmosDB as the database. I might consider moving to using Azure blob store later.

# Original requirements

When I set up this public proxy, I wanted to give anyone interested an opportunity to try Athens without installing anything. I think it did exactly that, and it turns out the proxy also got Gophers to start thinking about modules and get their local environment set up.

This deployment was the first of its kind, too! I had no intention of making it "production ready" though. My goal was to encourage teams and companies to roll out their own hosted servers, so the Athens team could focus more on empowering teams to run their own _internal_ Athens servers. I needed a deployment that would mostly run itself.

Those above technologies let me build an architecture that took care of mostly that. Most of the work I had to put into this deployment was managing the transition from Go 1.11 to 1.12 and then to 1.13.

I generally spend no more than an hour a week taking care of this infrastructure, and I've talked to a nontrivial amount of folks using `athens.azurefd.net` in their production CI/CD pipelines. That's a testament to how well this architecture worked over time.

# New requirements

These days, there are [three](https://proxy.golang.org) [good](https://gocenter.io) [proxies](https://gonexus.dev) that I know of which are public and mostly production ready proxy. In theory, my deployment isn't really necessary anymore. That being said, I firmly believe that Athens needs to live on the public internet as a testbed for our software.

I still won't commit to it being production ready, so use it at your own risk (I'm sure some folks will!) But it's not going anywhere.

Not only is it going to stick around for testing, I'm also going to use the infrastructure as an educational tool. Instead of running Athens on ACI, I'll be moving the hosting infrastructure to Kubernetes. This move will open up a lot of flexibility in how I run, debug and deploy the service. Here's a high level diagram for what I'm planning to do:

```console
you ---> Azure Front Door
                |
                |
                |
Kubernetes ingress controller in AKS ---> Athens deployment (pods)
                                                |
                                                |
                                             CosmosDB
```

And here are some things you don't see:

- This simple [script](https://github.com/arschles/athens-azure/blob/71932e2df1c226163c9c62c0024e0809aca27b1b/aci.sh) deploys to a (few) regions that Athens currently runs. I'll be using Terraform to deploy to one (and more in the future!) [AKS](https://docs.microsoft.com/en-us/azure/aks/) cluster
- I'll likely use a service mesh system to control 
- I'm planning to use Athens' built-in distributed tracing support to help with debugging
- I'm planning to use [lathens](https://github.com/arschles/athens-azure/tree/master/lathens) to cache the `/list` and `/latest` API calls for each module
- I'm planning to use [crathens](https://github.com/arschles/athens-azure/tree/master/crathens) to crawl upstream VCS repos and proxies to keep modules up to date
- I'm planning to centralize logs somewhere. Not sure where yet!

# What's coming

I've already gotten started. In fact, you can see a running prototype of an AKS-backed Athens at [k8s.goreg.dev](https://k8s.goreg.dev). It's missing a lot of the features in that above list, but it's built approximately the same as the architecture in that diagram above.

All the code to do all of this is [open source](https://github.com/arschles/athens-azure), and this is post #1 of a series that I'll use to chronicle my experience setting up and running Athens.

See you soon!
