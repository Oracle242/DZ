package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/mongo2"
	"GoNews/pkg/storage/postgres"
	"log"
	"net/http"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db := memdb.New()

	// Реляционная БД PostgreSQL.
	// pwd := os.Getenv("DBPASS")
	// if pwd == "" {
	// 	os.Exit(1)
	// }
	// connstr := "postgres://postgres:" + pwd + "@127.0.0.1/GoNews"
	// присвоение переменной типа интерфейс конкретной реализации БД
	db2, err := postgres.New("@127.0.0.1/GoNews")
	if err != nil {
		log.Fatal(err)
	}

	// Документная БД MongoDB.
	db3, err := mongo2.NewMongo("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db, db2

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db3

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router())
}
