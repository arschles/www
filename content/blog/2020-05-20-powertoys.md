---
author: "Aaron Schlesinger"
date: 2020-05-20T11:39:16-07:00
title: 'Windows, Now With More Toys'
slug: '2020-05-20-powertoys'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

If you were around way back in the Windows 95 days, you might remember [the PowerToys from then](https://socket3.wordpress.com/2016/10/22/using-windows-95-powertoys/). Years and years later, it's back in a similar form, but for Windows 10.

I just got started with it and it's another tool in my belt to help me feel comfortable using Windows every day after [many](https://deploy-preview-70--arschles-www.netlify.app/blog/coming-from-a-mac-to-windows-wsl-2/) [years](https://deploy-preview-70--arschles-www.netlify.app/blog/how-to-wsl-2/) of using Macs.

Let me take you through it.

>You might have read that I recently [came from a Mac](https://arschles.com/blog/coming-from-a-mac-to-windows-wsl-2/). I have a follow-up post to this on how to use PowerToys Keyboard Shortcuts to make your keyboard more mac-like

# 🤯❔

PowerToys is a [few things in one package]((https://github.com/microsoft/PowerToys#current-powertoy-utilities)), and different folks describe it differently.

The [official GitHub repository](https://github.com/microsoft/powertoys) says this:

"...set of utilities for power users to tune and streamline their Windows experience for greater productivity."

Well, I love productivity (if you've seen any of my screencasts, I hit the keyboard shortcuts **hard**), so I'm gonna go with that!

>😎 By the way, PowerToys is 100% open source at [github.com/microsoft.com/PowerToys](https://github.com/microsoft.com/PowerToys)

I just installed this thing a few hours ago and started messing around with it. It's at version v0.18.0, but the install was not as bad as I thought. Let's go over it!

## 🍟 Get a Package Manager

There are a [few ways to install it](https://github.com/microsoft/PowerToys#installing-and-running-microsoft-powertoys), but I decided to use [WinGet](https://cda.ms/1hP) because it's also new and I hadn't used it yet.

My colleague has [detailed instructions](https://www.thomasmaurer.ch/2020/05/how-to-install-winget-windows-package-manager/) for installing WinGet, but the gist of it is you need to be a [Windows Insider](https://cda.ms/1hQ) (free and open for anyone to join, despite how the name might sound 🎃) and then you can download it from the [Windows Store](ms-windows-store:/pdp/?productid=9nblggh4nns1). That's the easiest way.

>WinGet is also open source [here](https://github.com/microsoft/winget-cli)

## 🥨 Use WinGet to Install It

When you have it installed, getting PowerToys is a command:

```powershell
$ WinGet install powertoys
```

>Run this in PowerShell, not WSL2

## 🍕 Get Used to the Features 

The very first thing to do is bask in the awesomeness of "PowerToys Run". Seriously.

**Just hit left-alt+space and a search bar pops up.**

You can type anything in there and quickly get to "things" on your computer. That generally means a metric ton of things that you didn't even know you had. But it also means ... YOU CAN GET TO YOUR APPS WITH NO CLICKS!!!!🎉🤘🏄‍♀️. 

I really missed that from my Mac days.

Anyway, hit left-alt+space, type in "PowerToys", and hit enter. That will take you to something that looks like the standard Windows settings app, but it's actually the _added_ features that PowerToys gives.

Check out that left bar. 7 new features plus a "General" thing at the top. I'm gonna boil each complex feature down into a single sentence, Aaron style™. Here goes:

#### 🍚 General

Go here to do general stuff (😂), most importantly getting updates to PowerToys itself.

#### 🥧 FancyZones

Automatically resize each window in different patterns on the screen, like side by side or in rows.

>Yes, the name is amazing. The zones are indeed fancy 🍰! But the feature overall is pretty slick and powerful once you get the hang of it. You'll want to try out a few different setups to figure out the best one for you.

#### 🍣 File Explorer preview 

See rendered previews of SVGs and rendered Markdown in the file explorer. Optional second sentence 💰: I imagine there will be more formats in the future.

#### 🍤 Image resizer 

Resize any image with a right click from the file explorer.

#### 🍎Keyboard Manager

Remap any key or keyboard shortcut from a nice UI.

#### 🍉 PowerRename

Rename tons and tons of files by highlighting them all and right clicking.

#### 🍍 PowerToys run

The belle of the ball for me - basically Alfred or Spotlight for Mac, but for Windows.

#### 🥦 Shortcut guide

A full-screen pop-up that reminds you any time of the keyboard shortcuts you have set up.


## 🚢 Keep Calm and PowerToy On!

Of course, I'm not doing any of these items justice in a single sentence. You can learn a bit more about each one [here](https://github.com/microsoft/powertoys#current-powertoy-utilities), and follow the links in each section to go deeper.

Look out for my next post on using PowerToys to make your Mac keyboard work smoother Windows. See you soon!

👋