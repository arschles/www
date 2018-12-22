---
author: "Aaron Schlesinger"
date: 2018-12-20T17:45:40-08:00
title: 'Athens and Microsoft'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
meta_img: "/images/athens-gopher.png"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

As you might have read, Google recently [announced](https://blog.golang.org/modules2019) some of their plans for Go modules in 2019. And because I'm obsessed with module proxies/registries/repositories (more on the naming of these things in a future post), I paid special attention to the announcements about the services they're launching in 2019. I'm happy that Google is stepping up in a big way to help dependency management in the Go ecosystem.

I'll be writing more in future posts about how Athens works with this new stuff, but I'm focusing here on the Athens community.

Obviously none of the opinions in here necessarily reflect the policies or opinions of Microsoft (my employer).

# Athens is a Community Project

The Go team mentioned the Athens project very briefly in their post. I was obviously disappointed that we weren't mentioned a little more, but it's their announcement and their blog, so I respect their choice. I was, however, very disappointed in how they referred to the project as "Microsoft's Athens project."

Since that post came out, I've seen and personally gotten lots of questions on what they meant by that.

I regret that I have to write this article, but the Athens community means more than anything else to me. Because the Go team's post didn't acknowledge that the vast majority of amazing Athenians are not Microsoft engineers, I'd like to clear that up here.

>Quick note: I talk to lots of folks on the Go team regularly and I like and respect them. I'm not trying to call any one of them out here, instead I want to set the record straight on Athens and to acknowledge the amazing community we have.

# Athens is a Community Project, Not Microsoft's

I wrote the first prototype of Athens in my free time. After I moved it out of my personal organization into the `gomods` org, I focused on growing a diverse open source community outside of Microsoft to work on Athens. The `gomods` org is administered by everyone on the [Athens maintainers team](https://github.com/orgs/gomods/teams/maintainers/members).

Microsoft pays me now in part to work on Athens (I have other unrelated responsibilities as well), and I'm empowered to keep working on Athens and help grow the community as I have been.

There's nothing proprietary in Athens and it's MIT licensed. Both me and lots of other folks involved use Athens (or have plans to) and contribute bugfixes and features upstream to `gomods/athens`. Everything with the project happens on Github, in the open.

This is the part that I absolutely want you to read. If you want more details, read on.

# More Details

At the moment, me, [@bketelsen](https://twitter.com/bketelsen) and [@carolynvs](https://twitter.com/carolynvs) are the 3 folks from Microsoft on the [core maintainers team](https://github.com/orgs/gomods/teams/maintainers/members) (there are 8 total), and the MS maintainers don't have special voting rights or influence over the project.

We also have [clear guidelines](https://docs.gomods.io/contributing/community/participating/) on how anyone can advance from community member to contributor to maintainer, and there's no language in that document that indicates that you have to work at a specific company or organization to advance. We have a track record on both the contributors and maintainers teams to prove that as well.

One of our contributors, [chriscoffee](https://github.com/chriscoffee), submitted an [issue](https://github.com/golang/go/issues/29361) two days ago to `golang/go` asking that the "Microsoft" be removed from "Microsoft's Athens Project" on the blog. There's been a little bit of movement on it since then.

I hope that until Google removes that language from their post, this post will be enough to show that Athens is not Microsoft's project. It is the Go community's project, and Microsoft engineers contribute (some of them get paid to contribute) to the project.

# More to Come

I'll be writing more about the technical aspects of Athens and specifically how it fits into the awesome new module-ey stuff in the community, particularly JFrog's GoCenter and of course the three new Google announced technologies.
