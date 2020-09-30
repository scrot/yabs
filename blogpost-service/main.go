package main

import {
	"context"
	"log"
	"net"
	"sync"

	// Import the generated protobuf code
	pb "github.com/<YourUserName>/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
}