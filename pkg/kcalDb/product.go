package kcaldb

import (
	"ai-kcal-agent/pkg/appContext"
	"database/sql"
)

type Product struct {
	Id            int64
	Name          string
	Kcal          float64
	Proteins      float64
	Fats          float64
	Carbohydrates float64
}

type ProductAlias struct {
	Name      string
	ProductId int64
	Id        int64
}

func GetUserProductByName(name string, userId int64) (*Product, error) {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	ctx.Logger.Printf("[DB]: get product by name query <%s> <%d>\n", name, userId)
	row := db.QueryRow(
		"SELECT name, kcal, proteins, fats, carbohydrates, id FROM products WHERE name=$1 AND user_id=$2;",
		name,
		userId,
	)
	var product Product
	err = row.Scan(&product.Name, &product.Kcal, &product.Proteins, &product.Fats, &product.Carbohydrates, &product.Id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	ctx.Logger.Printf("[DB]: get product by name query success <%s> <%d>\n", name, userId)
	return &product, err
}

func GetUserProductByAlias(name string, userId int64) (*Product, error) {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	ctx.Logger.Printf("[DB]: get product by alias query <%s> <%d>\n", name, userId)
	row := db.QueryRow("SELECT name, product_id, id FROM products_aliases WHERE name=$1 AND user_id=$2", name, userId)
	var alias ProductAlias
	err = row.Scan(&alias.Name, &alias.ProductId, &alias.Id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var product Product
	row = db.QueryRow("SELECT name, kcal, proteins, fats, carbohydrates, id FROM products WHERE id=$1;", alias.ProductId)
	err = row.Scan(&product.Name, &product.Kcal, &product.Proteins, &product.Fats, &product.Carbohydrates, &product.Id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	ctx.Logger.Printf("[DB]: get product by alias query success <%s> <%d>\n", name, userId)
	return &product, err
}
