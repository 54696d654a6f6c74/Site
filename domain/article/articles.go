package articles

import (
	"fmt"
	"log"
	entities "main/entities"
	"main/services"
	"os"
	"slices"
)

type mdArticleRenderer struct{}

func (r mdArticleRenderer) GetRenderedContent(raw []byte) string {
	return string(services.MdToHTML(raw))
}

func GetArticle(name string) *entities.Article {
	articles, err := os.ReadDir("./articles/")

	if err != nil {
		return nil
	}

	targetArticleIndex := slices.IndexFunc(articles, func(e os.DirEntry) bool {
		return e.Name() == name
	})

	if targetArticleIndex < 0 {
		return nil
	}

	return entities.ArticleFrom(articles[targetArticleIndex], fmt.Sprintf("./articles/%s", name), mdArticleRenderer{})
}

func GetArticles() []*entities.Article {
	articleFolders, err := os.ReadDir("./articles")

	if err != nil {
		log.Println(err)
	}

	articleFolders = services.FilterForFolders(articleFolders, false)

	articles := []*entities.Article{}

	for _, file := range articleFolders {
		article := GetArticle(file.Name())
		if article == nil {
			log.Printf("Could not find article - %s", file.Name())
		}

		articles = append(articles, article)
	}

	sortArticles(articles)

	return articles
}

func sortArticles(articles []*entities.Article) {
	slices.SortStableFunc(articles, func(a, b *entities.Article) int {
		diff := b.Metadata.Modified - a.Metadata.Modified

		if diff > 0 {
			return 1
		} else if diff < 0 {
			return -1
		}
		return 0
	})
}
