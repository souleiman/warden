package parser

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
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
		err = json.Unmarshal(body, &sets)

		for _, s := range sets {
			s.clean()
		}
	} else {
		var s set
		err = json.Unmarshal(body, &s)
		s.clean()
	}

	fmt.Println(err)
}
