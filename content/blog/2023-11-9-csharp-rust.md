---
author: "Aaron Schlesinger"
date: 2023-11-09T21:58:32Z
title: 'Calling Rust from C#, and back again: Part 1'
slug: "csharp-rust"
tags: ['systems', 'languages', 'rust', 'csharp']
---

I'm working on a project at Microsoft called Hyperlight. The project is all about making very small-footprint virtual machines that start up specific kinds of applications very quickly. At a very high level, we're using modern virtualization technology and some concepts from Unikernels to make these things happen.

Mark Russinovich highlighted the project at [his keynote at Microsoft Build 2023](https://www.youtube.com/watch?v=Tz2SOjKZwVA), so if you want more background on the project, go [watch that keynote](https://www.youtube.com/watch?v=Tz2SOjKZwVA).

The project was prototyped initially with C#, but it turns out C# isn't the best choice for a variety of reasons. Two of the biggest of those reasons are as follows:

- It takes care and effort to call native code, like `ioctl` system calls, from C# (or really, all .Net) applictions
- It also takes care and effort to deal with memory not controlled by .Net -- often called _unmanaged memory_

## Background: Rewriting C# in Rust

In response to these challenges (and more), we decided to move the project to [Rust](https://rust-lang.org) because it's a much better fit for the needs of the project. Rust is a fast, safe, and modern systems programming language, and a great fit for this project.

We chose to move the C# codebase incrementally to Rust, and wanted to end up with both a C# and Rust SDK, with the former being a thin wrapper around the latter. Ironically, this work forced us to write a lot of temporary C# code that did precisely the kinds of things we were trying to avoid, like calling into native code and dealing with unmanaged memory.

As the rewrite progressed, we were able to remove some of that code, but two things became clear as we began to understand the mechanics of the final joint C#/Rust codebase:

1. We indeed would always need to call "system calls" (sometimes called [P-Invoke](https://learn.microsoft.com/en-us/dotnet/standard/native-interop/pinvoke) in C#) from C#, since that's how the Rust APIs are exposed
2. In some cases, we'd need to go the other way and _call C# functions from Rust_

In our experience so far, (1) is relatively well defined, but (2) is not. We figured out how to do both in a somewhat robust and reliable way. In this post, In this post, I'll describe (1), and save (2) for a follow-up.

## Calling Rust from C#

C# was designed from the start to be able to call "native" APIs, mainly so developers could still integrate with the many C/C++ APIs already present in Windows. This feature is really useful for us, because we can implement some Rust code, expose it as a C API, then call it from C#.

Assuming we design this system carefully, the C# code won't know it's calling Rust code, and the Rust code won't know it's being called from C#. Decoupling is a beautiful thing!

In other words, C# can call C APIs [1] and Rust can produce C-compatible APIs, so on either side of the "fence", we just need to figure out how to deal with C-compatible APIs.

### Creating C APIs in Rust

Rust has strong support for creating C-compatible APIs [2]. While there are a few gotchas, if you can compile Rust code like the following, it will be callable by C code:

```rust
#[no_mangle]
pub extern "C" fn callable_from_c(i: i64) -> i64 {
    i+1
}
```

There's not a lot of code here, but there are some important features to note, as follows:

- The `#[no_mangle]` is important, because it ensures the Rust compiler doesn't modify ("mangle") the name, so C code can call it without understanding the internals of Rust [name mangling](https://rust-lang.github.io/rfcs/2603-rust-symbol-name-mangling-v0.html)
- The `extern "C"` is also important, because it tells the Rust compiler to produce a C-compatible ABI for the function, so C code can call it without understanding the internals of Rust [ABIs](https://en.wikipedia.org/wiki/Application_binary_interface)
- Finally, the types of parameters and return types must all be `#[repr(C)]`
  - `#[repr(C)]` is code you put atop a Rust type that tells the Rust compiler the type must be binary-compatible with C code. It will not work on all types, but if you do use it and your code compiles, it's guaranteed to work
- Rust pointers (not references!) are compatible with C pointers, so if you have a type that is not `#[repr(C)]`, you can use a pointer to it in your C-compatible API
  - Warning: most of the Rust code you'll write that deals with pointers in any meaningful way, will be `unsafe`!

Once you have a bunch of these C-compatible functions, you need to put them in a library crate that _only_ has type `cdylib`, like the following:

```toml
[lib]
name = "my_c_lib"
crate-type = ["cdylib"]
```

>Note: if you have a `cdylib` crate, you can't also compile it as a "standard" Rust crate. If you want one of those too (hint: you probably will, eventually!), split your C-compatible code into a separate crate

### Calling C APIs from C#

The other side of this is calling C APIs from C#, which as I mentioned above is a well-known feature of C#. There are a lot of resources out there (reminder: search for "P-Invoke"), but I'll just give a quick example here:

```cs
[DllImport("my_c_lib", SetLastError = false, ExactSpelling = true)]
// You may need this, depending on how your code searches for shared libraries
// (i.e. DLLs on Windows or shared objects on Linux).
// If you do use it, the C# compiler will throw warnings 
[DefaultDllImportSearchPaths(DllImportSearchPath.AssemblyDirectory)]
[return: MarshalAs(UnmanagedType.U1)]
// An i64 in Rust corresponds to a long in C#
static extern long callable_from_c(long input);
```

## Calling C# from Rust

I've written a lot here, but believe it or not, we only covered the first half of this process. Going the other way -- calling C# from Rust -- involves more details, including some .Net runtime internals. I will cover those details in my next post. See you then!

---
_[1] Technically, the C# compiler knows how to call C APIs, and the runtime knows how to call C [ABI](https://en.wikipedia.org/wiki/Application_binary_interface)s_

_[2] Similar to what the C# compiler does for calling C API/ABIs, the Rust compiler knows how to expose ABI-compatible C APIs_
