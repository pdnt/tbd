## TBD

TBD allows you to find and fix common mistakes in writing.

Usage `tbd <file|glob>`. Examples:

```bash
# lint README.md file
tbd README.md

# lint all md files in the current directory
tbd *.md
```

`tbd` comes with a built-in set of [linters](#linters), and by default it will run all the available linters. You can modify this behavior by explicitly providing the set of linters you'd like to run via flags, for example:

```bash
# search only for duplicates
tbd -duplicates *.md

# search for passive voice and weasels
tbd -passive -weasels *md
```

### Linters

- **Weasels** (-weasels, `-w`)
- **Duplicates** (`-duplicates`, `-d`)
- **Passive voice** (`-passive`, `-p`)

### Install

Download the appropriate version for your platform from [releases](/releases) Once downloaded, the binary can be run from anywhere.

Ideally, you should install it somewhere in your `PATH` for easy use. `/usr/local/bin` is the most probable location.
