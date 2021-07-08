package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var client *redis.Client
var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

func main() {
	// 打开sql的链接
	db, err := sql.Open("mysql", "root:by1593118@tcp(localhost:3306)/test")

	if err != nil {
		fmt.Println("error opening database")
		panic(err.Error())
	}

	defer db.Close()

	// 尝试连接db
	err = db.Ping()

	if err != nil {
		fmt.Println("error connect to db")
		panic(err.Error())
	}

	fmt.Println("connected to mysql")

	// redis 服务器
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// mux的路由
	router := Router()

	// 获取环境变量
	ENV_PORT := os.Getenv("PORT")

	if ENV_PORT == "" {
		ENV_PORT = "8080"
	}

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("api server")
	// }).Methods("GET")

	// //路由表
	// router.HandleFunc("/api/v1", indexHandler).Methods("GET")
	// router.HandleFunc("/api/v1/user/login", loginGetHandler).Methods("GET")
	// router.HandleFunc("/api/v1/user/login", loginPostHandler).Methods("POST")

	fmt.Println("server running on port", ENV_PORT)

	// 监听http服务器
	http.ListenAndServe(":8080", router)

}
