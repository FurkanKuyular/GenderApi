package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"strings"
)

type GenderPayload struct {
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	CountryCode string `json:"country"`
}

type Gender struct {
	Success       bool          `json:"success"`
	GenderPayload GenderPayload `json:"payload"`
}

type ErrorPayload struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Success      bool         `json:"success"`
	ErrorPayload ErrorPayload `json:"error"`
}

func checkName(res http.ResponseWriter, req *http.Request) bool {
	res.Header().Set("Content-Type", "application/json")
	name := req.FormValue("name")

	if name == "" {
		ErrorPayload := ErrorPayload{}
		ErrorPayload.Message = "Name is not exist"
		_ = json.NewEncoder(res).Encode(ErrorResponse{Success: false, ErrorPayload: ErrorPayload})

		return false
	}

	return true
}

func getName(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	name := req.FormValue("name")

	name = trToEn(name)
	name = strings.ToUpper(name)

	_ = godotenv.Load(".env")
	var myEnv map[string]string
	myEnv, _ = godotenv.Read()

	db, err := sql.Open("mysql", fmt.Sprintf("%s/%s", myEnv["DATABASE_DSN"], myEnv["DATABASE_NAME"]))

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusServiceUnavailable)
		ErrorPayload := ErrorPayload{}
		ErrorPayload.Message = "Technical error has occurred"
		_ = json.NewEncoder(res).Encode(ErrorResponse{Success: false, ErrorPayload: ErrorPayload})

		return
	}

	result, err := db.Query(fmt.Sprintf("SELECT name, gender, country_code FROM gender WHERE name = '%s' LIMIT 1", strings.ToUpper(name)))

	if err != nil {
		res.WriteHeader(http.StatusServiceUnavailable)
		ErrorPayload := ErrorPayload{}
		ErrorPayload.Message = "Technical error has occurred"
		_ = json.NewEncoder(res).Encode(ErrorResponse{Success: false, ErrorPayload: ErrorPayload})

		return
	}

	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {
			ErrorPayload := ErrorPayload{}
			ErrorPayload.Message = "Technical error has occurred"
			_ = json.NewEncoder(res).Encode(ErrorResponse{Success: false, ErrorPayload: ErrorPayload})

			return
		}
	}(result)

	var genderPayload GenderPayload
	var gender Gender

	for result.Next() {
		err = result.Scan(&genderPayload.Name, &genderPayload.Gender, &genderPayload.CountryCode)
		if err != nil {
			ErrorPayload := ErrorPayload{}
			ErrorPayload.Message = "Technical error has occurred"
			_ = json.NewEncoder(res).Encode(ErrorResponse{Success: false, ErrorPayload: ErrorPayload})

			return
		}
	}

	gender.Success = true
	gender.GenderPayload = genderPayload

	if genderPayload.Name == "" {
		ErrorPayload := ErrorPayload{}
		ErrorPayload.Message = "Name not found in database"
		_ = json.NewEncoder(res).Encode(ErrorResponse{Success: false, ErrorPayload: ErrorPayload})

		return
	}

	_ = json.NewEncoder(res).Encode(gender)
}

func handleRequests() {
	http.HandleFunc("/gender", func(res http.ResponseWriter, req *http.Request) {
		if checkName(res, req) {
			getName(res, req)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func trToEn(name string) string {
	trChars := []string{"ı", "ğ", "İ", "Ğ", "ç", "Ç", "ş", "Ş", "ö", "Ö", "ü", "Ü"}
	enChars := []string{"i", "g", "I", "G", "c", "C", "s", "S", "o", "O", "u", "U"}

	for i, toreplace := range trChars {
		r := strings.NewReplacer(toreplace, enChars[i])
		name = r.Replace(name)
	}

	return name
}

func main() {
	handleRequests()
}
