package budget

import (
	"database/sql"
)

type Category struct {
	catId int
	name string
	typeId int
}

func GetCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query("SELECT * FROM Category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []Category

	for rows.Next() {
		var cat Category

		err := rows.Scan(&Category.catId, &Category.name, &Category.typeId)
		if err != nil {
			return nil,  err
		}
		cats = append(cats, cat)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cats, nil
}
// uses - Dependency injection -> more details(https://www.alexedwards.net/blog/organising-database-access)
// Define all db tables model here