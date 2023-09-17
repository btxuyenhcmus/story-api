package crawler

import "encoding/json"

func GetStoriesPagination(typeStr string, page int) (CrawlList, error) {
	req, err := GetRequestList(typeStr, page)
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
