package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Company struct {}

func (c *Company) SendPay() {
	fmt.Println("It is time to send pay!")
	delay := c.SomethingWrong(false)
	time.Sleep(time.Hour * 24 * time.Duration(delay))
	fmt.Println("...")
}

func (c *Company) SomethingWrong(asked bool) int {
	delay := rand.Intn(10)
	if asked {
		fmt.Println("something goes wrong.")
	}
	return delay
}

func main() {
	c := new(Company)
	c.SendPay()
}

