---
layout: single
title: "Bandit Level 18 Level 19"
date: 2023-12-17 03:42:09 -0900
categories: ctf, writeup, overthewire, bandit
---
![Bandit Level 18 Level 19]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-18-level-19-00.png)
------------
 In this level, we'll learn about SSH connections, command execution, and file retrieval.

Before we dive into this level, let's review the flag we obtained in the previous challenge:

`hga5tuuCLF6fFzUpnagiMN8ssu9LFrdg`

### Step 1: SSH into the Server

Our first step is to SSH into the server. We'll use the following command to do so:

```
ssh bandit18@bandit.labs.overthewire.org -p 2220 -t '/bin/sh'
```

This command connects us to the server, and we execute a shell to interact with it.

### Step 2: Retrieve the Readme File

Once we've accessed the server, we need to retrieve the "readme" file. You can do this by executing:

```
cat readme
```

This command will display the content of the "readme" file and the flag for the next level.

Flag: `awhqfNnAbc1naukrpqDYcF95h7HoMTrC`
