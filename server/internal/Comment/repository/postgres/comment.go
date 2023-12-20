package postgres

import (
	"database/sql"
	"fmt"
	"server/internal/Comment/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type commentRepo struct {
	DB *sql.DB
}

//NewCommentRepo creates comment repo
func NewCommentRepo(db *sql.DB) repository.CommentRepositoryI {
	return &commentRepo{
		DB: db,
	}
}

func (repo *commentRepo) Create(comment *dto.DBReqCreateComment) (*entity.Comment, error) {
	insertComment := `INSERT INTO comment (content, restaurant_id, user_id, rating)
					  VALUES ($1, $2, $3, $4)
					  RETURNING ID`
	var id uint
	err := repo.DB.QueryRow(insertComment, comment.Text, comment.RestaurantID, comment.UserID, comment.Rating).Scan(&id)

	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	respComment := &dto.DBRespCreateComment{
		Text:         comment.Text,
		UserID:       comment.UserID,
		RestaurantID: comment.RestaurantID,
		Rating:       comment.Rating,
	}
	getCommentDate := `SELECT created_at
					   FROM comment
					   WHERE id = $1`
	err = repo.DB.QueryRow(getCommentDate, id).Scan(&respComment.Date)
	fmt.Println(err)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}

	return respComment.FromDBRespToEntCreateComment(), nil
}

func (repo *commentRepo) Get(id uint) ([]*dto.RespGetComment, error) {
	getComment := `SELECT u.username, u.icon, c.content, c.rating, c.created_at
				   FROM comment c
				   JOIN users u ON c.user_id = u.id
				   WHERE c.restaurant_id = $1
				   ORDER BY created_at desc`
	rows, err := repo.DB.Query(getComment, id)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	defer rows.Close()

	var comments = []*dto.RespGetComment{}
	for rows.Next() {
		comment := &dto.DBRespGetComment{}
		err = rows.Scan(
			&comment.Username,
			&comment.Icon,
			&comment.Text,
			&comment.Rating,
			&comment.Date,
		)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		comments = append(comments, comment.FromDBtoDel())
	}
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound
		}
		return nil, entity.ErrInternalServerError
	}

	return comments, nil
}
