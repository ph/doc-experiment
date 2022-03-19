package typ

// Defines the datastream used by an input.
//
type DataStream struct {
	
	//   description: |
	//     The type of blah blah blah.
	Type      string `yaml:"type"`
	
	//   description: |
	//     The dataset to use to fetch and structure data.
	//     The dataset name must conform to the naming conventions for Elasticsearch indices.
	//     The name cannot contain dashes (-) or exceed 100 bytes.
	DataSet   string `yaml:"dataset"`
	
	//   description: |
	//     A user-configurable arbitrary grouping, such as an environment (dev, prod, or qa), a team, or a strategic business unit.
	//     A namespace can be up to 100 bytes in length (multibyte characters will count toward this limit faster).
	Namespace string `yaml:"namespace"`
}
