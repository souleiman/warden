package parser

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	kard "github.com/souleiman/warden/card"
)

func Parse(url string, all bool) {
	request, _ := http.Get(url)
	defer request.Body.Close()

	fmt.Println("Fetching resource...")
	body, _ := ioutil.ReadAll(request.Body)
	fmt.Println("Resource has been fetched, processing...")

	var err error
	if all {
		var sets map[string]set
		var conv_sets []kard.Set

		err = json.Unmarshal(body, &sets)

		for _, set := range sets {
			set.clean()
			conv_sets = append(conv_sets, set.convert())
		}
	} else {
		var set set
		err = json.Unmarshal(body, &set)
		set.clean()
		//conv_set := set.convert()
	}

	fmt.Println(err)
}
