package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"pb"

	"google.golang.org/grpc"
)

func main() {
	CallMe()
}
func CallMe() {
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	c := pb.NewPhoneBookServiceClient(cc)

	ReadPerson(c)
	// createPerson(c)

	// updatePerson(c)
	//deletePerson(c)
	// listPerson(c)
}

func ReadPerson(c pb.PhoneBookServiceClient) {
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
}

func createPerson(c pb.PhoneBookServiceClient) {
	fmt.Println("Creating the person...")
	person := &pb.Person{
		Name:  "Guga Zimmermann",
		Email: "gugazimmermann@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "+55 47 98870-4247",
				Type:   pb.Person_MOBILE,
			},
			{
				Number: "+55 47 XXXXX-XXXX",
				Type:   pb.Person_HOME,
			},
		},
	}
	res, err := c.CreatePerson(context.Background(), &pb.PersonRequest{Person: person})
	if err != nil {
		fmt.Printf("Error while creating the person: %v\n", err)
	}
	fmt.Printf("Person Created: %v\n", res)
}

func updatePerson(c pb.PhoneBookServiceClient) {
	// CHANGE TO THE ID THAT YOU RECEIVED WHEN CREATE THE PERSON
	// YOU CAN TRY 605812e409be8dac8d59b5af TO SEE code = NotFound
	// AND xxxx TO SEE code = InvalidArgument
	personId := "62c0457f6b784ce8eb5a092c"
	fmt.Printf("Update person with ID: %v\n", personId)
	person := &pb.Person{
		Id:    personId,
		Name:  "José Augusto Zimmermann de Negreiros",
		Email: "jose.augusto@x-team.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "+55 47 98870-4247",
				Type:   pb.Person_WORK,
			},
		},
	}
	res, err := c.UpdatePerson(context.Background(), &pb.PersonRequest{Person: person})
	if err != nil {
		fmt.Printf("Error while updating the person: %v\n", err)
	}
	fmt.Printf("Person: %v\n", res)
}

func deletePerson(c pb.PhoneBookServiceClient) {
	// CHANGE TO THE ID THAT YOU RECEIVED WHEN CREATE THE PERSON
	// YOU CAN TRY 605812e409be8dac8d59b5af TO SEE code = NotFound
	// AND xxxx TO SEE code = InvalidArgument
	personId := "62c0457f6b784ce8eb5a092c"
	fmt.Printf("Deleting person with ID: %v\n", personId)
	res, err := c.DeletePerson(context.Background(), &pb.PersonIdRequest{PersonId: personId})
	if err != nil {
		fmt.Printf("Error while deleting the person: %v\n", err)
	}
	fmt.Printf("Person: %v\n", res)
}

func listPerson(c pb.PhoneBookServiceClient) {
	fmt.Println("listPerson...")
	stream, err := c.ListPerson(context.Background(), &pb.ListPersonResquest{})
	if err != nil {
		fmt.Printf("Error while calling ListPerson RPC: %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened while receive stream: %v\n", err)
		}
		fmt.Println(res.GetPerson())
	}
}
