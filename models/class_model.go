package models

import (
	"database/sql"
	"encoding/json"
	"goworkspace/entities"
	"net/http"
)

type ClassModel struct {
	Db *sql.DB
}

func (classModel ClassModel) FindAll() (class []entities.Class, err error) {
	rows, err := classModel.Db.Query("select * from class")
	if err != nil {
		return nil, err
	} else {
		var classes []entities.Class
		for rows.Next() {
			var id int64
			var name string
			var roll int64
			var age int64
			var email string
			err2 := rows.Scan(&id, &name, &roll, &age, &email)
			if err2 != nil {
				return nil, err2
			} else {
				class := entities.Class{
					Id:    id,
					Name:  name,
					Roll:  roll,
					Age:   age,
					Email: email,
				}
				classes = append(classes, class)
			}
		}
		return classes, nil
	}
}

func (classModel ClassModel) FindEmail() (email []string, err error) {
	rows, err := classModel.Db.Query("select email_id from class")
	if err != nil {
		return nil, err
	} else {
		var emails []string
		for rows.Next() {
			var email string
			err2 := rows.Scan(&email)
			if err2 != nil {
				return nil, err2
			}
			// } else {
			// 	class := entities.Class{
			// 		Id:   id,
			// 		Name: name,
			// 		Roll: roll,
			// 		Age:  age,
			// 	}
			emails = append(emails, email)
		}

		return emails, nil
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
