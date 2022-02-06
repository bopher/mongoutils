package mongoutils

import "time"

// base model with timestamp and util functions
type Model struct {
	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at"`
}

// IsEditable check if document is editable
func (this Model) IsEditable() bool {
	return true
}

// IsDeletable check if document is deletable
func (this Model) IsDeletable() bool {
	return false
}

// Cleanup document before save
func (this *Model) Cleanup() {}

// PrepareInsert document before save
func (this *Model) PrepareInsert() {
	this.CreatedAt = time.Now().UTC()
}

// PrepareUpdate document before save
//
// in ghost mode UpdatedAt field not changed
func (this *Model) PrepareUpdate(ghost bool) {
	if !ghost {
		now := time.Now().UTC()
		this.UpdatedAt = &now
	}
}

// PrepareDelete update/delete related document before delete
func (this *Model) PrepareDelete() {}
