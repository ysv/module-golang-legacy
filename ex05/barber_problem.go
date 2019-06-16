package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	CUTTING_TIME = 20
	BARBERS      = 1
	HALL_SITS    = 5
)

type Barber struct {
	numb   int
	status int
}

type Client struct {
	numb   int
	status int
}

func NewBarber(numb int) *Barber {
	return &Barber{numb, 0}
}

func NewClient(numb int) *Client {
	return &Client{numb, 1}
}

func ClientQueue(clients chan *Client) {
	for i := 0; ; i++ {
		amt := time.Duration(rand.Intn(1250))
		time.Sleep(time.Millisecond * amt)
		clients <- NewClient(i)
	}
}

func cutting(barber *Barber, client *Client, finished chan *Barber) {
	time.Sleep(CUTTING_TIME * time.Millisecond)
	client.status = 0
	finished <- barber
}

func BarberShop(clients chan *Client) {
	Barbers := []*Barber{}
	waitingClient := []*Client{}
	BarberChan := make(chan *Barber)
	for i := 0; i < BARBERS; i++ {
		Barbers = append(Barbers, NewBarber(i))
	}
	for {
		select {
		case client := <-clients:
			if len(Barbers) == 0 {
				if len(waitingClient) < HALL_SITS {
					waitingClient = append(waitingClient, client)
					fmt.Printf("Client is waiting in hall (%v)\n", len(waitingClient))
				} else {
					fmt.Println("No free space for client")
				}
			} else {
				barber := Barbers[0]
				Barbers = Barbers[1:]
				barber.status = 1
				fmt.Println("Client goes to barber")
				go cutting(barber, client, BarberChan)
			}
		case barber := <-BarberChan:
			if len(waitingClient) > 0 {
				client := waitingClient[0]
				waitingClient = waitingClient[1:]
				fmt.Printf("Take client from room (%v)\n", len(waitingClient))
				go cutting(barber, client, BarberChan)
			} else {
				fmt.Println("Zzz..")
				barber.status = 0
				Barbers = append(Barbers, barber)
			}

		}
	}
}

func main() {
	clients := make(chan *Client)
	go ClientQueue(clients)
	go BarberShop(clients)
	time.Sleep(2 * time.Second)
}
