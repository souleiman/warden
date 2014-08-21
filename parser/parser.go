package main

import (
	docopt "github.com/docopt/docopt-go"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/souleiman/warden/card"
)

var usage string = `Usage:
	test -u <url>

-u, --url   The resource URL`

func main() {
	args, _ := docopt.Parse(usage, nil, true, "Sampler 0.1", true)
	var url string = args["<url>"].(string)

	request, _ := http.Get(url)
	defer request.Body.Close()

	fmt.Println("Fetching resource...")
	body, _ := ioutil.ReadAll(request.Body)
	fmt.Println("Resource has been fetched, processing...")

	var sets map[string]card.Set
	err := json.Unmarshal(body, &sets)

	for _, set := range sets {
		set.Transform()
	}

	fmt.Println(err)
}
