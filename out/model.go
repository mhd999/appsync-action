package out

type (
	version struct {
		Ref string `json:"ref"`
	}

	InputJSON struct {
		Params  map[string]string `json:"params,omitempty"`
		Source  map[string]string `json:"source,omitempty"`
		Version version           `json:"version,omitempty"`
		// in case of using Github actions
		Ci              string `json:"ci,omitempty"`
		ResolversFile   string `json:"resolvers_file,omitempty"`
		SchemaFile      string `json:"schema_file,omitempty"`
		AccessKeyId     string `json:"access_key_id,omitempty"`
		SecretAccessKey string `json:"secret_access_key,omitempty"`
		SessionToken    string `json:"session_token,omitempty"`
		ApiID           string `json:"api_id,omitempty"`
		RegionName      string `json:"region_name,omitempty"`
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
