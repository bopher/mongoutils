package mongoutils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// In generate $in map {k: {$in: v}}
func In(k string, v interface{}) primitive.M {
	return primitive.M{k: primitive.M{"$in": v}}
}

// Set generate simple set map  {$set: v}
func Set(v interface{}) primitive.M {
	return primitive.M{"$set": v}
}

// SetNested generate nested set map {$set: {k: v}}
func SetNested(k string, v interface{}) primitive.M {
	return primitive.M{"$set": primitive.M{k: v}}
}

// Match generate nested set map {$match: v}
func Match(v interface{}) primitive.M {
	return primitive.M{"$match": v}
}
