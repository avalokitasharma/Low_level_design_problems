// lets say there is an auction and bunch of subscribers to that auction,
// whenever a higher bid is placed all the subscribers get notfied.
// make use of channels in go build this observer pattern
//  - simple in memory demo

package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	Id      string
	Message chan string
	Quit    chan struct{}
}

func NewUser(id string) *User {
	return &User{
		Id:      id,
		Message: make(chan string),
		Quit:    make(chan struct{}),
	}
}
func (u *User) Listen(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case msg := <-u.Message:
			fmt.Printf("Subscriber %s received notification: %s\n", u.Id, msg)
		case <-u.Quit:
			fmt.Printf("Subscriber %s unsubscribed.\n", u.Id)
			return
		}
	}
}

type Auction struct {
	Id             string
	Subscribers    map[string]*User
	Mu             sync.Mutex
	HighestBid     float64
	HighestBidder  string
	Starting_price float64
}

func NewAuction(Id string, startingPrice float64) *Auction {
	return &Auction{
		Id:             Id,
		Subscribers:    make(map[string]*User),
		Starting_price: startingPrice,
	}
}
func (a *Auction) Subscribe(u *User) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.Subscribers[u.Id] = u
}
func (a *Auction) Unsubscribe(u *User) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	delete(a.Subscribers, u.Id)
	close(u.Quit)
}

func (a *Auction) NotifyAll(msg string) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	for _, sub := range a.Subscribers {
		select {
		case sub.Message <- msg:
		default:
			fmt.Printf("Subscriber %s is not ready, skipping.\n", sub.Id)
		}
	}
}

func (a *Auction) PlaceBid(bidder string, amount float64) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	if a.HighestBid < amount {
		a.HighestBid = amount
		a.HighestBidder = bidder
		go a.NotifyAll(fmt.Sprintf("New highest bid by %s is $%f", bidder, amount))
	} else {
		fmt.Printf("Bid by %s of $%f is too low.\n", bidder, amount)
	}

}

func Init() {
	auction := NewAuction("1", 100)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		sub := NewUser(fmt.Sprintf("User%d", i))
		auction.Subscribe(sub)
		wg.Add(1)
		go sub.Listen(&wg)
	}
	auction.PlaceBid("Mimi", 110)
	time.Sleep(500 * time.Millisecond)
	auction.PlaceBid("Avi", 120)
	time.Sleep(500 * time.Millisecond)
	auction.PlaceBid("Nidhi", 115)
	time.Sleep(500 * time.Millisecond)

	auction.Unsubscribe(auction.Subscribers["User2"])
	auction.PlaceBid("Avi", 122)
	time.Sleep(500 * time.Millisecond)

	for _, sub := range auction.Subscribers {
		close(sub.Quit)
	}

	wg.Wait()
}
