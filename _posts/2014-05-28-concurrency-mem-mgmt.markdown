---
layout: post
title: Concurrency is the New Memory Management
summary: A brief but related interlude to the ideal programming language series
---

I want to talk about concurrency in depth as it relates to the evolution of
programming languages.

# History

I tend to classify programming languages into generations. Like people, not
everyone fits cleanly into a generation or any other kind of category, but
we have general "buckets" that we can put people in.

Just like we talk about the baby boomers, generation Y, etc... I believe we can
talk about programming languages in terms of coarse grained groups.

## Close To the Metal

These languages have been around for as long as we've had compilers. These
languages abstract away almost none of the operating system. Engineers
who use those languages get great control over the behavior of their programs
at the cost of greater complexity of their programs. Here's a non-exhaustive
list of things you'll need to keep in your head if you write programs
"close to the metal":

- allocating and deallocating your memory
- managing concurrently executing threads of control
