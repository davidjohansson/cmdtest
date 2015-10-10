package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Inspect(fields string, objectids []string) {
	for _, id := range objectids {
		InspectOne(fields, id)
	}
}

func InspectOne(fields string, objectid string) {
	url := fmt.Sprintf("https://objectapi-stage.app.svt.se/object-api/article/%s", objectid)
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		panic(err)

	}

	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var jsonData interface{}
	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	ma := jsonData.(map[string]interface{})

	if err != nil {
		panic(err)
	}

	responseMeta := ma["response"].(map[string]interface{})
	article := responseMeta["article"].(map[string]interface{})

	fmt.Println(fmt.Sprintf("id:%s", objectid))
	if fields != "" {
		fieldsslice := strings.Split(fields, ",")
		for _, rawfield := range fieldsslice {
			field := strings.TrimSpace(rawfield)
			fmt.Print(field)
			fmt.Print(": ")
			fmt.Print(article[field])
			fmt.Print("\n")
		}
	} else {
		for k, v := range article {
			if v != "" {
				k := strings.TrimSpace(k)
				fmt.Print(k)
				fmt.Print(": ")
				fmt.Print(v)
				fmt.Print("\n")
			}
		}
	}
	fmt.Print("\n")
}

