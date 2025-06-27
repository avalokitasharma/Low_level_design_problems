package ratingervice

import (
	"fmt"
	"time"
)

func Init() {
	rs := NewRatingService()

	// Adding/updating ratings
	rs.AddOrUpdateRating("item1", "user1", 5)
	rs.AddOrUpdateRating("item1", "user2", 4)
	rs.AddOrUpdateRating("item1", "user1", 3) // Update user1's rating

	// Fetching average rating (real-time computation)
	average, total, err := rs.GetAverageRating("item1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Real-time Average Rating: %.2f, Total Ratings: %d\n", average, total)
	}

	// Fetching cached average rating
	time.Sleep(6 * time.Second) // Wait for the background thread to update
	cachedAverage, err := rs.GetCachedAverageRating("item1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cached Average Rating: %.2f\n", cachedAverage)
	}

	// Deleting a rating
	if err := rs.DeleteRating("item1", "user2"); err != nil {
		fmt.Println("Error:", err)
	}

	// Fetching cached average rating after deletion
	time.Sleep(6 * time.Second) // Wait for the background thread to update
	cachedAverage, err = rs.GetCachedAverageRating("item1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cached Average Rating After Deletion: %.2f\n", cachedAverage)
	}
}
