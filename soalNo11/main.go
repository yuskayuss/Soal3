package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type Article struct {
	Title       *string `json:"title"`
	StoryTitle  *string `json:"story_title"`
	NumComments *int    `json:"num_comments"`
}

type Response struct {
	Page    int       `json:"page"`
	PerPage int       `json:"per_page"`
	Total   int       `json:"total"`
	Data    []Article `json:"data"`
}

func topArticles(author string, limit int) ([]string, error) {
	var allArticles []Article
	page := 1

	for {
		url := fmt.Sprintf("https://hackerrank/articles?author=%s&page=%d", author, page)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}

		var response Response
		if err := json.Unmarshal(body, &response); err != nil {
			return nil, err
		}

		if len(response.Data) == 0 {
			break
		}

		allArticles = append(allArticles, response.Data...)

		
		if page*response.PerPage >= response.Total {
			break
		}
		page++
	}

	
	filtered := []Article{}
	for _, art := range allArticles {
		if (art.Title != nil && strings.TrimSpace(*art.Title) != "") ||
			(art.StoryTitle != nil && strings.TrimSpace(*art.StoryTitle) != "") {
			filtered = append(filtered, art)
		}
	}

	
	sort.Slice(filtered, func(i, j int) bool {
		numCommentsI := 0
		if filtered[i].NumComments != nil {
			numCommentsI = *filtered[i].NumComments
		}
		numCommentsJ := 0
		if filtered[j].NumComments != nil {
			numCommentsJ = *filtered[j].NumComments
		}

		if numCommentsI != numCommentsJ {
			return numCommentsI > numCommentsJ
		}

		
		getTitle := func(a Article) string {
			if a.Title != nil && strings.TrimSpace(*a.Title) != "" {
				return *a.Title
			}
			return *a.StoryTitle
		}

		return getTitle(filtered[i]) < getTitle(filtered[j])
	})

	
	result := []string{}
	for i, art := range filtered {
		if i >= limit {
			break
		}
		if art.Title != nil && strings.TrimSpace(*art.Title) != "" {
			result = append(result, *art.Title)
		} else {
			result = append(result, *art.StoryTitle)
		}
	}

	return result, nil
}

func main() {
	articles, err := topArticles("someauthor", 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, title := range articles {
		fmt.Println(title)
	}
}
