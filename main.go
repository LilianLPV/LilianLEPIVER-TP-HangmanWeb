package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {

	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf("ERREUR => %s", err.Error())
		os.Exit(02)
	}
	var nombredevu int 
	type User struct {
		FirstName string
		LastName  string
		Age       int
		Sex       string
	}
	type Promo struct {
		Title string
		Users []User
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Promo{"Liste des étudiants",
			[]User{{"Bastien", " JOUFFRE", 16, "Masculin"},
				{"Lilian", " LE PIVER", 21, "Masculin"},
				{"Nans", " MOLL", 20, "Masculin"},
				{"Alexis", " DUPIN", 30, "Masculin"},
				{"Yoan", " COMMEAU", 28, "Masculin"}}}
		temp.ExecuteTemplate(w, "promo.html", data)
	})

	// ============= Exemples template ==================

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		
		temp.ExecuteTemplate(w, "change", nil)
	})

	// ============= Exemples variables ==================

	http.HandleFunc("/user/form", func(w http.ResponseWriter, r *http.Request) {
		// Appel du template nommé "exempleVariable" avec les données stockées dans la variable dataPage
		temp.ExecuteTemplate(w, "form", nil)
	})

	// Déclaratrion d'une structure correspondant aux champs du template

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {

		// Appel du template nommé "exempleVariableComplexe" avec les données stockées dans la variable dataPage
		temp.ExecuteTemplate(w, "exempleVariableComplexe", nil)
	})

	//============= Exemples conditions ================

	// Déclaratrion d'une structure correspondant aux champs du template

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {

		// Appel du template nommé "exempleConditionSimple" avec les données stockées dans la variable dataPage
		temp.ExecuteTemplate(w, "exempleConditionSimple", nil)
	})

	http.ListenAndServe("localhost:8080", nil)

}
