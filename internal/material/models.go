package material

import "time"

type Material struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	CategoryID  int       `json:"category_id"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
}
