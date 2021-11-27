package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	userpb "github.com/andreoav/protocolbuffers/types/go"
)

func main() {
	user := &userpb.User{
		Name:  "Andreo Vieira",
		Email: "contato@andreovieira.com.br",
		Age:   31,
		Phones: []*userpb.Phone{
			{Number: "+5551993554327", Type: userpb.Phone_MOBILE},
			{Number: "+5551999999999", Type: userpb.Phone_WORK},
			{Number: "+5551999999999", Type: userpb.Phone_HOME},
		},
		Addresses: []*userpb.Address{
			{
				Type:         userpb.Address_DELIVERY,
				Street:       "Address 0001",
				StreetNumber: "9999",
				City:         "Porto Alegre",
				State:        "RS",
				Country:      "Brazil",
				ZipCode:      "99999-999",
			},
			{
				Type:         userpb.Address_HOME,
				Street:       "Address 0001",
				StreetNumber: "9999",
				City:         "Porto Alegre",
				State:        "RS",
				Country:      "Brazil",
				ZipCode:      "99999-999",
			},
			{
				Type:         userpb.Address_OTHER,
				Street:       "Address 0001",
				StreetNumber: "9999",
				City:         "Porto Alegre",
				State:        "RS",
				Country:      "Brazil",
				ZipCode:      "99999-999",
			},
		},
		IsActive:  true,
		CreatedAt: timestamppb.Now(),
	}

	bytes, err := proto.Marshal(user)

	if err != nil {
		log.Fatal("failed to marshal user", err)
	}

	if err := ioutil.WriteFile("./data/user.bin", bytes, 0644); err != nil {
		log.Fatal("failed to save user.bin", err)
	}

	fmt.Println(user)
	fmt.Println(protojson.Format(user))
}
