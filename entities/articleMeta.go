package entities

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type ArticleMeta struct {
	Created  int64  `from:"created"`  // Seconds
	Modified int64  `from:"modified"` // Seconds
	Name     string `from:"name"`
	Path     string `from:"path"`
}

func (metadata *ArticleMeta) GetCreated() time.Time {
	return time.Unix(metadata.Created, 0)
}

func (metadata *ArticleMeta) GetModified() time.Time {
	return time.Unix(metadata.Modified, 0)
}

func (metadata *ArticleMeta) WriteMeta(target string) error {
	marshaled, err := json.Marshal(metadata)

	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("%s/metadata.json", target), marshaled, 0644)
}
