package cracker

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func remove_duplicates(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

func get_MD5_hash(plaintext string) string {
	h := md5.Sum([]byte(plaintext))
	return hex.EncodeToString(h[:])
}

func format_it(hash string, plaintext string) string {
	s := fmt.Sprintf("%s:%s", hash, plaintext)
	return s
}
