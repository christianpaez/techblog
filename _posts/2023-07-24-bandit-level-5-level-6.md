---
layout: single
title: "Bandit Level 5 Level 6"
date: 2023-07-24 18:03:25 -2500
categories: writeup, ctf, overthewire, bandit
---

![Bandit Level 5 Level 6]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-5-level-6-00.png)
------------
Bandit 6 is the seventh level of the OverTheWire Bandit wargame, we are tasked with finding a file within a directory hierarchy and extracting a password from that file.

### Steps

1. Connect to the remote server using SSH by running the following command in your terminal:
    
    ```
    ssh bandit5@bandit.labs.overthewire.org -p 2220
    ```
    
    You will be prompted to enter the password for Bandit level 6, which is `lrIWWI6bB37kxfiCQZqUdOIYfr6eEeqR`. Enter the password and hit enter.
    
2. Once you have successfully logged in, navigate to the `inhere` directory by running the following command:
    
    ```
    cd inhere/
    ```
    
    This directory contains a number of subdirectories and files, and our task is to find the file containing the password.
    
3. To find the file, we can use the `find` command with a combination of flags to search for files that meet certain criteria. In this case, we want to find files that are not executable and have a size of exactly 1033 bytes. We can use the following command:
    
    ```
    find . -type f -not -executable -size 1033c
    ```
    
    This will return a list of files that meet the criteria.
    
4. Next, we want to check the contents of each file in the list to see if it contains the password. We can use the `file` command to check if each file is an ASCII text file, and then use `cat` to print the contents of any file that matches the criteria. We can use the following command:
    
    ```
    find . -type f -not -executable -size 1033c | xargs file | grep "ASCII text" | awk -F: '{print $1}' | xargs cat
    ```
    
    This will print the contents of any file that is an ASCII text file and has a size of 1033 bytes and is not executable.
    
5. The password is located in the file `/home/bandit6/inhere/maybehere07/.file2`. Use the following command to print the contents of the file:
    
    ```
    cat /home/bandit6/inhere/maybehere07/.file2
    ```
    
    This will output the password, which is `P4L4vucdmLnm8I7Vl7jG1ApGSfjYKqJU`.
    

Congratulations! You have successfully completed Bandit level 6.
