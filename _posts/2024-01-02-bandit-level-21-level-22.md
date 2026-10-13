---
layout: single
title: "Bandit Level 21 Level 22"
date: 2024-01-02 00:00:43 -4300
categories: ctf, writeup, overthewire, bandit
---

![Bandit Level 21 Level 22]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-21-level-22-00.png)
------------
## Introduction

Welcome back, fearless hacker, to the Bandit challenges! In this level, we'll learn to exploit cron jobs and bashscript files.

## Previous Flag

```
NvEJF7oVjkddltPSrdKEFOllh9V1IBcq
```

### Exploring Cron Jobs

Our path to the next flag begins with exploring the cron jobs on the system. Let's list the contents of the `/etc/cron.d/` directory:

```bash
ls -la /etc/cron.d/
```

This reveals the existence of a cron job named `cronjob_bandit22`.

### Analyzing Cron Job Configuration

Let's examine the configuration of the `cronjob_bandit22`:

```bash
cat /etc/cron.d/cronjob_bandit22
```

The output indicates that there's a scheduled job running every minute as `bandit22`:

```
* * * * bandit22 /usr/bin/cronjob_bandit22.sh &> /dev/null
```

### Understanding the Script

```bash
cat /usr/bin/cronjob_bandit22.sh

```

The script does two things: it changes the permissions of a file in `/tmp/` and then copies the password for Bandit level 22 into that file.

```bash
#!/bin/bash
chmod 644 /tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv
cat /etc/bandit_pass/bandit22 > /tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv
```

### Retrieving the Flag

Now, let's check the contents of the file in `/tmp/`:

```bash
ls -la /tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv
cat /tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv
```

This should unveil the password for Bandit level 22:

Flag:

```
WdDozAdTM2z9DiFEQ2mGlwngMfj4EZff
```
