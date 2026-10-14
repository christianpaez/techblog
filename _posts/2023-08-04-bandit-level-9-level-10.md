---
layout: single
title: "Bandit Level 9 Level 10"
date: 2023-08-04 03:26:40 -4000
categories: bandit, ctf, overthewire, writeup
---

![Bandit Level 9 Level 10]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-9-level-10-00.png)
------------
Bandit10 is the eleventh level in the OverTheWire Bandit wargame. In this level, we are given a file named "data.txt" and we are required to find a string of text that occurs only once in the file and contains only letters and spaces.

### Steps

1. Open your terminal application.
2. Enter the following command to ssh into the remote server using the credentials for Bandit9:
    
    ```
    ssh bandit9@bandit.labs.overthewire.org -p 2220
    ```
    
3. Enter the password for Bandit9 `EN632PlfYiZbn3PhVK3XOGSlNInNE00t` when prompted.
4. Once you are logged in, enter the following command to read the contents of the "data.txt" file:
    
    ```
    cat data.txt
    ```
    
5. You will see that the file contains a lot of random text.
6. Next, enter the following command to extract all the printable strings from the file:
    
    ```
    cat data.txt | strings
    ```
    
7. You will see that the output includes some text surrounded by multiple equal signs.
8. To filter out only the strings that contain letters and spaces and are surrounded by equal signs, enter the following command:
    
    ```
    cat data.txt | strings | grep '^=\\{2,\\}' | awk -F " " '{print $2}' | tr '\\n' ' ' | xargs echo
    ```
    
Let's break down this command:
    
- `cat data.txt`: display the contents of the file data.txt.
- `strings`: extract human-readable strings from the binary file data.txt.
- `grep '^=\{2,\}'`: search for lines that start with two or more equal signs (==).
- `awk -F " " '{print $2}'`: print the second field of each matching line, which is the password.
- `tr '\n' ' '`: replace the newline character with a space character.
- `xargs echo`: pass the output of the previous command to echo.
    
 After running the command, the password for the next level `G7w8LIi6J3kTb8A7j9LgrywtEUlyyp6s` will be displayed.

Congratulations! You have successfully completed Bandit10 and found the flag.
