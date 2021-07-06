package out

type (
	version struct {
		Ref string `json:"ref"`
	}

	InputJSON struct {
		Params  map[string]string `json:"params,omitempty"`
		Source  map[string]string `json:"source,omitempty"`
		Version version           `json:"version,omitempty"`
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
