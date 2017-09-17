package j2t

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSON2HtmlTable convert json string to html table string
func JSON2HtmlTable(jsonStr string, customTitles []string) (bool, string) {
	htmlTable := ""
	jsonArray := []map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonStr), &jsonArray)
	if err != nil || 0 == len(jsonArray) {
		fmt.Println("invalid json string")
		return false, htmlTable
	}

	titles := customTitles
	if nil == customTitles || 0 == len(titles) { // if custom titles are not provided, use json keys as titles
		titles = getKeys(jsonArray[0])
	} else { // if custom titles are provided, sort json array
		for tid, title := range titles {
			swapped := true
			for swapped {
				swapped = false
				for i := 0; i < len(jsonArray)-1; i++ {
					va, oka := jsonArray[i][title].(string)
					vb, okb := jsonArray[i+1][title].(string)
					if !oka || !okb {
						swapped = false
						break
					}
					if strings.Compare(va, vb) > 0 {
						if tid != 0 {
							va, _ := jsonArray[i][titles[tid-1]].(string)
							vb, _ := jsonArray[i+1][titles[tid-1]].(string)
							if va != vb {
								continue
							}
						}
						tmp := jsonArray[i]
						jsonArray[i] = jsonArray[i+1]
						jsonArray[i+1] = tmp
						swapped = true
					}
				}
			}
		}
	}
	// convert table headers
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
