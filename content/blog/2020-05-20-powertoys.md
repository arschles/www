---
author: "Aaron Schlesinger"
date: 2020-05-20T11:39:16-07:00
title: 'Windows PowerToys Part 1'
slug: '2020-05-20-powertoys'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

If you were around for Windows 95, you might remember [the PowerToys from then](https://socket3.wordpress.com/2016/10/22/using-windows-95-powertoys/). Well, it's baaaack!

Let's talk about what it's all about today.

>You might have read that I recently [came from a Mac](https://arschles.com/blog/coming-from-a-mac-to-windows-wsl-2/). I have a follow-up post to this on how to use PowerToys Keyboard Shortcuts to make your keyboard more mac-like

# 🤯❔

PowerToys is a [few things in one package]((https://github.com/microsoft/PowerToys#current-powertoy-utilities)), and different folks describe it differently.

The [official GitHub repository](https://github.com/microsoft/powertoys) says this:

"...set of utilities for power users to tune and streamline their Windows experience for greater productivity."

Well, I love productivity (if you've seen any of my screencasts, I hit the keyboard shortcuts **hard**), so I'm gonna go with that!

>😎 By the way, PowerTools is 100% open source at [github.com/microsoft.com/PowerToys](https://github.com/microsoft.com/PowerToys)

I just installed this thing a few hours ago and started messing around with it. It's at version v0.18.0, but the install was not as bad as I thought. Let's go over it!

## 🍟 Get a Package Manager

There are a [few ways to install it](https://github.com/microsoft/PowerToys#installing-and-running-microsoft-powertoys), but I decided to use [WinGet](https://cda.ms/1hP) because it's also new and I hadn't used it yet.

My colleague has [detailed instructions](https://www.thomasmaurer.ch/2020/05/how-to-install-winget-windows-package-manager/) for installing WinGet, but the gist of it is you need to be a [Windows Insider](https://cda.ms/1hQ) (free and open for anyone to join, despite how the name might sound 🎃) and then you can download it from the [Windows Store](ms-windows-store:/pdp/?productid=9nblggh4nns1). That's the easiest way.

>WinGet is also open source [here](https://github.com/microsoft/winget-cli)

## Use WinGet to Install It

When you have it installed, getting PowerToys is a command:

```powershell
$ WinGet install powertoys
```

>Run this in PowerShell, not WSL2

## Get Used to the Features

The very first thing to do is bask in the awesomeness of "PowerToys Run". Seriously. Just hit left-alt+space and a search bar pops up. You can type anything in there and quickly get to "things" on your computer (hint lots of stuff you don't actually need to get to, or even knew you had) or more importantly, YOUR APPS 🎉🤘🏄‍♀️. 

Anyway, hit left-alt+space, type in "PowerToys", and hit enter. You're in the PowerToys settings! It looks like the standard Windows settings app, but it's actually the _added_ features that PowerToys gives.

Check out that left bar. 7 new features plus a "General" thing at the top. I'm gonna boil each complex feature down into a single sentence, Aaron style™. Here goes:

- General - go here to do general stuff (😂), most importantly getting updates to PowerToys itself.
- FancyZones - Automatically resize each window in different patterns on the screen, like side by side or in rows
- File Explorer preview - See rendered previews of SVGs and rendered Markdown in the file explorer. Optional second sentence 💰: I imagine there will be more formats in the future
- Image resizer - resize any image with a right click from the file explorer
- Keyboard Manager - re-map any key or keyboard shortcut with a nice UI
- PowerRename - rename tons and tons of files by highlighting them all and right clicking
- PowerToys run - basically Alfred or Spotlight for Mac, but for Windows
- Shortcut guide - a full-screen pop-up that reminds you of the keyboard shortcuts you have set up

