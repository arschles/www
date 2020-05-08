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

If you don't remember much about it, don't worry. I didn't either... it's been over 6 months! I jumped back into the project, so let's recap, and this time with plenty of emojis.

The Athens that you see at [athens.azurefd.net](https://athens.azurefd.net) is running a modern [Athens](https://github.com/gomods/athens) version, but a global deployment of the service demands some more things than just a database and a server. You'd use something like that if you were running your server for your team.

More on that soon.

## First, the Architecture! 🏠

Here's the (very) high level architecture right now:

```
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

And this is the architecture I'm moving it to:

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

That's it. That's this section.

Moving right along ➡

## Why Kubernetes? 🤔

I asked myself that question. Kubernetes is complicated. That's the first thing. The current thing I'm using is pretty reliable, but I want to play with new toys. 😉

Oh wait, no. There is _actually_ a good reason I'm doing this. I need more flexibility!

I'm running the Athens containers inside of [ACI](https://azure.microsoft.com/en-us/services/container-instances/) right now. it's a reliable service but with that comes limitations. I'm pretty much stuck with running only the Athens servers in a container. With that limitation, I'm stuck doing some things that don't jive with running a reliable server:

- Using `git` to ask GitHub for a list of tags all the time
- Relying on a `git` operation for potentially thousands of requests at one time
- Trying to `git clone` a repo while someone is waiting for their dependency to come back

See a pattern? Pro-tip: don't use `git` for the backbone of your global service

Most of those things above "don't happen that often" because a lot of the stuff that people need is behind a caching proxy. But when they do, the containers either run out of memory and crash, or just crash because they feel like it. And then I have to hope the cloud restarts them quickly. 🤞

## So, What Are You Gonna Do About It?

Fair question! In short, we're gonna use Kubernetes to take the current stuff from 😭 to 👍🎂🎉🥳. Ok, it's not that bad right now but you get the point.

I need to run more services to add some reliability to the deployment.

>Hint: Kubernetes is really good at running, maintaining, and connecting services together.

The most important service I need is an [external storage server](https://github.com/gomods/athens/tree/master/pkg/storage/external) for Athens. You don't need to know specifically what that is, just that it's an unreleased feature and that I'm a cutting edge genius for knowing about it.

![i am so smrt](/images/homer-i-am-so-smrt.jpg)

Or, one of the other Athens maintainers told me about it. It's hard to say which one 😛.

## Progress!

That server is gonna be easy to build. Most of the work is [copypasta](https://en.wikipedia.org/wiki/Copypasta)-ing code from Athens and making it my own. That's how all open source projects succeed, right?

On to deploying it! If you haven't deployed to Kubernetes before, just know that the whole system is written in YAML. I'm pretty sure they invented their own CPU in YAML.

So before you get started, make sure you update your LinkedIn profile before you do. Your old skills don't matter. It's all "YAML engineer" from here on out.

Ok just kidding. You can also write [Terraform](https://www.terraform.io/) files and that makes it a little easier.

So naturally I chose to write YAML.

![cmon man!](/images/ditka-cmon-man.jpg)

## Terraform + GitHub Actions 🍤

Just kidding. I went with Terraform 😅.

>If you want to cut to the chase, [here are all my Terraform scripts](https://github.com/arschles/athens-azure/tree/tf-azure/tf)

Here's how they work:

- They have a bunch of variables in them so that I can control where I deploy to from [GitHub Actions](https://github.com/features/actions)
- [GitHub Actions](https://github.com/features/actions) runs on every push I do to that repo. That would be the way I deploy a new version of Athens, for example
- The GitHub actions YAML file figures out if it's a pull request or the master branch. If it's a PR, it deploys to a staging AKS cluster. Otherwise, it deploys to a production one
    - This part is in progress. Check out [my Terraform file](https://github.com/arschles/athens-azure/blob/tf-azure/.github/workflows/terraform-pr.yml), and shield your eyes from my sheer Terrafomrm brilliance 😂
- ???
- **Profit**

## That's It?

Yup! Terraform, AKS, and all the cloud stuff that's backing this make it pretty easy. My Terraform scripts and GitHub actions are just tying them all together. You already know that I'm using some of these, but if you're interested, below are all the cloud services I'm tying together, along with all the fancy acronyms so you can sound like a boss.

![like a boss](/images/like-a-boss.jpg)

- [Azure container registry (ACR)](https://azure.microsoft.com/en-us/services/container-registry/) - for storing all the Docker containers of my code. Not Athens, though. That's already [on DockerHub](https://hub.docker.com/r/gomods/athens)
- [Azure Storage](https://azure.microsoft.com/en-us/services/storage/) - for storing the current state of my cluster. Terraform uses that to figure out what to update
- [Azure Kubernetes Service (AKS)](https://docs.microsoft.com/en-us/azure/aks/intro-kubernetes) - for hosting the actual Athens and supporting services
- [Azure Front Door](https://azure.microsoft.com/en-us/services/frontdoor/) - for doing all the fancy TLS/SSL stuff and caching module ZIP files so my poor little Athens instances don't have to serve them up over and over again
- [Azure Monitor](https://azure.microsoft.com/en-us/services/monitor/) - for keeping track of all the metrics, tracing, and logging that Athens (and the other services) are going to be spitting out. I haven't integrated this in yet, though

aaaand that should do it. My work continues! Check back soon for an update. It's gonna be fun to see this thing take shape 😊

Til next time 👋
