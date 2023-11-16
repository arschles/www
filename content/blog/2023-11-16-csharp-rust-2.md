---
author: "Aaron Schlesinger"
date: 2023-11-16T21:58:32Z
title: 'Calling Rust from C#, and back again: Part 2'
slug: "csharp-rust-2"
tags: ['systems', 'languages', 'rust', 'csharp']
---

Back in [part 1](/blog/csharp-rust), I talked about Hyperlight, its history, and some of the challenges involved with the project, including calling [Rust](https://rust-lang.org) from [C#](https://learn.microsoft.com/en-us/dotnet/csharp/).

At the very end of that post, I promised another talking about the other direction -- calling C# from Rust. So, here we go!

## The Challenge

Fundamantally, we want to have C# call into Rust, which we covered in part 1, but in some cases, we want to pass a C# function to Rust, and have the Rust code call back into C#.

To do that second part -- the callback -- we have to dive into a topic the C#/.Net world calls _unmanaged_. For our purposes, "unmanaged" means C APIs/ABIs implemented with Rust. The process to make C# code call unmanaged code is well documented, and the C# compiler and .Net runtime give us explicit support for doing so.

It's not much of a stretch to go one step further and _implement_ those C APIs with Rust. We essentially use C as an abstraction layer between C# and Rust, and it works remarkably well.

On the other hand, it's far less clear how to get that callback functionality working. We have to put together several different C# and .Net features and wire those things up with Rust to make this work, and I'll show herein how to do it.

## Pieces of the puzzle

The process to get this callback system working is three-fold, as follows:

1. Convert a C# function (usually called a "delegate") into a C function pointer
2. Consume that C function pointer in Rust and call it
3. Finally, we need to clean up the function pointer's memory when we're done with it

I'll describe these steps below, each in more detail.

## Converting a C# delegate to a C function pointer

In this first step, the overall goal is to convert a C# [`delegate`](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/delegates/) into an unmanaged function pointer. "Unmanaged function pointers" in C# roughly correspond to C function pointers. Given a C# [`Action`](https://learn.microsoft.com/en-us/dotnet/api/system.action-1?view=net-7.0) -- a function closure with no return value -- the code for converting it to a function pointer isn't obvious, but relatively straightforward, as shown below.

```c#
// this is the action we want to convert to a C-compatible function pointer.
Action<bool> myAction = (myBool) => {
  Console.WriteLine(myBool);
};

// establish a delegate type, and convert our action to it
private delegate void MyDelegate(bool param);
MyDelegate myActionWrapper = myAction;

// then tell the .Net runtime, including the garbage collector, to "forget"
// about this function pointer. we need to do this so we can pass the 
// function pointer to Rust, and be sure that C# won't delete it at some
// point before Rust needs it.
GCHandle myActionHdl = GCHandle.Alloc(myActionWrapper);

// finally, convert our delegate to a function pointer.
IntPtr myActionPtr = Marshal.GetFunctionPointerForDelegate<MyDelegate>(
  myActionWrapper
);
// we can now pass myActionPtr to Rust via a C API
```

>This code could also work with C# [`Func`](https://learn.microsoft.com/en-us/dotnet/api/system.func-2?view=net-7.0)s, which are function closures with return values. In either case, you have to ensure your parameter types and return types (where applicable) are compatible with C.

## Passing the C#-generated function pointer to Rust

Now, in Rust-land, we have to accept the `IntPtr myActionPtr` we just created in C#-land. Rust has very good, nearly zero-cost [^1] support for both calling C functions and defining C-callable functions, but this topic is nevertheless very nuanced.

Collectively these features are called "FFI" (foreign-function interface), but we're interested in only the latter feature. Defining a function callable by C code is as simple as the following:

```rust
/// accept a boolean and print it out
#[no_mangle]
pub extern "C" fn my_func(val: bool) {
  println!("hello! the given value is {}", val);
}
```

Take note of a few important highlights of this function prototype.

- **`#[no_mangle]`** - by default, `rustc` (the Rust compiler) reserves the right to "mangle" the names of functions, methods, and other types. The reasons for mangling are beyond the scope of this article, but we want to turn this feature off for this function so we can call it from C code
- **`extern "C"`** - this nomenclature tells Rust to generate code for this function that adheres to the C [calling convention](https://en.wikipedia.org/wiki/Calling_convention)
  - You can read more about calling conventions in the [Rust nomicon](https://doc.rust-lang.org/nomicon/ffi.html#foreign-calling-conventions). 
- **Argument and return types** - the Rust compiler will make sure all types in the public function prototype -- parameters and return values -- are FFI-compatible. This feature is invaluable for writing FFI functions in Rust because you can be sure your FFI function is at least technically callable [^2] from C if it compiles.
- **Header file and linkage** - The above function prototype tells the compiler what to output for that function, but doesn't guarantee C code can compile or link against it. I can't go into more detail here without doubling the length of this post, but here's the very high-level way to get that working:
  - Use [cbindgen](https://github.com/mozilla/cbindgen) to generate C header files from your `extern "C"` functions
  - Set your crate type to [`cdylib`](https://doc.rust-lang.org/reference/linkage.html) so your C programs can link against your code

#### Functions as parameters

Armed with this ability to define C functions in Rust, we can now create higher-order functions as follows:

```rust
pub extern "C" fn outer_func(inner_func: extern "C" fn(bool)) {
  // we likely want to save our inner_func value for later usage, but
  // that's outside the scope.
  //
  // we can effectively call into C# by simply calling inner_func:
  inner_func(true);
  inner_func(false);
}
```

## Cleaning everything up

Cleaning up Rust-owned function pointers generated from C# can be a complex topic that requires some care.

#### Rust vs. C# ownership

First, Rust doesn't truly "own" the function pointer, since it can't reliably and correctly clean up its backing memory (only C# can do that). This fact implies we need to make sure Rust is "done" with it before cleaning it up in C#. There are a lot of different possible ways to do that, varying from just "eyeballing it" to using reference-counting or other more formalized strategies. Inside Hyperlight, for example, we use a home-grown system vaguely similar to [reference counting](https://en.wikipedia.org/wiki/Reference_counting).

#### C# cleanup code

Next, after we're sure Rust is "done" with the function pointer, we can clean it up in C#. Recall in the first step ("Converting a C# delegate to a C function pointer", above) we called `GCHandle.Alloc` to get the garbage collector to forget a `delegate`, which means we need to manually clean it up. The following code does the cleanup work:

```c#
// First, if the type to which myActionHdl points is an IDisposable, clean
// that up by calling Dispose() on it
var tgt = myActionHdl.Target as IDisposable;
if (null != tgt) {
  tgt.Dispose();
}
// Finally, free the actual memory for the delegate to which myActionHdl points
myActionHdl.Free();
```

>Recall that `myActionHdl` is a `GCHandle` returned from `GCHandle.Alloc`

#### Caveat: closures

And finally, a big caveat. The C# `delegate` from which we created `myActionHdl` is a _closure_, which means it can _close over_ and/or reference memory outside of its definition. In either case, the garbage collector may still clean up that memory.

If that happens, we'll run into use-after-free issues that will likely be very difficult to debug, since the function is invoked from Rust. There is no perfect way to solve this issue, since we're fundamentally bridging two completely different environments (C#/.Net and Rust) over a low-fidelity abstraction (C), but once again, reference counting can help here.

## Conclusion

Putting the concepts from [part 1](/blog/csharp-rust) and this post together, we now have the tools to make C# code call Rust and vice-versa. If you build out a joint C# and Rust codebase, you'll likely encounter several practical issues, the biggest of which will likely be memory management.

Throughout this post, I showed how to manually reconcile the .Net memory management system -- the garbage collector -- with the Rust one -- ownership, borrowing and deterministically freeing memory without any runtime overhead. Doing this reconciliation usually requires a lot of error-prone work, so in many cases it's helpful to build a memory management system to unify the two environments.

As I implied several times herein, Hyperlight uses a system that uses something similar to reference-counting to do so. Stay tuned for part 3, where I'll detail that system.

---
_[^1] Rust's foreign-function-interface (FFI; basically, its ability to call C code) is officially a [zero-cost abstraction](https://doc.rust-lang.org/beta/embedded-book/static-guarantees/zero-cost-abstractions.html) (also see [this document](https://blog.rust-lang.org/2015/04/24/Rust-Once-Run-Everywhere.html#:~:text=To%20communicate%20with%20other%20languages,performance%20to%20C%20function%20calls.)), but for many practical applications of `extern "C"` functions, you will need to do additional memory allocations, which incur a cost._

_[^2] There is unfortunately a big distance between theory and practice in this case but regardless, knowing your Rust code can theoretically be called from C is a big step forward._
