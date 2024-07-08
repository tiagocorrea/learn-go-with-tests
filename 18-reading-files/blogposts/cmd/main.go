package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tiagocorrea/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		fmt.Printf("Title: %v", post.Title)
		fmt.Println()
		fmt.Printf("Description: %v", post.Description)
		fmt.Println()
		fmt.Printf("Tags: %v", post.Tags)
		fmt.Println()
		fmt.Println("---")
		fmt.Println(post.Body)
		fmt.Print("============\n\n")
	}
}
