package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
	Role     string `json:"role"`
	Location string `json:"location"`
}

func main() {
	user := User{
		Name:     "Aleksander Bogdan",
		Age:      21,
		Email:    "s25867@pjstk.edu.pl",
		Active:   true,
		Role:     "user",
		Location: "Gdynia",
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json error:", err)
		return
	}

	fmt.Println("JSON:")
	fmt.Println(string(userJSON))

	currentTime := time.Now()

	fmt.Println("Aktualny czas:", currentTime)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Operacja zakończona")
	}()
	wg.Wait()

	err = os.Setenv("GO_ENV", "production")
	if err != nil {
		fmt.Println("ev error:", err)
		return
	}
	envValue := os.Getenv("GO_ENV")
	fmt.Println("GO_ENV:", envValue)
}
