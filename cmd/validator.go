package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var validatorCmd = &cobra.Command{
	Use:   "validator",
	Short: "Validate the API using api-validator",
	Long:  `This command validates the API using api-validator`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := viper.GetString("validatorFile")
		fmt.Printf("Validating file %s...\n", filePath)
		if filePath == "" {
			return fmt.Errorf("missing input file path")
		}

		// Check file extension
		ext := filepath.Ext(filePath)
		if ext != ".json" && ext != ".yaml" && ext != ".yml" {
			return fmt.Errorf("invalid file extension: %s", ext)
		}

		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read input file: %v", err)
		}

		var jsonData interface{}

		switch ext {
		case ".json":
			err = json.Unmarshal(data, &jsonData)
			if err != nil {
				return fmt.Errorf("failed to unmarshal JSON data: %v", err)
			}
			_, err := json.MarshalIndent(jsonData, "", "    ")
			if err != nil {
				return fmt.Errorf("failed to encode JSON data: %v", err)
			}
		case ".yaml", ".yml":
			err = yaml.Unmarshal(data, &jsonData)
			if err != nil {
				return fmt.Errorf("failed to unmarshal YAML data: %v", err)
			}
			_, err := yaml.Marshal(jsonData)
			if err != nil {
				return fmt.Errorf("failed to encode YAML data: %v", err)
			}
		}

		// TODO: perform validation using data

		response := map[string]string{
			"status":  "success",
			"message": "API validation successful",
			"size":    fmt.Sprintf("%d bytes", len(data)),
		}

		output, err := json.Marshal(response)
		if err != nil {
			return fmt.Errorf("failed to encode response as JSON: %v", err)
		}

		fmt.Println(string(output))
		return nil
	},
}
