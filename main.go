package main

import (
	"database/sql"
	"internal/infrastructure/database"
	"internal/interface/handler"
	"internal/interface/router"
	"internal/usecase"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := &database.UserRepository{DB: db}
	workRepository := &database.WorkRepository{DB: db}

	userUsecase := &usecase.UserUsecase{UserRepository: userRepository}
	workUsecase := &usecase.WorkUsecase{WorkRepository: workRepository}

	userHandler := &handler.UserHandler{UserUsecase: userUsecase}
	workHandler := &handler.WorkHandler{WorkUsecase: workUsecase}

	r := router.NewRouter(userHandler, workHandler)

	// データベースの初期化（ここからは追加したもの）
	db = database.InitDB("test.db")
	defer db.Close()

	// ハンドラーの設定
	userHandler = handler.NewUserHandler(db)
	workHandler = handler.NewWorkHandler(db)

	// ルーターの設定
	r = router.NewRouter(userHandler, workHandler)

	log.Println("Server running at http://localhost:8080")
	log.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
