package mongoutils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mDoc struct {
	data primitive.D
}

func (this *mDoc) Add(k string, v interface{}) MongoDoc {
	this.data = append(this.data, primitive.E{Key: k, Value: v})
	return this
}

func (this *mDoc) Doc(k string, cb func(d MongoDoc) MongoDoc) MongoDoc {
	return this.Add(k, cb(NewDoc()).Build())
}

func (this *mDoc) Array(k string, v ...interface{}) MongoDoc {
	return this.Add(k, v)
}

func (this *mDoc) Nested(root string, k string, v interface{}) MongoDoc {
	return this.Add(root, primitive.M{k: v})
}

func (this *mDoc) NestedDoc(root string, k string, cb func(d MongoDoc) MongoDoc) MongoDoc {
	return this.Add(root, primitive.M{k: cb(NewDoc()).Build()})
}

func (this *mDoc) NestedArray(root string, k string, v ...interface{}) MongoDoc {
	return this.Add(root, primitive.M{k: v})
}

func (this *mDoc) Regex(k string, pattern string, opt string) MongoDoc {
	return this.Add(k, primitive.Regex{Pattern: pattern, Options: opt})
}

func (this mDoc) Map() primitive.M {
	return this.data.Map()
}

func (this mDoc) Build() primitive.D {
	return this.data
}
