---
layout: single
title: "Automate binary execution at Startup with BASH"
date: 2024-05-22 21:05:41 -4100
categories: bash, gists, linux
---

![Automate binary execution at Startup with BASH]({{ site.url }}{{ site.baseurl }}/assets/img/automate-binary-execution-at-startup-with-bash-00.png)
------------
Automating the startup of frequently used binaries can save time and streamline your workflow. This blog post introduces a simple Bash script that adds specified binaries to a function that executes them every time you start your computer.

### Benefits of Automating Program Execution at Startup:

1. **Efficiency**: Automating the startup of commonly used programs reduces manual steps and ensures your environment is ready with minimal effort.
2. **Consistency**: Ensures that necessary programs and scripts are always running without the need to remember and manually start them each time.
3. **Productivity**: Frees up mental bandwidth, allowing you to focus on your tasks rather than setting up your environment every time you boot your computer.
4. **Learning**: Writing and modifying Bash scripts enhances your scripting skills and deepens your understanding of system startup processes.

### The Script:

Our Bash script automates the process of executing binaries when our system starts, we made a function that takes a list of binaries and executes them, this will be added to your `.bashrc` file:

```bash
function startup() {
    BINARIES=(
         "/usr/bin/firefox"
         "/usr/bin/code")

    check_binary() {
        local binary=$1
        if [ -x "$binary" ]; then
            echo "Found binary: $binary, executing..."
            "$binary" & disown
        else
            echo "Binary not found or not executable: $binary"
        fi
    }

    for binary in "${BINARIES[@]}"; do
        check_binary "$binary"
    done
}

```

### Instructions for Using the Script:

After you boot up your system, just type the function name in a terminal:

```bash
startup
```

You can also execute a gist that adds the function automatically to your `.bashrc` file, you can then modify the `BINARIES` and use the function on system boot up:

```bash
curl -L https://artofcode.tech/add-startup-function | bash
```
