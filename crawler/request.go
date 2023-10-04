package crawler

import (
	"fmt"
	"io"
	"net/http"
)

const HOST_TEMPLATE_URL string = "https://api.truyenfull.vn/v1/%s/%s"

func DownloadStory(storyID int, page int) ([]byte, error) {
	param := fmt.Sprintf("download?page=%d", page)
	path := fmt.Sprintf("story/detail/%d", storyID)
	url := fmt.Sprintf(HOST_TEMPLATE_URL, path, param)

	return getRequest(url)
}

func DetailStory(storyID int) ([]byte, error) {
	param := fmt.Sprintf("detail/%d", storyID)
	url := fmt.Sprintf(HOST_TEMPLATE_URL, "story", param)

	return getRequest(url)
}

func CategoryStory(category string, page int) ([]byte, error) {
	param := fmt.Sprintf("cate?cate=%s&type=story_new&page=%d", category, page)
	url := fmt.Sprintf(HOST_TEMPLATE_URL, "story", param)

	return getRequest(url)
}

func AllStory(typeStr string, page int) ([]byte, error) {
	param := fmt.Sprintf("all?type=%s&page=%d", typeStr, page)
	url := fmt.Sprintf(HOST_TEMPLATE_URL, "story", param)
	return getRequest(url)
}

func getRequest(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpResponse, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return responseData, err
}
