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

type CrawlDataList struct {
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

type CrawlDetail struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Meta       CrawlMeta   `json:"meta"`
	Data       []CrawlData `json:"data"`
}

type CrawlList struct {
	Status     string          `json:"status"`
	Message    string          `json:"message"`
	StatusCode int             `json:"status_code"`
	Meta       CrawlMeta       `json:"meta"`
	Data       []CrawlDataList `json:"data"`
}

type CrawlStoryDetail struct {
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

type CrawlStoryData struct {
	Status     string           `json:"status"`
	Message    string           `json:"message"`
	StatusCode int              `json:"status_code"`
	Data       CrawlStoryDetail `json:"data"`
}
