## Rating Service Requirements
 - Implement a service to handle ratings for items (e.g., products, services, or movies).
 - Support CRUD operations for ratings: Create, Read (average rating for an item), Update, and Delete.
 - Ratings should have attributes: item_id (unique identifier), user_id (unique identifier), rating (1 to 5 scale).
 - Ensure that each user can rate an item only once, with updates overwriting the existing rating.
 - Provide an API to fetch the average rating and the total number of ratings for an item.
 - Handle edge cases such as invalid inputs, items with no ratings, and concurrent updates gracefully.
 - Optimize for scalability with a focus on efficient read operations for fetching average ratings.