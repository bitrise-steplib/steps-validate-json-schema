package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	schemas "github.com/bitrise-io/bitrise-json-schemas"
	"github.com/bitrise-io/go-steputils/stepconf"
	"github.com/bitrise-io/go-utils/env"
	"github.com/bitrise-io/go-utils/filedownloader"
	"github.com/bitrise-io/go-utils/log"
)

type Inputs struct {
	SchemaURL       string `env:"schema_url,required"`
	YAMLPath        string `env:"yaml_path,required"`
	WarningPatterns string `env:"warning_patterns,required"`
}

func main() {
	var inputs Inputs
	if err := stepconf.NewInputParser(env.NewRepository()).Parse(&inputs); err != nil {
		panic(err)
	}
	stepconf.Print(inputs)

	fmt.Println()
	log.Infof("Validating: ", inputs.YAMLPath)

	var schema string
	downloader := filedownloader.New(http.DefaultClient)
	if strings.HasPrefix(inputs.SchemaURL, "http") {
		b, err := downloader.GetRemoteContents(inputs.SchemaURL)
		if err != nil {
			panic(err)
		}
		schema = string(b)
	} else if strings.HasPrefix(inputs.SchemaURL, "file://") {
		pth := strings.TrimPrefix(inputs.SchemaURL, "file://")
		b, err := downloader.ReadLocalFile(pth)
		if err != nil {
			panic(err)
		}
		schema = string(b)
	} else {
		panic("invalid schema path, should start with http or file://")
	}

	validator, err := schemas.NewJSONSchemaValidator(schema)
	if err != nil {
		panic(err)
	}

	var warningPatters []string
	split := strings.Split(inputs.WarningPatterns, "\n")
	for _, s := range split {
		if s != "" {
			warningPatters = append(warningPatters, s)
		}
	}

	warnings, errors, err := validator.Validate(inputs.YAMLPath, warningPatters...)
	if err != nil {
		panic(err)
	}

	log.Warnf("Warnings: ", len(warnings))
	for _, warning := range warnings {
		fmt.Println("- ", warning)
	}

	log.Errorf("Errors: ", len(errors))
	for _, error := range errors {
		fmt.Println("- ", error)
	}

	if len(errors) > 0 {
		log.Errorf("Invalid step.yml")
		os.Exit(1)
	}
	log.Donef("Valid step.yml")

}
