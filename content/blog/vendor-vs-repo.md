---
author: "Aaron Schlesinger"
date: 2019-02-12T17:18:08-08:00
title: 'Should you use a vendor directory?'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Big changes these days in the Go dependencies ecosystem! Let's review:

- We have a new dependency management system called [modules](https://blog.golang.org/modules2019)!
- The new system includes a download protocol and that's **important** because we can build servers to serve up modules to our programs
- These servers are called [Module Repositories](https://jfrog.com/blog/naming-is-hard-the-quest-for-the-right-name-for-go-module-repository/)
- There are already some Module Repositories on the scene:
  - [Athens](https://docs.gomods.io)
  - [GoCenter](https://gocenter.jfrog.com)

So we have servers that hold our dependencies, but lots of projects use vendor directories to store their dependencies. What gives, and how do we know what to use?
