---
layout: single
title: "Bandit Level 7 Level 8"
date: 2024-02-23 15:36:54 -5400
categories: bandit, ctf, overthewire, writeup
---

![Bandit Level 7 Level 8]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-7-level-8-00.png)
------------
### Introduction

Bandit level 8 is all about finding a specific word in a file and extracting its value. In this level, we are given a file called `data.txt` and are tasked with finding the value of a specific word.

### Steps

1. Connect to the Bandit server using SSH with the following command:

```
ssh bandit7@bandit.labs.overthewire.org -p 2220
```

1. Enter the password for Bandit level 8 when prompted: `z7WtoNQU2XfjmMtWA8u5rN4vzqu4v99S`.
2. We are given a file called `data.txt`. Use the `cat` command to view the contents of the file:

```
cat data.txt
```

1. We are tasked with finding the value of the word "millionth" in the file. To accomplish this, we can use the `grep` command to search for the word, and then use the `awk` command to extract its value. The command to do this is as follows:

```
cat data.txt | grep millionth | awk -F " " '{print $2}'
```

1. After running the command, the value of the word "millionth" will be displayed in the terminal.

```
TESKZC0XvTetK0S9xNwm25STk5iWrBvP
```

### Conclusion

In this level, we learned how to search for a specific word in a file using the `grep` command and how to extract its value using the `awk` command. With this knowledge, we were able to successfully complete the level and obtain the password for the next level.
