package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Struktur table
type GroupBusiness struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// Menampilkan 1 group business saja
func (gb *GroupBusiness) GetGroupBusiness(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id, name, created_at FROM group_businesses WHERE id=%d", gb.ID)
	return db.QueryRow(statement).Scan(&gb.ID, &gb.Name, &gb.CreatedAt)
}

// Update data group business
func (gb *GroupBusiness) UpdateGroupBusiness(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE group_businesses SET name='%s' WHERE id=%d", gb.Name, gb.ID)
	_, err := db.Exec(statement)
	return err
}

// Menghapus data group business
func (gb *GroupBusiness) DeleteGroupBusiness(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM group_businesses WHERE id=%d", gb.ID)
	_, err := db.Exec(statement)
	return err
}

// Simpan data group business
func (gb *GroupBusiness) CreateGroupBusiness(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO group_businesses(name) VALUES('%s')", gb.Name)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&gb.ID)
	if err != nil {
		return err
	}

	return nil
}

// Show all group business
func GetGroupBusinesses(db *sql.DB, start, count int) ([]GroupBusiness, error) {
	statement := fmt.Sprintf("SELECT id, name FROM group_businesses LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	groupBusinesses := []GroupBusiness{}
	for rows.Next() {
		var gb GroupBusiness
		if err := rows.Scan(&gb.ID, &gb.Name); err != nil {
			return nil, err
		}

		groupBusinesses = append(groupBusinesses, gb)
	}

	return groupBusinesses, nil
}
