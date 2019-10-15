---
author: "Aaron Schlesinger"
date: 2019-09-12T13:37:57-07:00
title: 'Go 1.13 for Private Repositories'
image: "/images/gopher-university.png"

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Go 1.13 is a new beast. It's a patch release in semver terms, but there are some big changes going on here. Most of them are really good, but if you have private dependencies, your build is gonna break out of the box.

I'm gonna explain here what's going on and how to fix it, using only two (I know, right??) shameless plugins for my favorite project -- which happens to solve these problems and I also am on the core team for -- [Athens](https://docs.gomods.org).

So, here's the big thing that changed and is gonna cause you headaches:

>This is the first time that the `go` toolchain has out-of-the-box integration with a hosted service

In more detail:

1. Modules are now turned on _everywhere_ by default
1. The default way you download new dependencies is from [proxy.golang.org](https://proxy.golang.org) - not hosted version control systems like GitHub
1. Go now consults the [checksum database](https://sum.golang.org) every time you ask for a new module not in your `go.sum` file

Let's go through these points one by one.

# Modules are now turned on everywhere by default

Everything else being equal, this is amazing. There's been debate, accusations, fighting in the community over how to do Go dependency management. I stayed out of almost all of that because I don't have any more hairs on my head to lose.

Not everyone is happy with modules, but just like every change, that will pass. More importantly, we have a good enough system, that we can make work, and that we can _agree on_.

Standardization is good in software systems! Also, part of the formal modules standard is what makes Athens and all the other module proxies possible.

Speaking of proxies ....

# The default way you download new dependencies is from [proxy.golang.org](https://proxy.golang.org)

This one can be good or bad, depending on the type of codebase you're working in. My explanation is long, so I'm not even gonna be that sad if you just scroll down to the bold text below.

If you want details, read on!

The module spec has a small section that details a module download API, which is fancy way to say "a few GET requests to download source code and `go.mod` files." 

But the API means we can build servers for modules now. So as of 1.11, you didn't have to use `git clone` for you package manager anymore. By the way, how did we last for almost 10 years using `git clone` to install dependencies?

Anyway, I started building one of the [first module servers](https://github.com/gomods/athens) and one of the most widely used open source ones. I also gave a talk about it at GopherCon US this year. My big message was that Athens is one of the tools we're going to use to keep the modules ecosystem decentralized - just like dependencies always have been in Go. I know I already bashed `git clone` (see what I did there?) as a package manager, but it's a distributed versioning system, and we co-opted it to build a distributed package index.

That takes me to the change in 1.13. Before 1.13, you had to set a `GOPROXY` environment variable to tell Go to use this module API and to fetch modules from the server you specify. Otherwise, it would keep doing `git clone`.

Now things are different! `GOPROXY` defaults to `https://proxy.golang.org` as of 1.13. `proxy.golang.org` is a module repository run by the Go team at Google.

This change means that, out of the box, all of your `go build`, `go get`, `go install`, etc... commands will be communicating with Google-run servers to fetch code.

I don't mean to make that sound ominous, I'm not trying to throw any shade at Google, and Google didn't pull a fast one as you might think (or have read?) - this change was on the roadmap like all the others. 

So that's me softening the blow. If you're working in a private repository, this new thing sucks for you. By default, you'll be sending Google the names of your private modules, _and_ your builds will break because the public Google proxy can't find your private modules.

What's the solution? I'm glad you asked, friend! Here's how to fix it, in bold print for people who just skimmed the article, and with a shameless plug for Athens:

**If you have a private codebase, set your `GOPROXY` environment variable to a private module proxy. Where do you get a private module proxy, you ask? Check out the biggest open source one if you're on the market: [Athens](https://docs.gomods.io)**

Ok, and there's one more thing you need to look at...

# Go now consults the [checksum database](https://sum.golang.org) every time you ask for a new module

This section's gonna be really familiar if you read everything in the last one. If you do a `go get` (or similar) and the module isn't in your `go.sum` file (yes, the checksum file), then the Go tool is gonna ask `sum.golang.org` -- the checksum database -- what the checksum is supposed to be.

That's to stop a proxy from giving you the wrong stuff the first time you ask for it. Since they're using their proxy by default, they're actually protecting you from ... themselves. The Go team is the hero Gotham deserves, not the hero it needs.

But seriously, the idea behind the sum database is good - it helps you not get corrupted or tampered-with code, whichever proxy you decide to use. It's a really good tool to encourage a secure, diverse, and decentralized ecosystem.

The problem is you pesky private repositories! The same send-your-data-to-Google bug is gonna bite you here too. There are two things you can do about it, and I'm not gonna bold them because I already let you lazy readers off the hook last section, dammit!

1. Set `GONOSUMDB` to have a comma-separated-list of module prefixes you don't want Go to send to the sum DB. For example, `export GONOSUMDB="myvcs.internal/*,secretevilplans.com/modules/*"`
1. Use Athens (shameless plug #2) to set a global `GONOSUMDB` variable for your team. This will prevent anyone on the team who is misconfigured from spilling the beans to the Goog

---

So that's it for today. Make sure you go understand what's up with the [Google proxy](https://proxy.golang.org) and [Google Checksum DB](https://sum.golang.org) and how they affect your code.

If you have private repositories, _really_ understand them. Your builds will break out of the box if you don't.
