package db

import (
	"log"
	_ "database/sql"
	"github.com/jmoiron/sqlx"
)

type Category struct {
	CatId int64 `db:"catId"`
	Name string `db:"name"`
	TypeId int64 `db:"typeId"`
}

func GetCategories(db *sqlx.DB) ([]Category, error) {
	rows, err := db.Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []Category

	for rows.Next() {
		var cat Category
		err := rows.Scan(&cat.CatId, &cat.Name, &cat.TypeId)
		if err != nil {
			return nil,  err
		}
		log.Printf("cat %v", cat)
		cats = append(cats, cat)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// look at cats then
	// find a way to return []*proto.GetCategoriesResponse_Category
	return cats, nil
}
// uses - Dependency injection -> more details(https://www.alexedwards.net/blog/organising-database-access)
// Define all db tables model here