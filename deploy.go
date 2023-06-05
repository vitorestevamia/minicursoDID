package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// get endpoint argument
	if len(os.Args) == 1 {
		log.Fatal("Endpoint argument is empty")
	}
	endpoint := os.Args[1]

	//generate a seed
	seed, _ := ProvideDid()

	//prepare docker command
	command := fmt.Sprintf("docker compose -f ./docker-compose.yml up")
	parsedCommand := strings.Split(command, " ")
	cmd := exec.Command(parsedCommand[0], parsedCommand[1:]...)

	//prepare env variables
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("WALLET_SEED=%s", seed))
	cmd.Env = append(cmd.Env, fmt.Sprintf("ENDPOINT=%s", endpoint))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//execute
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
