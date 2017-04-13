package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"regexp"

	"github.com/flosch/pongo2"
	"github.com/go-openapi/loads/fmts"
	"github.com/go-openapi/spec"
)

func regexReplace(src string, pattern string, repl string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(src, repl)
}

func prettyJSON(arg interface{}) string {
	data, _ := json.MarshalIndent(arg, "", "  ")
	return string(data)
}

func main() {
	data, err := fmts.YAMLDoc("/Users/yandongdong/works/test/contract.yaml")
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println(string(data))

	var sw = &spec.Swagger{}

	err = sw.UnmarshalJSON(data)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = spec.ExpandSpec(sw, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}

	data, err = sw.MarshalJSON()

	log.Println(string(data))

	in, err := ioutil.ReadFile("api_doc.md")
	if err != nil {
		log.Fatalln(err)
		return
	}

	tpl, err := pongo2.FromBytes(in)
	if err != nil {
		log.Fatalln(err)
		return
	}

	out, err := tpl.Execute(pongo2.Context{"sw": sw, "op": map[string]interface{}{
		"prettyjson":   prettyJSON,
		"regexReplace": regexReplace,
	}})
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println(out)
}
