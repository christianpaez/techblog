---
layout: single
title: "Bandit Level 11 Level 12"
date: 2023-10-04 19:33:49 -4900
categories: bandit, ctf, overthewire, writeup
---

![Bandit Level 11 Level 12]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-11-level-12-00.png)
------------
### Introduction

In Level 11 of the Bandit wargame, we're faced with a cryptographic challenge that requires some clever manipulation. Let's jump right into the steps to crack it.

### **Steps.**

1. **Access the Remote Server**: Open your terminal and connect to the remote server using this command:
    
    ```bash
    ssh bandit10@bandit.labs.overthewire.org -p 2220
    ```
    
    Input the password from the previous level: `6zPeziLdR2RKNdNYFNb6nVCKzphlXHBM`.
    
2. **Decrypt with `tr`**: Once connected, run this command to decrypt the data:
    
    ```bash
    cat data.txt | tr 'A-Za-z' 'N-ZA-Mn-za-m'
    ```
    
    The `tr` command shifts each letter by 13 positions, revealing the concealed message.

The output of the command contains the password for the next level:

```
JVNBBFSmZwKKOP0XbFXOoW8chDz5yVRv
```
