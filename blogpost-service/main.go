package main

import {
	"context"
	"log"
	"net"
	"sync"

	// Import the generated protobuf code
	pb "https://github.com/scrot/yabs/blogpost-service/gen/blogposts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
}