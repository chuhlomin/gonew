// GitHub Action that does something
package main

import (
	"bytes"
	"fmt"
	"os"
	_ "time/tzdata"

	"github.com/caarlos0/env/v9"
)

// type vars map[string]interface{}

type config struct {
	Key  string `env:"INPUT_KEY"`
	Key2 string `env:"INPUT_KEY2" envDefault:".kube.yml"`
	// Vars vars   `env:"INPUT_VARS"` // custom parser
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("::error::%v", err)
		os.Exit(1)
	}
}

func run() error {
	var c config
	// parsers := map[reflect.Type]env.ParserFunc{
	// 	reflect.TypeOf(vars{}): varsParser,
	// }
	// if err := env.ParseWithFuncs(&c, parsers); err != nil {
	// 	return err
	// }

	if err := env.Parse(&c); err != nil {
		return err
	}

	output, err := doSomething(c.Key)
	if err != nil {
		return fmt.Errorf("failed to render template: %v", err)
	}

	if err = writeOutput(output); err != nil {
		return err
	}

	return nil
}

// func varsParser(v string) (interface{}, error) {
// 	m := map[string]interface{}{}
// 	err := yaml.Unmarshal([]byte(v), &m)
// 	if err != nil {
// 		return nil, fmt.Errorf("unable to parse Vars: %v", err)
// 	}
// 	return m, nil
// }

func doSomething(key string) (string, error) {
	return key, nil
}

func writeOutput(output string) error {
	githubOutput := formatOutput("result", output)
	if githubOutput == "" {
		return nil
	}

	path := os.Getenv("GITHUB_OUTPUT")

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf(
			"failed to open result file %q: %v. "+
				"If you are using self-hosted runners "+
				"make sure they are updated to version 2.297.0 or greater",
			path,
			err,
		)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("failed to close result file %q: %v", path, err)
		}
	}()

	if _, err = f.WriteString(githubOutput); err != nil {
		return fmt.Errorf("failed to write result to file %q: %v", path, err)
	}

	return nil
}

func formatOutput(name, value string) string {
	if value == "" {
		return ""
	}

	// if value contains new line, use multiline format
	if bytes.ContainsRune([]byte(value), '\n') {
		return fmt.Sprintf("%s<<OUTPUT\n%s\nOUTPUT", name, value)
	}

	return fmt.Sprintf("%s=%s", name, value)
}
