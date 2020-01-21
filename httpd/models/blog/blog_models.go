package blogmodels

// Post represents an entry to the blog. Posts have a one-to-many relationship to Replies
type Post struct {
	id       int
	title    string `json:"title"`
	body     string `json:"body"`
	comments string `json:"comment"`
}

// ValidatePost - pretty self explanatory. Guard clauses for blog posts.
func validatePost() {
	//add validation here - might need to import pq
}

/*
So the question is how to structure an interface for PostgreSQL.
First impressions:
1. define your structured datatype as a struct
2. define methods for that datatype

Should methods call pq directly?
should I format the Queries myself as strings
*/

/*
func insertPost(post Post, body string) (post Post, err error) {
	query := `
	INSERT INTO post (
		title, body, created_date
	) VALUES (
		?,     ?,    NOW()
	)`
	res, err := db_w.Exec(query,
		title,
		body,
	)
	if err != nil {
		return Post{}, err
	}

	id, _ := res.LastInsertId()
	return Post{
		id,
		title,
		body,
		time.Now().UTC().Format(time.RFC3339),
		1,
		time.Now().UTC().Format(time.RFC3339),
		0,
	}
}
*/
