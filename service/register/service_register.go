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
	service     = os.Args[1]
)

func serviceAssembler() *consul_api.AgentServiceRegistration {
	serviceCheckerAgent := consul_api.AgentServiceCheck{
		TCP:      "localhost:22",
		Interval: "10s",
		Timeout:  "1s",
	}
	serviceRegistrarAgent := consul_api.AgentServiceRegistration{
		Name:              service,
		ID:                service,
		Tags:              []string{service},
		Address:           "",
		Port:              0,
		EnableTagOverride: false,
		Check:             &serviceCheckerAgent,
	}
	return &serviceRegistrarAgent
}

func registerService(client consul_api.Client) {
	err := client.Agent().ServiceRegister(serviceAssembler())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Service registered with name %s\n", service)
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
	registerService(*client)
}
