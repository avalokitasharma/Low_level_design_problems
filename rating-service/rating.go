package ratingervice

import (
	"errors"
	"sync"
	"time"
)

type Rating struct {
	ItemID string
	UserID string
	Rating int
}

type RatingService struct {
	mu             sync.RWMutex
	ratings        map[string]map[string]int // item_id -> (user_id -> rating)
	averageRatings map[string]float64        // item_id -> average rating
}

func NewRatingService() *RatingService {
	rs := &RatingService{
		ratings:        make(map[string]map[string]int),
		averageRatings: make(map[string]float64),
	}
	go rs.updateAverageRatings()
	return rs
}

// AddOrUpdateRating adds or updates a rating for a given item by a user.
func (rs *RatingService) AddOrUpdateRating(itemID, userID string, rating int) error {
	if rating < 1 || rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	rs.mu.Lock()
	defer rs.mu.Unlock()

	if _, exists := rs.ratings[itemID]; !exists {
		rs.ratings[itemID] = make(map[string]int)
	}

	rs.ratings[itemID][userID] = rating
	return nil
}

// DeleteRating deletes a user's rating for an item.
func (rs *RatingService) DeleteRating(itemID, userID string) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if _, exists := rs.ratings[itemID]; !exists {
		return errors.New("item not found")
	}
	if _, exists := rs.ratings[itemID][userID]; !exists {
		return errors.New("rating not found for user")
	}

	delete(rs.ratings[itemID], userID)
	if len(rs.ratings[itemID]) == 0 {
		delete(rs.ratings, itemID)
		delete(rs.averageRatings, itemID)
	}
	return nil
}

// GetAverageRating returns the average rating and total number of ratings for an item.
func (rs *RatingService) GetAverageRating(itemID string) (float64, int, error) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	if _, exists := rs.ratings[itemID]; !exists {
		return 0, 0, errors.New("item not found")
	}

	sum := 0
	numRatings := 0
	for _, rating := range rs.ratings[itemID] {
		sum += rating
		numRatings++
	}

	average := float64(sum) / float64(numRatings)
	return average, numRatings, nil
}

// GetCachedAverageRating returns the precomputed average rating for an item.
func (rs *RatingService) GetCachedAverageRating(itemID string) (float64, error) {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	average, exists := rs.averageRatings[itemID]
	if !exists {
		return 0, errors.New("item not found")
	}

	return average, nil
}

// updateAverageRatings updates the average ratings in the background.
func (rs *RatingService) updateAverageRatings() {
	for {
		func() {
			rs.mu.Lock()
			defer rs.mu.Unlock()

			for itemID, userRatings := range rs.ratings {
				sum := 0
				count := 0
				for _, rating := range userRatings {
					sum += rating
					count++
				}
				if count > 0 {
					rs.averageRatings[itemID] = float64(sum) / float64(count)
				} else {
					delete(rs.averageRatings, itemID)
				}
			}
		}()

		time.Sleep(1 * time.Second)
	}
}
