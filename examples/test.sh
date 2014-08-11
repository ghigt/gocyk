#!/bin/bash

go build -o examples ex_go_gram.go go_gram.go && (
        (./examples "`cat test/test0.go.g`" >  log && echo test0 PASS || echo test0 failed) &&
        (./examples "`cat test/test1.go.g`" >> log && echo test1 PASS || echo test1 failed) &&
        (./examples "`cat test/test2.go.g`" >> log && echo test2 PASS || echo test2 failed) &&
        (./examples "`cat test/test3.go.g`" >> log && echo test3 PASS || echo test3 failed) &&
        (./examples "`cat test/test4.go.g`" >> log && echo test4 PASS || echo test4 failed) &&
        (./examples "`cat test/test5.go.g`" >> log && echo test4 PASS || echo test4 failed) &&
        (./examples "`cat test/test.go.g`"  >> log && echo test. PASS || echo test. failed) &&

        (./examples "`cat test/test`"       >> log && echo test  PASS || echo test  failed) 
)
