package main

import (
	"sync"

	// Import the generated protobuf code
	pb "github.com/scrot/yabs/blogpost-service/gen/blogposts"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Blogpost) (*pb.Blogpost, error)
}

type Repository struct {
	mu        sync.RWMutex
	blogposts []*pb.Blogpost
}

func (repo *Repository) Create(blogpost *pb.Blogpost) (*pb.Blogpost, error) {
	repo.mu.Lock()
	updated := append(repo.blogposts, blogpost)
	repo.mu.Unlock()
	return blogpost, nil
}

type service struct {
	repo repository
}
