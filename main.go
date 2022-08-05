package main

import (
	"example/invisifox"
	"fmt"
	"time"
)

func main() {
	client := invisifox.New("v1exC65dG4A7bbE1/2N17OJmPEnKn8sRTj3tS.Uo0NTd4lq6XwzKy")

	balance, err := client.GetBalance()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Your balance: %v\n", balance.CaptchaBalance)

	task, err := client.SolveCaptcha("discord.com", "4c672d35-0701-42b2-88c3-78380b0db560", "LilForkiAv9:29KAb6Gg5ZDX74ls_sticky-3@142.202.220.242:6484", "", "", "", "false")
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
