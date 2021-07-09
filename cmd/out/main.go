package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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

	if whichCi == "github" {
		log.Println("Start listing files")
		files, err := ioutil.ReadDir("/github/workspace")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			log.Println(f.Name())
		}

		var gihubFiles []string

		root := "github"
		errg := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			gihubFiles = append(gihubFiles, path)
			return nil
		})
		if errg != nil {
			panic(errg)
		}
		log.Println("files in github")
		for _, file := range gihubFiles {
			log.Printf("The file %s", file)
		}

		sessionToken := os.Getenv("INPUT_SESSION_TOKEN")
		secretAccessKey := os.Getenv("INPUT_SECRET_ACCESS_KEY")
		accessKeyId := os.Getenv("INPUT_ACCESS_KEY_ID")
		regionName := os.Getenv("INPUT_REGION_NAME")
		apiID := os.Getenv("INPUT_API_ID")
		schemaFile := os.Getenv("INPUT_SCHEMA_FILE")
		resolversFile := os.Getenv("INPUT_RESOLVERS_FILE")

		input.Source = map[string]string{
			"api_id":            apiID,
			"access_key_id":     accessKeyId,
			"secret_access_key": secretAccessKey,
			"session_token":     sessionToken,
			"region_name":       regionName,
		}

		input.Params = map[string]string{
			"schema_file":    schemaFile,
			"resolvers_file": resolversFile,
		}
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
