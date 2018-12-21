---
author: "Aaron Schlesinger"
date: 2018-12-20T17:45:40-08:00
title: 'Athens and Microsoft'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

As you may have read, Google recently [announced](https://blog.golang.org/modules2019) some of their plans for Go modules in 2019. And because I'm obsessed with hosted module services, I noticed that they're planning to launch a few important services:

- A module index that provides a public log of new modules in the community
- A module notary that provides authentication for any module in the community
    - This uses the log to proactively calculate certificates to verify module integrity
- A module mirror that provides a server that 
