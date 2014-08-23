package main

import (
	docopt "github.com/docopt/docopt-go"
	"github.com/souleiman/warden/parser"
)

var usage string = `Usage:
	warden [-a] -u <url>

-u, --url   The resource URL
-a   Process every set`

func main() {
	args, _ := docopt.Parse(usage, nil, true, "Warden 0.1", true)
	var url string = args["<url>"].(string)

	parser.Parse(url, args["-a"].(bool))
}
