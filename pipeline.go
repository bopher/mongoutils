package mongoutils

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoPipeline mongo pipeline (mongo.Pipeline) builder
type MongoPipeline interface {
	// Add add new Doc
	Add(cb func(d MongoDoc) MongoDoc) MongoPipeline
	// Match add $match stage
	Match(filters interface{}) MongoPipeline
	// In add $in stage
	In(key string, v interface{}) MongoPipeline
	// Limit add $limit stage (ignore negative and zero value)
	Limit(limit int64) MongoPipeline
	// Skip add $skip stage (ignore negative and zero value)
	Skip(skip int64) MongoPipeline
	// Sort add $sort stage (ignore nil value)
	Sort(sorts interface{}) MongoPipeline
	// Unwind add $unwind stage
	Unwind(path string, prevNullAndEmpty bool) MongoPipeline
	// Lookup add $lookup stage
	Lookup(from string, local string, foreign string, as string) MongoPipeline
	// Unwrap get first item of array and insert to doc using $addFields stage
	Unwrap(field string, as string) MongoPipeline
	// LoadRelation load related document using $lookup and $addField
	LoadRelation(from string, local string, foreign string, as string) MongoPipeline
	// Group add $group stage
	Group(cb func(d MongoDoc) MongoDoc) MongoPipeline
	// ReplaceRoot add $replaceRoot stage
	ReplaceRoot(v interface{}) MongoPipeline
	// MergeRoot add $replaceRoot stage with $mergeObjects operator
	MergeRoot(fields ...interface{}) MongoPipeline
	// UnProject generate $project stage to remove fields from result
	UnProject(fields ...string) MongoPipeline
	// Build generate mongo pipeline
	Build() mongo.Pipeline
}
