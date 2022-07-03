package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dhananjayksharma/grpc-example/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	grpcClient := pb.NewPhoneBookServiceClient(cc)

	// Set up a http server.
	r := gin.Default()
	r.GET("/listbyid", func(c *gin.Context) {
		data := ReadPerson(grpcClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(data),
		})
	})

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func ReadPerson(c pb.PhoneBookServiceClient) interface{} {
	// CHANGE TO THE ID THAT YOU RECEIVED WHEN CREATE THE PERSON
	// YOU CAN TRY 605812e409be8dac8d59b5af TO SEE code = NotFound
	// AND xxxx TO SEE code = InvalidArgument
	personId := "62c0529c1f80803fb380d642"
	fmt.Printf("Reading person with ID: %v\n", personId)
	res, err := c.ReadPerson(context.Background(), &pb.PersonIdRequest{PersonId: personId})
	if err != nil {
		fmt.Printf("Error while reading the person: %v\n", err)
	}
	fmt.Printf("Person: %v\n", res)
	var data interface{}
	data = res
	return data
}
