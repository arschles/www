---
layout: post
title: Going from Java to Scala
summary: Making the jump isn't as hard as you might think
---

Scala, in my opinion, is a wonderful language. I have refrained from saying that publicly until after the "honeymoon period" was over between me and it. Now that it is, I feel comfortable with that statement.

Perhaps the greatest feature of Scala is that it can work "seamlessly" with Java. But, that whole "seamlessly" assertion comes with plenty of caveats, because as some know, sometimes not even Java can work seamlessly with Java. But it is possible to ship production quality, stable code that is mixed Java & Scala. And although I think Scala is great, I won't spend time in this post defending that opinion, so if you disagree or otherwise don't want to use Scala stop reading now. Otherwise, read on!

# Why you'll need to tangle with Java

Java is everywhere. If you have never written a line of Java and you're staring fresh with Scala, welcome to [the JVM](http://en.wikipedia.org/wiki/Java_Virtual_Machine). Yes, it is an open virtual machine on which almost anything can run (yay Scala, Clojure, JRuby, Jython, etc...), but **Java is king on this platform**. Well, actually [Java is king everywhere right now](http://www.tiobe.com/index.php/content/paperinfo/tpci/index.html).

Also, as you go, don't forget that Java is still an excellent language regardless of how much Scala may outshine it in certain areas. It *might* be a better fit sometimes for what you're trying to do (I haven't found any places where it is yet, though).

# Lesson of the First: Java Conversions

[Java's Collections Framework](http://download.oracle.com/javase/6/docs/technotes/guides/collections/index.html) and [Scala's Collections Framework](http://www.scala-lang.org/api/current/index.html#scala.collection.package) are very different in practice. They don't even pretend to work together. Instead, Scala provides ```scala.collection.JavaConversions``` to convert between them.

JavaConversions is good at converting Scala mutable collections to their Java counterparts and back, and can do so pretty efficiently, but for immutable data structures, you're on your own. That's significant because immutable data structures give you saner and more thread safe code most of the time. My rule of thumb for deciding whether to use an immutable data structure: if you can't do it much better with mutable data structures, use immutable.

If you do, write a set of implicits that you keep around for when you need to use them for some Java API. Here's what I start with:

    package com.my.app.implicits.collection

    import java.util.{Map => JMap, HashMap => JHashMap, List => JList, Set => JSet}
    import scala.collection.JavaConversions._

    object JavaConversions {  
      implicit def JListToImmutable[T](l:JList[T]): List[T] = asScalaBuffer(l).toList
      implicit def JMapToImmutable[T, U](m:JMap[T, U]): Map[T, U] = mapAsScalaMap(m).toMap
      implicit def jSetToScalaImmutableSet[T](s:JSet[T]) = asScalaSet(s).toSet
    }

Pretty straightforward functions. They have 2 major features: (1) to convert Java collections into their Scala immutable counterparts, and (2) they take advantage of Scala-native JavaConversions functions to get decent performance.

So, in your code, you can ```import collection.JavaCoversions._``` and ```import com.my.app.implicits.collection.JavaConversions._``` and you'll have a much easier time going between Java and Scala immutable collections.

# Lesson of the Second: DSLs for a More Functional World

This is a portion of Java code from a real codebase, with names modified:

    //...
    AtomicInteger requestCounter = new AtomicInteger(0);
    //...
    Map<String, String> requestMap = new HashMap<String, String>();
    requestMap.put("objectID", "ff3300");
    requestMap.put("objectType", "red");
    String body = JSONLibrary.encode(requestMap);
    Request req = new Request(counter, HttpRequestMethod.POST, "http://some.domain.com/doMethod", body);
    Map<String, String> cookies = new HashMap<String, String>();
    cookies.put("SessionID", "sdfasdfaskgasgmnaasgsgsdfgsdf24q23krjqbwjeb12jb");
    req.setCookies(cookies);
    Map<String, String> headers = new HashMap<String, String>();
    headers.put("Content-Type", "application/json");
    req.setHeaders(headers);
    Response resp = executeRequest(req);
    //...


After a small addition of a Scala wrapper to build a DSL around the Request class, that code became the following. Again, names are modified. No Java was harmed in the making of this example:

    val req = request(HttpRequestMethod.POST) withURL("http://some.domain.com/doMethod") withHeader("Content-Type", "application/json") withJsonBody(Map("objectID" -> "ff3300", "objectType" -> "red"))
    val resp = executeRequest(req)

I don't need to say anything about these examples here. Scala & some basic functional techniques made this more readable, and way shorter. The way to do this is to implement a [DSL](http://en.wikipedia.org/wiki/Domain-specific_language) that wraps the Request object and transforms it from a procedural API into a functional one.

Once you understand Scala implicits, and how they enable the ["Pimp my library" pattern](http://www.artima.com/weblogs/viewpost.jsp?thread=179766), you're ready to shred all your Java container classes and the Builder classes that help wrap them. This is a common pattern that I fondly call the DSL pattern in Scala.

We'll call everything ```RequestDSL```, so we can ```import my.app.RequestDSL._``` in our code and start building ```Request```s. Here's what the code looks like:

    package my.app

    import java.util.concurrent.atomic.AtomicInteger
    import collection.JavaConversions._

    object RequestDSL {
      private val c = new AtomicInteger(0)

      def request(method:HttpRequestMethod, url:String, qString:Map[String, String], body:String): Request =
        new Request(c.getAndIncrement.toString,
            method,
            url,
            qString.toList.flatMap {tup:(String, String) => List[String](tup._1 + "=" + tup._2)}.mkString("&"),
            body)

      def request(method:HttpRequestMethod, url:String, qString:Map[String, String]): String => Request =
        Request(verb, url, qString, _:String)

      def request(method:HttpRequestMethod, url:String): (Map[String, String], String) => Request =
        Request(verb, url, _:Map[String, String], _:String)

      def request(method:HttpRequestMethod):(String, Map[String, String], String) => Request =
        Request(verb, _:String, _:Map[String, String], _:String)

      class F1Container(fn:String => Request) {
        def withBody(body:String): Request = fn(body)
        def withoutBody(): Request = withBody("")
      }

      class F2Container(fn:(Map[String, String], String) => Request) {
        def withQueryParams(params:Map[String, String]):(String => Request) = fn(params, _:String)
        def withoutQueryParams(): String => Request = withQueryParams(Map[String, String]())
      }

      class F3Container(fn:(String, Map[String, String], String) => Request) {
        def withURL(url:String): (Map[String, String], String) => Request = fn(url, _:Map[String, String], _:String)
      }

      implicit def f1ToContainer(f:String => Request) = new Function1Holder(f)
      implicit def f2ToContainer(f:(Map[String, String], String) => Request) = new Function2Holder(f)
      implicit def f3ToContainer(f:(String, Map[String, String], String) => Request) = new Function3Holder(f)
      implicit def RequestToWrapper(req:Request) = new RequestWrapper(req)

      class RequestWrapper(req:Request) {
        val oldHeaders:Map[String, String] = req.getHeaders

        def withHeader(key:String, value:String) = {
          req.setHeaders(oldHeaders + ((key, value)))
          req
        }

        def withHeaders(additionalHeaders:Map[String, String]) = {
          req.setHeaders(oldHeaders ++ additionalHeaders)
          req
        }

        def withCookie(key:String, value:String) = {
          req.setCookies(Map(key -> value))
          req
        }
      }
    }

Let me briefly describe what's going on here:

1. Each of the ```request(...)``` methods allows you to build up the request, starting with the ```HttpRequestMethod```, and going on from there.
2. The intermediate request methods are the ones that return a function that takes in some arguments and returns a Request. Each intermediate partially applies the final request builder in order to return a partially applied function, which is a really nice way in Scala to implement [curring](http://en.wikipedia.org/wiki/Currying).
3. We've added some pimps (container classes & implicit conversions to those containers) for the functions that the intermediate request methods return so that you don't have to do the conversions yourself. The containers are called ```F*Container``` and the implicit conversions are called ```f*ToContainer```.
4. We've added a pimp for Request so that you can chain methods together to operate on the Request. The container is called RequestWrapper and the implicit conversion is RequestToWrapper.

First, the bad news about the DSL code: if you haven't noticed, we've traded 11 total lines of Java for 64 total lines of Scala. That will be common with DSLs. You have to do a decent amount of work to wrap a procedural Java API, so it's not worth it unless you use that API a lot. Probably goes without saying.

But there's more good news than bad:

1. Scala implicits, partial function application, & immutable data structure initializers let us build ```Request```s in **one line** of code (or, if the line gets too long, 2 or 3 lines of code) like we saw above. Fuck yea
2. Although the Scala code is longer, it is extremely extensible (you can easily add complex operations on Request objects to your DSL) and does not use any "exotic" features or patterns of the Scala language. In fact, [pimp my library](http://docs.scala-lang.org/sips/pending/implicit-classes.html) may even become part of the core Scala language.
3. Code that uses this DSL reads closer to English, and therefore is more self-documenting.

# Continuing on

The community is always developing better & easier ways to achieve Scala - Java interoperability. As Scala gets more popular this area will be explored more, so someday soon, the code and approaches here will become outdated. I encourage you to start with these and add your own.