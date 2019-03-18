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
- The module repository concept is pretty 

So we have servers that hold our dependencies, but lots of projects use vendor directories to store their dependencies. What gives, and how do we know what to use?

# What's the Point of a Vendor Directory?

If you've seen me speak about Athens, I always talk about GitHub not being a CDN, and I give a horror story: what is a maintainer changes their code out from under you??? Your app will break! O noes!!!

Well, even though it

d how if any maintainer changes their code, it can break your app. All of that is true, and before module repository servers came along the best way to prevent those situations was to download all of your app's dependencies - the code it needs in order to compile - and literally check in all that code. 

Where do you put all that code? In the `vendor/` directory, of course! When you run a `go build`, Go will check first for the dependencies your code needs in `vendor/` (it moves on to other places if it doesn't find them in vendor).

Soooo, we're in a pretty good situation now. We have all our dependencies right next to our code. They only change if we decide they should (and there are tons of great tools in the community to help us manage the changes if we want), and our builds can be fast because we don't have to make any network requests to get the code we need.


# Vendoring is the worst kind of forking
