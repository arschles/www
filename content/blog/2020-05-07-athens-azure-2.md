---
author: "Aaron Schlesinger"
date: 2020-05-07T16:18:19-07:00
title: 'Athens on Azure Kubernetes Service... Part 2'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Howdy all! It's been quite a while since I wrote about [Athens on Azure Kubernetes Service](https://arschles.com/blog/athens-on-azure-kubernetes-service/).

If you don't remember much about it, don't worry. I didn't either... it's been over 6 months! I jumped back into the project, so let's recap and look to the future, and this time with plenty of emojis.

## The Deployment Right Now

The Athens that you see at [athens.azurefd.net](https://athens.azurefd.net) is running a modern [Athens](https://github.com/gomods/athens) version, and that's about it.

It's supposed to be a global deployment, but it does some things that aren't great for something like that behind the scenes. Right now, it's a proxy, container and a database. We need to be doing more than that.

## Current Architecture Limitations

I'm gonna pull an old PR trick here and hit you with the bad news first, and end on a good note. This section is all about what we _can't_ do right now.

In case you forgot, here's the current architecture of the deployment: 

```
you ---> Azure Front Door
                |
                |
                |
        ACI Container Running Athens
                |
                |
                |
        Azure Cosmos DB
```

In other words, I'm running the Athens containers inside of [ACI containers](https://azure.microsoft.com/en-us/services/container-instances/) right now. Each container uses some other cloud services to do the rest of their work.

The whole system is a semi-reliable service but comes with some limitations, and when the containers crash, those limitations usually rear their ugly head. When I'm fixing a crash, I usually say something like "I could fix this if I could set up `$USEFUL_NEW_THING` ... but ACI doesn't let you".

>[ACI](https://azure.microsoft.com/en-us/services/container-instances/) is generally good for quickly launching one-off containers, specialized workloads, or using as "burst" computing capacity from a Kubernetes cluster (you can do that with [Virtual Kubelet](https://github.com/virtual-kubelet/virtual-kubelet)). They're less good for long-running servers that need to scale

Here are some of the worst things that Athens _has to do_ right now:

- Use `git` to ask GitHub for a list of tags all the time
- Rely on a `git` operation for potentially thousands of requests at one time
- Try to `git clone` a repo while someone is waiting for their dependency to come back

See a pattern? (Pro-tip: don't use `git` as the backbone of your public server. It wasn't made for that)

Most of those things above "don't happen that often" because a lot of the stuff that people need is behind a caching proxy, so Athens doesn't get the requests that trigger the `git` operations. But when you run a service long enough, the "don't happen that often" things eventually happen, often.

>All those bad things that might happen in production will happen, just like water will eventually carve out the grand canyon.

When problems crop up, the containers either run out of memory and crash, or just crash because they feel like it. And then I have to hope the cloud restarts them quickly. 🤞

## Why Kubernetes? 🤔

I asked myself that question. First off, Kubernetes (AKA "k8s") is complicated. If you get over that complexity, you can get some real power, though.

I've been working in the Kubernetes space in some capacity since early 2016, so I tried to look at this decision as pragmatically as possible.

- 🍏 PRO: The current setup has outages that greatly concern me
- 🍏 PRO: k8s gives me more architectural options than ACI
- 🔴 CON: more options = more moving parts = more complexity = more things to break in production with k8s
- 🔴 CON: There's always something changing or something new to learn in the ecosystem. It's confusing!
- 🍏 PRO: I'm a bit of a grouchy old man and I'm confident I'll avoid most of the "bleeding edge" stuff that's almost certain to break/change. But I _will_ nerd out on it.... 😏
- 🔴 CON: There's a lot of YAML to learn (more on this later)
- 🍏 PRO: I already know a lot of it, enough to not write any this time around (did I mention I'm grouchy??). More on this later
- 🍏 PRO: I read a [Basecamp article](https://m.signalvnoise.com/seamless-branch-deploys-with-kubernetes/) recently describing their usage of Kubernetes for their new email service. I highly respect their pragmatic view on new technologies

That's 5 pros to 3 cons. Honestly, the Basecamp article tipped me over the edge though. I respect the technical decisions that team makes over almost any other group of people.

Let's get to work ⚒.

## The Plan

In case you forgot the high level architecture from last post, here ya go!

```
you ---> Azure Front Door
                |
                |
                |
Kubernetes ingress controller in AKS ---> Athens deployment (pods)
                                                |
                                                |
                                             CosmosDB
```

Let's zoom in on the `Athens deployment (pods)` part in that diagram. I need to run more services to add some reliability to the system, and that part is where it gets interesting.

>Hint: Kubernetes is really good at running, maintaining, and connecting services together.

## What New Services Do We Need?

The most important service I need is an [external storage server](https://github.com/gomods/athens/tree/master/pkg/storage/external) for Athens. You don't need to know specifically what that is, just that it's an unreleased feature and that I'm a cutting edge genius for knowing about it.

![i am so smrt](/images/homer-i-am-so-smrt.jpg)

Or, I worked with the other Athens maintainers on it. I actually didn't do that much of the work. I'm gonna go with genius though 😛.

There's another service, too, called the "crawler for Athens", also known as [crathens](https://github.com/arschles/athens-azure/tree/master/crathens). Same thing here, you don't really need to know what that's all about for now. I'll go deeper into architecture in a future post.

So now we have Athens, the storage service, and crathens. 

Dare I say we're building _microservices_????

## Progress!

That storage server is gonna be easy to build. Most of the work is [copypasta](https://en.wikipedia.org/wiki/Copypasta)-ing code from Athens and making it my own. That's how all open source projects succeed, right?

Crathens will need a little work, but it's pretty much good enough™ for now.

So at this point, I'm working on deploying this stuff. If you haven't deployed an app to Kubernetes before, just know that the whole system is written in YAML. I'm pretty sure they invented their own CPU in YAML.

>If you're gonna wade into the Kubernetes waters for the first eim, make sure you update your LinkedIn profile before you do. Your old skills don't matter. It's all "YAML engineer" from here on out.

Just kidding. You can also write [Terraform](https://www.terraform.io/) scripts if you want. I think that makes it a little easier.

So naturally I chose to write YAML.

![cmon man!](/images/ditka-cmon-man.jpg)

## Terraform + GitHub Actions 🍤

Just kidding. I went with Terraform 😅.

>If you want to cut to the chase, [here are all my Terraform scripts](https://github.com/arschles/athens-azure/tree/tf-azure/tf)

Here's how they work:

- They have a bunch of [input variables](https://www.terraform.io/docs/configuration/variables.html) in them so that I can control where I deploy to, all from [GitHub Actions](https://github.com/features/actions)
- [GitHub Actions](https://github.com/features/actions) runs on every push I do to that repo. That would be the way I deploy a new version of Athens, for example
- The GitHub actions YAML files figure out if it's a pull request or the master branch. If it's a PR, it deploys to a staging AKS cluster. Otherwise, it deploys to a production one
    - This part is in progress. Check out [my Terraform file](https://github.com/arschles/athens-azure/blob/tf-azure/.github/workflows/terraform-pr.yml), and shield your eyes from my sheer Terraform brilliance 😂
- ???
- **Profit**

## That's It?

Yup! Terraform, AKS, and all the cloud stuff that's backing this make it pretty easy. My Terraform scripts and GitHub actions are just tying them all together.

You already know some of what I'm using, but below is a complete list for you. I even incuded all the fancy acronyms so you can sound like a boss.

![like a boss](/images/like-a-boss.jpg)

- [Azure container registry (ACR)](https://azure.microsoft.com/en-us/services/container-registry/) - for storing all the Docker containers of my code. Not Athens, though. That's already [on DockerHub](https://hub.docker.com/r/gomods/athens)
- [Azure Storage](https://azure.microsoft.com/en-us/services/storage/) - for storing the current state of my cluster. Terraform uses that to figure out what to update
- [Azure Kubernetes Service (AKS)](https://docs.microsoft.com/en-us/azure/aks/intro-kubernetes) - for hosting the actual Athens and supporting services
- [Azure Front Door](https://azure.microsoft.com/en-us/services/frontdoor/) - for doing all the fancy TLS/SSL stuff and caching module ZIP files so my poor little Athens instances don't have to serve them up over and over again
- [Azure Monitor](https://azure.microsoft.com/en-us/services/monitor/) - for keeping track of all the metrics, tracing, and logging that Athens (and the other services) are going to be spitting out. I haven't integrated this in yet, though

That should do it! My work continues for now. I still have the storage server to do, and I need to finish up the Terraform scripts.

Check back soon for an update. It's gonna be fun to see this thing take shape 😊

Til next time 👋
