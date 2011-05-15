---
layout: post
title: Type Safety for Python Functions
summary: Add type safety to your Python functions
---

Have you ever written something like this Python code?
    
    {%highlight python%}
    def myFunction(a, b):
      if type(a) != str and type(b) != int:
        raise "a must be a string or b must be an int"
      #do some more stuff
    {%endhighlight%}
      
I find myself writing that all the time. I want the type flexibility of Python a lot of the time, but I also want to enforce types when I need to. Writing that if statement over and over at the top of all of your functions is not very DRY, and it's error prone and annoying. I solved that problem with a decorator.

Function Decorators FTW
-----------------------

Python decorators are functions that you tell Python to call & pass another function into it. Python will replace the passed function with whatever your decorator function returns. Here's a concrete example:

    {%highlight python%}
    def log(func):
      def inner(*argv):
        print "calling %s:"%(func.__name__)
        func(*argv)
        print "done calling %s"%(func.__name__)
      return inner
  
    @log
    def myFunction():
      print "this is my function"
    {%endhighlight%}

Python calls the log function only once before I call myFunction for the first time. When it calls log, it passes it myFunction. When log returns, Python replaces myFunction with its return value, so that whenever I write code to call myFunction, I'm actually calling whatever log returned.

Every decorator has to return an object that's callable (ie: a function or an object with the __call__ method implemented) in order for this whole thing to work. Here, I've got log returning a function that takes in any number of arguments, prints some stuff, calls the original function with the arguments that were passed to it, and then prints some more stuff. Simple but powerful. You can now hijack any function without the caller knowing.

Checking Types
--------------

The decorator for checking types is a bit more complex, but follows the same rules as the simpler one above. Here it is:
    
{%highlight python%}
    def typesrequired(*types):
      def outer(func):
        def inner(*args):
          if len(args) != len(types):
            raise Exception("function %s must be called with %i arguments"%(func.__name__, len(types)))
      
          i = 0
          while i < len(types):
            if type(args[i]) != types[i]:
              raise Exception("argument %i must be of type %s"%(i, types[i].__name__))
            i += 1
          func(*args)
        return inner
      return outer

    @typesrequired(str, int, int)
    def printCoords(name, x, y):
      print "%s: (%i, %i)"%(name, x, y)
    
{%endhighlight%}
This decorator is more code because it takes in the expected types (in an argument list), and then must return another function that itself takes in printCoords. That second function has to return the function that will replace printCoords. The type checking happens inside of the innermost function (creatively called inner).

Using this decorator introduces a performance hit because it has to check the types of each argument every time printCoords is called (such is life with an interpreted, dynamically typed language), so if you don't like that, turn this decorator on in debugging environments only. In all other environments, it can just pass directly through and return the original function.

More Decorators
---------------

Decorators can do tons of stuff. They effectively provide a means to filter any function once before it's called the first time. You can use decorators to implement some pretty wide reaching changes with very little code and no refactoring. I'll cover some more uses in my next post.

For now, check out the Python Decorator Library (which includes a more complete implementation of this type checking decorator): http://wiki.python.org/moin/PythonDecoratorLibrary
