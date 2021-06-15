---
author: "Aaron Schlesinger"
date: 2021-02-08T14:25:54-08:00
title: 'KEDA-HTTP Alpha 1 Is Released'
slug: "keda-http-alpha-1"

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

I last wrote about [KEDA-HTTP](/blog/kedahttp-2) when the project was brought into the [kedacore organization](https://github.com/kedacore). That was a big deal and for me, a very promising sign for the project.

Since then, we've worked hard to realize a high quality cloud-native architecture and build a good technical and community foundation for the project.

We believe that, although this project is not production-ready, we can now signal to the greater community that we think it's reached a quality high enough for developers to try out.

_**Alpha 1 of the [KEDA http-add-on](https://github.com/kedacore/http-add-on) is released today**_

We encourage you to try the software out in your own non-production Kubernetes clusters and tell us what you think and/or contribute back.

Please read below to understand what to expect from this release.

## What This Release Means For You

Most people reading this will likely be familiar with what the term "alpha quality" implies, it's worth briefly outlining specifically what it means to this project. See below.

- The process for submitting and merging PRs and issues is approximately the same as it will be going forward
- The technical architecture has been settled, and will serve as a foundation for future features
- The software is not fully tested (in fact, there are very few tests. See [issue #13](https://github.com/kedacore/http-add-on/issues/13)) and there may be significant bugs. We recommend against using this software in production deployments
- Community members are welcome to contribute, but documentation may not be complete. There are two corollaries to this point:
  1. Community members are welcome to report difficulties getting the project running. These reports help us write better and more comprehensive documentation
  2. Community members are encouraged to help us write new or amend existing documentation
