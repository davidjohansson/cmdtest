package ecmdarea

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"
	"regexp"
)

func ListArea(area string, section string) {

	url := fmt.Sprintf("https://objectapi-stage.app.svt.se/object-api/section/%s/areas/%s", section, area)
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	var out bytes.Buffer
	json.Indent(&out, []byte(jsonDataFromHttp), "", "\t")

	output := out.String()

	r := regexp.MustCompile(`"id": (\d+)`)
	result_slice := r.FindAllStringSubmatch(output, -1)

	for _, v := range result_slice {
		fmt.Printf("%s\n", v[1])
	}
}
