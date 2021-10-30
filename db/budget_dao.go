package db

import (
	"log"
	_ "database/sql"
	"github.com/jmoiron/sqlx"
)
// uses - Dependency injection -> more details(https://www.alexedwards.net/blog/organising-database-access)
type Category struct {
	CatId int64 `db:"catid"`
	Name string `db:"name"`
	TypeId int64 `db:"typeid"`
}

type Transaction struct {
	TranId int64 `db:"tranid"`
	CatId int64 `db:"catid"`
	Amount int64 `db:"amount"`
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
	return cats, nil
}

func GetTransactions(db *sqlx.DB) ([]Transaction, error) {
	var trans = []Transaction{}
	// get is single, select is all
	// err = db.Get(&trans, "SELECT * FROM transaction")
	db.Select(&trans, "SELECT * FROM transaction")

	log.Printf("trans %v", trans[0])
	return trans, nil
}