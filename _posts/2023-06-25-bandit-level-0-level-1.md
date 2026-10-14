---
layout: single
title: "Bandit Level 0 Level 1"
date: 2023-06-25 00:36:08 -0800
categories: writeup, overthewire, bandit, ctf
---

![Bandit Level 0 Level 1]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-0-level-1-00.png)
------------
Bandit Level 0 → Level 1 is the second level of the OverTheWire Bandit wargame. In this level, we will learn how to use the "find" command to search for files and how to use the "cat" command to read the contents of a file. By completing this level, we will gain access to the password for the next level.

### Steps

1. Open your terminal application.
2. Enter the following command to ssh into the remote server:

```
ssh bandit0@bandit.labs.overthewire.org -p 2220
```

1. You will be prompted to enter a password. Enter the password from the previous level, "bandit0", and hit enter.
2. You are now connected to the remote server and are able to execute commands.
3. Enter the following command to search for the file "readme" in the /home directory:

```
find /home -name readme 2>/dev/null
```

1. The output should show the path to the "readme" file. In this case, the path is /home/bandit0/readme.
2. Enter the following command to read the contents of the "readme" file:

```
cat /home/bandit0/readme
```

1. The output should show the password for the next level: NH2SXQwcBdpmTEzi3bvBHMM9H66vVXjL.

Congratulations! You have successfully completed Bandit1 and gained access to the password for the next level.
