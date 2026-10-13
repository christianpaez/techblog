---
layout: single
title: "Bandit Level 19 Level 20"
date: 2023-12-17 19:39:57 -5700
categories: ctf, writeup, overthewire, bandit
---

![Bandit Level 19 Level 20]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-19-level-20-00.png)
------------
## Introduction

Greetings, Welcome back to the OverTheWire Bandit challenges. In this level, we are presented with a custom binary that has to be exploited to get secret information.

## Previous Flag

In the previous challenge, we successfully acquired the following flag:

Flag:

```
awhqfNnAbc1naukrpqDYcF95h7HoMTrC
```

### Exploring bandit20-do

Our next task involves the execution of the `./bandit20-do` command. Let's initiate it and see what it does:

```bash
./bandit20-do
```

It seems like this binary executes whatever bash command we provide as the user we want to escalate to, bandit20.

### Retrieving Bandit20 Password

Now that we have access to `bandit20-do`, let's utilize it to read the contents of the `/etc/bandit_pass/bandit20` file:

```bash
./bandit20-do cat /etc/bandit_pass/bandit20
```

Executing this command should reveal the password for the Bandit level 20. 

Flag:

```
VxCazJaVykI6W36BkBU0mJTCM8rR95XT
```
