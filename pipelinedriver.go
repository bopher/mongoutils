package mongoutils

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type mPipe struct {
	data mongo.Pipeline
}

func (this *mPipe) Add(cb func(d MongoDoc) MongoDoc) MongoPipeline {
	this.data = append(this.data, cb(NewDoc()).Build())
	return this
}

func (this *mPipe) Match(filters interface{}) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		return d.Add("$match", filters)
	})
}

func (this *mPipe) In(key string, v interface{}) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		return d.Nested(key, "$in", v)
	})
}

func (this *mPipe) Limit(limit int64) MongoPipeline {
	if limit > 0 {
		this.Add(func(d MongoDoc) MongoDoc {
			return d.Add("$limit", limit)
		})
	}
	return this
}

func (this *mPipe) Skip(skip int64) MongoPipeline {
	if skip > 0 {
		this.Add(func(d MongoDoc) MongoDoc {
			return d.Add("$skip", skip)
		})
	}
	return this
}

func (this *mPipe) Sort(sorts interface{}) MongoPipeline {
	if sorts != nil {
		this.Add(func(d MongoDoc) MongoDoc {
			return d.Add("$sort", sorts)
		})
	}
	return this
}

func (this *mPipe) Unwind(path string, prevNullAndEmpty bool) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		return d.Doc("$unwind", func(d MongoDoc) MongoDoc {
			return d.
				Add("path", path).
				Add("preserveNullAndEmptyArrays", prevNullAndEmpty)
		})
	})
}

func (this *mPipe) Lookup(from string, local string, foreign string, as string) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		return d.Doc("$lookup", func(d MongoDoc) MongoDoc {
			return d.
				Add("from", from).
				Add("localField", local).
				Add("foreignField", foreign).
				Add("as", as)
		})
	})
}

func (this *mPipe) Unwrap(field string, as string) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		return d.Doc("$addFields", func(d MongoDoc) MongoDoc {
			return d.Nested(as, "$first", field)
		})
	})
}

func (this *mPipe) LoadRelation(from string, local string, foreign string, as string) MongoPipeline {
	this.Add(func(d MongoDoc) MongoDoc {
		return d.Doc("$lookup", func(d MongoDoc) MongoDoc {
			return d.
				Add("from", from).
				Add("localField", local).
				Add("foreignField", foreign).
				Add("as", as)
		})
	})
	this.Add(func(d MongoDoc) MongoDoc {
		return d.Doc("$addFields", func(d MongoDoc) MongoDoc {
			return d.Nested(as, "$first", "$"+as)
		})
	})
	return this
}

func (this *mPipe) Group(cb func(d MongoDoc) MongoDoc) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		return d.Doc("$group", cb)
	})
}

func (this *mPipe) ReplaceRoot(v interface{}) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		return d.Doc("$replaceRoot", func(d MongoDoc) MongoDoc {
			return d.Add("newRoot", v)
		})
	})
}

func (this *mPipe) MergeRoot(fields ...interface{}) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		return d.Doc("$replaceRoot", func(d MongoDoc) MongoDoc {
			return d.Nested("newRoot", "$mergeObjects", fields)
		})
	})
}

func (this *mPipe) UnProject(fields ...string) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		return d.Doc("$project", func(d MongoDoc) MongoDoc {
			for _, v := range fields {
				d.Add(v, 0)
			}
			return d
		})
	})
}

func (this mPipe) Build() mongo.Pipeline {
	return this.data
}
