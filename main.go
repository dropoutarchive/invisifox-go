package main

import (
	"example/invisifox"
	"fmt"
	"time"
)

func main() {
	client := invisifox.New("steal")

	balance, err := client.GetBalance()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Your balance: %v\n", balance.CaptchaBalance)

	task, err := client.SolveCaptcha("x", "x", "x", "x", "x", "x", "false")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created task: %v\n", task.TaskID)

	var result *invisifox.SolutionResponse
	for {
		result, err = client.GetSolution(task.TaskID)
		if err != nil {
			panic(err)
		}

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
