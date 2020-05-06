---
author: "Aaron Schlesinger"
date: 2020-04-27T16:42:35-07:00
title: 'So you want to build a computer'
slug: build-a-computer

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

I once built a cheap tower with clearance parts from Micro Center. See if you can guess how long ago it was from the specs:

- CPU: Intel Pentium 3
- RAM: 16GB
- Disk: 128mb

I haven't done much DIY electronic stuff since then except for some IoT tinkering and playing around with drones. So when I realized I need a [Windows machine for work](https://dev.to/arschles/how-to-wsl-516d), I decided I was gonna get back into the DIY game and build one from scratch.

Here's my story.

## Why build from scratch?

If computers are your profession, it's nice to choose all your tools so you get exactly what you need. Just like if you look in a carpenter's toolbox you'll see tools from tons of different brands. They chose exactly what they need to do their best work.

Also, you have to put time and work in up front but you come out with a completely custom computer for cheaper than what you can buy off the shelf. And you know all the ins and outs of the thing, so you can upgrade individual pieces in the future.


## How to started

Luckily I knew the basics of a computer, so I sketched out the pieces I need. If you're unfamiliar, here they are with crude definitions:

- Motherboard - holds everything together (CPU, memory, hard drive, fans, peripherals)
- CPU - **C**entral **P**rocessing **U**nit - does all the fancy things that runs your OS, compiles code, etc...
- Memory (AKA: RAM) - your processor and programs you build and/or run use this to store temporary data
- Hard Disk - where your files live, forever
- Case/Fans - The actual place you put all this stuff. Fans are mounted to the case (and the CPU) to cool all these electronics off while they're running
- Video card - you plug this into the motherboard and the other side has hookups for your monitor. Some Motherboards have this built in.
- Sound card - just like the video card but for sound. Just about all consumer-grade motherboards have this built in
- PSU - **P**ower **S**upply **U**nit. This is what you plug into the wall. The other side has wires with different connectors that attach to the video card, fans, motherboard, etc...

## How to choose parts

I started by thinking about the main reasons for this computer. I don't game and this was gonna be for work, so my needs were mostly compiling code and running lots of VMs. That means more CPU cores and memory and you can skimp on the video card (which can be a big cost). More CPU cores usually means a more expensive motherboard too, so I had to budget for that too.

My big purchase was an AMD Ryzen Threadripper CPU. I had been saving up for a while so I could get one. Most of the work was figuring out what kind of motherboard I would need and how much power I would need from the PSU.

### Resources

Once I had a more specific idea of what I needed, I went to [NewEgg](https://newegg.com) to look at general prices for things. They also have some nifty tools like the [power supply calculator](https://www.newegg.com/tools/power-supply-calculator/) that can help you decide what to buy.

I also went to [`/r/buildapc` on Reddit](https://buildapc.reddit.com) to see reviews. You can also ask questions there if you want.

Amazon had cheaper prices so I bought everything there. I would have gone to a physical store like to look at parts, but there wasn't a good option in my area.

>During the COVID pandemic, please research and buy everything online so you don't have to go out, even if you do have a parts store near you that's open

## Do the build!

I made too many mistakes to list here, so I'll share the advice I learned from it all:

- Read through the entire step-by-step process for the motherboard, CPU and fan before doing anything
- Don't force anything except for the memory. You'll probably need to apply pressure there until it clicks
- Do everything on a clean table if you can. Don't kill your back doing it on the floor like I did 🤣
- Attach the PSU before anything else
- ... then attach the CPU and fan before you screw in the motherboard
- ... then plug in all your components
- ... then put in the hard disk, video card, etc...
- When you're plugging things in, the wires will come from the PSU up to the motherboard and other components (often the hard disk and video card). Try to run  everything below the motherboard
    - Some cases come with zip-ties and loops so you can bunch wires together and attach them in place to the case. Use everything you can to organize the wires. You'll be happy you did later!

## Enjoy!

I got all my components in January and I've had the thing running since February.  As I said, I made a bunch of mistakes along the way. I still have things to improve (liquid-cooling and a bigger PSU for example!), but I'm actually _happy_ I do. This is both a tool and a hobby!

I'm really excited to get creative and keep my machine healthy and up to date as time passes.

If you do a build, I hope you get the same amount of joy out of it as I am 😊