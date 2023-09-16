package crawler

import (
	"fmt"
	"io"
	"net/http"
)

const HOST_TEMPLATE_URL string = "https://api.truyenfull.vn/v1/%s/%s"

func GetRequest(storyID int, page int) ([]byte, error) {
	param := fmt.Sprintf("download?page=%d", page)
	path := fmt.Sprintf("story/detail/%d", storyID)
	url := fmt.Sprintf(HOST_TEMPLATE_URL, path, param)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpReponse, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(httpReponse.Body)
	if err != nil {
		return nil, err
	}

	return responseData, err
}
