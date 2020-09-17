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

- **Weasels** (`-weasels`, `-w`), finds [Weasel Words](https://en.wikipedia.org/wiki/Weasel_word)
- **Duplicates** (`-duplicates`, `-d`), finds unintentional word repetitions.
- **Passive voice** (`-passive`, `-p`), finds usage of [Passive Voice](https://en.wikipedia.org/wiki/Passive_voice)

### Special characters

You can also include special characters in the linting with:

- **Withespace** via the `--include-whitespace` flag
- **Punctuation** via the `--include-punctuation` flag

### Install

Download the appropriate version for your platform from [releases](https://github.com/pdnt/tbd/releases/) Once downloaded, the binary can be run from anywhere.

Ideally, you should install it somewhere in your `PATH` for easy use. `/usr/local/bin` is the most probable location.
