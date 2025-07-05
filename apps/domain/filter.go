package domain

type Filter struct {
	Limit int
	Page  int
	Order string
	Where map[string]any
}
