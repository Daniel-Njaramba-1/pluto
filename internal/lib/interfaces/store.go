package interfaces

// Store defines all methods needed for a data model
type Store interface {
	FeedGetID() *int
	
	// ✅ CRUD : All Models
	FeedCreateQuery() string
	FeedGetByIDQuery() string
	FeedGetAllQuery() string
	FeedUpdateDetailsQuery() string
	FeedDeleteQuery() string
	
	// ✅ Soft Delete (IsActive): Category, SubCategory, Product, Admin, Customer, Cart
	FeedDeactivateQuery() string
	FeedReactivateQuery() string
}
