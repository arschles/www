---
author: "Aaron Schlesinger"
date: 2020-06-16T11:53:00-07:00
title: 'How to Run Windows With a Mac Keyboard'
slug: "powertoys-mac"
draft: false

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

I wrote a bit about PowerToys [previously](/blog/2020-05-20-powertoys), but today we're gonna talk about my second-favorite feature: the keyboard manager.

>My first-favorite feature is Powertoys Run - the spotlight-like feature for Windows. I talked about this one last post.

# Make Windows Feel Like Mac 🍏

I [switched to Windows from Mac](/blog/how-to-wsl-2/) a few months back, and I'm used to just about everything except the keyboard. Muscle memory keeps kicking in and my thumbs reach for the `cmd` key instead of `ctrl`. There's no `cmd` on Windows - it just opens the start menu. I've opened the start menu a **lot**.
 
# The Keyboard Manager ⌨

The key feature (see what I did there?) of PowerToys that lets my keyboard be all apple-ey is the keyboard manager. It looks like this:

![powertoys keyboard manager](2020-06-16-powertoys-mac/powertoys-keyboard-manager.png)

There's a lot going on in that window. Let's break it down...

## Remap Keyboard 🗺

This one is simple in theory. Like the heading says, you get to map a single key. Take the first one in there - I mapped Page Up to the home key (that one wasn't Mac specific - I just kept accidentally hitting it 😉).

Here are the keys I remapped to Mac-ify:

- `Caps Lock` => `Win (Left)`
- `Win (Left)` => `Ctrl (Left)`

This is basically making the Windows key to the left of the spacebar into the Ctrl key, which does similar things as the Cmd key on a Mac. Since I was overriding the Windows key, I moved that up to Caps Lock, since I never use it anyway. You could also do this with the Windows key on the right, if you have one.

To do these mapping, you click "Remap a Key', then the Plus icon at the bottom. Click on each box and type the keys you want to map. The "before" key goes on the left, and the "after" key on the right.

>❗ Do the mapping from `Win` -> `Ctrl` _last_

## Remap Shortcuts 🍕

This one piggybacks on key remapping. As you know, Windows has lots of key combinations that "do stuff"™. For example, alt+tab cycles through Windows, ctrl+c copies text, ctrl+v pastes text, and so on.

Here are the shortcuts I set for mac-ifying:

- `Ctrl (Left) + Tab` => `Alt (Left) + Tab`
    - cmd+tab on Mac lets you cycle through Windows. This makes it feel the same. This still have a [bug](https://github.com/microsoft/PowerToys/issues/3331), but it's usable enough for now
- `Ctrl (Left) + Shift (Left) + [` => `Ctrl (Left) + Shift (Left) + Tab`
    - This lets me cycle through tabs quickly, just like on a Mac. This particular one cycles left.
- `Win (Right) + Shift (Right) + [` => `Ctrl (Left) + Shift + Tab`
    - Cycle to the left, from the right side of the keyboard
- `Ctrl (Left) + Shift (Left) + ]` => `Ctrl (Left) + Tab`
    - Cycle to the right, from the left side of the keyboard
- - `Win (Right) + Shift (Right) + ]` => `Ctrl (Left) + Shift + Tab`
    - Cycle to the right, from the right side of the keyboard

I hope those help! It's by far not _everything_ to make Windows look like a mac - some Windows things just don't apply to Mac and vice versa. But I think you'll find this a good setup.

Enjoy!