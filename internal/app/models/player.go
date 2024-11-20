package models

// Stores Players Login info
type InsertPlayer struct {
	Username string `json:"username"`
	Password int64  `json:"password"`
}

type GetPlayer struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
}

type ListPlayer struct {
	Players []GetPlayer `json:"players"`
	Total   int64       `json:"total"`
}

type UpdatePlayer struct {
	ID       int64 `json:"id"`
	Password int64 `json:"password"`
}

type DeletePlayer struct {
	ID int64 `json:"id"`
}
