---
layout: single
title: "Bandit Level 8 Level 9"
date: 2023-08-01 16:31:28 -2800
categories: bandit, ctf, overthewire, writeup
---

![Bandit Level 8 Level 9]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-8-level-9-00.png)
------------
### Introduction

Bandit9 is the tenth level of the OverTheWire Bandit wargame. In this level, we are given a file called `data.txt` and we have to find the line that occurs only once in the file.

### Steps

1. Open your terminal application.
2. Enter the following command to ssh into the remote server:

```
ssh bandit8@bandit.labs.overthewire.org -p 2220
```

1. Enter the password for Bandit8: `z7WtoNQU2XfjmMtWA8u5rN4vzqu4v99S`
2. Now, you are logged in as Bandit8 and can start solving the level.
3. Use the `cat` command to read the contents of the `data.txt` file:

```
cat data.txt
```

1. We need to find the line that occurs only once in the file. To do this, we can use the `sort` command to sort the lines of the file, and then the `uniq -u` command to show only the lines that occur once:

```
cat data.txt | sort | uniq -u
```

Congratulations! You have successfully solved Bandit9 and obtained the password for the next level `EN632PlfYiZbn3PhVK3XOGSlNInNE00t`.

### Conclusion

In this level, we learned how to use the `sort` and `uniq` commands to find a unique line in a file. These are powerful commands that can be used to analyze and manipulate large amounts of data.
