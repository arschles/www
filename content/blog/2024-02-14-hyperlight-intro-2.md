---
author: "Aaron Schlesinger"
date: 2024-01-08T21:58:32Z
title: 'What is Hyperlight (part II)?'
slug: "hyperlight-overview-2"
tags: ['systems', 'languages', 'rust', 'csharp']
---

In [part 1](/blog/hyperlight-overview-1) of this series, I detailed the motivations and background of the Hyperlight project. In doing so, I also mentioned some of the challenges we’ve set out to solve and how, at a high level, we've solved them. In this post, I’m going to talk about these challenges in much greater detail. 

Hyperlight is systems-level software, which means it’s concerned with the lowest-level details of how software runs on a computer. The concepts and challenges with which it's concerned are very nuanced and detailed, so I think it's especially critical to talk about them directly. Thus, you'll find this post contains much more technical content than the last. I make a point herein to explain some of the more complex terms and concepts as we progress.

Let’s get started! 

## Virtual CPUs: the workhorse of Hyperlight 

In part 1, I mentioned Hyperlight relies on hypervisor, or virtual machine (VM), isolation to ensure we provide the security guarantees to which we’ve committed. I also explained that our use of Hypervisor isolation provides a set of virtualized hardware devices we use to run native code inside a VM. For the purposes of this post today, it's helpful to think of a single VM as an individual "slice" of a physical computer running in the cloud [^1].

