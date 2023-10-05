package crawler

import (
	"encoding/json"
	"strings"

	"github.com/gosimple/slug"
)

func GetStoriesPagination(typeStr string, page int) (ResponseListStory, error) {
	req, err := AllStory(typeStr, page)
	if err != nil {
		return ResponseListStory{}, err
	}

	var respListStory ResponseListStory
	err = json.Unmarshal(req, &respListStory)
	if err != nil {
		return ResponseListStory{}, err
	}

	respListStory.Meta.Pagination.Links.Next = ""

	return respListStory, nil
}

func GetStoryDetail(storyID int) (ResponseDetailStory, error) {
	req, err := DetailStory(storyID)
	if err != nil {
		return ResponseDetailStory{}, err
	}

	var respDetailStory ResponseDetailStory
	err = json.Unmarshal(req, &respDetailStory)
	if err != nil {
		return ResponseDetailStory{}, err
	}

	return respDetailStory, nil
}

func GetStoryRelatedPagination(storyID int, page int) (ResponseListStory, error) {
	crawlStoryData, err := GetStoryDetail(storyID)
	if err != nil {
		return ResponseListStory{}, err
	}

	categories := strings.Split(crawlStoryData.Data.Categories, ",")
	categorySlug := slug.Make(categories[0])

	req, err := CategoryStory(categorySlug, page)
	if err != nil {
		return ResponseListStory{}, err
	}

	var respListStory ResponseListStory
	err = json.Unmarshal(req, &respListStory)
	if err != nil {
		return ResponseListStory{}, err
	}

	respListStory.Meta.Pagination.Links.Next = ""
	return respListStory, nil
}
