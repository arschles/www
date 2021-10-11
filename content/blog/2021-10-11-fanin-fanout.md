---
author: "Aaron Schlesinger"
date: 2021-10-11T19:46:49Z
title: 'Fan-in and fan-out with Go'
slug: "fanin-fanout"
---

Hacking on the [KEDA HTTP Addon](https://github.com/kedacore/http-add-on), I found myself having to do something familiar:

>Split some work into N pieces, do them all concurrently, wait for them all to be done, and then merge all the results together.

I've done this a bunch of times before, but this time I forgot how to do it. I took a few minutes away from the computer to gather my thoughts and came back to it. So I don't have to forget how to do it again, I want to write the algorithm down here!

## What are we doing here?

First thing's first - we need a problem we can break down into a bunch of pieces. Sometimes it's called an "embarassingly parallel" problem. 

>Note that concurrency and parallelism aren't equivalent, but I'm going to be using the word "parallel" hereafter because I'm hoping the machine you run this algorithm on will be able to execute the work on different cores simultaneously.

The primary goal here is to run each different piece of the work in a goroutine. That's pretty easy in Go - just put `go` before the function call that does the work. The tougher part is to get the results of the work, check for errors, and wait for them all to be done -- not necessarrily in that order 🤣.

Even though the concept is simple, there is a big-ish gotcha when the rubber hits the road. Below is some code that does "fake" work, annotated with comments to explain it:

```go
workToDo := []int{"do", "some", "work"}
for idx, work := range workToDo {
    // make sure you pass the index and work into the 
    // function that runs in the goroutine.
    // this mechanism makes sure that the goroutine
    // gets a (stack-allocated) _copy_ of the data.
    // if you don't do this, idx and work will change out
    // from under it as the loop progresses.
    go func(idx int, work string) {
        fmt.Println(idx, work)
    }(idx, work)
}
```

The biggest gotcha is in that comment inside the `for` loop.


## Waiting for the work to be done

Now that we've got goroutines running with the right parameters, let's add a [`sync.WaitGroup`](https://pkg.go.dev/sync?utm_source=godoc#WaitGroup) to the mix. This mechanism will let us wait for all these goroutines to finish.

```go
var wg sync.WaitGroup
workToDo := []int{"do", "some", "work"}
for idx, work := range workToDo {
    // add 1 to the waitgroup _before_ you start the goroutine.
    // you want to do this in the same goroutine as where
    // you call wg.Wait() so that you're sure that, even if
    // none of the goroutines started yet, you have the
    // right number of pending work.
    wg.Add(1)
    // make sure you pass the index and work into the 
    // function that runs in the goroutine.
    // this mechanism makes sure that the goroutine
    // gets a (stack-allocated) _copy_ of the data.
    // if you don't do this, idx and work will change out
    // from under it as the loop progresses.
    go func(idx int, work string) {
        // wg.Done() tells the WaitGroup that we're done in
        // this goroutine. In other words, it decrements
        // the internal WaitGroup counter, whereas wg.Add(1)
        // above increments it.
        // Most commonly, we just do this in a defer statement.
        defer wg.Done()
        // this is the "work". in the next section, we'll be
        // changing this to return a value, because we'll
        // need to send that value somewhere
        fmt.Println(idx, work)
    }(idx, work)
}
// wait for all the goroutines to finish. this call
// blocks until the WaitGroup's internal count goes 
// to zero
wg.Wait()
```

## Getting the results

So, now we know when all the work is done, we need to get the results. There are two kinds of results we need to get - the actual values of the work we're doing -- we'll call this the "success value" -- and the errors that it might have returned.

Let's focus on the success values first. We're going to use one group of channels, one "final" channel, and a clever way of shuttling data between the former and the latter:

_Read this code on the [Go Playground](https://play.golang.org/p/CM34_zkrmrg), if you prefer_

```go
// this is the channel that will hold the results of the work
resultCh := make(chan string)
var wg sync.WaitGroup
workToDo := []string{"do", "some", "work"}
for idx, work := range workToDo {
    // add 1 to the waitgroup _before_ you start the goroutine.
    // you want to do this in the same goroutine as where
    // you call wg.Wait() so that you're sure that, even if
    // none of the goroutines started yet, you have the
    // right number of pending work.
    wg.Add(1)
    // this is the loop-local channel that our first goroutine
    // will send its results to. we'll start up a second
    // goroutine to forward its results to the final channel.
    ch := make(chan string)
    // make sure you pass the index and work into the
    // function that runs in the goroutine.
    // this mechanism makes sure that the goroutine
    // gets a (stack-allocated) _copy_ of the data.
    // if you don't do this, idx and work will change out
    // from under it as the loop progresses.
    go func(idx int, work string) {
        // this is the "work". right now, it just returns an
        // int. in the next section, it will return both an int
        // and an error
        res := doSomeWork(idx, work)
        ch <- res
    }(idx, work)
    // start up another goroutine to forward the results from
    // ch to resultCh
    go func() {
        // we want to indicate that we're done after we forward
        // the result to the final channel, _not_ just when we're
        // done with the actual computation. this arrangement
        // will be useful below, in our final goroutine that
        // runs after the for loop is done
        defer wg.Done()
        res := <-ch
        resultCh <- res
    }()
}
// start up a final goroutine that is going to watch for
// the moment when all of the loop goroutines are both
//
// 1. done with their work
// 2. done sending their results to the final channel
//
// after that, we can close the resultCh. this closure is
// important for the following for loop, since ranging over
// a channel will only stop after that channel is closed
go func() {
    wg.Wait()
    close(resultCh)
}()

// now that we have that final goroutine running, we can
// be sure that this for loop will end after:
//
// 1. all goroutines are done with their work
// 2. all goroutines are done sending their work to resultCh
// 3. we have processed each result
//  (in this case, we just print it out)
for result := range resultCh {
    fmt.Println("result:", result)
}
```

The code is extensively commented, but notice a few more "meta" things about it:

- We're enlisting _another_ goroutine for each loop iteration, so now we're using `2N` goroutines rather than `N` (where `N` is the number of work items to do).
    - If you're worried about the extra goroutines, remember that a Go program can run hundreds of thousands of them comfortably on a relatively modern laptop. Plan accordingly with that in mind.
- We use one extra final goroutine to determine when the final goroutine should be closed
- We no longer use `wg.Wait()` in the main goroutine. Instead, we range over `resultCh` to both get the results _and_ determine when all the work items are done.

## A final wrinkle: error handling

Now that you (hopefully) have a decent grasp over the code in the previous section, consider that, for most workloads, you'll also have to deal with error handling. It doesn't take a lot of _additional_ code to do it, but it does add a bit more complexity. Let's see how it works:

_Read this code on the [Go Playground](https://play.golang.org/p/Bcv_XQwoAi6), if you prefer_

>Note that you can reduce complexity by using the [`errgroup`](https://pkg.go.dev/golang.org/x/sync/errgroup) package. The code herein implements functionality similar to that of `errgroup`.

```go
// this is the channel that will hold the results of the work
resultCh := make(chan string)
// this channel receives all the errors that occur.
// for each work item, either resultCh or errCh will receive
// precisely once. both channels will be closed immediately
// after all receives happen
errCh := make(chan error)
var wg sync.WaitGroup
workToDo := []string{"do", "some", "work"}
for idx, work := range workToDo {
    // add 1 to the waitgroup _before_ you start the goroutine.
    // you want to do this in the same goroutine as where
    // you call wg.Wait() so that you're sure that, even if
    // none of the goroutines started yet, you have the
    // right number of pending work.
    wg.Add(1)
    // this is the loop-local channel that our first goroutine
    // will send its results to. we'll start up a second
    // goroutine to forward its results to the final channel.
    ch := make(chan string)
    // this is the loop-local channel that our first goroutine
    // will send errors on. for each loop iteration, exactly
    // one of ch or errCh will receive
    eCh := make(chan error)
    // make sure you pass the index and work into the
    // function that runs in the goroutine.
    // this mechanism makes sure that the goroutine
    // gets a (stack-allocated) _copy_ of the data.
    // if you don't do this, idx and work will change out
    // from under it as the loop progresses.
    go func(idx int, work string) {
        // this is the "work". right now, it just returns an
        // int. in the next section, it will return both an int
        // and an error
        res, err := doSomeWork(idx, work)
        if err != nil {
            eCh <- err
        } else {
            ch <- res
        }
    }(idx, work)
    // start up another goroutine to forward the results from:
    //
    // - ch to resultCh
    // - eCh to errCh
    go func() {
        // we want to indicate that we're done after we do the
        // forward operation, similar to the code in the
        // previous section
        defer wg.Done()
        // only one forward operation will happen per loop
        // iteration, so we use a select to choose exactly
        // one of the channels - either the success or error
        // one.
        select {
            case res := <-ch:
                resultCh <- res
            case err := <-eCh:
                errCh <- err
        }
    }()
}
// start up a final goroutine that is going to watch for
// the moment when all of the loop goroutines are both
//
// 1. done with their work
// 2. done sending their results to the appropriate channel
//
// after that, we can close both resultCh and errCh.
go func() {
    wg.Wait()
    close(resultCh)
    close(errCh)
}()

// we're now at a point where we have two "final" channels:
//
// - one for the successful results
// - one for the errors
//
// we have a few choices on how to handle them, and it's
// largely up to your use case how you handle errors or success
// results. In our case, we'll loop through both channels,
// print out the result either way, and then exit when all
// receives happen.

// these two booleans are going to keep track of when 
// each channel is closed and done receiving
resultsDone := false
errsDone := false
// we're going to use an infinite loop and break out of it
// when both channels are done receiving
for {
    if resultsDone && errsDone {
        break
    }
    select {
    case res, valid := <-resultCh:
        if !valid {
            resultsDone = true
        } else {
            fmt.Println("result:", res)
        }
    case err, valid := <-errCh:
        if !valid {
            errsDone = true
        } else {
            fmt.Println("error:", err)
        }
    }
}
```

A few more things to note here:

- We've handled errors and success values with equal importance. In many cases, you might want to immediately exit if you see an error. In that case, make sure that you find a way to receive the rest of the errors and success values on `errCh` and `resultCh` (respectively), or tell the remaining goroutines to exit.
    - If you intend to do the latter, I highly recommend using [`context`](https://pkg.go.dev/context)
- There is a lot of code here! And for that reason, it's worth repeating that you can reduce complexity by using the [`errgroup`](https://pkg.go.dev/golang.org/x/sync/errgroup) package.
- The primary source of complexity is the parallelism (that's the reason this blog post exists!) If you're thinking of using this pattern, I encourage you to measure the serial (non-parallel) version of the algorithm first to determine whether you really need to take on this complexity
