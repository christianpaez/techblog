---
layout: single
title: "Bandit Level 10 Level 11"
date: 2023-08-09 18:54:08 -0800
categories: bandit, ctf, overthewire, writeup
---

![Bandit Level 10 Level 11]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-10-level-11-00.png)
------------
Bandit11 is the twelfth level of the OverTheWire Bandit wargame. In this level, we will learn how to decode base64 encoded data to retrieve the password for the next level.

### Steps

1. Open your terminal application.
2. Enter the following command to ssh into the remote server:

```
ssh bandit10@bandit.labs.overthewire.org -p 2220
```

1. You will be prompted to enter a password. Enter the password obtained from the previous level: `G7w8LIi6J3kTb8A7j9LgrywtEUlyyp6s` and hit enter.
2. You are now connected to the remote server and are able to execute commands.
3. To retrieve the password for the next level, enter the following command:

```
cat data.txt | base64 -d
```

1. The output of the command is the password for the next level `6zPeziLdR2RKNdNYFNb6nVCKzphlXHBM`.

Congratulations! You have successfully decoded the base64 encoded data and retrieved the password for Bandit12.
