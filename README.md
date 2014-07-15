GoCYK
=====

GoCYK is an implementation of an online parsing using the CYK algorithm.
The given result recognition table is a merge between these two resources : 
* [Parsing techniques: A practical guide (*ch. 4*)](http://port70.net/~nsz/articles/book/grune_jacobs_parsing_techniques_2008.pdf)
* [Parallel on-line parsing in constant time per word](http://doc.utwente.nl/18047/1/Sikkel93parallel.pdf)

Installation
=====

After having a Go environment installed, just call `go get`:
```bash
$ go get github.com/ghigt/gocyk
```

Usage
=====

One parameter corresponding to the string to be parsed, for example:
```bash
$ gocyk "12.3e+1"

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
