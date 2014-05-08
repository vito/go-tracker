package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"

	"github.com/xoebus/go-tracker"
)

func main() {
	token := readToken()
	client := tracker.NewClient(token)

	me, err := client.Me()
	if err != nil {
		log.Fatalf("could not get current tracker user: %s", err)
	}

	fmt.Printf("%+v", me)
}

func readToken() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("could not get current user: %s", err)
	}

	path := filepath.Join(user.HomeDir, ".trackertoken")
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not read .trackertoken file: %s", err)
	}

	return string(contents)
}
