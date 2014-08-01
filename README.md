GoCYK
=====

GoCYK is an implementation of an online parsing using the CYK algorithm.
The given result recognition table is a merge between these two resources: 
* [Parsing techniques: A practical guide (*ch. 4*)](http://port70.net/~nsz/articles/book/grune_jacobs_parsing_techniques_2008.pdf)
* [Parallel on-line parsing in constant time per word](http://doc.utwente.nl/18047/1/Sikkel93parallel.pdf)

[![GoDoc](https://godoc.org/github.com/ghigt/gocyk?status.svg)](https://godoc.org/github.com/ghigt/gocyk)

Installation
=====

After having a Go environment installed, just call `go get`:
```bash
$ go get github.com/ghigt/gocyk
```

Example
=====

Look at the examples folder. A default grammar file is present. Feel free to modify it. But be aware that it must respect the [CNF (Chomsky Normal Form)](http://en.wikipedia.org/wiki/Chomsky_normal_form).  
Build the example like as follow:
```bash
$ go build
```
Then, launch the program with one parameter corresponding to the string to be parsed, for example `12.3e+4` for the default grammar (-v is the option to see the recognition tree in action):
```bash
$ ./ex_num_gram -v "12.3e+4"

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
