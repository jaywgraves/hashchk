package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"sort"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage:  %s [-h=md5|sha1] files... \n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "This utility will hash all the files given on the command line using the given\n")
	fmt.Fprintf(os.Stderr, "hash function.\n")
	fmt.Fprintf(os.Stderr, "It will print the results two ways, sorted by hash digest and by file name.\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "  files...: 1 or more file name specifications\n")
}

type hashtype func() hash.Hash

func HashFile(filepath string, ht hashtype, digestlen int) string {
	if f, e := os.Open(filepath); e == nil {
		defer f.Close()
		hash := ht()
		if _, e := io.Copy(hash, f); e == nil {
			digest := make([]byte, 0, digestlen)
			return hex.EncodeToString(hash.Sum(digest))
		}
	}
	return ""
}

type FileDigest struct {
	fpth, dgst string
}

type ByFile []FileDigest

func (s ByFile) Len() int           { return len(s) }
func (s ByFile) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByFile) Less(i, j int) bool { return s[i].fpth <= s[j].fpth }

type ByDigest []FileDigest

func (s ByDigest) Len() int           { return len(s) }
func (s ByDigest) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByDigest) Less(i, j int) bool { return s[i].dgst <= s[j].dgst }

func output(sorteddata []FileDigest) {
	for _, fd := range sorteddata {
		fmt.Println(fd.dgst, "->", fd.fpth)
	}
}

func main() {

	hashflg := flag.String("h", "md5", "hash function to use: md5|sha1")
	flag.Usage = Usage
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "No file specifications given as arguments.\n")
		Usage()
		return
	}

	digestlen := 0
	var h func() hash.Hash
	switch *hashflg {
	case "md5":
		digestlen = 32
		h = md5.New
		//fmt.Printf("%T\n",h)
	case "sha1":
		digestlen = 40
		h = sha1.New
		//fmt.Printf("%T\n",h)
	}

	var data []FileDigest
	for _, a := range flag.Args() {
		matches, _ := filepath.Glob(a)
		for _, m := range matches {
			data = append(data, FileDigest{m, HashFile(m, h, digestlen)})
		}
	}
	sort.Sort(ByFile(data))
	fmt.Println("by file name")
	output(data)
	fmt.Println("\n\nby file digest")
	sort.Sort(ByDigest(data))
	output(data)

}
