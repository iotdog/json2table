package j2t

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSON2HtmlTable convert json string to html table string
func JSON2HtmlTable(jsonStr string) (bool, string) {
	htmlTable := ""
	jsonArray := []map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonStr), &jsonArray)
	if err != nil || 0 == len(jsonArray) {
		fmt.Println("invalid json string")
		return false, htmlTable
	}

	// convert table headers
	titles := getKeys(jsonArray[0])
	if 0 == len(titles) {
		fmt.Println("json is not supported")
	}
	tmp := []string{}
	for _, title := range titles {
		tmp = append(tmp, fmt.Sprintf("<th>%s</th>", title))
	}
	thCon := strings.Join(tmp, "")

	// convert table cells
	rows := []string{}
	for _, jsonObj := range jsonArray {
		tmp = []string{}
		for _, key := range titles {
			cell := fmt.Sprintf("<td>%v</td>", jsonObj[key])
			tmp = append(tmp, cell)
		}
		tdCon := strings.Join(tmp, "")
		row := fmt.Sprintf("<tr>%s</tr>", tdCon)
		rows = append(rows, row)
	}
	trCon := strings.Join(rows, "")

	htmlTable = fmt.Sprintf(`<table border="1" cellpadding="1" cellspacing="1">%s%s</table>`,
		fmt.Sprintf("<thead>%s</thead>", thCon), fmt.Sprintf("<tbody>%s</tbody>", trCon))
	return true, htmlTable
}

func getKeys(jsonObj map[string]interface{}) []string {
	keys := make([]string, 0, len(jsonObj))
	for k := range jsonObj {
		keys = append(keys, k)
	}
	return keys
}
