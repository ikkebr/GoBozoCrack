package main

import (
	"flag"
	"fmt"
	"github.com/ikkebr/gobozocrack/cracker"
)

func main() {
	var single = flag.String("single", "MD5HASH", "cracks a single hash")
	var file = flag.String("file", "HASHFILE", "cracks multiple hashes on a file")
	flag.Parse()

	if *single != "MD5HASH" {
		fmt.Println(cracker.Crack_single_hash(*single))
	} else if *file != "HASHFILE" {
		cracker.Crack(*file)
	} else {
		fmt.Println("Please select either -single YOUR_MD5HASH or -file file.ext")
	}
}
