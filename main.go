package main

import (
	"flag"
	"fmt"
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
	output := flag.String("output", "neo-datasource.xml", "Output XML file")
	templateDir := "templates"
	flag.Parse()
	if *seed == "" || *input == "" {
		flag.Usage()
	}
	inFile, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()
	d := yaml.NewDecoder(inFile)
	doc := datasources.NewDoc()
	err = d.Decode(doc)
	if err != nil {
		panic("Error reading datasource file " + err.Error())
	}

	start := "<wddxPacket version='1.0'><header/><data><array length='2'><struct type='coldfusion.server.ConfigMap'>"
	end := "</struct><struct type='coldfusion.server.ConfigMap'><var name='maxcachecount'><number>100.0</number></var></struct></array></data></wddxPacket>"

	outFile, err := os.Create(*output)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	templates := map[string]*template.Template{}

	if _, err = outFile.WriteString(start); err != nil {
		panic(err)
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
		err = templates[key].Execute(outFile, data)
		if err != nil {
			panic(err)
		}
	}
	if _, err = outFile.WriteString(end); err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d datasources\n", len(doc.Datasources))
}
