package yaml

import (
	"fmt"
	"strings"

	yml "github.com/ghodss/yaml"
	jsoniter "github.com/json-iterator/go"
)

// Marshal the object into JSON then converts JSON to YAML and returns the
// YAML.
func Marshal(o interface{}) ([]byte, error) {
	j, err := jsoniter.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("error marshaling into JSON: %v", err)
	}

	y, err := JSONToYAML(j)
	if err != nil {
		return nil, fmt.Errorf("error converting JSON to YAML: %v", err)
	}

	return y, nil
}

// JSONToYAML Convert JSON to YAML.
func JSONToYAML(j []byte) ([]byte, error) {
	return yml.JSONToYAML(j)
}

// Unmarshal converts YAML to JSON then uses JSON to unmarshal into an object,
func Unmarshal(y []byte, o interface{}) error {
	j, err := YAMLToJSON(y)
	if err != nil {
		return fmt.Errorf("error converting YAML to JSON: %v", err)
	}

	err = jsoniter.Unmarshal(j, o)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return nil
}

// YAMLToJSON Convert YAML to JSON. Since JSON is a subset of YAML, passing JSON through
// this method should be a no-op.
//
// Things YAML can do that are not supported by JSON:
//   - In YAML you can have binary and null keys in your maps. These are invalid
//     in JSON. (int and float keys are converted to strings.)
//   - Binary data in YAML with the !!binary tag is not supported. If you want to
//     use binary data with this library, encode the data as base64 as usual but do
//     not use the !!binary tag in your YAML. This will ensure the original base64
//     encoded data makes it all the way through to the JSON.
func YAMLToJSON(y []byte) ([]byte, error) {
	return yml.YAMLToJSON(y)
}

func IsYAMLSuffix(filename string) bool {
	return strings.HasSuffix(filename, ".yaml") || strings.HasSuffix(filename, ".yml")
}
