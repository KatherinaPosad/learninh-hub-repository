package reviews

import "time"

type Review struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	MaterialID int       `json:"material_id"`
	Rating     bool      `json:"rating"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}
