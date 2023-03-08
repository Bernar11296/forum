package repository

import (
	"database/sql"
	"fmt"
	"forum/models"
)

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (r *CommentRepo) CreateComment(comment models.Comment) error {
	query := `INSERT INTO comment (user_id, post_id, text, author, date) values($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, comment.UserId, comment.PostId, comment.Text, comment.Author, comment.Date)
	if err != nil {
		return fmt.Errorf("repository create comment: %w", err)
	}
	return nil
}

func (r *CommentRepo) GetCommentByPostId(id int) (*[]models.Comment, error) {
	rows, err := r.db.Query(`SELECT * FROM comment where post_id=$1`, id)
	if err != nil {
		return nil, fmt.Errorf("repository get comment by post id: %w", err)
	}
	var comments []models.Comment
	for rows.Next() {
		comment := models.Comment{}
		if err := rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Likes, &comment.Dislikes, &comment.Text, &comment.Author, &comment.Date); err != nil {
			return nil, fmt.Errorf("repository get comment row scan: %w", err)
		}
		comments = append(comments, comment)
	}
	return &comments, nil
}
