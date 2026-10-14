---
layout: single
title: "Bandit Level 2 Level 3"
date: 2023-07-12 00:47:00 -0000
categories: writeup, overthewire, ctf, bandit
---

![Bandit Level 2 Level 3]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-2-level-3-00.png)
------------
Bandit3 is the fourth level of the OverTheWire Bandit wargame. In this level, we will learn how to handle filenames with spaces in them and how to read the contents of a file with a space in its name. By completing this level, we will gain access to the password for the next level.

### Steps

1. Open your terminal application.
2. Enter the following command to ssh into the remote server:

```
ssh bandit2@bandit.labs.overthewire.org -p 2220
```

1. You will be prompted to enter a password. Enter the password from the previous level, "rRGizSaX8Mk1RTb1CNQoXTcYZWU6lgzi", and hit enter.
2. You are now connected to the remote server and are able to execute commands.
3. Enter the following command to search for files with spaces in their name:

```
find /home -name "spaces in this filename" 2>/dev/null
```

1. The output should show the path to the file with spaces in its name. In this case, the path is /home/bandit2/spaces in this filename.
2. Enter the following command to read the contents of the file:

```
cat "/home/bandit2/spaces in this filename"
```

1. The output should show the password for the next level: aBZ0W5EmUfAf7kHTQeOwd8bauFJ2lAiG.

Congratulations! You have successfully completed Bandit3 and gained access to the password for the next level.

### Conclusion

In this level, we learned how to handle filenames with spaces in them and how to read the contents of a file with a space in its name. These are important skills to have when working with files in Linux command line, and mastering them will be useful for solving future levels in the OverTheWire Bandit wargame.
