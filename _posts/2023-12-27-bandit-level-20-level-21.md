---
layout: single
title: "Bandit Level 20 Level 21"
date: 2023-12-27 21:44:51 -5100
categories: ctf, writeup, overthewire, bandit
---

![Bandit Level 20 Level 21]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-20-level-21-00.png)
------------
## Introduction

Welcome back tothe OverTheWire Bandit challenges. In this level, we encounter a scenario involving networking and data transfer.

## Previous Flag

Before we delve into the current challenge, let's input the flag from the previous level:

Flag:

```
VxCazJaVykI6W36BkBU0mJTCM8rR95XT
```

### Netcat

The provided command involves using Netcat (`nc`) to transfer the contents of the `/etc/bandit_pass/bandit20` file over a network connection. Let's break it down:

```bash
cat /etc/bandit_pass/bandit20 | nc -lvp 1234 &
```

This command reads the password from the file and sends it over a network connection. The `nc -lvp 1234` part sets up a Netcat listener on port 1234, waiting to receive the data.

### Connecting with suconnect

Now, let's use the `./suconnect` command to establish a connection and retrieve the password:

```bash
./suconnect 1234
```

This command initiates a connection to the Netcat listener on port 1234, allowing us to receive the contents of the password file.

### Flag Unveiled

Upon successful execution, the password for Bandit level 21 should be displayed, providing us with the key to the next stage:

Flag:

```
NvEJF7oVjkddltPSrdKEFOllh9V1IBcq
```
