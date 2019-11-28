---
author: "Aaron Schlesinger"
date: 2019-11-27T17:09:38-08:00
title: 'Athens on Azure Kubernetes Service'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

I've been working on Athens for well over a year now, and I've mentioned before that I have been running the `athens.azurefd.net` server on my own for almost all that time. The infrastructure that powered that server looked like this:

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

If you're not familiar with the acronyms, here are some short explanations for you:

- [ACI](https://azure.microsoft.com/en-us/services/container-instances/) stands for "Azure Container Instances". It's a handy way to run a group of Docker containers together in Azure, and get a public IP & hostname for them
- [Front Door](https://azure.microsoft.com/en-us/services/frontdoor/) is Azure's reverse proxy. It does a bunch of stuff, but the things that I care about are TLS termination and global caching. I've written and spoken about front door a lot in the past, and I'm a big believer in it. I think it's one of Azure's most handy services
- [Cosmos DB]

Pretty simple, but it worked for all this time!
