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
	for k, v := range article {
		key := ""
		if fields != "" {
			fieldsslice2 := strings.Split(fields, ",")
			for _, rawfield := range fieldsslice2 {
				field := strings.TrimSpace(rawfield)
				if (strings.EqualFold(field, strings.TrimSpace(k))) {
					key = k
				}
			}
		} else {
			key = k
		}

		if key != "" {
			kp := strings.TrimSpace(key)
			fmt.Print(kp)
			fmt.Print(": ")
			fmt.Print(v)
			fmt.Print("\n")
		}
	}

	fmt.Print("\n")
}

