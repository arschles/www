---
author: "Aaron Schlesinger"
date: 2020-04-21T11:13:56-07:00
title: 'How to make your app harder to build'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

> This is the first post in a series about [Dapr](https://github.com/dapr/dapr)

Software is very easy to write. For lots of us programmers, it's really exciting to create something _new_.

Then it grows and changes over time. Like the foundation of a house, it looks indestructible at first, but the earth has forever to try and break it. It moves, causes cracks and eventually water leaks. The same thing happens to software.

>Software is easy to write but hard to get right

## Challenges of distributed systems

Look at any non-trivial, modern app, and you might be able to find 25% of the bugs that it has in it. You'll have no idea about the rest of them, because time and pressure can only find them.

Your app gets twice as hard to get right if it has to talk to other APIs, databases or microservices. Then you'll hit network blips, slow APIs, out-of-memory errors, and so on. In other words, the world will eventually find a way to break your app. You have to find a way to handle all of that.

## Frameworks

That's where distributed frameworks come in. For modern distributed systems, these are sidecar processes/containers that give your app a simple API that helps you deal with these things. 

They help you:

- Save things to a database
    - Even if there's a problem with the database
- Talk to other services
    - Even when you have to try again
- Tell you what services your app is talking to
    - Even when you can't tell from the code
- Responding to events like pub/sub
    - Even when you don't expect that many events

[Dapr](https://github.com/dapr/dapr) a new framework that can help with these things. Stay tuned for my next post to learn more!


