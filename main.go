package mongoutils

import "go.mongodb.org/mongo-driver/bson/primitive"

// NewPipe new mongo pipe builder
func NewPipe() MongoPipeline {
	return new(mPipe)
}

// NewDoc new mongo doc builder
func NewDoc() MongoDoc {
	return new(mDoc)
}

// ParseObjectID parse object id from string
func ParseObjectID(id string) *primitive.ObjectID {
	if oId, err := primitive.ObjectIDFromHex(id); err == nil && !oId.IsZero() {
		return &oId
	}
	return nil
}

// IsValidObjectId check if object id is valid and not zero
func IsValidObjectId(id *primitive.ObjectID) bool {
	return id != nil && !id.IsZero()
}
