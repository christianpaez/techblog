---
layout: single
title: "Bandit Level 6 Level 7"
date: 2023-07-26 20:45:23 -2300
categories: bandit, ctf, overthewire, writeup
---

![Bandit Level 6 Level 7]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-6-level-7-00.png)
------------
### Introduction

Bandit 7 is the eigth level of the OverTheWire Bandit wargame. In this level, we will learn how to search for files by owner and size.

### Steps

1. Open your terminal application.
2. Enter the following command to ssh into the remote server:

```
ssh bandit6@bandit.labs.overthewire.org -p 2220
```

1. You will be prompted to enter a password. Enter the password from the previous level: `P4L4vucdmLnm8I7Vl7jG1ApGSfjYKqJU`.
2. Once you're logged in, run the following command to search for files that are 33 bytes in size and owned by user `bandit7` and group `bandit6`:

```
find / -size 33c -user bandit7 -group bandit6 2>/dev/null
```

1. This will output the location of the file. Use the `cat` command to read the contents of the file:

```
cat /var/lib/dpkg/info/bandit7.password
```

1. You should now see the password for the next level which is `z7WtoNQU2XfjmMtWA8u5rN4vzqu4v99S`.

Congratulations! You have successfully completed Bandit 7 and obtained the password for the next level.
