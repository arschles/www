---
author: "Aaron Schlesinger"
date: 2019-07-02T16:38:36-07:00
title: 'Tools for Developers on Kubernetes, Part 1'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Kubernetes is hard for app developers to wrap their head around. The big thing that makes it hard is that coming in, you don't know what you don't know. Before you know it, you've dockerized your app and you're reading up on pods and so on.

Put another way, choosing Kubernetes is taking out a second mortgage on your house for more [innovation tokens](http://boringtechnology.club/#17), and going back into debt tomorrow.

# What Makes K8s Hard

The Kubernetes API and deployment tools are low level (on purpose, by the way!), but we write modern software with high level languages, and that requires a big context switch every time we need to deploy. That's what we struggle with.

It's true that you need to spend a lot of energy writing the initial manifests, and that's tolerable! But like any technology, there's an ongoing maintenance cost. Look at all the space below the flat line (an integral if you're math-ey), and you'll see what I mean:

![kube starting and deployment](/images/kube-hard-energy.png)

Think about the risk & cost of adding a new production database to your stack. Kubernetes is about the same risk & cost because both technologies sit between your users and your code.

# Making the Context Switch Smaller

I have a big rant that nitpicks at all the specific 

There's no edit/test cycle for developers

In everyday life as a developer, you find yourself tweaking yaml, submitting it to Kubernetes, and then doing some `kubectl` / `helm` commands to see what happened. That workflow is miles away from the standard edit/test loop that all of us know and love. So that's one thing we need to fix.

The other thing we need to fix is production deployments. If you're a big team, you have time to 

 _We need more tools like that._ But then we also need tools to help deploy our stuff to production.

I'm biased and I love [draft](https://draft.sh) because it  (I'm biased), so I always recommend that folks check it out



