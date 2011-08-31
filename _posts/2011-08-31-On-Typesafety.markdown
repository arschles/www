---
layout: post
title: Function Maps in Scala
summary: Naming groups of related functionality
---

I find myself using the Function Map pattern a lot in my Scala code. Generally, when I need to group a set of related functions together, I use one of these.
My favorite example is in actors. When passed a label that identifies an action, the actor looks up the action in the function map and then executes is. Here's the code:
    
    ...
    case class BaseMessage(s:String)
    case class MessageOne(s:String) extends BaseMessage(s)
    case class MessageTwo(s:String) extends BaseMessage(s)
    case class MessageThree(s:String) extends BaseMessage(s)
    
    val FunctionMap = Map(
        MessageOne -> ((a:String) => println(a)),
        MessageTwo -> ((a:String) => doSomethingWithString(a)),
        MessageThree -> ((a:String) => doSomethingElseWithString(a))
    )
    
    val MessageProcessor = actor {
        loop{
            react {
                case b:BaseMessage => {
                    MessageProcessor.get(e) match {
                        case Some(f:((a:String) => Unit)) => f(b.s)
                        case None => throw new Exception("unrecognized BaseMessage " + b)
                    }
                }
                case a => throw new Exception("unrecognized message " + a)
            }
        }
    }
    ...

If I'm trying to send a message to the MessageProcessor actor, I can look up what the actor does in the FunctionMap variable first, and then
decide if I want to call the actor.

FunctionMaps can be used in a bunch of other ways also. Try using one next time you want to group a set of related functions.