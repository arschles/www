---
layout: post
title: Map (no reduce) with Ripeline
summary: Sometimes, map is all you need.
---

I went to MongoSF 2011 and attended a talk called "MongoDB's New Aggregation Features - A Sneak Peek" by Chris Westin. As of writing, I don't believe that slides are posted, but they will be soon.

In the talk, Chris said that they're working on pipeline-based aggregation for Mongo, as well as a rich set of pre-built operators that can be applied in any order in that pipeline. This is a distributed pipeline system. Step 1: define your data (a Mongo collection), step 2: write your data pipeline, step 3: profit.

In functional terms, this is the same as applying multiple mappers to a list (ruby's version: [http://www.ruby-doc.org/core/classes/Array.html#M000249](http://www.ruby-doc.org/core/classes/Array.html#M000249)), and streaming results from each mapper to the next. That model fits lots of problems too. Sometimes you just don't need to reduce.

Turns out someone wrote a [pipeline for appengine](http://code.google.com/p/appengine-pipeline/) which can be used to implement the streaming mappers system.

But I wanted to do it on totally open infrastructure so it can be used anywhere, so I started implementing the same idea and so far I have [ripeline (ruby + pipeline = ripeline)](http://github.com/arschles/ripeline). Ripeline uses redis for communication between the mappers, and ruby to write each mapper.

My ideal is to be able to write mappers in ruby once and then tell the system the order in which to run them, and it takes care of the rest: monitoring, communication, streaming, failure detection, etc...

On top of that, I'd love to build futures on top of the pipeline so that, for pipelines that will complete in a reasonable amount of time, you can start a pipeline and then hold a future for that pipeline's result.

Although MapReduce is a really rich & proven computation model, you sometimes don't need the reduce step, and so you can go to a simpler solution. I'd like Ripeline to be that.