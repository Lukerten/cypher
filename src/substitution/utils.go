package substitution

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type SubstitutionKey struct {
	Char  string `yaml:"char"`
	Value string `yaml:"value"`
}

func LoadSubstitutionKeys(keysPath string) ([]SubstitutionKey, error) {
	data, err := os.ReadFile(keysPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read keys file: %w", err)
	}

	var keys []SubstitutionKey
	err = yaml.Unmarshal(data, &keys)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal keys: %w", err)
	}

	return keys, nil
}
