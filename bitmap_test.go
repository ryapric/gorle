package main

import (
	"reflect"
	"testing"
)

func TestBitmap(t *testing.T) {
	data := []byte("aaabbb")

	/// 97 is byte "a", 98 is byte "b"
	want := bitmap{
		97: []int{1, 1, 1, 0, 0, 0},
		98: []int{0, 0, 0, 1, 1, 1},
	}

	got := getBitmap(data)

	if !reflect.DeepEqual(want, got) {
		t.Errorf(
			"getBitmap() did not make a proper bitmap out of '%s' (want %v, got %v)\n",
			data, want, got,
		)
	}

	data = []byte("abaabb")
	want = bitmap{
		97: []int{1, 0, 1, 1, 0, 0},
		98: []int{0, 1, 0, 0, 1, 1},
	}
	got = getBitmap(data)
	if !reflect.DeepEqual(want, got) {
		t.Errorf(
			"getBitmap() did not make a proper bitmap out of '%s' (want %v, got %v)\n",
			data, want, got,
		)
	}
}

func TestCompressBitmap(t *testing.T) {
	bm := bitmap{
		97: []int{1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
		98: []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		99: []int{1, 0, 0, 1, 1, 0, 1, 1, 1, 1},
	}

	want := bitmapRLE{
		OriginalLength: 10,
		RLE: map[byte]string{
			97: "0,5,5",
			98: "5,5",
			99: "0,1,2,2,1,4",
		},
	}

	got := compressBitmap(bm)

	if !reflect.DeepEqual(want, got) {
		t.Errorf(
			"bitmap compression doesn't look right (want %v, got %v)\n",
			want, got,
		)
	}
}

func TestDecompress(t *testing.T) {
	want := []byte("abaabb")

	bmRLE := bitmapRLE{
		OriginalLength: 6,
		RLE: map[byte]string{
			97: "0,1,1,2,2",
			98: "1,0,2,2",
		},
	}
	got := decompress(bmRLE)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("bitmapRLE decompression doesn't look right (want %v, got %v\n", want, got)
	}
}
