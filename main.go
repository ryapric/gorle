package main

import (
	"fmt"
)

// func encodeOld(data string) {
// 	contents, err := ioutil.ReadFile(data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var valKeyDict []string

// 	// Map the byte positions to each byte by pasting them together (we'll split
// 	// afterwards on the delimeter), making sure the position is second so the
// 	// value sorting is consistent. Then, sort.
// 	for i, e := range contents {
// 		valKeyDict = append(valKeyDict, fmt.Sprintf("%v: %v", string(e), i))
// 	}
// 	sort.Strings(valKeyDict)

// 	// Split the position keys from the values
// 	var keys []string
// 	var vals []string
// 	for _, e := range valKeyDict {
// 		valKey := strings.Split(e, ": ")
// 		vals = append(vals, valKey[0])
// 		keys = append(keys, valKey[1])
// 	}

// 	// Run-length encode the values
// 	var run int = 1
// 	var valsOut []string
// 	for i, e := range vals {
// 		if i == 0 {
// 			continue
// 		} else if e == vals[i-1] {
// 			run++
// 		} else {
// 			valsOut = append(valsOut, fmt.Sprintf("%v:%v", vals[i-1], run))
// 			run = 1
// 		}
// 	}

// 	// Write out both the file with sorted position keys, and the values file
// 	keyFile, err := os.Create("./keyfile.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	valFile, err := os.Create("./valfile.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer keyFile.Close()
// 	defer valFile.Close()

// 	keyWriter := bufio.NewWriter(keyFile)
// 	valWriter := bufio.NewWriter(valFile)

// 	// Key positions are written out delimited
// 	for _, s := range keys {
// 		_, err := keyWriter.WriteString(s + "\n")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if err := keyWriter.Flush(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	for _, s := range valsOut {
// 		_, err := valWriter.WriteString(s + "\n")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if err := valWriter.Flush(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// // func decode(keyFile, valFile *string) {
// //	keyData, err := ioutil.ReadFile(*keyFile)
// //	if err != nil {
// //		log.Fatal(err)
// //	}
// //	valData, err := ioutil.ReadFile(*valFile)
// //	if err != nil {
// //		log.Fatal(err)
// //	}

// // }

func main() {
	// var keyFile = flag.String("k", "", "File containing decryption key data")
	// var valFile = flag.String("v", "", "File containing encrypted data values")
	// flag.Parse()

	// if *keyFile == "" {
	//	log.Fatal("ERROR: yo wtf where the keyfile")
	// }
	// if *valFile == "" {
	//	log.Fatal("ERROR: yo wtf where the valfile")
	// }

	// args := flag.Args()
	// if len(args) == 0 {
	// 	log.Fatal("ERROR: You must provide a file to encrypt!")
	// }
	// encodeViaDB(args[0])
	// decode(keyFile, valFile)

	x := getBitmap([]byte("the quick brown fox jumped over the lazy dog"))
	y := compressBitmap(x)
	fmt.Println(y)
}
