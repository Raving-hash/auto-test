package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db *sql.DB

func initDB() {
	var err error
	// Replace USERNAME:PASSWORD according to your MySQL credentials
	db, err = sql.Open("mysql", "root:yigeyy00@tcp(localhost:3307)/autotest")
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
func main() {
	initDB()
	r := mux.NewRouter()

	// 子功能相关路由
	r.HandleFunc("/subfeatures", GetSubFeatures).Methods("GET")
	r.HandleFunc("/subfeatures", CreateSubFeature).Methods("POST")
	r.HandleFunc("/subfeatures/{id}", UpdateSubFeature).Methods("PUT")
	r.HandleFunc("/subfeatures/{id}", DeleteSubFeature).Methods("DELETE")
	r.HandleFunc("/subfeatures/{id}/test", TestSubFeature).Methods("POST")

	// 启动服务器
	log.Println("Server starting on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
