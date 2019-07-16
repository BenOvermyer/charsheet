package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

// CharSheet is a character sheet
type CharSheet struct {
	Name    string  `json:"name"`
	Classes []Class `json:"classes"`
	Skills  []Skill `json:"skills"`
}

// Class is a character class
type Class struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// Skill is a character skill
type Skill struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
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

	// Open our jsonFile
	jsonFile, err := os.Open(*filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &charSheetData)

	htmlTemplate := getSheetTemplate()

	writer, err := os.Create("./" + *outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(charSheetData)

	t, err := template.New("sheet").Parse(htmlTemplate)
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
