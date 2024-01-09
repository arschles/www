# What is Hyperlight?

After I wrote those posts about C#-to-Rust and Rust-to-C# communication ([part 1](./2023-11-9-csharp-rust.md) and [part 2](./2023-11-16-csharp-rust-2.md)), a lot of people reached out with questions about what Hyperlight is in the first place. There's not a lot of information out there about the project besides what I've written, what's been presented at Microsoft Build (more on that in a second), and word-of-mouth, so I'm going to fill in the gaps here.

I'm going to take a break from the C# <-> Rust series (which I'll resume soon!) to talk about Hyperlight in more detail. I'll cover what it is, why we're building it, and how it works.

# Background

At the 2023 Microsoft Build conference, [Mark Russinovich](https://www.linkedin.com/in/markrussinovich/) showed a [demo of Hyperlight](https://www.youtube.com/watch?v=Tz2SOjKZwVA) in his Azure Innovations keynote. In that same keynote, he described Hyperlight as a solution for improving the management and security of [Web Assembly (Wasm)](https://webassembly.org) workloads on Azure, and explained the motivation of the project as follows: 

>“Wasm lacks the kind of isolation that we require for running a public cloud” 

In this post, I’m going to detail what he means by that and describe Hyperlight in greater detail. This is the first of two posts, and in the second, I’ll dive into much greater detail on how it works. 

# Built on the shoulders of giants 

Hyperlight borrows from, and builds on, substantial prior art, and it’s important to not only credit the technologies on top of which we’ve built this project, but also to establish a common foundation before we delve deeper into the project itself.

## Hypervisors

One of the most important of these technologies is the hypervisor. Simply put, a hypervisor is a system, often built right into your operating system and hardware stack (but not always! Read more about the various kinds of hypervisors [here](https://www.geeksforgeeks.org/hypervisor/)), that can create virtual “machines” (Hypervisors are also often called Virtual Machines or VMs), or computing devices, on top of which you run some other piece of software. 

These systems have a long [history](https://en.wikipedia.org/wiki/Timeline_of_virtualization_development) [of](https://dl.acm.org/doi/fullHtml/10.1145/3365199) [research](https://web.eecs.umich.edu/~prabal/teaching/eecs582-w11/readings/Gol74.pdf) and are one of the critical foundational technologies behind cloud computing. These days, VMs are flexible enough to run software ranging from the most demanding industrial workloads in the cloud to full-blown operating systems on a laptop for a consumer. In fact, I’m writing this article on Windows running in a virtual machine on my Mac as we speak! 

Importantly for our purposes, hypervisors provide virtualized hardware devices -- including CPUs, hard disks, networking devices, and memory -- on which our “guest” software (the software that runs Wasm workloads; more on these mechanics below) runs. This hardware virtualization layer provides an important security isolation boundary we can use to guarantee a specific level of security to our users’ running applications. 

VMs often ship with a lot of associated technologies, many of which we don’t use directly, but are important to understand for context. We’ll talk more about them when we dive deeper in part 2. 

## Rust

Another important technology we use is the [Rust programming language](https://rust-lang.org), or just “Rust”. I’ll talk less about Rust in this initial post in the series than in post 2, but I do want to cover one of Rust’s most important features for our use case: memory safety. 

Traditionally, a limited set of programming languages were available to do systems programming – writing programs to interact directly with the computer’s hardware. The most popular were the C and C++ languages, and both required developers to manually manage the memory they wanted to use. Situations where a developer manages memory incorrectly – like accidentally trying to use it after they’ve already freed it – have led to [many](https://www.cisa.gov/news-events/news/urgent-need-memory-safety-software-products) [security](https://msrc.microsoft.com/blog/2019/07/a-proactive-approach-to-more-secure-code/) [vulnerabilities](https://www.chromium.org/Home/chromium-security/memory-safety/) across the entire computer software industry, including in Microsoft software. 

Problems like these have become so well-known and serious that our very own (and aforementioned) Mark Russinovich has [tweeted](https://twitter.com/markrussinovich/status/1571995117233504257?lang=en) his support for deprecating use of C/C++ in new projects in favor of Rust. He’s not alone; many large companies and well-known projects are [investing in Rust](https://rustmagazine.org/issue-1/2022-review-the-adoption-of-rust-in-business/) because it’s simultaneously a powerful systems programming language and a memory safe one. The way it achieves this memory safety is somewhat novel and very interesting but I don’t have space today to detail it. Anyway, with this powerful combination of features, Rust fills a niche not filled by nearly any other languages, including the mighty C and C++. 

There’s much more to say about Rust and Hypervisors. We’ll discuss both at length in part 2 and beyond, but we have enough to set a foundation for part 1. Onward! 

## WebAssembly is executed by a virtual machine 

When we describe Hyperlight, we often need to differentiate between two different kinds of virtual machines (VMs). While we’ve described hypervisors above, we also need to consider the virtual machines that execute Wasm (called “Wasm VMs” hereafter)! Fortunately, there’s a very clear line between the two, and it’s all about how hardware is represented and utilized. 

As mentioned in the previous section, hypervisor-based VMs provide a set of virtual hardware devices, but current Wasm VMs do not. Importantly, the latter provide something that looks like a virtual “CPU”, but it only executes the web assembly instruction format and provides no support for any hardware devices. There are APIs built atop Wasm, like WASI, that abstract over hardware devices (like doing network calls and file system accesses), but it is up to the user to implement these APIs in terms of any hardware they see fit. 

In other words, Wasm VMs don’t provide hardware virtualization, while hypervisor-based VMs do. The latter generally provides additional flexibility and a different kind of security isolation boundary. 

These differences also imply important differences in how these platforms execute code. Many (but not all) programs can be compiled to Wasm, and any system with a Wasm VM can then execute it, regardless of the underlying hardware architecture. If you’re familiar with .Net or the Java Virtual Machine (or JVM), you can think of Wasm similarly to how you think of CLR or JVM bytecode. Alternatively, I can compile any program directly to native code and run it directly on the hardware (or virtualized hardware!)  

## A VM inside a VM? 

I just spent a whole section describing the difference between the Wasm VM and a Hypervisor-based VM, and you may be surprised to learn that Hyperlight runs a Wasm VM inside a hypervisor-based VM. Let’s discuss here why we use both. 

If you recall Mark’s key comment in his Azure Innovations keynote, he said that while, Wasm provides an isolation scheme, it lacks the isolation we need to run a secure public cloud. This statement has several important points in it, and I want to elaborate on them. 

## The “kind of isolation” Wasm does have 

We already discussed how Wasm and hypervisor-level isolation differ, so let’s discuss how those differences apply to the needs of a public cloud like Azure. You may not know that Microsoft’s online infrastructure is one of the most attacked in the world. This kind of malicious activity isn’t new to us, so we’ve been investing heavily in the security of our software for a long time. One of our most powerful tools to keep our systems secure, and maintain our customers’ trust, is [defense-in-depth](https://azure.microsoft.com/en-us/blog/microsoft-azures-defense-in-depth-approach-to-cloud-vulnerabilities/).

If you haven’t heard of this defense-in-depth approach, it’s simply a way to ensure sensitive data and systems sit behind as many security boundaries as possible, while still ensuring suitable performance, efficiency, and more. We’re big believers in this approach at Microsoft, and we employ it wherever possible to achieve at least the following two major objectives: 

1. To present an attack surface for any would-be attackers that is ever more difficult to breach
2. To limit the scope (also called the “blast radius”) of any potential breach if one were to happen

While Wasm does provide a [provably safe virtualization technology](https://www.usenix.org/system/files/sec22-bosamiya.pdf), it’s only “safe” if the engine we use to execute Wasm is always bug-free (it’s not). Since we usually can’t guarantee any software is bug-free (this is why we work to improve our software, software construction methods, and more!), we don’t want to rely only on a Wasm VM to provide the security boundaries we need, so we add another layer of a different type. This new layer is a hypervisor-based VM.  

Now, we have two VMs, each possibly with bugs in them, with one running inside of the other. This arrangement ensures that security bugs in one VM have a lower impact because of the additional security boundary provided by the other VM layer. In other words, the impact of a security bug in one VM is reduced because of the additional security boundary provided by the other VM layer. 

## How do we do all this? 

Well, as I mentioned at the beginning of this post, I’m going to save most of the “how” for part 2, but I want to wrap this post up with a preview. 

Running software inside of a Wasm VM is relatively easy these days – there are several high quality Wasm VMs available (Microsoft uses both [Wasmtime](https://wasmtime.dev) and [WAMR](https://github.com/bytecodealliance/wasm-micro-runtime), and Visual Studio Code uses [Node](https://nodejs.org/en), backed by the [V8 engine](https://v8.dev)), and most have very good APIs we can use to load and run a Wasm module. 

Our challenge is twofold, as follows: 

1. Bundle a Wasm VM and associated executable customer code (called a module), into a format that can run inside a hypervisor-based VM we build, and...
2. Build the aforementioned hypervisor-based VM that can load the Wasm VM + customer module, run a specified function, get the return value, and tear everything down, all with latency measured in microseconds (not milliseconds) 

As you’ll see in part 2, the first task is much simpler than the second from the above list, and we’ve spent a lot of time ensuring both low latencies and high “density”. One strategy we use to achieve both things, and reduce the attack surface even further, is to eliminate the operating system kernel entirely from the native code running inside the VM. If you’re familiar with [Unikernels](http://unikernel.org), you’ll notice some similarities. 

I’m looking forward to continuing this discussion. I hope to see in part 2! 
