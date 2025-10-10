package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

type Product struct {
	ID int
	Name string
	Price int
}

func HandleIndexController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		println("/products")

		rows, err := db.Query("SELECT id, name, price FROM products")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		var products []Product

		for rows.Next() {
			var p Product
			err := rows.Scan(&p.ID, &p.Name, &p.Price)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			products = append(products, p)
		}

		f := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(f)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]any{
			"Products" : products,
		}

		tmpl.Execute(w, data)
	}
}

func HandleProductCreate(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		println("Masuk ke /products/create", r.Method)

		if r.Method == "GET" {
			f := filepath.Join("views", "create.html")
			tmpl, err := template.ParseFiles(f)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return 
			}

			err = tmpl.Execute(w, nil)

			if err != nil {
				println(err.Error())
			}
			return
		}

		if r.Method == "POST" {
			name := r.FormValue("name")
			priceStr := r.FormValue("price")
			price, _ := strconv.Atoi(priceStr)

			_, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", name, price)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return 
			}

			http.Redirect(w, r, "/products", http.StatusSeeOther)
		}
	}
}

func HandleProductEdit(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		println("/products/edit", r.Method)

		if r.Method == "GET" {
			
			idStr := r.URL.Query().Get("id")
			id, _ := strconv.Atoi(idStr)

			var p Product

			err := db.QueryRow("SELECT id, name, price FROM products WHERE id = $1", id).
				Scan(&p.ID, &p.Name, &p.Price)

			if err != nil {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}

			f := filepath.Join("views", "edit.html")
			tmpl, _ := template.ParseFiles(f)
			tmpl.Execute(w, p)
			return
		}

		if r.Method == "POST" {
			id := r.FormValue("id")
			name := r.FormValue("name")
			priceStr := r.FormValue("price")
			price, _ := strconv.Atoi(priceStr)

			_, err := db.Exec("UPDATE products SET name = $1, price = $ WHERE id = $3", name, price, id)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/products", http.StatusSeeOther)
		}
	}
}

func HandleDeleteProduct(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Query().Get("id")

		_, err := db.Exec("DELETE FROM products WHERE id = $1", id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}