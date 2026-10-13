---
layout: single
title: "Bandit Level 13 Level 14"
date: 2023-10-26 02:59:16 -1600
categories: ctf, writeup, overthewire, bandit
---

![Bandit Level 13 Level 14]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-13-level-14-00.png)
------------
### Introduction

Greetings, We're back with another step in the Bandit wargame journey. Level 13 introduces us to Secure Shell (SSH) and the concept of privilege escalation. Let's dive right into the details with no extra flair.

### **Starting Point**

Before we proceed, have the password from Level 12 handy: `wbWdlBxEir4CaE8LaPhauuOo6pwRmrDw`.

### **Steps**

1. **SSH Access with a Private Key**: This time, you're using a private SSH key named `sshkey.private`. Let's get in:
    
    ```bash
    ssh -i sshkey.private bandit14@localhost -p 2220
    ```
    
    This command establishes a secure connection with the key and the provided username and port.
    
2. **Retrieve the Flag**: Once you're in, read the flag:
    
    ```bash
    cat /etc/bandit_pass/bandit14
    ```
    

By doing so, you will get the password for the next level:

```
fGrHPx402xGC7U7rXKDaxiWFTOiF0ENq
```
