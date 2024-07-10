package blogrenderer

import "strings"

// Post is a representation of a post
type Post struct {
	Title, Description, Body string
	Tags                     []string
}

// SanitisedTitle returns the title of the post with spaces replaced by dashes for pleasant URLs
//
//	post := Post{Title: "Hello World"}
//	post.SanitisedTitle()
//
// Will output:
//
//	hello-world
func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
