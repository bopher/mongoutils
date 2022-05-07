package mongoutils

import "go.mongodb.org/mongo-driver/bson/primitive"

type MtCounter interface {
	// Add new meta
	Add(_col, _meta string, id *primitive.ObjectID, amount int) MtCounter
	// Result get combined meta with query
	Result() []MetaQuery
}

type MetaQuery struct {
	Col string
	Ids []primitive.ObjectID
	// data to update
	Update map[string]int
}
