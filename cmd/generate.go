package cmd

import (
	"flag"
	"fmt"
	"git.jcu.edu.au/cft/cfds/crypt"
	"git.jcu.edu.au/cft/cfds/datasources"
	"github.com/ansel1/merry"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	input  string
	output string
)

func init() {
	generateCmd.Flags().StringVarP(&input, "input", "i", "datasources.yml", "")
	generateCmd.Flags().StringVarP(&output, "output", "o", "neo-datasource.xml", "")
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the neo-datasource.xml file that contains the Coldfusion datasources.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generateXmlFile()
	},
}

func generateXmlFile() error {
	if seed == "" {
		return merry.New("Requires seed parameter")
	}
	templateDir := "templates"
	flag.Parse()
	if seed == "" || input == "" {
		flag.Usage()
	}
	inFile, err := os.Open(input)
	if err != nil {
		return merry.Wrap(err)
	}
	defer inFile.Close()
	d := yaml.NewDecoder(inFile)
	doc := datasources.NewDoc()
	err = d.Decode(doc)
	if err != nil {
		return merry.New("Error reading datasource file").WithCause(err)
	}

	start := "<wddxPacket version='1.0'><header/><data><array length='2'><struct type='coldfusion.server.ConfigMap'>"
	end := "</struct><struct type='coldfusion.server.ConfigMap'><var name='maxcachecount'><number>100.0</number></var></struct></array></data></wddxPacket>"

	outFile, err := os.Create(output)
	if err != nil {
		return merry.Wrap(err)
	}
	defer outFile.Close()

	templates := map[string]*template.Template{}

	if _, err = outFile.WriteString(start); err != nil {
		return merry.Wrap(err)
	}
	for ds, data := range doc.Datasources {
		data.Name = ds
		data.Password, err = crypt.Encrypt(data.Password, seed)
		if err != nil {
			return merry.New("Could not encrypt password").WithCause(err)
		}
		key := strings.ToLower(data.Type)
		templateFile := filepath.Join(templateDir, key+".xml")
		if _, ok := templates[key]; !ok {
			// Load template
			content, err := ioutil.ReadFile(templateFile)
			if err != nil {
				return merry.New("Cannot load template " + templateFile).WithCause(err)
			}
			templates[key], err = template.New(key).Parse(string(content))
			if err != nil {
				return merry.Wrap(err)
			}
		}
		err = templates[key].Execute(outFile, data)
		if err != nil {
			return merry.Wrap(err)
		}
	}
	if _, err = outFile.WriteString(end); err != nil {
		return merry.Wrap(err)
	}
	fmt.Printf("Wrote %d datasources\n", len(doc.Datasources))
	return nil
}
