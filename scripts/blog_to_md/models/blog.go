package models

type Blog struct {
	Title      string `json:"title"`
	Content    string `json:"body_markdown"`
	UpdatedAt  string `json:"updated_at"`
	MainImage  string `json:"main_image"`
	Categories string `json:"cached_tag_list"`
}
