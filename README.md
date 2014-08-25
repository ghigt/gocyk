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

Be sur to have a [Go environment installed](https://golang.org/doc/install). Then, just call:
```bash
$ go get github.com/ghigt/gocyk
```

Examples
========

Look at the examples folder. A numeric grammar file is present. Feel free to modify it. But be aware that it must respect the [CNF (Chomsky Normal Form)](http://en.wikipedia.org/wiki/Chomsky_normal_form).
Launch the program with one parameter corresponding to the string to be parsed, for example `12.3e+4` for the numeric grammar. Several options can be used (-v is the option to see the recognition tree in action):
```bash
$ cd examples/num_gram
$ go build
$ ./num_gram -v "12.3e+4"

+--------------------------------------------------------------+
|Digit   |Number  |        |Number  |        |        |Number  |
|Number  |Integer |        |N1      |        |        |        |
|Integer |        |        |        |        |        |        |
+--------------------------------------------------------------|
         |Digit   |        |Number  |        |        |Number  |
         |Number  |        |N1      |        |        |        |
         |Integer |        |        |        |        |        |
         +-----------------------------------------------------|
                  |T1      |Fraction|        |        |        |
                  +--------------------------------------------|
                           |Number  |        |        |        |
                           |Integer |        |        |        |
                           |Digit   |        |        |        |
                           +-----------------------------------|
                                    |T2      |N2      |Scale   |
                                    +--------------------------|
                                             |Sign    |        |
                                             +-----------------|
                                                      |Digit   |
                                                      |Number  |
                                                      |Integer |
                                                      +--------+
+--------------------------------------------------------------+
|1       |2       |.       |3       |e       |+       |4       |
+--------------------------------------------------------------+

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

Tests
=====

In order to test the difference between the time spent for the normal/online/incremental parsing algorithm, a bunch of tests has been made in the `tests/` directory.

To test it, all you need is running this command line (`GOMAXPROCS` refers to the number of CPUs on you machine, see [here](http://golang.org/doc/effective_go.html#parallel) to know more):

```bash
$ go build -o tests *.go && GOMAXPROCS=4 ./tests
```

The result might normal look something like this:
```bash
computeBegNotOnline      :    1.109938532s    :    It works :)
computeBegOnline         :    1.10281264s     :    It works :)
computeBegIncremental    :    26.763131ms     :    It works :)
computeMidNotOnline      :    1.679127562s    :    It works :)
computeMidOnline         :    1.039250382s    :    It works :)
computeMidIncremental    :    623.758154ms    :    It works :)
computeEndNotOnline      :    1.112078578s    :    It works :)
computeEndOnline         :    25.341187ms     :    It works :)
computeEndIncremental    :    25.843226ms     :    It works :)
...
```

It uses a home made "go like" grammar in CNF available in `examples/go_gram/go_gram.go`.

Context
=======

This project is produced as part of a MSc. Computer Science dissertation in [Kent University](http://www.kent.ac.uk/).
