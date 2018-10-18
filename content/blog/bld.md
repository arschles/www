---
author: "Aaron Schlesinger"
date: 2018-10-08T16:19:36-07:00
title: 'az acr build or docker build?'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

I had heard about [`acr build`](https://cda.ms/HB) when it came out and didn't pay much attention to it. I was like "I don't wanna send my whole build context over the internet that's crazy." That was kinda a facepalm in retrospect because it's kinda sorta my job to keep up on Azure things that are container related.

Anyway, now that I actually _looked_ at this thing, I have some feels and codes!

# How I Un-Facepalmed Myself

I was looking at [Brian](https://twitter.com/bketelsen)'s [acrtasks repo](https://github.com/bketelsen/acrtasks) the other day and _really_ learned what Azure Container Registry tasks are about. You run `acr build` and the CLI zips up your local build context, sends it to the server, and that does the `docker build` for you. Not that exciting, but as Brian said, it'll automatically re-build when base images change. Rad!

Here's some more rad things that I like even better than that:

1. I don't need a local Docker daemon!
1. I don't have to push image layers up to a registry
1. I don't have to do registry auth - just auth to Azure
1. I don't need a local Docker daemon
1. I don't need a local Docker daemon :P

Basically, I can just write my `Dockerfile` and run `acr build` without setting up the daemon. I'd want to on my dev machine, but I don't have to figure it out in other places:

- CI/CD systems like [Buildkite](https://buildkite.com/) or [Travis](https://travis-ci.org/). Travis has Docker support but you have to spin up a whole VM. You can run [Buildkite agents on Kubernetes](https://github.com/helm/charts/tree/master/stable/buildkite), but talking to the Docker daemon from inside a pod is crazytown (I'll die on that hill, btw :P) 
- I repeat: don't let pods talk to the external Docker daemon. [DinD](https://github.com/jpetazzo/dind) isn't much better on Kubernetes
- Underpowered systems - Ever try building [Buffalo](https://gobuffalo.io) on a Raspberry Pi? Not fast, kinda painful

Basically, I just feel like not having to run a Docker daemon sometimes is one less thing to worry about. And it's a big thing.

# Enter [`bld`](https://github.com/arschles/bld)

Getting real again, I basically wanted to do `docker build`s on my dev machines and `az acr build`s playing with Buildkite and stuff. I pretty quickly got tired of writing the same bash script to decide which one to use. So I wrote this little, super basic tool that just does the `if` for me.

Basically you do this:

```console
bld -t my/image .
```

It does `docker build` if the Docker CLI is on the system, otherwise does `az acr build`. It's _suuuper_ basic and that's kinda the point (at least right now). You decide whether you want `docker build` or `az acr build` based on whether you have `docker` (the CLI) installed.

## Other Stuff Already Exists

Yup, and honestly lots of it is better and I might steal it :P

The coolest thing I know of is part of [`draft`](https://draft.sh). The rad [Matt](https://twitter.com/bacongobbler) saw this and DM'd me with [this code](https://github.com/Azure/draft/tree/master/pkg/builder) in the project. The basic difference between `bld` and `draft` is that the former is just for building Docker images. The latter is for everything from coding to deploying to Kubernetes. By the way, Draft is rad and you should check it out :)

I want to keep this thing simple and I'm not sure whether to ~~shamelessly steal~~ use that code or just keep relying on `az`. [To be determined](https://github.com/arschles/bld/issues/1) I guess!

Keep on rockin'

