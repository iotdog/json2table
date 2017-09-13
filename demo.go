package main

import (
	"fmt"

	"github.com/iotdog/json2table/j2t"
)

func main() {
	jsonStr := `[{"title1": "hello", "title2": "world"}, {"title1": "have", "title2": "fun"}]`
	ok, html := j2t.JSON2HtmlTable(jsonStr)
	if ok {
		fmt.Println(html)
	} else {
		fmt.Println("failed to convert json to html table")
	}
}