With this "slice" metaphor in place, we can visualize a physical machine as a pie with a specified number of slices, each of which can run arbitrary code -- which we call "guests" in Hyperlight -- that has been compiled to the underlying hardware architecture of the physical machine. In many cases, this architecture is [x86](https://en.wikipedia.org/wiki/X86), but [ARM](https://en.wikipedia.org/wiki/ARM_architecture_family) is becoming more widely used, and it’s very possible [RISC-V](https://en.wikipedia.org/wiki/RISC-V) will become relevant in the cloud as well. 

Just as a CPU generally executes code on your laptop or desktop computer, the virtual CPU (vCPU) executes code inside a Hyperlight `Sandbox`. `Sandbox`es can run atop KVM or Hyper-V, but in either case, the **virtual CPU (vCPU)** plays a central role.

### “Low level” code execution basics

On nearly all systems, the CPU/vCPU is the most important part of the hardware/software stack involved in making software run [^2]. However, while the vCPU is important, there are at least several more components involved including virtual hardware devices and the virtual memory hierarchy. Since there are so many components involved, an operating system (OS) kernel is generally required to manage them all and provide something like a "runtime" to abstract them away.

This abstraction is presented as features like filesystems, a process model, OS threads, and more, and as you might be aware, they're usually critically important to power general-purpose software. Recall, however, from part 1 that our use case is currently primarily a functions-as-a-service (FaaS) model that requires nearly none of them. Thus, we’ve done extensive work to _avoid_ these features and not provide any abstractions whatsoever. The result is we're left with only a vCPU, its virtual registers, and a single linear block of memory we map into the virtual machine.

## Results of stripping away the OS

The results of this aggressive simplification are important for us, and are two fold as follows:

- **No boot-up process**: because we have stripped away most features traditionally provided by the kernel, we can remove the OS entirely. The result of this strategy is significantly reduced latencies to execute an individual function in the FaaS system
- **Reduced attack surface**: because there’s operating system at-large, the attack surface inside the VM – the amount of software that can be compromised by an attacker – is exactly as big as the running binary 

This approach may seem very exotic or maybe even weird, but it's actually been used several times throughout history. Most recently, a family of technologies called "unikernels" has been developed to provide similar benefits. These systems, like Hyperlight, require you compile against or link directly to the OS-level abstractions your code requires, rather than requiring a pre-installed OS on some raw hardware. The result is that Unikernel-based applications ship _only_ with what they actually use and nothing more. The advantages are similar to the ones in the list above.

## Executing code without an operating system 

That being said, unikernel technologies are, generally speaking, not very popular. If executing code directly on a vCPU, without the help of an OS, makes you scratch your head a bit, you’re not alone – I had the same reaction for the first few weeks I worked on this project. The biggest question I kept returning to was - how can we support anything more than just simple math operations on a vCPU?

The key insight allowing us to support richer functionality in this very spartan environment is simple: we run a very stripped-down binary, called a "guest," directly in the Hypervisor-based VM -- a `Sandbox` in Hyperlight terminology -- and that binary in turn runs a managed language inside its VM.

In part 1, we talked about running a _different kind of VM_ inside the `Sandbox` to execute the [Wasm](https://webassembly.org) instruction set. We could run other, similar VMs as well, like Javascript, Python, or others.

While the nomenclature here is confusing, the idea is more clearly explained below, in terms of what an OS does.

## Wasm, and a different kind of VM

If a user compiles their code to Wasm, we don’t actually execute the resulting Wasm instructions directly on the vCPU. Instead, an aforementioned guest binary -- called hyperlight-wasm -- runs directly on the vCPU, then does the following completely within the Sandbox:

- Loads the Wasm instructions into memory 
- Launches a Wasm VM (we currently use [WAMR](https://github.com/bytecodealliance/wasm-micro-runtime) due to its small footprint and controlled feature set)
- Executes the Wasm instructions inside the VM
- Returns the result to the host 

Essentially, our hyperlight-wasm binary is a very, very lightweight approximate equivalent to an operating system, hence our earlier reference to Unikernels. We also unlock a very important benefit to this overall approach. Since the Wasm instruction set and associated VMs are very lightweight and require very little “infrastructure” -- like hardware devices, a kernel, time sharing facilities, and so on -- we can start up code very quickly with relatively few computing resources. We can also start up some of the other "infrastructure" we need, like filesystems or virtual networking devices, on demand, which further optimizes performance.

The results are exciting. In most cases, we can start a sandbox and begin executing a user’s Wasm module in approximately 800 microseconds (that’s 4/5 of 1 millisecond). 

This latency is competitive with some other hosted FaaS systems, which generally don’t currently provide hardware-level isolation, and an order-of-magnitude improvement over other FaaS systems, which generally do so. We believe the key innovation Hyperlight provides, especially when used in such a FaaS scenario, is the combination of the hardware-level isolation from the latter system and the low latencies of the former system. 

## Rust

Because Hyperlight is such foundational technology due to its position in the software/hardware stack, the tools we use to build it are critically important. I'm going to change gears here and move from talking about how we designed Hyperlight to the tools we used to actually build it. I believe the former is just as important as the latter.

Recall from part 1 that Wasm provides [mathematically-provable safety guarantees](https://www.usenix.org/system/files/sec22-bosamiya.pdf), but these proofs are only valid if the software that implements the math is bug-free. I believe the innovations described above can truly make a big,positive change in the efficiency of cloud workloads, but the bug-free requirement holds as much in this project as it does in Wasm VMs or any other security-focused systems software.

It's hard (but possible) to guarantee bug-free software, so we at least have the goal to eliminate several _classes_ (or, types) of bugs. On the top of our list is several types of memory safety bugs like [buffer overflows](https://en.wikipedia.org/wiki/Buffer_overflow) and [use-after-free](https://owasp.org/www-community/vulnerabilities/Using_freed_memory).

To accomplish these goals, we must implement the system described herein [^3] in a robust, secure and efficient manner. The Rust programming language has proven to be the best tool for the job. 

### Background on Rust

Rust is a relatively new language, especially in the group of systems languages that let the programmer access low-level functionality on the machine. The language differs from most of these other systems languages, however, by providing some very strong memory safety guarantees.  

To build Hyperlight, we need almost exactly those features. Recall from part 1 that memory safety is an area of focus for the software technology industry at-large, and Rust is being adopted and supported in large part because it does a good job of providing memory safety guarantees. In some cases, it even can provide formal mathematical guarantees, and we've in fact been able to reap some of those benefits in our codebase.

Bugs are still possible, but we believe Rust gives us the best chance to build safe, reliable and generally high-quality software (see: clippy, other strict code requirements like disallowing unused functions etc...) 

## Conclusion 

Overall, I'm very excited about the future of Hyperlight along two dimensions:

1. **The use-cases for the system**: while FaaS is the first and primary use case, I think there are other use cases for which a system like Hyperlight could be a great fit
2. **The technology**: our use of Rust is expanding within the project, we're developing new guests to run different VMs, and we're expanding the security features of the system

Overall, I think Hyperlight represents a new way to run specific kinds of software in the cloud, and I'm excited to see how it gets used over time.

---

[^1] Note this metaphor is not perfect because the physical machine is not always divided equally or even logically, and in many cases may be oversubscribed on vCPUs or memory

[^2] These days, GPUs and other specialized computing devices are becoming more and more common, especially in modern AI/ML and computer graphics applications

[^3] And not just this system as specified. We also need to write robust testing and monitoring/alerting code as well as everything you’d need to run software like this in an enormous distributed system in the cloud
