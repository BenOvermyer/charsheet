package main

import (
	"flag"
	"fmt"
	"ioutil"
	"json"
	"html/template"
	"os"
)

// CharSheet is a character sheet
type CharSheet struct {
	Name string `json:"name"`
	Classes struct  {
		Name string `json:"name"`
		Level int `json:"level"`
	} `json:"classes"`
	Skills struct {
		Name string `json:"name"`
		Rank int `json:"rank"`
	} `json:"skills"`
}

func getSheetTemplate() string {
	return `<!DOCTYPE html>
	<html>
	<head>
	  <title>Character Sheet</title>
	</head>
	<body>
		<h1>{{.Name}}</h1>
		<h2>Classes</h2>
		{{range .Classes}}
		<div><p><strong>{{.Name}}:</strong> Level {{.Level}}</p></div>
		{{end}}
		<h2>Skills</h2>
		<p>Trained skills. Ranks go from 1 (some familiarity) to 5 (expert).</p>
		<table>
			<thead>
			<tr><th>Skill</th><th>Ranks</th></tr>
			</thead>
			<tbody>
			{{range .Skills}}
			<tr><th>{{.Name}}</th><td>{{.Rank}}</td></tr>
			{{end}}
			</tbody>
		</table>
	</body>
	</html>`
}

func main() {
	var charSheetData CharSheet

	filePath := flag.String("f", "character.json", "JSON data file to use")
	outputFile := flag.String("o", "sheet.html", "Output file to create")
	flag.Parse()

	byteValue, _ := ioutil.ReadAll(*filePath)

	json.Unmarshal(byteValue, &charSheetData)

	htmlIndexTemplate := getSheetTemplate()

	writer, err := os.Create("./" + *outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := template.New("htmlIndex").Parse(htmlIndexTemplate)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Execute(writer, charSheetData)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer writer.Close()
}