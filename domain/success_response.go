package domain

type SuccessResponse struct {
	Message string `json:"message"`
}

type StorySuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
