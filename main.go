package mongoutils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

// FindOption generate find option with sorts params
func FindOption(sort interface{}, skip int64, limit int64) *options.FindOptions {
	opt := new(options.FindOptions)
	opt.SetAllowDiskUse(true)
	opt.SetSkip(skip)
	if limit > 0 {
		opt.SetLimit(limit)
	}
	if sort != nil {
		opt.SetSort(sort)
	}
	return opt
}

// AggregateOption generate aggregation options
func AggregateOption() *options.AggregateOptions {
	return new(options.AggregateOptions).
		SetAllowDiskUse(true)
}
