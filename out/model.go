package out

type (
	version struct {
		Ref string `json:"ref"`
	}

	InputJSON struct {
		Params  map[string]string `json:"params"`
		Source  map[string]string `json:"source"`
		Version version           `json:"version"`
		// in case of using Github actions
		Ci string `json:"ci"`
		ResolversFile  string `json:"resolvers_file"`
		SchemaFile  string 	`json:"schema_file"`
		AccessKeyId  string `json:"access_key_id"`
		SecretAccessKey  string `json:"secret_access_key"`
		SessionToken  string `json:"session_token"`
		ApiID  string `json:"api_id"`
		RegionName  string `json:"region_name"`
	}

	metadata struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	checkOutputJSON []version
	inOutputJSON    struct {
		Version  version    `json:"version"`
		Metadata []metadata `json:"metadata,omitempty"`
	}
	outOutputJSON inOutputJSON
)
