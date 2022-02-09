package mongoutils

import "time"

// base model with timestamp and util functions
type Model struct {
	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at"`
}

// IsEditable check if document is editable
//
// by default returns true
func (this Model) IsEditable() bool {
	return true
}

// IsDeletable check if document is deletable
//
// by default returns false
func (this Model) IsDeletable() bool {
	return false
}

// BeforeInsert function to call before insert
func (this *Model) BeforeInsert() {}

// AfterInsert function to call after insert
func (this *Model) AfterInsert() {}

// BeforeUpdate function to call before update
func (this *Model) BeforeUpdate() {}

// AfterUpdate function to call after update
func (this *Model) AfterUpdate() {}

// BeforeDelete function to call before delete
func (this *Model) BeforeDelete() {}

// AfterDelete function to call after delete
func (this *Model) AfterDelete() {}

// Cleanup document before save
//
// e.g set relation document to nil for ignore saving
func (this *Model) Cleanup() {}

// PrepareInsert fill created_at before insert
func (this *Model) PrepareInsert() {
	this.CreatedAt = time.Now().UTC()
}

// PrepareUpdate fill updated_at before insert
//
// in ghost mode updated_at field not changed
func (this *Model) PrepareUpdate(ghost bool) {
	if !ghost {
		now := time.Now().UTC()
		this.UpdatedAt = &now
	}
}
