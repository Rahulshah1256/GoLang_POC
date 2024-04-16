package db

import (
	"database/sql"
	"fmt"
)

type SQLStore struct {
	db *sql.DB
}

func NewStore(config util.Config) (*SQLStore, error) {
	db, err := sql.Open(config.SQLDriver, config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return &SQLStore{db: db}, nil
}

func (s *SQLStore) GetItems() ([]Item, error) {
	rows, err := s.db.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SQLStore) GetItem(id string) (*Item, error) {
	query := "SELECT * FROM items WHERE id = ?"
	row := s.db.QueryRow(query, id)

	var item Item
	if err := row.Scan(&item.ID, &item.Name, &item.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("item not found")
		}
		return nil, err
	}

	return &item, nil
}

func (s *SQLStore) CreateItem(item *Item) error {
	query := "INSERT INTO items (id, name, price) VALUES (?, ?, ?)"
	_, err := s.db.Exec(query, item.ID, item.Name, item.Price)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLStore) UpdateItem(id string, item *Item) error {
	query := "UPDATE items SET name = ?, price = ? WHERE id = ?"
	_, err := s.db.Exec(query, item.Name, item.Price, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLStore) DeleteItem(id string) error {
	query := "DELETE FROM items WHERE id = ?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
