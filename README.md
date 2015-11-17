# lang

`lang` is a small programming language. It is inspired by Go, JavaScript and Python.

## Code Example

```
x int = 22

if x == 33 {
    print(x)
}

for i in range(5) {
    print(i)
}
```

## Motivation

Over the (Australian) summer holidays I attended a summer camp called [NCSS](http://ncss.edu.au), (It was amazing, I highly recommend attending if you have the opportunity). One of the things we learned about was Finite State Automata, over a day or so we each created our own basic regex implementations... I thought this was awesome. Since then, I've tried my hand at writing a lexer + parser for [mathemtic expressions](https://github.com/paked/algebra), and [markdown](http://github.com/paked/down). I was never really content, and always ended up leaving the project with a *large* amount of bugs and *really* bad design.

`lang` essentially encompasses everything I have learned thus far, and embodies it into a programming language which I would like to use.

## For Reading

If you're a hacker looking at this and curious as to how it works, I suggest you begin by reading the [lexer.go](lexer.go) file, and then proceeding to [parser.go](parser.go).

## License

This repository is released under an `MIT` license found in the [LICENSE.md](LICENSE.md) file.
