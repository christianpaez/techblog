---
layout: single
title: "Bandit Level 22 Level 23"
date: 2024-01-11 00:41:42 -4200
categories: 
---

![Bandit Level 22 Level 23]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-22-level-23-00.png)
------------
## Introduction

In this level, we'll navigate through the uses of cron jobs and bash scripting to find hidden files on a linux system.

### Previous Flag

```
WdDozAdTM2z9DiFEQ2mGlwngMfj4EZff
```

### Exploring Cron Jobs

Our first clue leads us to investigate the cron jobs for the bandit23 user. Execute the following command to reveal the contents of the cron file:

```bash
cat /etc/cron.d/cronjob_bandit23
```

This exposes the details of a scheduled task that executes a script at specified intervals.

### Decoding the Script

Now, let's inspect the contents of the script executed by the cron job:

```bash
cat /usr/bin/cronjob_bandit23.sh
```

The script is unveiled:

```bash
#!/bin/bash

myname=$(whoami)
mytarget=$(echo I am user $myname | md5sum | cut -d ' ' -f 1)

echo "Copying passwordfile /etc/bandit_pass/$myname to /tmp/$mytarget"

cat /etc/bandit_pass/$myname > /tmp/$mytarget

```

Breaking it down:

- The script identifies the current user and generates an MD5 hash using a specific string.
- It then echoes the process of copying the password file from `/etc/bandit_pass/$myname` to `/tmp/$mytarget`.

### Calculating the Hash

To unveil the target file and extract the password, we need to calculate the MD5 hash of "I am user bandit23":

```bash
echo I am user bandit23 | md5sum | cut -d ' ' -f 1

```

This command yields the hash we seek.

### Reaping the Rewards

Now, let's retrieve the password from the calculated hash:

```bash
cat /tmp/8ca319486bfbbc3663ea0fbe81326349
```

And there it is, the flag to the next level:

```
QYw0Y2aiA672PsMmh9puTQuhoz8SyR2G
```
