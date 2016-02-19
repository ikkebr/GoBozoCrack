package cracker

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func load_cache(cache_file string) map[string]string {

	cache := make(map[string]string)

	file, err := os.OpenFile(cache_file, os.O_CREATE|os.O_RDONLY, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		line := scanner.Text()
		//log.Println(line)
		kv := strings.Split(line, ":")
		cache[kv[0]] = kv[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	return cache
}

func append_to_cache(cache_file string, hash string, value string) {

	file, err := os.OpenFile(cache_file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}

	bufferedWriter := bufio.NewWriter(file)
	_, err = bufferedWriter.WriteString(format_it(hash, value) + "\n")

	if err != nil {
		log.Fatal(err)
	}

	if err = bufferedWriter.Flush(); err != nil {
		log.Println(err)
	}

	file.Close()

}

func dictionary_attack(h string, wordlist *[]string) string {

	for _, word := range *wordlist {
		if get_MD5_hash(word) == h {
			return word
		}
	}

	return ""
}

var cache = load_cache("cache.db")

func Crack(hash_file string) error {

	file, err := os.OpenFile(hash_file, os.O_RDONLY, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
		return err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(Crack_single_hash(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	file.Close()

	return nil
}

func Crack_single_hash(h string) string {

	if val, ok := cache[h]; ok {
		return h+":"+val
	}

	resp, err := http.Get(fmt.Sprintf("http://www.google.com/search?q=%v", h))

	if err != nil {
		log.Fatal(err)
		return "Could not connect to Google"
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return "Could not retrieve results from Google"
	}

	sbody := string(body)

	sbody = strings.Replace(sbody, ".", " ", -1)

	wordlist := strings.Split(sbody, " ")

	remove_duplicates(&wordlist)

	plaintext := dictionary_attack(h, &wordlist)

	if plaintext != "" {
		cache[h] = plaintext
		append_to_cache("cache.db", h, plaintext)
	} else {
		plaintext = "ERROR - Hash could not be cracked."
	}

	return h+":"+plaintext
}
