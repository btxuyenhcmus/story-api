package crawler

import "encoding/json"

func DownloadChapter(storyID int, page int) (CrawlDetail, error) {
	req, err := DownloadStory(storyID, page)
	if err != nil {
		return CrawlDetail{}, err
	}

	var crawlDetail CrawlDetail
	err = json.Unmarshal(req, &crawlDetail)
	if err != nil {
		return CrawlDetail{}, err
	}

	crawlDetail.Meta.Pagination.Links.Next = ""

	return crawlDetail, nil
}
