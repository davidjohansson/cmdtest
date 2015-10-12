package area

import (
	"fmt"
	"github.com/davidjohansson/ecmd/util"
)

func ListArea(area string, section string, listAreas bool) {

	if listAreas {
		url := fmt.Sprintf("https://objectapi-stage.app.svt.se/object-api/section/%s/areas", section)
		util.PrintRegexInRespBody(`"([a-z]+-[a-z]+)"`, url)
	} else {
		url := fmt.Sprintf("https://objectapi-stage.app.svt.se/object-api/section/%s/areas/%s", section, area)
		util.PrintRegexInRespBody(`"id": (\d+)`, url)
	}
}
