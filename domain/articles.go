package articles

import (
	"log"
	"main/services"
	"os"
	"sort"
)

func GetArticle(name string) []byte {
	content, err := os.ReadFile("./articles/" + name + "/article.md")

	if err != nil {
		log.Println("Could not find file: ", name)
	}

	return content
}

func GetArticles() [][]byte {
	articleFolders, err := os.ReadDir("./articles")

	if err != nil {
		log.Println(err)
	}

	articleFolders = services.FilterForFolders(articleFolders)

	sort.Slice(articleFolders, func(a, b int) bool {
		aInfo, err := articleFolders[a].Info()

		if err != nil {
			log.Println(err)
		}

		bInfo, err := articleFolders[b].Info()

		aTime := services.GetFileChangeTime(aInfo)
		bTime := services.GetFileChangeTime(bInfo)

		return aTime.Unix() > bTime.Unix()
	})

	articles := [][]byte{}

	for _, file := range articleFolders {
		content, err := os.ReadFile("./articles/" + file.Name() + "/article.md")

		if err != nil {
			log.Println(err)
		}

		articles = append(articles, content)
	}

	return articles
}

func GetArticleNames() []string {
	articleFolders, err := os.ReadDir("./articles")

	if err != nil {
		log.Println(err)
	}

	articleFolders = services.FilterForFolders(articleFolders)

	sort.Slice(articleFolders, func(a, b int) bool {
		aInfo, err := articleFolders[a].Info()

		if err != nil {
			log.Println(err)
		}

		bInfo, err := articleFolders[b].Info()

		aTime := services.GetFileChangeTime(aInfo)
		bTime := services.GetFileChangeTime(bInfo)

		return aTime.Unix() > bTime.Unix()
	})

	articleNames := []string{}

	for _, articleFodler := range articleFolders {
		articleNames = append(articleNames, articleFodler.Name())
	}

	return articleNames
}
