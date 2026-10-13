---
layout: single
title: "Bandit Level 1 Level 2"
date: 2023-07-03 15:32:55 -5500
categories: writeup, overthewire, bandit, ctf
---

![Bandit Level 1 Level 2]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-1-level-2-00.png)
------------
### Introduction

Bandit2 is the third level of the OverTheWire Bandit wargame. In this level, we will learn how to handle special characters in filenames and how to read the contents of a file that has a hyphen in its name. By completing this level, we will gain access to the password for the next level.

### Steps

1. Open your terminal application.
2. Enter the following command to ssh into the remote server:

```
ssh bandit1@bandit.labs.overthewire.org -p 2220

```

1. You will be prompted to enter a password. Enter the password from the previous level, "NH2SXQwcBdpmTEzi3bvBHMM9H66vVXjL", and hit enter.
2. You are now connected to the remote server and are able to execute commands.
3. Enter the following command to search for files with a hyphen in their name:

```
find /home -name "-" 2>/dev/null

```

1. The output should show the path to the file with a hyphen in its name. In this case, the path is /home/bandit1/-.
2. Enter the following command to read the contents of the file:

```
cat /home/bandit1/-

```

1. The output should show the password for the next level: rRGizSaX8Mk1RTb1CNQoXTcYZWU6lgzi.

Congratulations! You have successfully completed Bandit2 and gained access to the password for the next level.

### Conclusion

In this level, we learned how to handle special characters in filenames and how to read the contents of a file that has a hyphen in its name. These are important skills to have when working with files in Linux command line, and mastering them will be useful for solving future levels in the OverTheWire Bandit wargame.
