package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// user name
type UserName struct {
	gorm.Model()
	Name string
}

// for CORS
func forCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){
		w.Header().Set('Access-Control-Allow-Origin', '*')
		w.Header().Set('Access-Control-Allow-Headers', '*')

		if r.Method == 'OPTONS' {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
		return

	})
}



// データベース操作の関数群
// データベース初期化
func dbInit(){
	db, err :=  gorm.Open('sqlite3', '3chiku.sqlite3')

	// エラー処理
	if err != nil {
		panic('database can not be opened ! (dbInit)')
	}
	db.AutoMigrate(&UserName{})
	defer db.Close()
}

// データベースにメンバーを追加
func dbAdd()  {
	db, err := gorm.Open('sqlite3', '3chiku.sqlite3')

	if err != nil {
		panic('database can not be opend! (dbAdd)')
	}
}

func main(){

	// データベース初期化
	dbInit()

	// ルーター起動
	router := mux.NewRouter()
	// CORS対策
	router.Use(forCORS)
	// ルーティング
	router.HandleFunc('/useradd').Methods('POST', 'OPTIONS')
}