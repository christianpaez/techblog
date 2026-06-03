---
title:  "Golang for Rubyists, the need for a second language"
date:   2026-06-02 00:00:00 -0500
categories: web
layout: single
classes:
  - landing
  - dark-theme
---
![Link text]({{ site.url }}{{ site.baseurl }}/assets/img/gopher.png)
------------
Like many software developers, I also have a language that I enjoy using no matter what, and it never seems to be a puzzle to write, debug or read it. In my case, it’s Ruby. It nicely solves most of the things I need to do at work (or outside of it), which is nice, so, why did I want to try something else?

# There are no golden hammers

No matter what language or tool you use, it is not designed to solve all problems (even though you probably can) and you might be surprised at how easy it is to do something in a simpler way just by changing your tool.

# But there are a lot of hammers

Each language was probably designed with something in mind, and if you have tried a couple you probably noticed they are very easy and comfortable to use as long as you find the right problem for them. If you have used enough software, you probably have more than one kind of problem and the need for a second language (even if it’s a distant second) becomes real.

So, from a Ruby perspective, why Golang?

# A web focused hammer

Golang was designed, among other things, to be used on the web. It is very natural or almost automatic for most Ruby programmers to have worked on the web to some extent. A second good use case for both Ruby and Golang is scripting and batch/CLI applications.

# Simplicity

Ruby is very simple and elegant in terms of expression. The main focus of the language is to remove friction in terms of communication between humans. When working with Golang, the same seems to happen but also with the added benefit of simplicity in terms of computing. Despite Go having some guardrails like Garbage Collection and some type inference, it does not feel like the language is hiding a lot from you in terms of what a computer needs to know to do some work.

# Speed aka ‘DX’

Ruby on Rails is an awesome tool for making servers. In a couple of minutes, you can leverage RoR’s CLI and get a working server (deployment included). When I first tried Golang, I used Termux, which is a very limited Linux emulator, and I was able to set up the language, build tools, formatter and lsp in minutes. Even though I did not really write anything useful or impressive at the time, getting the tool working was really intuitive and fast.

# Little or Zero dependencies

Working with dependencies and package managers is hell. I began working for SPA’s and as soon as I knew I needed to do `yarn install, I knew I need to get some coffee. However, even outside of Rails, Ruby’s stdlib is very complete and has everything you need to build stuff for the web or CLIs. Maybe you need rack/puma if you need a proper server with a ProcessPool, but Ractors are looking to improve this area. Golang’s stdlib feels a lot better nonetheless. I have written a couple of projects and I do not think I have used more than 2 or 3 dependencies. It really has anything you need.

> *"Learning of a foreign language and culture is important. It helps you think and be a better person. With programming languages it helps you become a better programmer"*. 
— Bjarne Stroustrup
>
