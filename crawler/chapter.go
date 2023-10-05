package crawler

import "encoding/json"

func DownloadChapter(storyID int, page int) (ResponseDownloadStory, error) {
	req, err := DownloadStory(storyID, page)
	if err != nil {
		return ResponseDownloadStory{}, err
	}

	var respDownloadStory ResponseDownloadStory
	err = json.Unmarshal(req, &respDownloadStory)
	if err != nil {
		return ResponseDownloadStory{}, err
	}

	respDownloadStory.Meta.Pagination.Links.Next = ""

	return respDownloadStory, nil
}

func ListChapter(storyID int, page int) (ResponseListChapter, error) {
	req, err := FetchStory(storyID, page)
	if err != nil {
		return ResponseListChapter{}, err
	}

	var respListChapter ResponseListChapter
	err = json.Unmarshal(req, &respListChapter)
	if err != nil {
		return ResponseListChapter{}, err
	}

	respListChapter.Meta.Pagination.Links.Next = ""

	return respListChapter, nil
}

func FetchChapter(chapterID int) (ResponseDetailChapter, error) {
	req, err := DetailChapter(chapterID)
	if err != nil {
		return ResponseDetailChapter{}, err
	}

	var respDetailChapter ResponseDetailChapter
	err = json.Unmarshal(req, &respDetailChapter)
	if err != nil {
		return ResponseDetailChapter{}, err
	}

	return respDetailChapter, nil
}
