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
	parser [-a] -u <url>

-u, --url   The resource URL
-a   Process every set`

func main() {
	args, _ := docopt.Parse(usage, nil, true, "Sampler 0.1", true)
	var url string = args["<url>"].(string)

	request, _ := http.Get(url)
	defer request.Body.Close()

	fmt.Println("Fetching resource...")
	body, _ := ioutil.ReadAll(request.Body)
	fmt.Println("Resource has been fetched, processing...")

	var err error
	if args["-a"].(bool) {
		var sets map[string]card.Set
		err = json.Unmarshal(body, &sets)

		for _, set := range sets {
			set.Clean()
		}
	} else {
		var set card.Set
		err = json.Unmarshal(body, &set)
		set.Clean()
	}

	fmt.Println(err)
}
