package area

import (
	"fmt"
	"github.com/davidjohansson/ecmd/util"
)

func ListArea(area string, section string) {
	url := fmt.Sprintf("https://objectapi-stage.app.svt.se/object-api/section/%s/areas/%s", section, area)
	util.PrintRegexInRespBody(`"id": (\d+)`, url)
}
