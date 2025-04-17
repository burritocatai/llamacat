![Static Badge](https://img.shields.io/badge/mission-quick_second_brain_notation-purple)

![Static Badge](https://img.shields.io/badge/openai-support-green?logo=openai)
![Static Badge](https://img.shields.io/badge/ollama-support_soon-red?logo=ollama)
![Static Badge](https://img.shields.io/badge/anthropic-support-green?logo=anthropic)
![Static Badge](https://img.shields.io/badge/groq-support-green)
![Static Badge](https://img.shields.io/badge/mistral-support-green?logo=mistral)


# `llamacat`

`llamacat` is an application designed to capture the output of LLMs from your input and prompts and seamlessly integrate them into your Second Brain: Obsidian, or into STDOUT for piping into other terminal apps or files. Even better, it's extensible, so others can add further targets to `llamacat`. (This is a process that needs figured out.)

Basic Usage:

```shell
cat interestingfile.txt | llamacat -p default:extract_ideas -m openai:chatgpt-4o-mini -o workvault:Notes
```

Yup, that's it.
## The What and Why
### What
`llamacat` simply takes in content and applies a prompt to it, then it then dumps that in your second brain: Obsidian.

### Why
I've become an enthusiastic user of Obsidian for organizing not only work-related tasks and notes but also for personal journaling, home and tech projects, and tracking books. In my work vault, I track Linux commands and code snippets that I want to remember. However, without careful management, these can turn into a jumbled heap of commands and code with no explanations. Creating an Obsidian note, adding an explanation, and formatting the code or command into Markdown require significant time and effort.

I found inspiration in the remarkable [Fabric CLI application](https://github.com/danielmiessler/fabric) created by Daniel Miessler and his visionary approach to AI.

>AI isn't a thing; it's a _magnifier_ of a thing. And that thing is **human creativity**.
>*Daniel Miessler*

With that in mind, I wanted to create an AI tool that would consume content that I wanted to keep track of and automatically format it and provide quick explanations for me. This would help in future searches as well as head-scratching when trying to remember what exactly that code or command does and why I would save it.

## Installation
You must have Go installed in order to run the following command:

```shell
go install github.com/burritocatai/llamacat@latest
```

Run it again at anytime to upgrade `llamacat`.

`llamacat` is currently in pre-release, so please use it at your own risk. Remember to check for updates frequently!

## Configure
`llamacat` will soon have an easy to configure interface.

```shell
llamacat config
```


## pbpaste

> The below instructions are taken from Fabric's Readme as it highlights how to get pbpaste on non-macOS better than I could have written it

`pbpaste` is not available on Windows or Linux, but there are alternatives.

On Windows, you can use the PowerShell command `Get-Clipboard` from a PowerShell command prompt. If you like, you can also alias it to `pbpaste`. If you are using classic PowerShell, edit the file `~\Documents\WindowsPowerShell\.profile.ps1`, or if you are using PowerShell Core, edit `~\Documents\PowerShell\.profile.ps1` and add the alias,

```powershell
Set-Alias pbpaste Get-Clipboard
```

On Linux, you can use `xclip -selection clipboard -o` to paste from the clipboard. You will likely need to install `xclip` with your package manager. For Debian based systems including Ubuntu,

```shell
sudo apt update
sudo apt install xclip -y
```

You can also create an alias by editing `~/.bashrc` or `~/.zshrc` and adding the alias,

```shell
alias pbpaste='xclip -selection clipboard -o'
```

