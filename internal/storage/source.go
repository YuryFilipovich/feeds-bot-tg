package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"news-feed-bot-tg/internal/model"
	"time"
)

type SourcePostgresStorage struct {
	db *sqlx.DB
}

type dbSource struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	FeedURL   string    `db:"feed_url"`
	CreatedAt time.Time `db:"created_at"`
}

// Sources list retrieval
func (s *SourcePostgresStorage) Sources(ctx context.Context) ([]model.Source, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer func(conn *sqlx.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	var sources []dbSource
	if err := conn.SelectContext(ctx, &sources, `SELECT * FROM sources`); err != nil {
		return nil, err
	}

	return lo.Map(sources, func(source dbSource, _ int) model.Source { return model.Source(source) }), nil

}

func (s *SourcePostgresStorage) SourceByID(ctx context.Context, id int64) (*model.Source, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer func(conn *sqlx.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	var source dbSource
	if err := conn.GetContext(ctx, &source, `SELECT * FROM sources WHERE id = ?`, id); err != nil {
		return nil, err
	}

	return (*model.Source)(&source), nil
}

func (s *SourcePostgresStorage) Add(ctx context.Context, source model.Source) (int64, error) {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return 0, err
	}
	defer func(conn *sqlx.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	var id int64

	row := conn.QueryRowxContext(
		ctx, `INSERT INTO sources(name, feed_url, created_at) VALUES (?, ?, ?) RETURNING id`,
		source.Name,
		source.FeedURL,
		source.CreatedAt,
	)

	if err := row.Err(); err != nil {
		return 0, err
	}

	if err := row.StructScan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SourcePostgresStorage) Delete(ctx context.Context, id int64) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer func(conn *sqlx.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	if _, err := conn.ExecContext(ctx, `DELETE FROM sources WHERE id = ?`, id); err != nil {
		return err
	}
	return nil
}
