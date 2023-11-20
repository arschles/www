---
author: "Aaron Schlesinger"
date: 2019-12-20T17:14:30-08:00
title: 'The new Athens, part II'
tags: ["Athens", "Athens Hosting", "Kubernetes", "DevOps"]

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Welcome to the second issue of the transformation of the hosted Athens! In my [very first post](./athens-azure-1), I talked about the "old" (technically, current) infrastructure behind `athens.azurefd.net`, why it's worked for so long, and what architecture I'm moving it to (hint: you can see a working preview of it at `k8s.goreg.dev`).

Today I'm going to zoom in on one of the first details I've tackled to get to where we are now - deployment.

I started `athens.azurefd.net` with a script that called `az container create` to deploy some [ACI](https://azure.microsoft.com/en-us/services/container-instances/) containers to different regions. I had manually set up [Front Door](https://azure.microsoft.com/en-us/services/frontdoor/) to route to those containers, and [CosmosDB](https://docs.microsoft.com/en-us/azure/cosmos-db/introduction) to store all the module metadata.

While I was using newer cloud technologies like containers and NoSQL technologies, I consider this "boring technology" because it passed the smell test. This architecture kept working while I was busy doing other things. I didn't have to think about it. And I certainly didn't have to do "DevOps" on it.
