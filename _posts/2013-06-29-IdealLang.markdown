---
layout: post
title: My Ideal Programming Language, Part 1
summary: The most important things I'd most like to see in a programming language
---

Among the many things Bjarne Stroustrup has said that someone has published on the internet, my favorite is this: "There are only two kinds of languages: the ones people complain about and the ones nobody uses."

As of this writing, I've been writing code as my full time job, as well as for some open source projects, for a few days over 5 years. Before that, I majored in computer science, where I wrote lots of code, and I worked for the [UM Computer Aided Engineering Network](http://www.engin.umich.edu/caen/), where I wrote lots of code. Before college, I wrote code for fun.

Most of the work I've done in my adult life has been writing code, so I've worked with a wide variety of programming languages. I'm writing here what my ideal programming language would look like. I'm writing this post as a time capsule. In 5 more years, I want to compare my then-ideal language with my now-ideal language. I think the difference will effectively show me what I've learned and how my thinking has changed.

I am neither a language designer nor expert. I have never written a compiler or designed a programming language more complicated than a stripped down version of MIPS assembly. This design is based entirely on judgements made from my experience in the world.

# Language Shapes Your View on the World

There's plenty of research discussing the theories that spoken languages shape cultures and thinking. I believe the same is true for programming languages. The effect of a language on thinking and culture is especially apparent in these areas:

* packaging code
* running code
* testing and debugging code
* type safety
* polymorhpism
* missing values (see [the billion dollar mistake](http://qconlondon.com/london-2009/presentation/Null+References:+The+Billion+Dollar+Mistake))
* error handling and recovery
* memory management
* concurrent programming
* distributed programming
* I/O
* native extension

Every existing language has a library or built in feature to address each of the areas in that list, so my ideal will too.

To start, Ideal is a compiled, strongly typed language. The Ideal compiler builds your code into a native executable, which can be immediately run on the system for which it was compiled. There is no VM.

# Packaging code

The Ideal compiler recognizes a group of code files in a directory as a module which has the same name as the directory. Any code file not in the same module can use the `import` function to pull in module top-level names to the namespace of the file.

When you compile your code, you can compile it as a library or a a runnable.

## Runnables

If you compile your code as a runnable, you build an executable that starts by running the `main` function. If there's no main function, compilation fails.

## Libraries

If you compile your code as a library, the compiler builds up a package that looks like a module to other programs. The module's name is specified in a special `library` file that must be in the top level directory. If the file isn't there, the compilation will fail.

If your code (library or runnable) wants to use other libraries, you put them into a `dependencies` file. On each line, you specify the name of the module, where to get it from, and the version. Ideal has a central module repository, and others can exist. The module repository looks like a simple REST API, with no authentication. For example, to get the "aaron" module version 2, from the (fictional) ideal repository, the format is GET http://repo.ideal-lang.org/aaron/2

In the `dependencies` file, you can rename incoming modules to avoid name conflicts, and Ideal isolates all transitive dependencies, so if module A depends on modules B and C, and both of those depend on different versions of module D, no error occurs. Internally, the packager/linker/whatever stores modules by their name and version, so the right code is used in the right place. Circular dependencies are not allowed and error at compile time.

There's no enforced versioning scheme. The packager caches dependencies locally, and uses [ETags](http://en.wikipedia.org/wiki/HTTP_ETag) for local cache invalidation. A dependency should fail if it is not etag enabled.

# Running code

All code is compiled into intermediate bytecode. Libraries contain only a metadata header (which contains at least the name and a hash of its original code) and the intermediate bytecode that results in combining all dependencies' bytecode, similar to static linking. Runnables are compiled into similar intermediate bytecode with a header, with all dependencies similarly compiled inline.

Dependencies are always deduplicated by the hash of their code. This de-duplication scheme means that two different libraries of the same name and version may not be deduplicated because they're actually different, and libraries of different names and versions may be deduplicated because they're actually the same code. The result is that naming and versioning are irrelevant to the compiler. The hash of the code (not bytecode) determines duplication.

Runnables come packaged in an tar file that contains a native executable (for the target host) that translates, caches and executes the native code that resulted from translation. The structure of the tarball looks like this:

```
metadata
runner
bytecode
```

# Testing and Debugging code

Any library or runnable can have tests in it. They must all be in the same root directory, and you tell the compiler where they are in a `tests` file, at the top level of your project.

## Running and Debugging Tests

Ideal has test runner and debugger, which loads your project into a [REPL](http://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop) that has all of your project's dependencies in the namespace. You can run all of your tests by typing `test`, all of the tests in a set of subdirectories by typing `test subdir1 subdir1`, or a single test by typing `test subdir1/filename/testname`.

Debugging is done in the test runner as well. Type `debug-linenum dir/file linenum` to set a breakpoint on a line, or `debug-function dir/file funcname` to set a breakpoint when a function is invoked. Watchpoints and other features are be supported too.

## Writing tests

Ideal has a test harness built in. The test harness pre-builds a `main` function for any group of tests that you decide to run, so you never have to write one. A source file has 1 or more tests in it, and each has a name (description of the behavior that it's testing). Since names need not be unique, typing `test subdir1/filename/testname` might execute more than one test. Test syntax should be like [specs2](http://etorreborre.github.io/specs2/) or other frameworks from which it was inspired:

```
"myFunction should fail when you pass a negative number" ! myFuncFailsWithNegative
...
```

# Type Safety

Since Ideal is a compiled and strongly typed language, there's a good foundation for type safety.
