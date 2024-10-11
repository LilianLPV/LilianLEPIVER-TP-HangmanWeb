package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
)

func Erreur(Nom, Prenom string) bool {
	NomPrenomRegex := `^[A-Za-zÀ-ÿ][a-zà-ÿ]+([-'\s][A-Za-zÀ-ÿ][a-zà-ÿ]+)?$`

	regex := regexp.MustCompile(NomPrenomRegex)

	if !(regex.MatchString(Nom)) {
		return false
	}
	if !(regex.MatchString(Prenom)) {
		return false
	}

	return true
}

var Nombredevu int

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

func main() {

	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf("ERREUR => %s", err.Error())
		os.Exit(02)
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Promo{"Liste des étudiants",
			[]User{{"Bastien", " JOUFFRE", 16, "Masculin"},
				{"Lilian", " LE PIVER", 21, "Masculin"},
				{"Nans", " MOLL", 20, "Masculin"},
				{"Alexis", " DUPIN", 30, "Masculin"},
				{"Yoan", " COMMEAU", 28, "Masculin"}}}
		err := temp.ExecuteTemplate(w, "promo", data)
		if err != nil {
			http.Error(w, "ERREUR Promo", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		Nombredevu++
		Texte := ""
		if Nombredevu%2 == 0 {
			Texte += "Nombre de vu Pair"
		} else {
			Texte = "Nombre de vu Impaire"
		}
		Texte += fmt.Sprintf(" : %d", Nombredevu)
		err := temp.ExecuteTemplate(w, "change", Texte)
		if err != nil {
			http.Error(w, "ERREUR Change", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/user/form", func(w http.ResponseWriter, r *http.Request) {
		err := temp.ExecuteTemplate(w, "form", nil)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "ERREUR Formulaire", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {

		err := temp.ExecuteTemplate(w, "display", nil)
		if err != nil {
			http.Error(w, "ERREUR Display", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		err := temp.ExecuteTemplate(w, "display", nil)
		if err != nil {
			http.Error(w, "ERREUR Display", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/user/error", func(w http.ResponseWriter, r *http.Request) {
		err := temp.ExecuteTemplate(w, "error", nil)
		if err != nil {
			http.Error(w, "ERREUR ERREUR", http.StatusInternalServerError)
		}
	})
	http.ListenAndServe("localhost:8080", nil)

}
