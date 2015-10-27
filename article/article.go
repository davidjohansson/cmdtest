package article

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Inspect(fields string, responsedata string, relations string, objectids []string) {
	for _, id := range objectids {
		if id != ""{
			InspectOne(fields, responsedata, relations, id)
		}
	}
}

func InspectOne(fields string, responsedata string, relations string, objectid string) {

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
	linksMeta := ma["links"].(map[string]interface{})

	//print data
//	fmt.Println(fmt.Sprintf("id:%s", objectid))
	printResponseData(responsedata, objectid, responseMeta)
	printArticle(fields, objectid, responseMeta)
	printRelations(relations, objectid, linksMeta)

	fmt.Print("\n")
}

type Args struct {
	ArticleFields  []string
	MetaFields     [] string
	RelationFields [] string
}

func printArticle(fields string, objectid string, response map[string]interface{}) {
	article := response["article"].(map[string]interface{})

	printFields(article, fields)

}

func printRelations(relations string, objectId string, links map[string]interface{}) {

	if relations != "" {
		relatedMeta := links["related"].(map[string]interface{})

		teasers := relatedMeta[relations].([]interface{})

		for _, v := range teasers{
			teaser := v.(map[string]interface{})
			fmt.Println(int(teaser["id"].(float64)))
		}
	}
}

func printResponseData(responseData string, objectid string, response map[string]interface{}) {

	if responseData == "_all" {
		responseData = "contentType,relativeUri,url"
	}

	printFields(response, responseData)
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

			//		} else {
			//			key = k
		}

		if key != "" {
			kp := strings.TrimSpace(key)
			fmt.Print(kp)
			fmt.Print(": ")

			if key == "id"{
				fmt.Print(int(v.(float64)))
			} else{
				fmt.Print(v)
			}

			fmt.Print("\n")
		}
	}
}

