package agent

type Agent struct {
	Outputs        map[string]Output
	Inputs         []Input
	ManagementMode string
}

type Output struct {
	Type string
}

type Elasticsearch struct {
	Hosts    []string
	ApiKey   string
	Username string
	Password string
}

type Input struct {
	Id string
}
