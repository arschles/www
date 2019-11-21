---
author: "Aaron Schlesinger"
date: 2019-11-21T13:19:52-08:00
title: 'Building Full Stack Applications with Go, Vecty & WebAssembly'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
meta_img: "/images/vecty.png"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---


# Building full stack web apps with Go, Vecty and WebAssembly

Welcome folks! Today, we're going to talk about writing full stack web applications in a brand new way.

>I wrote this post a while ago and I'm reposting it on my blog now. It originally appeared [here](https://blog.logrocket.com/building-full-stack-web-apps-with-go-vecty-and-webassembly/)

Many of us have heard of -- and maybe written -- full stack web applications. We do them in a variety of different ways, but the common denominator is usually [Javascript](https://developer.mozilla.org/en-US/docs/Web/javascript) & [Node.js](https://nodejs.org/en/).

Today, we're going to break with that tradition and write a complete web application - front end and backend - without writing a line of Javascript. We’ll be comparing the developer experience to JavaScript along the way, but we’re going to be only writing [Go](https://golang.org) for this entire project.

We'll learn how to build a single page link shortener application with just Go, and we'll end up with working code that shows it in action.

# Prerequisites

Today, we're going to be focusing on Go so make sure you've [installed](https://golang.org/dl/) the tooling on your machine. I'm going to assume you have basic knowledge of Go, so check out the free [Tour of Go](https://tour.golang.org) to brush up (or learn!) if you need to.

All the shell commands that I'm going to be showing work on a Mac, but should also work on most Linux systems (including [WSL](https://docs.microsoft.com/en-us/windows/wsl/about)!).

Finally, make sure to clone the [repository](https://github.com/arschles/vectyshortener) with the code for this article.

And then you're good to go, so let's get started!

# Getting started 

First, we're going to get the application running locally.

Coming from [Webpack](https://webpack.js.org/) and surrounding technologies -- which you'd use to build a web app with Javascript -- building and running this application is embarrassingly easy. There's a front-end and a backend part (more on that below), and you compile both of them with the `go` tool, which requires no configuration.

First, run the backend server:

```console
$ go run .
```

Next, build the frontend in a new terminal window:

```console
$ cd frontend
$ GOOS=js GOARCH=wasm go build -o ../public/frontend.wasm
```

Finally, go to https://localhost:8081 in your browser to see the app in action.

# How this all works

Like most web apps, our link shortener has a front end and backend piece. In our app, the backend is just a static server written in Go. All of the magic is in the front-end directory, so let's start there!

If you're familiar with [React](https://reactjs.org) or the [DOM](https://www.w3.org/TR/WD-DOM/introduction.html), you'll recognize lots of the concepts we'll cover. If not, this stuff should come naturally as we go through it.

We're using a new Go framework called [Vecty](https://github.com/gopherjs/vecty) to organize our application. Vecty makes you break down your app into components and arrange them into a tree. The whole scheme is really similar to HTML and the DOM or React.

Here’s what our app’s high level components would look like if they were HTML:

- A `h2` for the title of the page
- A `form` to enter the link to shorten
- A `div` to hold the shortened link
    - This value is dynamically updated as the user types the link into the above
- An `a` to save the short link

Vecty components are so similar to React that they look like the Go equivalent of [JSX](https://reactjs.org/docs/introducing-jsx.html), except that they have more parentheses =P.

Let’s zoom in on one and see how it works. Here's the code for the `form` component:

```go
elem.Form(
    elem.Input(vecty.Markup(
        event.Input(func(e *vecty.Event) {
            short := uuid.NewV4().String()[0:5]
            h.shortened = short
            vecty.Rerender(h)
        }),
    )),
)
```

First, `elem.Form` and `elem.Input` on lines 1 and 2 are for the `<form>` and `<input>` tags, respectively. Those are both Go functions that take one or more arguments. Each argument is something that goes between the opening and closing HTML tags. For example, the stuff we pass to `elem.Form` goes in between `<form>` and `</form>`. This is what the above Go code would look like in HTML:


```html
<form>
    <input></input>
</form>
```

Pretty simple, right?

The last piece of code we didn't look at is that `event.Input` function. This is an event handler just like in HTML/JavaScript. This function takes in *another* function, which in this case is roughly an `[onchange](https://www.w3schools.com/jsref/event_onchange.asp)` handler. Just like you'd expect, that `*vecty.Event` argument the handler takes in is roughly the same as the JavaScript event.

The logic to actually shorten the link is all inside this handler, and it's fairly simple. Here is that code commented thoroughly, to explain what's going on:

```go
// First, make a new UUID and take the first 5 characters of it.
// This will be our new shortcode
short := uuid.NewV4().String()[0:5]
// Next, write the shortcode to a variable. This variable is shared
// with the <div>, so when we re-render this component, the <div> will
// get updated
h.shortened = short
// Finally, re-render the component so that the <div> gets the new shortcode.
// Unlike React, there's no automatic diff functionality. We tell Vecty
// explicitly which components to re-render.
vecty.Rerender(h)
```

## You Get Web Assembly for Free 

Vecty can scale to big applications because of this component structure, and we can scale our app as big as we want by adding more components as needed. For example, we can add a component *above* our current top level to dynamically route to different sub-components based on the URL. This would be similar to [react-router](https://github.com/ReactTraining/react-router) for the React community, for example.

## WASM == HTML?

Nope! WASM is a full departure from the DOM and everything HTML.

I compared all the components in the last section to HTML tags, but they aren't! That's where the big difference between Vecty / WASM and React comes in. We're compiling our Go code *straight* to WASM, which represents these components differently from the DOM.

## WASM == Javascript?

The same thing goes for Javascript! That Go code we saw in the event handler doesn't ever turn into Javascript, and the browser doesn't run the Javascript interpreter to run it. Instead, the Go compiler translates it into a [WebAssembly program](https://en.wikipedia.org/wiki/WebAssembly#Wasm_program), and the browser runs that _directly_.

You get some performance wins this way because your stuff doesn't need to go through the Javascript runtime. On the other hand, Go-based web assembly binaries can get pretty big (they're smaller with other langauges though), so there's a tradeoff there.

## Conclusion 

At the end of the day, you get some big benefits from using Go and Vecty to build apps:


1. You get to think in terms of components and nesting, just like with React and the DOM
2. You can write as much dynamic logic as you want, right next to your components, all in pure Go
3. You can share code between the server and client, similar to writing a React client and a Node.js server
4. You get to take advantage of WASM
    1. Or you can compile your Vecty code to HTML too if you want! That's a whole other article ;)

