package mongoutils

import (
	"context"
	"time"
)

// base model with timestamp and util functions
type Model struct {
	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at"`
}

// IsEditable check if document is editable
//
// by default returns true
func (me Model) IsEditable() bool {
	return true
}

// IsDeletable check if document is deletable
//
// by default returns false
func (me Model) IsDeletable() bool {
	return false
}

// BeforeInsert function to call before insert
func (me *Model) BeforeInsert(ctx context.Context) {}

// AfterInsert function to call after insert
func (me Model) AfterInsert(ctx context.Context) {}

// BeforeUpdate function to call before update
func (me *Model) BeforeUpdate(ctx context.Context) {}

// AfterUpdate function to call after update
func (me Model) AfterUpdate(old any, ctx context.Context) {}

// BeforeDelete function to call before delete
func (me *Model) BeforeDelete(ctx context.Context) {}

// AfterDelete function to call after delete
func (me Model) AfterDelete(ctx context.Context) {}

// Cleanup document before save
//
// e.g set relation document to nil for ignore saving
func (me *Model) Cleanup() {}

// PrepareInsert fill created_at before insert
func (me *Model) PrepareInsert() {
	me.CreatedAt = time.Now().UTC()
}

// PrepareUpdate fill updated_at before insert
//
// in ghost mode updated_at field not changed
func (me *Model) PrepareUpdate(ghost bool) {
	if !ghost {
		now := time.Now().UTC()
		me.UpdatedAt = &now
	}
}
