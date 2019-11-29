package post

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mintutu/restful-golang/model"
	"github.com/mintutu/restful-golang/repository"
)

func NewMySQLPostRepo(Conn *sql.DB) repository.PostRepo {
	return &MySQLPostRepo{
		Conn: Conn,
	}
}

type MySQLPostRepo struct {
	Conn *sql.DB
}

func (m *MySQLPostRepo) Fetch(ctx context.Context) ([]*model.Post, error) {
	query := "SELECT id, title FROM db.post"
	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		fmt.Printf("%v \n", err)
		return nil, err
	}
	defer rows.Close()

	payload := make([]*model.Post, 0)
	for rows.Next() {
		data := new(model.Post)

		err := rows.Scan(
			&data.ID,
			&data.Title,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *MySQLPostRepo) Create(ctx context.Context, p *model.Post) (int64, error) {
	query := "INSERT post SET title=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Title)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
