package entities

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Renderer interface {
	GetRenderedContent(raw []byte) string
}

type Article struct {
	rawContent []byte
	Metadata   *ArticleMeta
	renderer   Renderer
}

// Has to be the folder of the article containing the article.md file
// Otherwise returns nil
func ArticleFrom(dirEntry os.DirEntry, root string, renderer Renderer) *Article {
	if !dirEntry.IsDir() {
		return nil
	}

	content, err := os.ReadFile(fmt.Sprintf("%s/article.md", root))
	if err != nil {
		return nil
	}

	var metadata *ArticleMeta

	rawMetadata, err := os.ReadFile(fmt.Sprintf("%s/metadata.json", root))

	if err != nil {
		metadata = generateDefaultMetadata(dirEntry, root)
	} else {
		err := json.Unmarshal(rawMetadata, &metadata)

		if err != nil {
			metadata = generateDefaultMetadata(dirEntry, root)
		}
	}

	return &Article{
		rawContent: content,
		Metadata:   metadata,
		renderer:   renderer,
	}
}

func (article *Article) GetRenderedContent() string {
	return article.renderer.GetRenderedContent(article.rawContent)
}

// Default metadata assumes (mod time == creation time)
func generateDefaultMetadata(dirEntry os.DirEntry, root string) *ArticleMeta {
	inf, err := dirEntry.Info()

	if err != nil {
		log.Printf("Cannot load folder info for %s\nNo metadata commited", root)
		return nil
	}

	modTime := inf.ModTime()

	meta := &ArticleMeta{
		Created:  modTime.Unix(),
		Modified: modTime.Unix(),
		Name:     inf.Name(),
		Path:     inf.Name(),
	}

	err = meta.WriteMeta(root)

	if err != nil {
		log.Printf("Failed tot write metadata to root %s\nErr: %s", root, err)
	}

	return meta
}
