package postgres

import (
	"GoNews/pkg/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func New(connString string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{db}, nil
}

func (p *PostgresDB) Posts() ([]storage.Post, error) {
	fmt.Println("Выводит все посты")
	post, err := p.db.Query("SELECT title,content FROM posts;")
	if err != nil {
		return nil, err
	}
	var pos []storage.Post
	for post.Next() {
		var t storage.Post
		err = post.Scan(
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		pos = append(pos, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return pos, post.Err()
}

func (p *PostgresDB) AddPost(post storage.Post) error {
	fmt.Println("Добовляет пост")
	add, err := p.db.Query("INSERT INTO posts author_id,title,content VALUES 2,'НЛО','Как то раз я увидел НЛО, печаль.';")
	if err != nil {
		return err
	}
	defer add.Close()
	return nil
}

func (p *PostgresDB) UpdatePost(post storage.Post) error {
	fmt.Println("Изменяет пост")
	upd, err := p.db.Query("UPDATE posts SET  title = 'Привет', content = 'обычно так говорят при встрече' WHERE id = '5';")
	if err != nil {
		return err
	}
	defer upd.Close()
	return nil

}

func (p *PostgresDB) DeletePost(post storage.Post) error {
	fmt.Println("Удаляет пост")
	del, err := p.db.Query("DELETE FROM posts WHERE id = 3;")
	if err != nil {
		return err
	}
	defer del.Close()
	return nil
}
