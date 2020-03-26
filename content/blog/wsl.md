---
author: "Aaron Schlesinger"
date: 2020-03-26T10:00:41-07:00
title: 'Coming from a Mac to Windows & WSL 2'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

I've spent most of my life as a professional programmer using a potpurri of Apple laptops and desktops. The first machine I used to do paid development work with was a MacBook in 2004. I've cycled through a few Mac Minis and MacBooks since then, and I'm typing this on yet another Mac Mini.

I don't love Mac OS or Apple hardware per se, but I'm just so _used_ to the Apple ecosystem. I'm talking about everything from knowing the ins and outs of an iPhone, all the way down to details like the many ways to use the `cmd` key on the terminal.

Basically, Apple products rule everything around me

>I recently even wrote about my [Mac Dev Setup](https://arschles.com/blog/my-mac-dev-setup). This is a story about why I'm using WSL 2 more and more for my daily work. I'm not going to be ditching the Mac any time soon though :)

## Windows...

Until now!

Apple has been taking a little bit of flack in the developer ecosystem. Lately it's been the keyboards and there have been some other minor hardware issues in the past. They've also made some software decisions that affect me in my day-to-day life.

My list of grievances isn't huge, but it got me thinking that what I really want is a Linux environment to do coding in and macs aren't the only game in town. I'd go with Ubuntu, but there's something shiny out there that I wanted to try - WSL2!

I haven't done anything on Windows since 2008, so I figured I would flail around, but it turned out to be fairly easy to get my day to day development etc... done.

>Doesn't hurt that I work at Microsoft, and unsurprisingly, Windows works _really_ well on all of our internal work stuff.

## My setup

First off, I built a new machine just to run Windows on. I'll let a picture do the talking. I love this thing.

![windows tower](/images/wsl/windows-tower.jpg)

After I installed windows, the very first thing I noticed was muscle memory. I'm still catching myself reaching for the `cmd` key and I'm retraining myself to hit `ctrl` instead. C'est la vie!

### VS Code

I was never an Emacs or Vim person. I like a nice looking GUI that I can write code in and I'm used to VS Code, so I picked it up and it feels just as comfortable as it did on Mac. Sometimes I feel homesick and I just open VSCode to feel better =P

![VSCode Screenshot](/images/wsl/VSCode-Screenshot.png)

>VS Code has an [extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl) that bridges the Windows world with the WSL 2 world (I'll get to that next), so you can use the Windows app to write code that runs on Linux...

### WSL 2

This is the big one. My whole development environment aside from VS Code relies on a Linux terminal. [WSL 2](https://docs.microsoft.com/en-us/windows/wsl/wsl2-index) is a Linux VM that runs on Windows.

I'm not a VM expert by any means, but from my experience with [VirtualBox](https://www.virtualbox.org/) on Mac and cloud VMs, the WSL2 VMs aren't the same thing. Startup times are just about the same as regular apps, and resource utilization grows and shrinks just like a regular app.

### Windows Terminal

WSL 2 VMs are distributed as apps for each Linux distro. For example, right now I have an "Ubuntu" app and a "Debian" app (I'm writing this from the Debian one). If you open one of the apps, you get a terminal in the VM of your choice. It works but you can't customize it and your distros are all scattered each in their own app.

[Windows terminal](https://github.com/Microsoft/Terminal) is the open source terminal that gathers them all together and it supports the Windows console and Powershell too (hint: you need Powershell to set up WSL 2!).

This + WSL 2 + VSCode feels just about the same as my [Mac setup](https://arschles.com/blog/my-mac-dev-setup/) and I can safely say that my muscle memory is the only thing killing my productivity at this point :)

### My dotfiles

Last thing! I use [yadm](https://thelocehiliosan.github.io/yadm/) to set up my new Mac and Linux environments. I store all of the scripts that yadm uses in my [dotfiles repo](https://github.com/arschles/dotfiles), so I went back and gave everything a shine, so it works really nicely on any Debian/Ubuntu VM on WSL 2.

I tend to tear down my WSL 2 VMs a lot, so it's nice to set up a new one, set up SSH keys and just run `yadm clone git@github.com:arschles/dotfiles.git` to have my dev setup all ready to go.

If you're new to WSL 2 (or even Debian, for that matter), I encourage you to go check out -- and steal! -- the code in that repository.

>I've gone through about 5 WSL 2 VMs so far! I'll be posting details on how I set everything up next.

Happy coding!