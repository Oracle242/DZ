package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	db *mongo.Client
}

func NewMongo(opts string) (*MongoDB, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(opts))
	if err != nil {
		log.Fatal(err)
	}
	// не забываем закрывать ресурсы
	defer db.Disconnect(context.Background())
	// проверка связи с БД
	err = db.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return &MongoDB{db}, nil
}

func (m *MongoDB) Posts() ([]storage.Post, error) {
	fmt.Println("Выводит все посты")
	var posts []storage.Post
	return posts, nil
}

func (m *MongoDB) AddPost(post storage.Post) error {
	return nil
}

func (m *MongoDB) UpdatePost(post storage.Post) error {
	fmt.Println("Изменяет пост")
	return nil
}

func (m *MongoDB) DeletePost(post storage.Post) error {
	fmt.Println("Удаляет пост")
	return nil
}
