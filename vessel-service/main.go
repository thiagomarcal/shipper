// vessel-service/main.go
package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	micro "github.com/micro/go-micro"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	k8s "github.com/micro/kubernetes/go/micro"
	userService "github.com/thiagomarcal/shipper/user-service/proto/user"
	pb "github.com/thiagomarcal/shipper/vessel-service/proto/vessel"
	"golang.org/x/net/context"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	log.Println("DB_HOST3", host)

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Fatalf("Error connecting to datastore: %v", err)
	}

	repo := &VesselRepository{session.Copy()}

	createDummyData(repo)

	srv := k8s.NewService(
		micro.Name("shipper.vessel"),
		micro.Version("latest"),
		// micro.WrapHandler(AuthWrapper),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		authClient := userService.NewUserServiceClient("shipper.user", client.DefaultClient)
		_, err := authClient.ValidateToken(ctx, &userService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
