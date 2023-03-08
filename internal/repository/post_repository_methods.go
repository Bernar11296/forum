package repository

import (
	"database/sql"
	"fmt"
	"forum/models"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) CreatePost(post *models.Post) error {
	query := `INSERT INTO post (author_id, title, category, content, author, date) values ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, post.AuthorID, post.Title, post.Category, post.Content, post.Author, post.Date)
	if err != nil {
		return fmt.Errorf("repository create post: %w", err)
	}
	return nil
}

func (r *PostRepo) GetAllPost() (*[]models.Post, error) {
	rows, err := r.db.Query(`SELECT * FROM post`)
	if err != nil {
		return nil, fmt.Errorf("repository get all post: %w", err)
	}
	var posts []models.Post
	for rows.Next() {
		post := &models.Post{}
		if err := rows.Scan(&post.ID, &post.AuthorID, &post.Likes, &post.Dislikes, &post.Title, &post.Category, &post.Content, &post.Author, &post.Date); err != nil {
			return nil, fmt.Errorf("repository get all post scan: %w", err)
		}
		posts = append(posts, *post)
	}
	return &posts, nil
}
func (r *PostRepo) GetPostByCategory(category string) (*[]models.Post, error) {
	rows, err := r.db.Query(`SELECT * FROM post where category like '% ` + category + `'`)
	if err != nil {
		return nil, fmt.Errorf("repository get post by category: %w", err)
	}
	var posts []models.Post
	for rows.Next() {
		post := &models.Post{}
		if err := rows.Scan(&post.ID, &post.AuthorID, &post.Likes, &post.Dislikes, &post.Title, &post.Category, &post.Content, &post.Author, &post.Date); err != nil {
			return nil, fmt.Errorf("repository get post by category rows scan: %w", err)
		}
		posts = append(posts, *post)
	}
	return &posts, nil
}

func (r *PostRepo) GetPostByID(string) (*models.Post, error) {
	query := `SELECT * FROM post where id=$1`
	row := r.db.QueryRow(query)
	var post models.Post
	if err := row.Scan(&post.ID, &post.AuthorID, &post.Likes, &post.Dislikes, &post.Title, &post.Category, &post.Content, &post.Author, &post.Date); err != nil {
		return nil, fmt.Errorf("repository get post by id: %w", err)
	}
	return &post, nil
}

func (r *PostRepo) MyPosts(id string) (*[]models.Post, error) {
	rows, err := r.db.Query(`SELECT * FROM post where author_id=` + id + ``)
	if err != nil {
		return nil, fmt.Errorf("repository my posts: %w", err)
	}
	var posts []models.Post
	for rows.Next() {
		post := &models.Post{}
		if err := rows.Scan(&post.ID, &post.AuthorID, &post.Likes, &post.Dislikes, &post.Title, &post.Category, &post.Content, &post.Author, &post.Date); err != nil {
			return nil, fmt.Errorf("repository my posts rows scan: %w", err)
		}
		posts = append(posts, *post)
	}
	return &posts, nil
}

func (r *PostRepo) MyFavourites(id int) (*[]models.Post, error) {
	rows, err := r.db.Query(`SELECT * FROM post where user_id=$1 AND post_id != 0 AND active=1`, id)
	if err != nil {
		return nil, fmt.Errorf("repository my favourites: %w", err)
	}
	var postsID []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("repository my favourites: scan post_id: %w", err)
		}
		postsID = append(postsID, id)
	}
	var posts []models.Post
	for _, id := range postsID {
		post := new(models.Post)
		query := `SELECT * FROM post WHERE id=$1 ORDER BY id DESC`
		row := r.db.QueryRow(query, id)
		if err = row.Scan(&post.ID, &post.AuthorID, &post.Likes, &post.Dislikes, &post.Title, &post.Category, &post.Content, &post.Author, &post.Date); err != nil {
			return nil, fmt.Errorf("repository my favourites: scan post: %w", err)
		}
		posts = append(posts, *post)
	}
	return &posts, nil
}
