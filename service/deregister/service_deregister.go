package main

import (
	"fmt"
	"log"
	"os"

	consul_api "github.com/hashicorp/consul/api"
)

var (
	consulAddr  = os.Getenv("CONSUL_ADDR")
	consulToken = os.Getenv("CONSUL_TOKEN")
	serviceId   = os.Args[1]
)

func deregisterService(client consul_api.Client) {
	err := client.Agent().ServiceDeregister(serviceId)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Service deregistered with name %s\n", serviceId)
}

func main() {
	config := consul_api.Config{
		Address: consulAddr,
	}

	client, err := consul_api.NewClient(&config)
	if err != nil {
		log.Fatalf("Unable to init consul client: %v", err)
	}
	client.AddHeader("X-Consul-Token", consulToken)

	deregisterService(*client)
}
