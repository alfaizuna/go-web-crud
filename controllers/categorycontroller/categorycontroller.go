package categorycontroller

import (
	"go-web-native/entities"
	categoryservices "go-web-native/services/categoryservices"
	"net/http"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categoryservices.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/categories/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/categories/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category
		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categoryservices.Create(category); !ok {
			temp, _ := template.ParseFiles("views/categories/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)

	}
}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
