---
author: "Aaron Schlesinger"
date: 2024-02-14T21:58:32Z
title: 'What is Hyperlight (part II)?'
slug: "hyperlight-overview-2"
tags: ['systems', 'languages', 'rust', 'csharp']
---

In [part 1](/blog/hyperlight-overview-1) of this series, I detailed the motivations and background of the Hyperlight project. In doing so, I also mentioned some of the challenges we’ve set out to solve and, at a high level, how we've solved them. In this post, I’m going to talk about these challenges in much greater detail.

Hyperlight is systems-level software, which means it interacts with some of the lowest-level details of how software runs on a computer. The concepts and challenges with which it's concerned are very nuanced and detailed, so I think it's especially critical to talk about them directly, right alongside the high-level description from part 1.

Thus, you'll find this post contains much more technical content than the last, and I try to take time to explain some of the more complex terms and concepts throughout.

Let’s get started! 

## Virtual CPUs: the workhorse of Hyperlight 

In part 1, I mentioned Hyperlight relies on hypervisor isolation to ensure we provide the security guarantees to which we’ve committed. I also explained that our use of Hypervisor isolation -- also called virtual machines or VMs -- provides a set of virtualized hardware we use to run native code inside a VM. For the purposes of this post today, it's helpful to think of a single VM as an individual "slice" of a physical computer running in the cloud [^1].

With this "slice" metaphor in mind, we can visualize a physical machine as a pie with a specified number of slices, each of which can run arbitrary code, which we call a "guest" in Hyperlight parlance, that has been compiled to the underlying hardware architecture of the physical machine [^2].

Just as a CPU generally executes code on your laptop or desktop computer, a VM's virtual CPU (vCPU) executes code inside a Hyperlight `Sandbox`, which is the term we use to describe a "slice." 

`Sandbox`es abstract away the underlying VM and vCPU, and currently support KVM or Hyper-V across Windows and Linux.

### “Low level” code execution basics

On nearly all systems, the CPU/vCPU is the most important part of the hardware/software stack involved in making software run [^3]. There are, however, at least several more components involved in making most running code work, including virtual hardware devices, the virtual memory hierarchy, and many more. Since there are so many components involved, an operating system (OS) kernel is usually required to manage them all and provide a manageable abstraction for programs to compile (and link) against.

