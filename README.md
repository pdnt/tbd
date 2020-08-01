## TBD

Linter for common mistakes in writing, work in progress.

Usage `tbd <file|glob>` (for now). Examples:

```bash
# lint README.md file
$ tbd README.md
# lint all md files in the current directory
$ tbd *.md
```

### TODO

- [ ] Add a `Reporter(fileName, text string, tokens []Token, w *io.Writer)` method, that takes a string of tokens and the text and pretty prints the result
  - [x] Add a `reporters` package, and add a default reporter `console.go` (use any name you want)
  - [x] Add ability to recognize repeated words with simbols between them
  - [ ] *LATER*: Call `Fprint(w)` instead of `fmt.Print()`
- [ ] Add missing linters
  - [ ] linter.Weasels() []Token
  - [ ] linter.Passive() []Token
- [ ] Define CLI API
  - [ ] Option to choose what parsers to use (runs all by default)
  - [ ] Option to ignore whitespace repetition (example: multiple space bars)
- [ ] Add a JSON reporter

```
Duplicated word in file **filePath* [Line:x , Row:y]:

ccccc *ccccc* ccccc ccccc ccccc
      ^

Duplicated word in file **filePath* [Line:x , Row:y]:

ccccc ccccc *ccccc* ccccc ccccc
             ^
```
