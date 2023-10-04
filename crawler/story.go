package crawler

import (
	"encoding/json"
	"strings"

	"github.com/gosimple/slug"
)

func GetStoriesPagination(typeStr string, page int) (CrawlList, error) {
	req, err := AllStory(typeStr, page)
	if err != nil {
		return CrawlList{}, err
	}

	var crawlList CrawlList
	err = json.Unmarshal(req, &crawlList)
	if err != nil {
		return CrawlList{}, err
	}

	crawlList.Meta.Pagination.Links.Next = ""

	return crawlList, nil
}

func GetStoryDetail(storyID int) (CrawlStoryData, error) {
	req, err := DetailStory(storyID)
	if err != nil {
		return CrawlStoryData{}, err
	}

	var crawlStoryData CrawlStoryData
	err = json.Unmarshal(req, &crawlStoryData)
	if err != nil {
		return CrawlStoryData{}, err
	}

	return crawlStoryData, nil
}

func GetStoryRelatedPagination(storyID int, page int) (CrawlList, error) {
	crawlStoryData, err := GetStoryDetail(storyID)
	if err != nil {
		return CrawlList{}, err
	}

	categories := strings.Split(crawlStoryData.Data.Categories, ",")
	categorySlug := slug.Make(categories[0])

	req, err := CategoryStory(categorySlug, page)
	if err != nil {
		return CrawlList{}, err
	}

	var crawlList CrawlList
	err = json.Unmarshal(req, &crawlList)
	if err != nil {
		return CrawlList{}, err
	}

	crawlList.Meta.Pagination.Links.Next = ""
	return crawlList, nil
}
