package main

import (
	"flag"
	"fmt"
	"time"
)

type Test struct {
	Name string // Name + "_" + "beg|mid|end" + ".go.g"
	Beg  File
	Mid  File
	End  File
}

type File struct {
	Pos int
	Sub string
}

var tests []Test = []Test{
	Test{
		Name: "small",
		Beg:  File{Pos: 0, Sub: "package"},
		Mid:  File{Pos: 14, Sub: "ioutil"},
		End:  File{Sub: "}"},
	},
	Test{
		Name: "medium",
		Beg:  File{Pos: 0, Sub: "package"},
		Mid:  File{Pos: 35, Sub: "if"},
		End:  File{Sub: "}"},
	},
	Test{
		Name: "big",
		Beg:  File{Pos: 0, Sub: "package"},
		Mid:  File{Pos: 133, Sub: "if"},
		End:  File{Sub: "}"},
	},
}

func main() {

	flag.Parse()

	for _, t := range tests {

		beg := getFile(t.Name, "beg")
		mid := getFile(t.Name, "mid")
		end := getFile(t.Name, "end")
		origin := getFile(t.Name, "origin")

		{
			computeBegNotOnline(t.Beg, beg)
			computeBegOnline(t.Beg, beg)
			computeBegIncremental(t.Beg, beg)

			computeMidNotOnline(t.Mid, mid)
			computeMidOnline(t.Mid, mid)
			computeMidIncremental(t.Mid, mid)

			computeEndNotOnline(t.End, end)
			computeEndOnline(t.End, end)
			computeEndIncremental(t.End, end)
			fmt.Println()
		}

		{
			computeBegIncrementalNC(t.Beg, beg)
			computeMidIncrementalNC(t.Mid, mid)
			fmt.Println()
		}

		{
			now := time.Now()

			t1 := now
			cyk := computeCYK(origin)
			for i := 0; i < 10; i++ {
				t := (computeTreeNormal(cyk))
				t1 = t1.Add(t)
			}
			t2 := now
			for i := 0; i < 10; i++ {
				t := computeTreeConcurrently(cyk)
				t2 = t2.Add(t)
			}
			fmt.Println("diff=", t1.Nanosecond()-t2.Nanosecond(), "us")
			fmt.Println()
		}

	}
}
