---
layout: single
title: "Bandit Level 12 Level 13"
date: 2023-10-21 18:18:42 -4200
categories: bandit, ctf, overthewire, writeup
---

![Bandit Level 12 Level 13]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-12-level-13-00.png)
------------
### Introduction

Greetings, As we step into **Level 12** of the Bandit wargame, we're faced with a unique challenge involving data manipulation and repeated compression. Get ready to roll up your sleeves and dive into the intricate world of command-line operations.

### **Steps**

Before we embark on our journey, make sure you've successfully acquired the password for the previous level: `JVNBBFSmZwKKOP0XbFXOoW8chDz5yVRv`.

1. **Creating a Workspace**: Begin by creating a temporary directory named `/tmp/hexdump`. Navigate to this new directory using the following commands:
    
    ```bash
    mkdir /tmp/hexdump
    cd /tmp/hexdump
    ```
    
    This step ensures you have a clean workspace to carry out the tasks.
    
2. **Preparing the Data**: Now, let's grab the `data.txt` file from the previous level and place it into your temporary directory:
    
    ```bash
    cp /home/bandit12/data.txt .
    ```
    
    This sets the stage for our data manipulation journey.
    
3. **Decoding Hexdump**: We'll start by converting the hexdump back to its original form. Execute the following commands to perform this transformation:
    
    ```bash
    xxd -r data.txt > hex
    file hex
    ```
    
    The `xxd -r` command reverses the hexdump, reconstructing the original data. The `file` command is used to inspect the type of data in the `hex` file.
    
4. **Compressing and Decompressing**: In the next steps, we'll play with compression. These commands help us understand the data and its transformations:
    
    ```bash
    mv hex hex.gz
    gunzip hex.gz
    file hex
    ```
    
    Here, we rename `hex` to `hex.gz`, then unzip it using `gunzip`. The `file` command confirms the type of data.
    
5. **Further Compression**: Our exploration continues with more compression:
    
    ```bash
    mv hex hex.bz2
    bunzip2 hex.bz2
    file hex
    ```
    
    We rename `hex` to `hex.bz2` and use `bunzip2` to decompress it. Once again, the `file` command provides insight into the data's nature.
    
6. **Final Steps**: The puzzle nears its completion. Execute these commands to perform more decompression given the working file’s data type:
    
    ```bash
    mv hex hex.gz
    gunzip hex.gz
    file hex
    tar -xf hex
    file data5.bin
    tar -xf data5.bin
    file data6.bin
    mv data6.bin data6.bin.bz2
    bunzip2 data6.bin.bz2
    file data6.bin
    tar -xf data6.bin
    file data8.bin
    mv data8.bin data8.bin.gz
    cat data8.bin
    
    ```
    
    These commands navigate through decompression, untar, and finally, reveal the contents of `data8.bin` which is the password for the next level:
    

```
wbWdlBxEir4CaE8LaPhauuOo6pwRmrDw
```
