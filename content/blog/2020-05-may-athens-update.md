---
author: "Aaron Schlesinger"
date: 2020-05-05T16:29:07-07:00
title: 'May 2020 Athens Update'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
meta_img: "/images/athens-banner.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

You haven't heard from me in a while about [Athens](https://docs.gomods.io). If you use Athens, you're probably all like 😒, but hopefully this post will help you be more like 😊. Plus, we're more than 4 months into the year and I want to share ✔.

>Before you read on, here's the Athens Gopher for you, because why not!

![Athens Gopher](/images/athens-gopher.png)

>If you just want a summary, we've already had 4 releases this year and I'm excited about some bigger features coming!

## Why are you writing this?

Good question! Well, I've been writing code and some docs on [the site](https://docs.gomods.io) for a while. I haven't actually written about the project in a while 😨. Hopefully this post is helpful 😎.

## Releases

We've had 4 releases in 2020, starting with [v0.7.1](https://github.com/gomods/athens/releases/tag/v0.7.1). That means mostly bug fixes. The "big" feature is in [v0.8.0](https://github.com/gomods/athens/releases/tag/v0.8.0) where we added support for [Redis Sentinel](https://redis.io/topics/sentinel) as a backend for the distributed coordination system (AKA "single flight").

The biggest bugfix we made this year was to make it possible to set multiple values in the `ATHENS_GO_BINARY_ENV_VARS` environment variable ([#1404](https://github.com/gomods/athens/issues/1404))

## Coming Soon to and Athens Near You

I'm happy with things so far but I'm really excited for what's coming. For a while now, we've been thinking about how to build two major new features:

- A completely-offline mode
- Pluggable storages

The first one is still [in progress](https://github.com/gomods/athens/discussions/1538) (click that link to go join the discussion).

The second one is .... [done](https://github.com/gomods/athens/pull/1587)! 🥳🥳🥳🥳

We haven't released pluggable storages officially, so if you want to try it out, grab one of the recent Docker images from the [athens-dev Docker repo](https://hub.docker.com/r/gomods/athens-dev) and go for it!

When we do release it (soon), anyone will be able to write a simple HTTP server for their database of choice, and Athens will use it instead of its own built-in storage systems. The server is the [standard download API](https://docs.gomods.io/intro/protocol/) plus a few other APIs.

Athens isn't ditching any of its built-in storage systems, but we'll likely be building future storage backends as external ones.

Most importantly, **this means that you can use any storage system you want - including internal, proprietary ones - without maintaining your own fork of Athens**

## A Personal Note

I'm really happy with how things are going. Honestly, by the end of 2019, I was burned out. I had a reflex to get involved with everything in the project. That wasn't sustainable for me or the project.

Speaking of sustainability, I keep this proverb in my back pocket:

>If you want to go fast, go alone. If you want to go far, go together

It helps me remember to _listen more than I talk_ and _ask for help_. Both of those things are easy to do with everyone I've met in the Athens community. You are all smart (whether you think it or not), helpful, and good people.

>If I haven't met you, feel free to say hi to me. I'm [@arschles](https://twitter.com/arschles) on Twitter and my DMs are open. Same username on GitHub and the Gophers Slack. Any of those channels work. And, of course I'd love to see you [contribute to the project](https://docs.gomods.io/contributing/community/participating/)!

I'll do more of these updates from here on out, so until next time!

- Aaron💚