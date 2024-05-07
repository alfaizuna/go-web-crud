package categoryservices

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`select * from categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}
	}

	return categories
}
