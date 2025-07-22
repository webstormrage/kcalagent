package kcaldb

import (
	"ai-kcal-agent/pkg/appContext"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"time"
)

type Product struct {
	Id            int64
	Name          string
	Kcal          float64
	Proteins      float64
	Fats          float64
	Carbohydrates float64
}

func GetProductByName(name string) (*Product, error) {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	row := db.QueryRow("SELECT * FROM products WHERE name=$1;", name)
	var product Product
	err = row.Scan(&product.Name, &product.Kcal, &product.Proteins, &product.Fats, &product.Carbohydrates, &product.Id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &product, err
}

type ProductAlias struct {
	Name      string
	ProductId int64
	Id        int64
}

func GetProductByAlias(name string) (*Product, error) {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	row := db.QueryRow("SELECT * FROM products_aliases WHERE name=$1;", name)
	var alias ProductAlias
	err = row.Scan(&alias.Name, &alias.ProductId, &alias.Id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var product Product
	row = db.QueryRow("SELECT * FROM products WHERE id=$1;", alias.ProductId)
	err = row.Scan(&product.Name, &product.Kcal, &product.Proteins, &product.Fats, &product.Carbohydrates, &product.Id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &product, err
}

type MealPayload struct {
	Name          string
	Weight        float64
	Kcal          float64
	Proteins      float64
	Fats          float64
	Carbohydrates float64
}

func SaveProduct(product *Product) error {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(
		"INSERT INTO products (name, kcal, proteins, fats, carbohydrates) VALUES ($1, $2, $3, $4, $5, $6)",
		product.Name,
		product.Kcal,
		product.Proteins,
		product.Fats,
		product.Carbohydrates,
	)
	return err
}

func SaveMeals(meal *MealPayload) error {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(
		"INSERT INTO meals VALUES ($1, $2, $3, $4, $5, $6)",
		meal.Name,
		meal.Weight,
		meal.Kcal,
		meal.Proteins,
		meal.Fats,
		meal.Carbohydrates,
	)
	return err
}

type MealSummary struct {
	Kcal          float64
	Proteins      float64
	Fats          float64
	Carbohydrates float64
}

func GetDailySummary(dayStartTime time.Time) (*MealSummary, error) {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(
		"SELECT SUM(kcal) AS kcal, SUM(proteins) AS proteins, SUM(fats) AS fats, SUM(carbohydrates) AS carbohydrates FROM meals"+
			" WHERE created_at > $1",
		dayStartTime,
	)
	var summary MealSummary
	err = row.Scan(&summary.Kcal, &summary.Proteins, &summary.Fats, &summary.Carbohydrates)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &summary, err
}

func SetupDb() error {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	if err != nil {
		return err
	}
	defer db.Close()
	fileData, err := os.ReadFile("sql/schema-1.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(fileData))
	return err
}
