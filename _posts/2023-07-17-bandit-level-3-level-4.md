---
layout: single
title: "Bandit Level 3 Level 4"
date: 2023-07-17 03:13:08 -0800
categories: writeup, ctf, overthewire, bandit
---

![Bandit Level 3 Level 4]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-3-level-4-00.png)
------------
### Introduction

Bandit4 is the fifth level of the OverTheWire Bandit wargame. In this level, we will learn how to search for files that start with a dot and how to read the contents of a hidden file. By completing this level, we will gain access to the password for the next level.

### Steps

1. Open your terminal application.
2. Enter the following command to ssh into the remote server:

```
ssh bandit3@bandit.labs.overthewire.org -p 2220
```

1. You will be prompted to enter a password. Enter the password from the previous level, "aBZ0W5EmUfAf7kHTQeOwd8bauFJ2lAiG", and hit enter.
2. You are now connected to the remote server and are able to execute commands.
3. Enter the following command to search for files that start with a dot in the "inhere" directory:

```
find /home/bandit3/inhere/ -name ".*"
```

1. The output should show a file path that ends with ".hidden". In this case, the path is /home/bandit3/inhere/.hidden.
2. Enter the following command to read the contents of the hidden file:

```
cat /home/bandit3/inhere/.hidden
```

1. The output should show the password for the next level: 2EW7BBsr6aMMoJ2HjW067dm8EgX26xNe.

Congratulations! You have successfully completed Bandit4 and gained access to the password for the next level.

### Conclusion

In this level, we learned how to search for files that start with a dot and how to read the contents of a hidden file. These are important skills to have when working with files in Linux command line, and mastering them will be useful for solving future levels in the OverTheWire Bandit wargame.
