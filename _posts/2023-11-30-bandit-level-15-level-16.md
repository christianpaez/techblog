---
layout: single
title: "Bandit Level 15 Level 16"
date: 2023-11-30 17:51:37 -3700
categories: ctf, writeup, overthewire, bandit
---

![Bandit Level 15 Level 16]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-15-level-16-00.png)
------------
## Introduction

Welcome to another level of the OverTheWire Bandit challenges! In this level, we will continue our journey to enhance our Linux and ethical hacking skills.

## Previous Flag

In the previous challenge, we successfully retrieved the flag:

Flag: `jN2kgmIXJ6fShzhT2avhotn4Zcka6tnt`

### Scanning for Open Ports

We will use the`nmap` command to scan for open ports on the local machine.

```
nmap -p 30001 localhost
```

The output should indicate that port 30001 is open and that there is an openSSL service running on it.

### Establishing an SSL Connection

With port 30001 identified as an open port, let's establish an SSL connection using the `openssl` command:

```
openssl s_client localhost:30001
```

This command initiates an SSL/TLS session and displays the flag for the next level.

Flag: `JQttfApK4SeyHwDlI9SXGR50qclOAil1`
