package main

import (
	"fmt"
	"strconv"
	"strings"
)

type bitmap map[byte][]int

type bitmapRLE struct {
	OriginalLength int
	RLE            map[byte]string
}

// getBitmap makes a bitmap
func getBitmap(data []byte) bitmap {
	seenBytes := make(map[byte]int)

	// First, find all the unique bytes
	for _, e := range data {
		seenBytes[e]++
	}

	// Then, make your bitmap
	bm := make(bitmap)
	for k := range seenBytes {
		for _, e := range data {
			if k == e {
				bm[k] = append(bm[k], 1)
			} else {
				bm[k] = append(bm[k], 0)
			}
		}
	}

	return bm
}

/*
compressBitmap iterates through a bitmap, and run-length encodes it. The
function checks, in order, the count of zeroes seen, followed by ones, and
repeats until done.
*/
func compressBitmap(bm bitmap) bitmapRLE {
	bmRLE := make(map[byte]string)

	// One iteration to get original length of data, since maps don't have
	// indices to check against directly
	var n int
	for k := range bm {
		n = len(bm[k])
		break
	}

	for k := range bm {
		run := 1
		for i := range bm[k] {
			// Wrap up before you go out of range
			if i == n-1 {
				if bm[k][i] == bm[k][i-1] {
					bmRLE[k] += fmt.Sprintf("%d", run)
				} else {
					bmRLE[k] += fmt.Sprintf("%d", 1)
				}
				break
			}

			// If first value isn't 0, record zero runs and continue
			if i == 0 && bm[k][i] != 0 {
				bmRLE[k] += fmt.Sprintf("%d,", 0)
			}

			if bm[k][i] == bm[k][i+1] {
				run++
				// fmt.Printf(
				// 	"k: %d, i: %d, bm[k][i]: %d, run: %d\n",
				// 	k, i, bm[k][i], run,
				// )
			} else {
				bmRLE[k] += fmt.Sprintf("%d,", run)
				run = 1
				continue
			}
		}
	}

	bmRLEOut := bitmapRLE{
		OriginalLength: n,
		RLE:            bmRLE,
	}

	return bmRLEOut
}

func decompress(bmRLE bitmapRLE) []byte {
	bm := bitmap{}
	for b := range bmRLE.RLE {
		split := strings.Split(bmRLE.RLE[b], ",")
		for i := range split {
			v, err := strconv.Atoi(split[i])
			if err != nil {
				panic(err)
			}
			bm[b] = append(bm[b], v)
		}
		fmt.Println(split)
	}
	fmt.Println(bm)

	fmt.Println("not implemented")
	return []byte("not implemented")
}
