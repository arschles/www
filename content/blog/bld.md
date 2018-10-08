---
author: "Aaron Schlesinger"
date: 2018-10-08T16:19:36-07:00
title: 'Build Docker Images Locally or in the Cloud'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

I was looking at the Brian's [acrtasks repo](https://github.com/bketelsen/acrtasks) the other day and got introduced to `acr build`. I really got interested in it because I could replace `docker build` if I wanted to. No local Docker daemon needed. I could literally just write a `Dockerfile` and run `acr build` to build my image, and store it in ACR, all in one command. And I don't have to futz with Docker.

Obviously it takes some time for the entire directory to be zipped up and sent (over the internet!) to Azure, so this isn't a good strategy for local development almost ever.

But I've run into several situations where it is:

- Building Docker images in CI/CD systems. I've had to jump through hoops to get Docker daemons running in various CI/CD systems, just so I can build & push a Docker image. ACR builds would solve that
- Building Docker images from insde Kubernetes pods. There's a ton of literature on how to build Docker images in userspace, which lets you do this, but you have to read stuff and learn stuff and install stuff. I found it interesting, but I just plain don't want to deal with that. Enter ACR builds! I use all the same stuff, but just replace `docker build` with `az acr build`. Yay!
- Underpowered machines. Ever try building a decent size Docker image on a Raspberry Pi? I've tried it for a Buffalo app, and it took forever. This probably isn't the biggest use case ever but it's interesting. You can install the Azure CLI and do one less thing on the hardware and one more thing in the cloud. That's always nice when you're on an ARM processor

# Enter `bld`


