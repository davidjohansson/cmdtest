package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Inspect(fields string, responsedata string, relations string, meta string, objectids []string) {
	for _, id := range objectids {
		if id != "" {
			InspectOne(fields, responsedata, relations, meta, id)
		}
	}
}

func InspectOne(fieldkeys string, responsekeys string, relationkeys string, metakeys string, objectid string) {

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
	jsonResponse := jsonData.(map[string]interface{})

	if err != nil {
		panic(err)
	}

	responseData := jsonResponse["response"].(map[string]interface{})
	linksData := jsonResponse["links"].(map[string]interface{})
	metaData := jsonResponse["meta"].(map[string]interface{})

	printResponseData(responsekeys, objectid, responseData)
	printArticle(fieldkeys, objectid, responseData)
	printRelations(relationkeys, objectid, linksData)
	printMeta(metakeys, objectid, metaData)

	fmt.Print("\n")
}

func printResponseData(responseData string, objectid string, response map[string]interface{}) {

	if responseData == "_all" {
		responseData = "contentType,relativeUri,url"
	}

	printFields(response, responseData)
}

func printArticle(fields string, objectid string, response map[string]interface{}) {
	article := response["article"].(map[string]interface{})
	printFields(article, fields)

}

func printRelations(relations string, objectId string, links map[string]interface{}) {

	if relations != "" {
		relatedMeta := links["related"].(map[string]interface{})
		relatedContent := relatedMeta[relations].([]interface{})

		for _, v := range relatedContent {
			teaser := v.(map[string]interface{})
			fmt.Println(int(teaser["id"].(float64)))
		}
	}
}

func printMeta(metaFields string, objectid string, metaData map[string]interface{}) {
	printFields(metaData, metaFields)
}

func printFields(datamap map[string]interface{}, fields string) {
	for k, v := range datamap {
		key := ""
		fieldsslice := strings.Split(fields, ",")
		if fields != "_all" {
			for _, rawfield := range fieldsslice {
				field := strings.TrimSpace(rawfield)
				if (strings.EqualFold(field, strings.TrimSpace(k))) {
					key = strings.TrimSpace(k)
				}
			}
		} else {
			key = k
		}

		if key != "" {
			kp := strings.TrimSpace(key)
			fmt.Print(kp)
			fmt.Print(": ")

			if key == "id" {
				fmt.Print(int(v.(float64)))
			} else {
				fmt.Print(v)
			}

			fmt.Print("\n")
		}
	}
}

