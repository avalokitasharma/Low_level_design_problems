package shoppingcartservice

import "time"

type expiresAt time.Time

type ReservedItems struct {
	locks map[CartItem]expiresAt
}

func CreateLocks() *ReservedItems {
	return &ReservedItems{
		locks: make(map[CartItem]expiresAt),
	}
}

func (l *ReservedItems) AddLock(item CartItem) {}

func (l *ReservedItems) RemoveLock(item CartItem) {}
func (l *ReservedItems) CleanUp()                 {}
