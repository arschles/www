---
author: "Aaron Schlesinger"
date: 2018-10-31T11:25:58-07:00
title: 'Give Away Your OSS Power'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
meta_img: "/images/give-away-power/bus-factor.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Pretend you're an OSS maintainer - how do you figure out if your project is doing well? Some ideas:

- Number of contributors
- Forks
- Stars
- PRs closed
- Average issues opened per day
- Twitter followers to your project account
- Views to your project website
- All of the above?

And the list goes on...

# The Bus Factor

There's one more big one to add to the list. How many people have to drop out before the project is totally screwed and how hard would those people be to replace? In other words, what is your project's [bus factor](https://en.wikipedia.org/wiki/Bus_factor)?

![bus-factor](/images/give-away-power/bus-factor.jpg)

That should be the first measure of any project's success and health, because everything else is a vanity number if if can all go away in the blink of an eye.

I learned this lesson with [Athens](https://docs.gomods.io) sorta recently when I had to drop out of the project for a week. I had to go to the hospital and get my appendix removed on Monday, and didn't start doing work again til Friday. I wrote on the [THAT Conference blog](https://medium.com/that-conference/on-stepping-away-4a1ab23be68d) how I learned to step away from a project, and how valuable that can be for everyone. But I also learned how crucial it is to not hold all the power on the project.

Early on with Athens, I wanted to set a precedent that I would not be the one with all the power. That meant I wasn't the only one with the login to a thing, and also I wasn't the only decision maker. I was lucky that the precedent stuck early, and we haven't lost it. When we get a new core maintainer, they immediately get the keys to everything:

- On the "maintainers" Github team in the `gomods` org
- Administrative privilege for the `gomods` org
- Admin on our Dockerhub org
- Admin on our Netlify account (for our docs site)

There's another big thing, too. All the other maintainers work really hard to make sure they actually start making decisions in public. It's really important that they and everyone else feel comfortable with that.

If we didn't have that process, I would have probably come back from the hospital to a bunch of stagnant PRs and smart folks just stuck waiting to make progress.
