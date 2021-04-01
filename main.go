package main

import (
	"flag"
	"git.jcu.edu.au/cft/cfds/crypt"
	"git.jcu.edu.au/cft/cfds/datasources"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	seed := flag.String("seed", "", "Seed from seed.properties. 16 chars")
	input := flag.String("input", "datasources.yml", "Input yaml data sources")
	templateDir := "templates"
	flag.Parse()
	if *seed == "" || *input == "" {
		flag.Usage()
	}
	inFile, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	d := yaml.NewDecoder(inFile)
	doc := datasources.NewDoc()
	err = d.Decode(doc)
	_ = inFile.Close()

	templates := map[string]*template.Template{}
	if err != nil {
		panic("Error reading datasource file " + err.Error())
	}
	for ds, data := range doc.Datasources {
		data.Name = ds
		data.Password, err = crypt.Encrypt(data.Password, *seed)
		if err != nil {
			panic("Could not encrypt password")
		}
		key := strings.ToLower(data.Type)
		templateFile := filepath.Join(templateDir, key+".xml")
		if _, ok := templates[key]; !ok {
			// Load template
			content, err := ioutil.ReadFile(templateFile)
			if err != nil {
				panic("Cannot load template " + templateFile)
			}
			templates[key], err = template.New(key).Parse(string(content))
			if err != nil {
				panic(err)
			}
		}
		err = templates[key].Execute(os.Stdout, data)
		if err != nil {
			panic(err)
		}
	}

	//tmpl := template.New("ds")
}