This abstraction is commonly presented as a group of [system calls](https://en.wikipedia.org/wiki/System_call) ("syscalls") that allow programs to safely interact with the OS kernel, which in turn interacts with the hardware. Through syscalls, a program can interact with a filesystem, process model, OS threads, and more. Most general-purpose software requires all of these features and more.

## Low level code execution for a specific use case

Recall, however, from part 1 that we're not primarily aiming to power general-purpose software with Hyperlight. Instead, we're trying to run smaller pieces of software intended to execute repeatedly in response to an event. Typically, this model is called "functions-as-a-service," abbreviated to "FaaS". The FaaS model presents requirements that are much different from those of general-purpose software.

In particular, this model requires a set of features much, much smaller than the those provided by an OS kernel. In many cases, we can get away with providing none of them. Thus, we default to _avoiding these features altogether_ inside Hyperlight `Sandbox`es. Thus, a `Sandbox` provides a runtime environment that by default has only a vCPU, its virtual registers, and a single linear block of memory.

### What you get when you run without an OS

While it might seem like no program could ever execute in this stripped-down environment (and maybe that this is a dumb idea!), some can. And, for our specific use case, this idea turns out to work very well.

First, the main benefits to this approach are twofold, as follows:

- **No boot-up process**: because we've removed the OS entirely, there's no kernel to boot up and we can start running a user's code with significantly reduced latency
- **Reduced attack surface**: because there’s no OS, the attack surface inside the VM – the amount of running software that can be compromised by an attacker – is exactly as big as the running binary, and usually much smaller than that of a running OS

### Borrowing from existing approaches

Second, this approach is not really novel; it's been used several times throughout history, and its most recent incarnation is called [unikernels](http://unikernel.org).

These systems, like Hyperlight, require you to compile and link against a library, in your language of choice, that provides the abstractions you need to run your code. Thus, instead of calling OS-provided syscalls, you call a function provided by a library, which is then directly responsible for interacting with the underlying system to do what you need.

Since modern compilers and linkers are very good at optimizing code, the result of this approach is that unikernel applications ship with _only_ the abstractions they need and nothing more. There are drawbacks of this approach (no shared implementation, harder to build software, etc...), but the benefits approximately match those in the above list.

## Unikernels: an unpopular but useful (for us) technology

If you've heard of unikernels, you're probably in the minority. They're not very popular at all, and for good reason; they are a poor choice for running much of the software we use in our everyday lives.

As I write this, for example, I'm running at least 9 different applications on my laptop, each of which is relying on a wide range of syscalls. The OS, which implements all of those syscalls, is in turn utilizing all the different features of the hardware inside my laptop.

This is a very common scenario, and it thus makes sense to have one common, underlying operating system to manage that hardware in one place and provide a single, managed abstraction to all the applications that run atop it.

In the cloud, though, we have a very different situation. Instead of running 10 or 20 different applications, we're running millions or more, each with wildly different requirements. Some need to run for a long time, some need to run for a very short time, some need to use a lot of memory, some need to use a lot of CPU, and so on.

In this kind of environment, it pays to specialize as much as possible. As I mentioned above, Hyperlight targets applications that need to run for a very short time and need to use a very small set of resources.

For specifically this use case, unikernel-like technologies like Hyperlight are a great fit.

## How you actually execute code without an operating system 

We now know some important details behind Hyperlight and we get why it's important. Let's now look at how we accomplish Hyperlight's core goal of running _arbitrary customer code_.

First, it would be a bad product experience to require our customers compile their code against some unikernel library we provide. It's also not feasible to rewrite their code to use the unikernel library and then compile it to a unikernel binary.

The solution to this problem is (as with many problems in computer science) to add another layer of abstraction. Back in part 1, I talked about two different kinds of VMs, and how we're running one VM inside another VM. The inner VM is the abstraction in question.

We build a relatively small number of guest binaries, each of which runs a different VM. We then load the appropriate guest binary (based on the code the customer furnishes) into a `Sandbox`, load their code into the right place in memory, and finally execute the `Sandbox`.

We currently have a guest binary called `hyperlight-wasm` that can execute the [Wasm](https://webassembly.org) instruction set, and we've prototyped others.

### Wasm, and a different kind of VM

The `hyperlight-wasm` guest is an important one because it can accommodate any programming language that can be compiled to Wasm. That list of languages is large and growing, which means Hyperlight can easily support a wide array of languages and their associated ecosystems with very little extra work from us.

If a user gives us their code in some language, we need to compile it to Wasm and load the `hyperlight-wasm` binary into a `Sandbox`'s memory. The binary then does roughly the following:

- Loads the compiled Wasm instructions into memory 
- Launch a Wasm VM (we currently use [WAMR](https://github.com/bytecodealliance/wasm-micro-runtime) due to its small footprint and controlled feature set)
- Execute the Wasm instructions inside the VM
- Return the result to the host (the code that started the `Sandbox`)

Essentially, our `hyperlight-wasm` binary is a very, very lightweight approximate equivalent to an operating system that only executes Wasm instructions.

Also, the Wasm instruction set supports only a simple processing unit and a linear memory space -- a very close match to the capabilities of a default `Sandbox`. Wasm does, however, have facilities like [WASI](https://wasi.dev) that go beyond the defaults and utilize some of the familiar OS-level abstractions. If we encounter those use cases, we can initialize and provide those abstractions in a lazy or "on-demand" fashion and preserve the fast startup times I described above.

## Results and conclusion

The results of this approach are exciting. In most cases, we can start a sandbox and begin executing a user’s Wasm module in less than 1 millisecond (1/1000 of a second).

This latency is competitive with some other hosted FaaS systems, which generally don’t currently provide hardware-level isolation, and an order-of-magnitude improvement over other FaaS systems, which generally do so. I believe the key innovation Hyperlight provides, especially when used in such a FaaS scenario, is the combination of the hardware-level isolation from the latter system and the low latencies of the former system. 

Because we've achieved this combination, I'm very excited about the future of Hyperlight for the following two reasons:

1. **The use-cases**: while FaaS is the first and primary use case, I think there are other use cases for which a system like Hyperlight could be a great fit
2. **The technology**: our use of Rust is expanding within the project (I talk about our use of Rust below in an addendum), we're developing new guests to run different VMs, and we're expanding the security features of the system

Overall, I think Hyperlight represents a new way to run very specific kinds of software in the cloud, and I'm excited to see how it grows.

## Addendum: Rust powers everything

Because Hyperlight is such foundational technology due to its position in the software/hardware stack, the tools we use to build it are critically important for its security, speed, reliability, and overall success. I'm going to move from discussing how we designed Hyperlight to discussing some of the tools we used to actually build it. I believe the former is as important as the latter in this case.

Recall from part 1 that Wasm provides [mathematically-provable safety guarantees](https://www.usenix.org/system/files/sec22-bosamiya.pdf), but these proofs are only valid if the software that implements the math is bug-free. The same requirement applies to Hyperlight itself.

It's hard (and in many cases, impossible) to guarantee bug-free software, so we aim to eliminate several _classes_ (or, types) of bugs and have taken aim at several types of memory safety bugs like [buffer overflows](https://en.wikipedia.org/wiki/Buffer_overflow) and [use-after-free](https://owasp.org/www-community/vulnerabilities/Using_freed_memory) issues.

To accomplish these goals, we must implement Hyperlight [^4] in a robust, secure and efficient manner. The [Rust programming language](https://www.rust-lang.org), or "Rust," has proven to be an invaluable tool to help us do that job.

### Why Rust for Hyperlight?

Rust is a relatively new language, especially in the group of systems languages that let the programmer access low-level functionality on the machine. The language differs from most of those other systems languages, however, by providing some very strong memory safety guarantees.  

To build Hyperlight, we need almost exactly those features. Recall from part 1 that memory safety is an area of focus for the software technology industry at-large, and Rust is being adopted and supported in large part because it does a good job of providing memory safety guarantees. In some cases, it even can provide formal mathematical guarantees, and we've in fact been able to reap some of those benefits in our codebase.

Bugs are still possible, but we believe Rust gives us the best chance to build safe, reliable and generally high-quality software [^5]. There is much more to say about specifically how we exploit the power of Rust, but you can get some ideas from the "Calling Rust from C# and back again" series ([part 1](/blog/csharp-rust) and [part 2](/blog/csharp-rust-2)) I wrote on this blog previously.

---

[^1] Note this metaphor is not perfect because the physical machine is not always divided equally or even logically, and in many cases may be oversubscribed on vCPUs or memory.

[^2] In many cases, this architecture is [x86](https://en.wikipedia.org/wiki/X86), but [ARM](https://en.wikipedia.org/wiki/ARM_architecture_family) is becoming more widely used, and it’s very possible [RISC-V](https://en.wikipedia.org/wiki/RISC-V) will become relevant in the cloud as well

[^3] These days, GPUs and other specialized computing devices are becoming more and more common, especially in modern AI/ML and computer graphics applications.

[^4] And not just this system as specified here and in part 1. We also need to write robust testing and monitoring/alerting code, security features, memory management features, and everything else you’d need to run software like this in an enormous distributed system in the cloud.

[^5] The benefits of Rust in this domain go beyond memory safety. A good representative example is [Clippy](https://doc.rust-lang.org/stable/clippy/index.html), which is a static analysis tool included by default with the toolchain that can either directly catch _before compile time_ or prevent code that is likely to lead to future bugs.
