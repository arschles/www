---
author: "Aaron Schlesinger"
date: 2019-03-19T17:28:46-07:00
title: 'Serverless? We've Seen This Before!'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Here we go folks, it's buzzword time: **serverless**.

It's a hotly debated topic. Some folks say it's the second coming, some say it's fake, some say it's nothing new. Well, we obviously know it's not fake. Marketing-wise, it's the second coming. Serverless is the topic of many a keynote at industry conferences.

But the most important development is that **experienced, prominent engineers are focusing on serverless technologies**.

## Why?

For me, the most important piece of serverless technologies is that you tell someone else what you want to run. Not where, not how. Just what.

The underlying technologies to enable that workflow are fascinating. Containers, lightweight VMs, cutting edge hypervisor technologies, kernel emulators, the list goes on! So cool, right? That's why we see these awesome engineers working in this space. The technologies and the new workflow opens up a new world of devops and app develompent.

## Where Have We Seen Serverless Before?

Let's focus on the workflow part. Imagine you're writing code with a group, maybe on a team developing a product. Could be a massive web application or "shrink wrap" software, it doesn't matter. And imagine you need to keep quality high, so you write some tests. That's pretty standard for all serious engineering teams.

But all this work on quality and testing is worthless unless you actually _run_ the tests. So we build CI systems. Continuous integration. A computer runs your tests for you and tells you if they pass or fail.

Let's break CI down. It runs your code (tests) in response to an event (opening a pull request or committing code to the central repository). You don't tell the CI system where to run the code, you don't set up any infrastructure on each commit, you just tell it to run the tests.

Now we're talking. Serverless!

I argue that CI was adopted so aggressively because it gives you the workflow and the benefits that we talk about so much today. CI is even such a vital part of development methodologies like agile because it _removes so much complexity_ from the standard software develoment lifecycle. You can move so much faster with CI systems. You can move so much faster with serverless. They're one in the same!

## CI Systems Are _Ahead_ Of Serverless Today

And here's the kicker. CI got to serverless first. Wayyy first. Years ahead. And it's still ahead. On the most advanced serverless systems, you write some code, annotate it, choose a runtime, set up some triggers, and probably write a config file to glue all this stuff together. And then you test some of it out (if you're lucky) locally.

Look at what you have to do with CI systems:

- Write your code (tests)
- Write a config file saying what command to run for the tests

From that point onward, every `git push` runs essentially some serverless code that runs your tests.

## Catching Up to CI/CD

CI/CD is so easy these days because it's specific. The incoming triggers are hard-coded. Incoming is `git push` and outgoing is a status badge or pull request status. There's nothing to configure there, and the triggers are always the hard part. As for the runtime, well that's just the language and tools you're used to anyway.

Serverless, like many early technologies, suffers from a lack of established use cases. Once major serverless systems (cloud providers, OSS Kubernetes frameworks, now.sh, and so on) identify use cases, they can take all this awesome underlying technologies and build high level systems that make it just as easy to set up your code to run in the cloud.

How about these?

- Re-encoding audio files every time they're saved to dropbox, and saving a new copy to the same folder
- Re-training a text-to-speech model when new voice prints are sent to a database
- Handling inventory control when a SKU is scanned at a POS (point of sale) system

See how specific those are? They're really not - they just hard-code triggers and use awesome underlying technology we already have.

When it's all said and done, they're 10% _new_ technology and 90% great documentation and polished developer experience.

We are at the point where we have great fundamental serverless technologies, but we have yet to build on the shoulders of giants.
