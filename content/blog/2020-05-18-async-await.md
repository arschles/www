---
author: "Aaron Schlesinger"
date: 2020-05-18T15:56:33-07:00
title: 'async/await, in 4 Minutes'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
# meta_img = "/images/image.jpg"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Heard of the `async`/`await` pattern? Javascript, C# and Rust (now! 🥳) have it. I'm probably missing some languages too. More are probably gonna adopt it because it solves a very specific and common pain point.

Let's look at why... using examples ... in Javascript.

>I'm a Javascript dabbler at best. I can write code that works, but it's never pretty! I like to think that I do know a lot about concurrency patterns, though, so don't go too hard on my JS code below 😅

# What It's About 🤔

`async`/`await` is all about writing concurrent code easily. More importantly, writing the code so it's easy to **read**.

## Solving the Problem Three Ways 🕒

We're gonna build up from callbacks to Promises (reminder: Futures in other languages, because naming) to `async`/`await`.

### Callbacks 😭

You've probably heard about callbacks in Javascript. If you haven't, they're a programming pattern that lets you schedule work to be done in the future, after something else finishes. Callbacks are also the foundation of what we're talking about here.

>The core problem we're solving in this entire article is how to _automatically_ run code after some other work is done

The syntax of callbacks is basically passing a function into a function:

```javascript
function doStuff(callback) {
    // do something
    // now it's done, call the callback
    callback(someStuff)
}

doStuff(function(result) {
    // when doStuff is done doing its thing, it'll pass its result
    // to this function.
    //
    // we don't know when that'll be, just that this function will run.
    //
    // That means that the rest of our ENTIRE PROGRAM needs to go in here
    // (most of the time)
    // 
    // Barf, amirite?
    console.log("done with doStuff");
});

// Wait, though... if you put something here ... it'll run right
// away. It won't wait for doStuff to finish. Confusing much?
```

... but yea, it's confusing and [exhausting to write](http://callbackhell.com/) and read 😞.

### Promises 🙌

I'll see your callbacks and raise you a `Promise`! No really, Promises are dressed up callbacks that make things easier to deal with. But you still pass functions to functions and it's still harder than it has to be.

```javascript
function returnAPromiseYall() {
    // do some stuff!
    return somePromise;
}

// let's call it and get our promise
let myProm = returnAPromiseYall();

// now we have to do some stuff after the promise is ready
myProm.then(function(result) {
    // the result is the variable in the promise that we're waiting for,
    // just like in callback world
    return anotherPromise;
}).then(function(newResult) {
    // We can chain these "then" calls together to build a pipeline of 
    // code. So it's a little easier to read, but still. 
    // Passing functions to functions and remembering to write your code inside
    // these "then" calls is sorta tiring
    doMoreStuff(newResult);
});
```

We got a few small wins:

- No more crazy nested callbacks
- This `then` function kind of implies a _pipeline_ of code. That can be easier to read and understand

But we still have a few sticky problems:

- You have to remember to put the rest of your program into a `then`
- You're still passing functions to functions. It still gets tiring to write and needlessly tiring to read

### async/await 🥇

Alrighty, we're here folks! The promised land 🎉🥳🍤. We can get rid of passing functions to functions, `then`, and all that forgetting to put the rest of your program into the `then`.

All with this 🔥 pattern. Check it:

```javascript
function doStuff() {
    // just like the last two examples, return a promise
    return myPromise;
}

// now, behold! we can call it with await
let theResult = await doStuff();

// IN A WORLD, WHERE THERE ARE NO PROMISES ...
// ONLY GUARANTEES
//
// In other words, the value is ready right here!
console.log("the result is ready!")
```

So, now we can read the code from top to bottom! All thanks to the `await` keyword. There's magic going on under the hood. The magic differs from language to language (in JS, it's `Promises` most of the time), but the results are the same:

- Programmers can read/write code from top to bottom, the way we're used to
- No passing functions into functions means less `})` syntax to ~~forget~~ write
- The `await` keyword is a clear symbol that `doStuff` returns a promise, which might mean it's doing something "expensive" like doing a REST API call or something

#### What about the `async` keyword???

In many languages including JS, you also have to mark a function `async` if it uses `await` inside of it. There are language-specific reasons to do that, but here are some that you should care about:

- To tell the caller that there's an `await` inside of it (which implicitly means there are `Promise`s going on behind the scenes in JS) and the caller should probably do an `await` on the result
- To tell the runtime (or compiler in other languages) to do its magic behind the scenes to "make it work"™

## 🏁

And that's it. I left a lot of implementation details out, but it's really important to remember that this pattern exists more for human reasons rather than technical. You can do everything with promises that you can do with `async`/`await`, but in most cases, `async`/`await` is going to make your life easier. Take it from me, someone who's broken many a build and written many a bug because he forgot how promises/futures work.

Enjoy your day!

-- Aaron
