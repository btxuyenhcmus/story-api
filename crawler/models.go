package crawler

// Meta Data
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

// Chapter Data
type CrawlShortChapter struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

type CrawlLargeChapter struct {
	Id          int    `json:"chapter_id"`
	StoryId     int    `json:"story_id"`
	StoryName   string `json:"story_name"`
	Name        string `json:"chapter_name"`
	Next        int    `json:"chapter_next"`
	Prev        int    `json:"chapter_prev"`
	HasImage    bool   `json:"has_image"`
	Position    int    `json:"position"`
	CurrentPage int    `json:"current_page"`
	Content     string `json:"content"`
}

// Story Data
type CrawlShortStory struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Image         string `json:"image"`
	IsFull        bool   `json:"is_full"`
	Time          string `json:"time"`
	Author        string `json:"author"`
	Categories    string `json:"categories"`
	CategoryIds   string `json:"category_ids"`
	TotalChapters int    `json:"total_chapters"`
}

type CrawlLargeStory struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Author      string `json:"author"`
	Source      string `json:"source"`
	Liked       bool   `json:"liked"`
	TotalLike   int    `json:"total_like"`
	TotalView   string `json:"total_view"`
	Categories  string `json:"categories"`
	CategoryIds string `json:"category_ids"`
	Description string `json:"description"`
}

// Response Data
type ResponseListChapter struct {
	Status     string              `json:"status"`
	Message    string              `json:"message"`
	StatusCode int                 `json:"status_code"`
	Meta       CrawlMeta           `json:"meta"`
	Data       []CrawlShortChapter `json:"data"`
}

type ResponseDetailChapter struct {
	Status     string            `json:"status"`
	Message    string            `json:"message"`
	StatusCode int               `json:"status_code"`
	Data       CrawlLargeChapter `json:"data"`
}

type ResponseListStory struct {
	Status     string            `json:"status"`
	Message    string            `json:"message"`
	StatusCode int               `json:"status_code"`
	Meta       CrawlMeta         `json:"meta"`
	Data       []CrawlShortStory `json:"data"`
}

type ResponseDownloadStory struct {
	Status     string              `json:"status"`
	Message    string              `json:"message"`
	StatusCode int                 `json:"status_code"`
	Meta       CrawlMeta           `json:"meta"`
	Data       []CrawlLargeChapter `json:"data"`
}

type ResponseDetailStory struct {
	Status     string          `json:"status"`
	Message    string          `json:"message"`
	StatusCode int             `json:"status_code"`
	Data       CrawlLargeStory `json:"data"`
}
