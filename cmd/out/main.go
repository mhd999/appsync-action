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

	whichCi := os.Getenv("INPUT_CI")
	log.Printf(" which ci: %s", whichCi)

	if whichCi == "github" {
		sessionToken := os.Getenv("INPUT_SESSION_TOKEN")
		secretAccessKey := os.Getenv("INPUT_SECRET_ACCESS_KEY")
		accessKeyId := os.Getenv("INPUT_ACCESS_KEY_ID")
		regionName := os.Getenv("INPUT_REGION_NAME")
		apiID := os.Getenv("INPUT_API_ID")
		schemaFile := os.Getenv("INPUT_SCHEMA_FILE")
		resolversFile := os.Getenv("INPUT_RESOLVERS_FILE")

		log.Printf(" build source and param from github sessionToken: %s", sessionToken)
		log.Printf(" build source and param from github secretAccessKey: %s", secretAccessKey)
		log.Printf(" build source and param from github accessKeyId: %s", accessKeyId)
		log.Printf(" build source and param from github regionName: %s", regionName)
		log.Printf(" build source and param from github apiID: %s", apiID)
		log.Printf(" build source and param from github schemaFile: %s", schemaFile)
		log.Printf(" build source and param from github resolversFile: %s", resolversFile)

		input.Source = make(map[string]string)
		input.Params = make(map[string]string)
		input.Source["api_id"] = apiID
		input.Source["access_key_id"] = accessKeyId
		input.Source["secret_access_key"] = secretAccessKey
		input.Source["session_token"] = sessionToken
		input.Source["region_name"] = regionName
		input.Params["schema_file"] = schemaFile
		input.Params["resolvers_file"] = resolversFile

		log.Printf(" build source and param from github input.Source: %v", input.Source)
		log.Printf(" build source and param from github input.Params: %v", input.Params)
	} else if err := decoder.Decode(&input); err != nil {
		logger.Fatalf("Failed to decode to stdin: %s", err)
	}

	output, err := out.Command(input, logger)
	if err != nil {
		logger.Fatalf("Error execute out command: %s", err)
	}

	if err := createOutput(output, encoder, logger); err != nil {
		logger.Fatalf("Failed to encode to stdout: %s", err)
	}

}
