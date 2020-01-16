package blogmodels

// Post represents an entry to the blog. Posts have a one-to-many relationship to Replies
type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Replies is a Replies owned by a specific Post. Replies have a many-to-one relationship with Posts
type Replies struct {
	Reply string `json:"reply"`
}

// ValidatePost - pretty self explanatory. Guard clauses for blog posts.
func ValidatePost() {
	//add validation here - might need to import pq
}
