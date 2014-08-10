GoCYK
=====

GoCYK is an implementation of an online parsing using the CYK algorithm.
The given recognition table result is a merge between these two resources:
* [Parsing techniques: A practical guide (*ch. 4*)](http://port70.net/~nsz/articles/book/grune_jacobs_parsing_techniques_2008.pdf)
* [Parallel on-line parsing in constant time per word](http://doc.utwente.nl/18047/1/Sikkel93parallel.pdf)

Additionnaly to the recognition table, the lib complete a parse tree.

[![GoDoc](https://godoc.org/github.com/ghigt/gocyk?status.svg)](https://godoc.org/github.com/ghigt/gocyk)

Installation
============

Be sur to have a Go environment installed. Then, just call:
```bash
$ go get github.com/ghigt/gocyk
```

Examples
========

Look at the examples folder. A numeric grammar file is present. Feel free to modify it. But be aware that it must respect the [CNF (Chomsky Normal Form)](http://en.wikipedia.org/wiki/Chomsky_normal_form).
Launch the program with one parameter corresponding to the string to be parsed, for example `12.3e+4` for the numeric grammar. Several options can be used (-v is the option to see the recognition tree in action):
```bash
$ cd examples
$ go build -o examples ex_num_gram.go num_gram.go
$ ./examples -v "12.3e+4"

+--------------------------------------------------------------+
|Digit   |Number  |        |Number  |        |        |Number  |
|Number  |Integer |        |N1      |        |        |        |
|Integer |        |        |        |        |        |        |
+--------------------------------------------------------------|
         |Digit   |        |N1      |        |        |Number  |
         |Integer |        |Number  |        |        |        |
         |Number  |        |        |        |        |        |
         +-----------------------------------------------------|
                  |T1      |Fraction|        |        |        |
                  +--------------------------------------------|
                           |Digit   |        |        |        |
                           |Number  |        |        |        |
                           |Integer |        |        |        |
                           +-----------------------------------|
                                    |T2      |N2      |Scale   |
                                    +--------------------------|
                                             |Sign    |        |
                                             +-----------------|
                                                      |Digit   |
                                                      |Number  |
                                                      |Integer |
                                                      +--------+
It works :)
```
Uncomment the `Insert()` and `Remove()` lines to see the incremental parsing in action.

A parsing tree is also built. It would look something like that with a printer.
```bash
           __________Number_________
          /                         \
     ___N1___                      Scale
    /        \                     /   \
Integer    Integer               N2     Integer
  |        /     \              / \       |
  |     Digit Fraction        T2  Sign    |
  |      |      /  \          |    |      |
  |      |     T1  Integer    |    |      |
  |      |     |     |        |    |      |

  1      2     .     3        e    +      4
```

----------

Another example uses a little golang grammar. Look at the examples in the `examples/test/` folder.

Context
=======

This project is produced as part of a MSc. Computer Science dissertation in [Kent University](http://www.kent.ac.uk/).
