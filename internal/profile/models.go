package profile

type UpdateProfileRequest struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Profile struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
