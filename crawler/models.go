package crawler

type CrawlLinks struct {
	Next string `json:"next"`
}

type CrawlPagination struct {
	Total       int        `json:"total"`
	Count       int        `json:"count"`
	PerPage     int        `json:"per_page"`
	CurrentPage int        `json:"current_page"`
	TotalPages  int        `json:"total_pages"`
	Links       CrawlLinks `json:"links"`
}

type CrawlMeta struct {
	Pagination CrawlPagination `json:"pagination"`
}

type CrawlData struct {
	ChapterId   int    `json:"chapter_id"`
	StoryId     int    `json:"story_id"`
	StoryName   string `json:"story_name"`
	ChapterName string `json:"chapter_name"`
	ChapterNext int    `json:"chapter_next"`
	ChapterPrev int    `json:"chapter_prev"`
	HasImage    bool   `json:"has_image"`
	Position    int    `json:"position"`
	CurrentPage int    `json:"current_page"`
	Content     string `json:"content"`
}

type CrawlDetail struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Meta       CrawlMeta   `json:"meta"`
	Data       []CrawlData `json:"data"`
}
