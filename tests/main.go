package main

import "flag"

var (
	beg = flag.String("beg", "", "beginning missed")
	mid = flag.String("mid", "", "middle missed")
	end = flag.String("end", "", "end missed")
)

func main() {

	flag.Parse()

	computeBegNotOnline(*beg)
	computeBegOnline(*beg)
	computeBegIncremental(*beg)

	computeMidNotOnline(*mid)
	computeMidOnline(*mid)
	computeMidIncremental(*mid)

	computeEndNotOnline(*end)
	computeEndOnline(*end)
	computeEndIncremental(*end)
}
