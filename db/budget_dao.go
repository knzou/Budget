package db

import (
	"fmt"
	"log"
	"strings"
	_ "database/sql"

	proto "github.com/knzou/Budget/proto"

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
	TransDate string `db:"transdate"`
	Amount int64 `db:"amount"`
}

type Person struct {
	Pid int64 `db:"pid"`
	Name string `db:"name"`
}

func GetCategories(db *sqlx.DB) ([]Category, error) {
	rows, err := db.Query("SELECT * FROM category")
	stats := db.Stats()
	log.Printf("Pool Status \n Open Connections: %d \n InUse: %d \n Idle: %d", stats.OpenConnections, stats.InUse, stats.Idle)
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

func GetPeople(db *sqlx.DB, request *proto.GetPeopleRequest) ([]Person, error) {
	// This might be the long way for now, but as where clauses increase, it will scale
	contraints := map[string]string{"name": request.GetName()}
	var query strings.Builder
	query.WriteString("SELECT * FROM person")

	for _, k := range []string{"name"}{
		v, ok := contraints[k]
		fmt.Printf("%T", v)
		fmt.Println(ok)
		if ok && k == "name" {
			query.WriteString(fmt.Sprintf(" WHERE %s %% '%s'", k, v))
		} else {
			query.WriteString(fmt.Sprintf(" AND %s = '%s'", k, v))
		}
	}
	var people = []Person{}
	db.Select(&people, query.String())
	fmt.Println(query.String())
	return people, nil
}