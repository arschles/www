---
author: "Aaron Schlesinger"
date: 2020-03-18T17:04:19-07:00
title: 'My Mac Dev Setup'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Recently, I got into a really interesting [discussion on Twitter](https://twitter.com/jongalloway/status/1240165698506067970) with [@johngalloway](https://twitter.com/jongalloway), and it was just about 100% constructive and civil. On Twitter!

It's all about whether or not to use [Homebrew](https://brew.sh/) to set up and maintain a Mac for software development. I started programming on a Mac in late 2002, on Mac OS Jaguar, version 10.2. Since then I've used at least 10 different Macs, which means I've set up at least 10 different Macs.

We developers are picky about our setups, so setting up a brand new vanilla computer takes time and energy.

That got me thinking: I've used a Mac for almost 20 years as my main development machine. What do I do to set up and maintain _my_ machines?

Turns out I don't really, so I should figure that out, and that's what we're gonna do right here. Strap in for the ride!

## Command Line Tools

Like any good developer, I've tried a bunch to automate myself out of work, so a while back I tapped [yadm](https://yadm.io/) to install many of my tools and store, version, apply, and upgrade various configurations for them. I store all that stuff in [github.com/arschles/dotfiles](https://github.com/arschles/dotfiles), so let's start there!

- I use [Hyper](https://hyper.is) for my terminal. I just switched to it from [iTerm 2](https://iterm2.com/) primarily because it's configurable with a single json file. I also like the icon on the dock, not gonna lie
- I use [oh my zsh](https://ohmyz.sh/) (and the Z shell by extension) because it gives me plugins and a really uniform experience across different terminals. I switch to Linux a bunch and I can move my `~/.zshrc` file over and get the same plugins and everything else without any work
- I use [Homebrew](https://brew.sh) a _ton_. I was the one on Twitter advocating for it! Basically, I can give `brew` my [Brewfile](https://github.com/arschles/dotfiles/blob/master/.yadm/Brewfile) and it'll just go install everything I need with no fuss. It acts like `apt` on Debian. In fact, it now works on Linux (that's pretty new) so I can move my `Brewfile` over just like my `~/.zshrc`
- [direnv](https://direnv.net/) is a small little utility that lets me load/unload environment variables just in a single directory. It comes in handy **all over the place**. Ya know every time some getting started docs say "you need to have X, Y and Z variables in your environment"? That's when I reach for direnv and write up and `.envrc` file inside my project.
- [`hub`](https://github.com/github/hub) is a nice little CLI that lets me do GitHub "things". I do a lot of open source work ([athens](https://github.com/gomods/athens) a lot these days) so `hub` can take me from committing a change alllll the way to submitting a PR to the upstream repo. I really only use it for quick PRs because I like to go into the browser to really think through my PR descriptions, but it's really nice for quick things
- [`gh`](https://github.com/cli/cli) is sort of a newer version of `hub`. It's called the "GitHub's official command line too" so my money's on this thing implementing everything that `hub` has. It's not _quite_ there yet, but I'm starting to move away from `hub` to this. The maintainers of the project are [open to feedback](https://github.com/cli/cli#we-need-your-feedback) also, so I feel pretty good about the move 

## Desktop stuff

I try to stay in the command line as much as I can when I'm programming and only break out for really good reasons. I mentioned that I go into the browser when I want to submit pull requests. That's so I pull my head out of the sand and think about the stuff I just built.

Another good reason I pop out of the command line is for my editor. I never got into vim or emacs. I've always used IDEs. In recent-ish memory, I've used almost all of the [JetBrains](https://www.jetbrains.com/) IDEs, [Visual Studio](https://visualstudio.microsoft.com/) (I've written a good amount of C#!), [Sublime Text](https://www.sublimetext.com/), [Atom](https://atom.io/) and now I'm currently on [VS Code](https://code.visualstudio.com/).

I plan to stick with VS Code for a while for my main editor. It's really flexible and seems to have the biggest ecosystem of all of the open source editors/IDEs. Plus there are some cool "add on" features like [VS Online](https://online.visualstudio.com/). Actually, check that one out if you haven't. It's the VS Code GUI on your machine but you run/debug/whatever the code on a cloud VM

And the last one! I use [GitHub desktop](https://desktop.github.com/) a ton. Technically I can do just about everything on the command line instead of using this app, but I have lots of good reasons why a lot of the time I go into this app instead:

- The colors! Not gonna lie, I like the green and red lines showing me in the fancy, shiny Mac color scheme what I'm about to commit
- Partial commits. Being able to commit only _some_ of a file is really easy in the app, and can be super useful. Especially when I implement 5 things at once, pick my head up, and realize I need to split stuff up into 5 separate commits
- Everything in one place, in a UI. This app shows everything that's going on in your project at once. You can see what branch you're on, whether you need to pull code down from the server, how many commits you still need to push, history of commits, etc... all at once. On the command line you have to run a different command for almost all of that stuff. It's nice to see the whole picture sometimes

## One more thing...

Those are the big things. My development life is a hodgepodge of CLIs and GUI tools and it works great for me. I can make quick, focused progress on a PR or issue I'm working on, without many distractions, and that's all I can ask for!

But! I'm about to start challenging myself, and change the most fundamental of all of my tools: the operating system. I'm starting to move to Windows as my main development OS. I'll be using [WSL 2](https://docs.microsoft.com/en-us/windows/wsl/wsl2-index) to install and use many of the same CLI tools I'm already used to, but the rest will be interesting to say the least.

I'm challenging myself to learn (and unlearn!) a lot to get comfortable with a brand new dev setup.

Interesting times ahead - stay tuned :)
