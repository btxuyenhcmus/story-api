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

func FetchChapter(storyID int, page int) (CrawlListChapter, error) {
	req, err := FetchStory(storyID, page)
	if err != nil {
		return CrawlListChapter{}, err
	}

	var crawlListChapter CrawlListChapter
	err = json.Unmarshal(req, &crawlListChapter)
	if err != nil {
		return CrawlListChapter{}, err
	}

	crawlListChapter.Meta.Pagination.Links.Next = ""

	return crawlListChapter, nil
}
