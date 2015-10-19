package solr

import (
	"fmt"
	"github.com/davidjohansson/ecmd/util"
)

func Search(searchtype string) {
	url := fmt.Sprintf("http://svt-stostage-search01:8080/solr/collection1/select?q=contenttype%%3A%s&wt=json&indent=false", searchtype)
	regex := `"objectid": "(\d+)"`
	util.PrintRegexInRespBody(regex, url)
}