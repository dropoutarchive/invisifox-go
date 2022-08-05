package main

import (
	"example/invisifox"
	"fmt"
	"time"
)

func main() {
	client := invisifox.New("token secret")

	balance := client.GetBalance()
	fmt.Printf("Your balance: %v\n", balance.CaptchaBalance)

	task := client.SolveCaptcha("pageurl", "sitekey", "username:password@ip:port", "rqdata", "user-agent", "cookies", "false")
	fmt.Printf("Created task: %v\n", task.TaskID)

	var result *invisifox.SolutionResponse
	for {
		result = client.GetSolution(task.TaskID)
		fmt.Printf("Task status: %v\n", result.Status)

		if result.Status == "WAITING" {
			time.Sleep(5 * time.Second)
		} else if result.Status == "OK" {
			break
		} else {
			panic(result.Status)
		}
	}

	fmt.Println(result.Solution)
}
