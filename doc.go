package mongoutils

import "go.mongodb.org/mongo-driver/bson/primitive"

// MongoDoc mongo document (primitive.D) builder
type MongoDoc interface {
	// Add add new element
	Add(k string, v interface{}) MongoDoc
	// Doc add new element with nested doc value
	Doc(k string, cb func(d MongoDoc) MongoDoc) MongoDoc
	// Array add new element with array value
	Array(k string, v ...interface{}) MongoDoc
	// Nested add new nested element
	Nested(root string, k string, v interface{}) MongoDoc
	// NestedDoc add new nested element with doc value
	NestedDoc(root string, k string, cb func(d MongoDoc) MongoDoc) MongoDoc
	// NestedArray add new nested element with array value
	NestedArray(root string, k string, v ...interface{}) MongoDoc
	// Regex add new element with regex value
	Regex(k string, pattern string, opt string) MongoDoc
	// Map creates a map from the elements of the Doc
	Map() primitive.M
	// Build generate mongo doc
	Build() primitive.D
}