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

func (this *mPipe) Lookup(from string, local string, foreign string, as string) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		d.Doc("$lookup", func(d MongoDoc) MongoDoc {
			d.
				Add("from", from).
				Add("localField", local).
				Add("foreignField", foreign).
				Add("as", as)
			return d
		})
		return d
	})
}

func (this *mPipe) Unwind(path string, prevNullAndEmpty bool) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		d.Doc("$unwind", func(d MongoDoc) MongoDoc {
			d.
				Add("path", path).
				Add("preserveNullAndEmptyArrays", prevNullAndEmpty)
			return d
		})
		return d
	})
}

func (this *mPipe) Unwrap(field string, as string) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		d.Doc("$addFields", func(d MongoDoc) MongoDoc {
			d.
				Nested(as, "$first", field)
			return d
		})
		return d
	})
}

func (this *mPipe) Group(cb func(d MongoDoc) MongoDoc) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		d.Doc("$group", cb)
		return d
	})
}

func (this *mPipe) ReplaceRoot(v interface{}) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		d.Doc("$replaceRoot", func(d MongoDoc) MongoDoc {
			d.Add("newRoot", v)
			return d
		})
		return d
	})
}

func (this *mPipe) MergeRoot(fields ...interface{}) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		d.Doc("$replaceRoot", func(d MongoDoc) MongoDoc {
			d.Nested("newRoot", "$mergeObjects", fields)
			return d
		})
		return d
	})
}

func (this *mPipe) UnProject(fields ...string) MongoPipeline {
	return this.Add(func(d MongoDoc) MongoDoc {
		d.Doc("$project", func(d MongoDoc) MongoDoc {
			for _, v := range fields {
				d.Add(v, 0)
			}
			return d
		})
		return d
	})
}

func (this mPipe) Build() mongo.Pipeline {
	return this.data
}
