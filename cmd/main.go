package main

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var temp *template.Template
var logger *zap.SugaredLogger

type API struct {
	Endpoint    string      `yaml:"endpoint"`
	Verb        string      `yaml:"verb"`
	Response    interface{} `yaml:"response"`
	ContentType string      `yaml:"content_type"`
	StatusCode  int         `yaml:"status_code"`
	Headers     []string    `yaml:"headers"`
}

type Config struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
	API  []API  `yaml:"api"`
}

func trim(s string) string {
	return strings.ReplaceAll(strings.Trim(strings.TrimSpace(s), "\"\\"), "\n", "")
}

// Using the init function to make sure the template is only parsed once in the program
func init() {
	l, _ := zap.NewProduction()
	defer l.Sync()
	logger = l.Sugar()

	// template.Must takes the reponse of template.ParseFiles and does error checking
	temp = template.Must(template.New("api.tmpl").Funcs(template.FuncMap{
		"trim": trim,
	}).ParseFiles("./templates/api.tmpl"))
}

func main() {
	// Build the path:
	outputPath := filepath.Join("./", "cmd/api/main.go")

	// Create the file:
	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	defer f.Close() // don't forget to close the file when finished.

	// Read the file
	data, err := os.ReadFile("api.yaml")
	if err != nil {
		logger.Errorln(err)
		return
	}

	// Create a struct to hold the YAML data
	var config Config

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.Errorln(err)
		return
	}

	// Execute myName into the template and print to Stdout
	err = temp.Execute(f, config)
	if err != nil {
		logger.Fatalln(err)
	}
}
