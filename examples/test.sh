#!/bin/bash

go build ex_go_gram.go go_gram.go &&
        (./ex_go_gram "`cat test/test0.go.g`" > log  || echo test0 failed) &&
        (./ex_go_gram "`cat test/test1.go.g`" >> log || echo test1 failed) &&
        (./ex_go_gram "`cat test/test2.go.g`" >> log || echo test2 failed) &&
        (./ex_go_gram "`cat test/test3.go.g`" >> log || echo test3 failed) &&
        (./ex_go_gram "`cat test/test4.go.g`" >> log || echo test4 failed) &&
        (./ex_go_gram "`cat test/test.go.g`"  >> log || echo test. failed) &&

        (./ex_go_gram "`cat test/test`"       >> log || echo test failed)
