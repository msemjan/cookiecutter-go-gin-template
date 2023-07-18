package models

type BaseModel struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (m *BaseModel) String() string {
	return fmt.Sprintf("(ID: %v, Created: %v, Modified: %v, Deleted: %v)",
		m.ID, m.Created.Format("2006-01-02 15:04:05"),
		m.Modified.Format("2006-01-02 15:04:05"),
		m.Deleted.Format("2006-01-02 15:04:05"),
	)
}

func (m *BaseModel) IsDeleted() bool {
	return !m.Deleted.IsZero()
}
