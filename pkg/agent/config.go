package agent

// Defines the Elastic Agent configuration.
// QUESTION: I'm worried that users will confuse this with settings under "agent."
// like agent.logging.
type Agent struct {

	//   description: |
	//     One or more output configurations that specify where to send data.
	//     You can pair specific inputs with specific outputs.
	Outputs        map[string]Output `yaml:"outputs"`

	//   description: |
	//     List of input configurations that specify how to locate and processes data.
	//     By default Elastic Agent collects system metrics, such as cpu, memory, network, and filesystem metrics, and sends them to the default output. 
	Inputs         []Input `yaml:"inputs"`

	//   description: |
	//     The management mode.
	//   values:
	//     - local
	//     - whatelse
	//   default: local
	ManagementMode string `yaml:"mode"`
}

// Defines an Elastic Agent output configuration.
// Can be an Elasticsearch or Logstash output configuration.
type Output struct { `yaml:"type"`
	Type string
}

// Defines the Elasticsearch output configuration.
//
type Elasticsearch struct {

	//   description: |
	//     A list of Elasticsearch nodes to connect to.
	//     The events are distributed to these nodes in round robin order.
	//     If one node becomes unreachable, the event is automatically sent to another node.	
	Hosts    []string `yaml:"hosts"`


	//   description: |
	//     The API key to use for token-based authentication.
	ApiKey   string `yaml:"api-key"`

	//   description: |
	//     The basic authentication username for connecting to Elasticsearch.
	Username string `yaml:"username"`

	//   description: |
	//     The basic authentication password for connecting to Elasticsearch.
	Password string `yaml:"password"`
}

// Defines an Elastic Agent input configuration.
//
type Input struct {
	Id string
}
