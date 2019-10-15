---
author: "Aaron Schlesinger"
date: 2019-03-11T14:06:10-07:00
title: 'Functional Programming in Go With dcode'

# For twitter cards, see https://github.com/mtn/cocoa-eh-hugo-theme/wiki/Twitter-cards
meta_img: "images/go-functional-gopher.png"
image: "images/go-functional-gopher.png"

# For hacker news and lobsters builtin links, see github.com/mtn/cocoa-eh-hugo-theme/wiki/Social-Links
# hacker_news_id = ""
# lobsters_id = ""
---

Back in 2017, I gave a [talk](https://www.youtube.com/watch?v=c8Fwb4KbVJM) at GopherCon about doing functional programming with Go. I went up on stage and talked about some FP basics like `Functor`s, higher-order functions and so on. I heard some feedback that doing FP in Go isn't practical until Go has generics (fret not, folks, [contracts](https://go.googlesource.com/proposal/+/master/design/go2draft-contracts.md) are on the table!).

It's true that contracts will be a huge boon for functional Go code, but there's lots of functional programming we can do in Go _without_ generics.

**You don't need generics to do FP in Go.**

# dcode

When I started writing that talk in 2017, I created the [github.com/go-functional](https://github.com/go-functional) GitHub organization to hold experimental projects for doing FP in Go.

Beyond the code I showed on the slides, which is in [core](https://github.com/go-functional/core), I've been prototyping some other functional libraries like [quicktest](https://github.com/go-functional/quicktest) and pure functional [SQL query builders](https://github.com/go-functional/query). The one area that I really latched onto was JSON, and that's how `dcode` was born.

## Inspiration

A long time ago, I built web services using Scala and [lift-json](https://github.com/lift/lift/tree/master/framework/lift-base/lift-json/). That library provides a pure functional tree representation for a JSON object. It's also relatively fast for deep traversals.

More recently, I found the [Elm language](https://elm-lang.org/), and particularly its [JSON parser](https://guide.elm-lang.org/effects/json.html). This parser interface is a concise DSL that lets you say you you want to traverse a JSON object. It's incredibly self-documenting, even when dealing with large JSON objects. In fact, the DSL tends to look similar to [JSONPath](http://jsonpath.com/).

The Elm API makes the lift-json tree representation accessible to anyone, without requiring them to write their own tree traversal code.

## Usage

You write one line of code to traverse any number of levels into a JSON object  as necessary.

For example, if you have this JSON (adapted from the JSONPath page):

```json
{
  "firstName": "John",
  "lastName" : "doe",
  "age"      : 26,
  "address"  : {
    "streetAddress": "naist street",
    "city"         : "Nara",
    "postalCode"   : "630-0192",
    "accessInstructions": {
        "gate": true,
        "message": "Please dial 0123 so I can buzz you into the gate"
    }
  }
}
```

And you want to get to `$.address.accessInstructions.message` (that is, the value of the `message` field inside the address access instructions object), you'd write this code to crate a decoder that, when it's called, will try to pull that value out of any JSON you give it.

```go
decoder := Field(
    "address",
    Field("accessInstructions"),
    Field("message", String()),
)
```

There's also an accumulator interface so you don't have to nest all those calls to `Field`:

```go
decoder := First("address").Then("accessInstructions").Then("message").Into(String())
```

# Why?

`dcode` obviously doesn't fit every use case. If you're expecting to decode a big JSON object with lots of keys, and the incoming JSON value is the same every time, you'll most likely be more successful writing a `struct` and using [`encoding/json`](https://godoc.org/encoding/json) to decode into it (although `dcode` does have functionality to decode into `struct`s).

This library does well, however, in a few cases:

- **You need to traverse deep JSON objects to fetch values**. This use case is common with programs that need to use a small part of a large JSON API
- **You don't know the "shape" of the object ahead of time**. Often, you'll want to try several different decoding strategies at a time. This use case is common with web applications that deal with legacy API clients

Since `dcode` is based on functional programming principles, it also comes with two helpful patterns:

#### Composable

Each `Decoder` that you construct (with `Field` or `First()...Into()`) is a stateless function that is approximately:

```go
func (d *Decoder) Decode(b []byte, i interface{}) error
```

>Note: I've written this function here as an example. It doesn't exist in the library, but you can use the [top-level `Decode` function](https://godoc.org/github.com/go-functional/dcode#Decode) instead to get the same result.

Of course, this function signature looks a lot like [`(encoding/json).Unmarshal`](https://godoc.org/encoding/json#Unmarshal), which is nice for familiarity.

More importantly, these `Decode` functions are built up by composing other `Decode` functions together. The `Field` or `First()...Into()` interfaces both enable that composition.

This API makes it obvious how to structure your JSON decoding logic, and trivial to refactor it. It also really encourages the same kind of composibility as `Decoder`s "leak" into other parts of your codebase

>Bonus: `dcode` itself takes heavy advantage of composition. Check out [`Map`](https://godoc.org/github.com/go-functional/dcode#Map) and [`OneOf`](https://godoc.org/github.com/go-functional/dcode#OneOf) for examples.

#### Loosely Coupled

`Decoder`s are instructions for decoding a JSON object, and they can exist alone, without the values that the JSON should be decoded into. This lets you pass `Decoder`s to any function so that it can decode any JSON into any value as it sees fit. In the future, we may consider adding functionality to serialize and deserialize `Decoder`s themselves, so you can pass them over the network.

To contrast, if you use `encoding/json` to decode JSON, you usually create a `struct` to define the "shape" of the JSON. This API requires that the values (struct fields) and the decoding instructions (the reflected field names, or struct tags) are tightly coupled.

# Get Started!

Go check out the [README](https://github.com/go-functional/dcode/blob/master/README.md) for more on how and why to use `dcode`. Of course this is a work in progress, so if you find something missing, wrong, or just plain broken, please don't be shy and [file an issue](https://github.com/go-functional/dcode/issues/new)!

Keep on rockin', Gophers
