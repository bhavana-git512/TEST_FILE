package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Courses struct {
	Duration int `json:"Duration"`
	Fee      int `json:"Fee"`
}

var db *gorm.DB

func main() {
	//var err error = nil

	db, _ := gorm.Open("mysql", "root:Bhavana@1998@tcp(127.0.0.1:3306)/code?charset=utf8&parseTime=True")

	//if err != nil {

	//}
	db.AutoMigrate(&Courses{10, 5})
	ret := mux.NewRouter()

	ret.HandleFunc("/go/{Duration}/{Fee}", func(res http.ResponseWriter, req *http.Request) {

		mapp := mux.Vars(req)

		d := mapp["Duration"]
		d1, err := strconv.Atoi(d)
		f := mapp["Fee"]
		f1, err := strconv.Atoi(f)

		if err != nil {

		}
		nandu := Courses{Duration: d1, Fee: f1}
		db.Save(&nandu)

		json.NewEncoder(res).Encode(nandu)

	}).Methods("POST")
	log.Fatal(http.ListenAndServe(":6666", ret))

}
