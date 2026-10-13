---
layout: single
title: "Bandit Level 4 Level 5"
date: 2023-07-23 03:02:52 -5200
categories: writeup, ctf, overthewire, bandit
---

![Bandit Level 4 Level 5]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-4-level-5-00.png)
------------
### Introduction

Bandit5 is the sixth level of the OverTheWire Bandit wargame. In this level, we will learn how to use the "file" command to determine the file type of a file and how to read the contents of a specific file. By completing this level, we will gain access to the password for the next level.

### Steps

1. Open your terminal application.
2. Enter the following command to ssh into the remote server:

```
ssh bandit4@bandit.labs.overthewire.org -p 2220
```

1. You will be prompted to enter a password. Enter the password from the previous level, "2EW7BBsr6aMMoJ2HjW067dm8EgX26xNe", and hit enter.
2. You are now connected to the remote server and are able to execute commands.
3. Enter the following command to view the file types of all files in the "inhere" directory:

```
file inhere/*
```

1. The output will show that all files in the directory are “data” files, except for one file which is a ASCII text file, this human-readable file should contain the password for the next level.
2. Enter the following command to read the contents of the file with the name "-file07" in the "inhere" directory:

```
cat inhere/-file07
```

1. The output should show the password for the next level: lrIWWI6bB37kxfiCQZqUdOIYfr6eEeqR.

Congratulations! You have successfully completed Bandit5 and gained access to the password for the next level.

### Conclusion

In this level, we learned how to use the "file" command to determine the file type of a file and how to read the contents of a specific file. These are important skills to have when working with files in Linux command line, and mastering them will be useful for solving future levels in the OverTheWire Bandit wargame.
