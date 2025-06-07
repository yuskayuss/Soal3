package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

type Article struct {
	Title       *string `json:"title"`
	StoryTitle  *string `json:"story_title"`
	NumComments *int    `json:"num_comments"`
}

type ApiResponse struct {
	Page       int       `json:"page"`
	PerPage    int       `json:"per_page"`
	Total      int       `json:"total"`
	TotalPages int       `json:"total_pages"`
	Data       []Article `json:"data"`
}

type NamedArticle struct {
	Name        string
	NumComments int
}

// Main function to test
func main() {
	articles := topArticles("olalonde", 1)
	for _, a := range articles {
		fmt.Println(a)
	}
}

func topArticles(author string, limit int) []string {
	var allArticles []NamedArticle
	page := 1

	for {
		url := "https://jsonmock.hackerrank.com/api/articles?author=" + author + "&page=" + strconv.Itoa(page)
		resp, err := http.Get(url)
		if err != nil {
			break
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		var apiResp ApiResponse
		json.Unmarshal(body, &apiResp)

		for _, article := range apiResp.Data {
			var title string
			if article.Title != nil {
				title = *article.Title
			} else if article.StoryTitle != nil {
				title = *article.StoryTitle
			} else {
				continue // skip if no valid title
			}

			comments := 0
			if article.NumComments != nil {
				comments = *article.NumComments
			}

			allArticles = append(allArticles, NamedArticle{
				Name:        title,
				NumComments: comments,
			})
		}

		if page >= apiResp.TotalPages {
			break
		}
		page++
	}

	// Sort by NumComments descending, then Name ascending
	sort.SliceStable(allArticles, func(i, j int) bool {
		if allArticles[i].NumComments == allArticles[j].NumComments {
			return allArticles[i].Name < allArticles[j].Name
		}
		return allArticles[i].NumComments > allArticles[j].NumComments
	})

	// Get top N articles
	result := []string{}
	for i := 0; i < limit && i < len(allArticles); i++ {
		result = append(result, allArticles[i].Name)
	}

	return result
}
