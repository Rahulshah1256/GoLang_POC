package db

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Store interface {
	GetItems() ([]Item, error)
	GetItem(id string) (*Item, error)
	CreateItem(item *Item) error
	UpdateItem(id string, item *Item) error
	DeleteItem(id string) error
}
