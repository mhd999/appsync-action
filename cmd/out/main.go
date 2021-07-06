package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/telia-oss/appsync-resource/out"
)

func createOutput(output interface{}, encoder *json.Encoder, logger *log.Logger) error {
	return encoder.Encode(output)
}

func main() {

	var (
		input   out.InputJSON
		decoder = json.NewDecoder(os.Stdin)
		encoder = json.NewEncoder(os.Stdout)
		logger  = log.New(os.Stderr, "resource:", log.Lshortfile)
	)
	logger.Println("input:", decoder)
	if err := decoder.Decode(&input); err != nil {
		logger.Fatalf("Failed to decode to stdin: %s", err)
	}

	if input.Ci == "github" {
		input.Source = make(map[string]string)
		input.Params = make(map[string]string)
		input.Source["api_id"] = input.ApiID
		input.Source["access_key_id"] = input.AccessKeyId
		input.Source["secret_access_key"] = input.SecretAccessKey
		input.Source["session_token"] = input.SessionToken
		input.Source["region_name"] = input.RegionName
		input.Source["session_token"] = input.SessionToken
		input.Params["schema_file"] = input.SchemaFile
		input.Params["resolvers_file"] = input.ResolversFile
	}

	output, err := out.Command(input, logger)
	if err != nil {
		logger.Fatalf("Error execute out command: %s", err)
	}

	if err := createOutput(output, encoder, logger); err != nil {
		logger.Fatalf("Failed to encode to stdout: %s", err)
	}

}
