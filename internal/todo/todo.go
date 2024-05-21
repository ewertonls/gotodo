package todo

import (
	"time"

	"github.com/uptrace/bun"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos"`
	ID            uint      `bun:",pk,autoincrement" json:"id"`
	Title         string    `json:"title,omitempty"`
	Description   string    `json:"description,omitempty"`
	Completed     bool      `json:"completed"`
	CompletedAt   time.Time `json:"completedAt,omitempty"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt,omitempty"`
}

func (t *Todo) MarkAsCompleted() {
	now := time.Now()
	t.Completed = true
	t.CompletedAt = now
	t.UpdatedAt = now
}

func (t *Todo) MarkAsUncompleted() {
	now := time.Now()
	t.Completed = false
	t.CompletedAt = time.Time{}
	t.UpdatedAt = now
}
