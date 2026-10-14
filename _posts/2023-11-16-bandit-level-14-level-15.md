---
layout: single
title: "Bandit Level 14 Level 15"
date: 2023-11-16 19:33:59 -5900
categories: ctf, writeup, overthewire, bandit
---
![Bandit Level 14 Level 15]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-14-level-15-00.png)
------------
Greetings, folks! We're back in the Bandit wargame, in this level where we'll deal with Secure Shell (SSH) and privilege escalation.

### Login via SSH

Now, let's set our sights on the previous flag and login as bandit14 with the password, `fGrHPx402xGC7U7rXKDaxiWFTOiF0ENq`.

### **Scanning the Environment**

Our initial step in the journey takes us into the domain of network exploration. By using the Nmap tool, we shall analyze the ports on our local host. Behold the command:

```bash
nmap -p 30000 localhost
```

This command shows a telnet service running on port 30000.

### **Connecting via Telnet**

With the port 30000 revealed, it's time to establish a connection:

```bash
telnet localhost 30000
```

After executing this command, the flag for the next level is printed in the output:

Flag: `jN2kgmIXJ6fShzhT2avhotn4Zcka6tnt`
