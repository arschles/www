+++
title = "Go Modules in 5 Minutes"
description = "A no-nonsense rundown on Go modules, in 5 minutes"
meta_img = "/images/modules_5_minutes.png"
type = "modules5"
twitter_card_type = "summary_large_image"
+++

Welcome, Gopher! You might be here because you have questions about [Go modules](https://github.com/golang/go/wiki/Modules), or maybe you're just looking to find out more.

Either way, welcome! I hope that this page helps you learn about the standard dependency manager for Go.

>I'm Aaron, by the way. I won't introduce myself here to keep this short! You can read more about me on my [about page](/about) if you'd like.

## ELI5: What Are Modules? 🤨

**Modules are the dependency management system for Go apps.**

That's pretty much it! You might have heard about prior dependency managers out there, but we're gonna focus just on modules here.

>If you're unfamiliar, `ELI5` stands for "explain like I'm 5"

## How Do I Get Started? 🚀

You most likely need to go into the root of your project and type the below command, substituting `YourModuleName` with ... your module's name! (use a VCS name like `github.com/arschles/assert`)

```console
$ go mod init YourModuleName
```

That should create a new `go.mod` file, which is where you'll keep track of all the modules that you `import` in your app.

You can also run this command inside an existing project to convert from older dependency management systems

>If that `go mod init` command doesn't work and you're doing it with an existing project, you might have to change a few things first. Your best option for now is to go ask in the `#modules` channel of the [Gophers Slack group](https://invite.slack.golangbridge.org/).

## How Do I Add A New Module? 🥳

You can use `go get`! Here's how to add a [popular testing package](https://github.com/stretchr/testify), at version `v1.5.1`:

```console
$ go get github.com/stretchr/testify@v1.5.1
```

>`go get` has been around forever, but now it supports versions and it knows how to update your dependency tracking files (see below)

## Ok, What About Deleting? 🧛‍♀️

_You don't have to explicitly delete a module from your project because modules aren't stored in your repository_

Instead of deleting, you run a `go get` with `@none` at the end, instead of the version number that we saw above:

```console
$ go get github.com/stretchr/testify@none
```

⚠ That `go get` command will remove the stretchr module from your project, and all of the modules in your project that depend on it!

>🦾Pro tip! Make sure you have a clean working directory before you remove a module. That way, if you don't like the post-removal world, you can always revert back to the way it was 🚢✌

## What's Up With These New Files? 🗃

Good eye! You caught the two new modules-specific files, `go.mod` and `go.sum`.

The `go.mod` has:

- Your app name (called `module` in the file)
- The version of Go you're using
- The list of modules that you import (the Go tool might put other `// indirect` ones in there too)

The `go.sum` has:

- A list of all the modules your app uses, including the [transitive](https://en.wikipedia.org/wiki/Transitive_dependency) ones (AKA: your dependency's dependency, their dependency, etc...)
- Every module's checksum

## Beyond 5 Minutes 🚀

Ok, so you have a lay of the land. You've probably got a feel for how things are going. Here are some other tips and tricks...

### The Global Cache 💵

Some programming languages store all your dependencies locally so you had to manually delete them when you’re done with them. Not the case with modern Go!

Go stores all the module code in a read-only central directory on your disk, so one version of a module isn't tied to just your project. If you have lots and lots of projects on your machine, that cache might get big. Delete it with this 🔥:

```console
$ go clean --modcache
```

>⚠ If you do this, you'll have to re-download all of your app's modules next time you build it

### Tidying Up 🧹

What if you forget to run that `go get ...@none` command from above in the "Deleting" section? You'll end up with modules in your `go.mod`/`go.sum` files that your code doesn't need.

Go has your back on that. You can always run this:

```console
$ go mod tidy
```

...to make sure that your mod files are in sync with the `import`s in your code.

>🦾Pro tip! Just as in the deleting section, make sure you have a clean working directory before you tidy up. That way, if you don't like the post-tidy world, you can always revert back to the way it was 🚢✌


### Seeing Your Dependencies 👀

All this talk about transitive dependencies, amirite??? You have two commands to help figure out why you're seeing modules in your `go.mod`/`go.sum` (most of the time, you'll have a question about why something is in the `go.sum`)

```console
$ go mod graph
```

⬆ shows you a list of modules that you can reconstruct into your dependency graph. The list looks like this:

```console
YourModule module1
YourModule module2
YourModule module3
YourModule module4
YourModule module5
YourModule module6
```

It's a little more complicated than that, but you get the idea.

>🦾 Pro tip! Each row is an [edge](https://en.wikipedia.org/wiki/Graph_(abstract_data_type)) (arrow) on the dependency graph

```
$ go mod why
```

This is kind of like the opposite of `go mod graph`. Graph shows you _everything_, but this one shows you why a module (that you give it) is in your app. 


## Beyond 5 Minutes: Where to Read More

I hope this document is enough to get you started and keep you going until you hit something really gnarly.

If/when you get to that point, check out the [wiki on modules](https://github.com/golang/go/wiki/Modules) to dive in to details. Things change from time to time with the underlying technology, and this wiki will be kept up to date as time goes on.
