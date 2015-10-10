package ecmdsolrsearch

import (
	"fmt"
	curl "github.com/andelf/go-curl"
	"regexp"
)

func Search(searchtype string) {
	easy := curl.EasyInit()
	defer easy.Cleanup()

	url := fmt.Sprintf("http://svt-stostage-search01:8080/solr/collection1/select?q=contenttype%%3A%s&wt=json&indent=false", searchtype)

	easy.Setopt(curl.OPT_URL, url)

	fooTest := func(buf []byte, userdata interface{}) bool {
		output := string(buf)
		r := regexp.MustCompile(`"objectid":"(\d+)"`)
		result_slice := r.FindAllStringSubmatch(output, -1)


		for _, v := range result_slice {
			fmt.Printf("%s\n", v[1])
		}

		return true
	}

	easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}