package domain

type MerchantID string

var (
	ErrRecentNameChanged = errors.New("recent name changed")
)

type Merchant struct {
	ID MerchantID
	Name string
	LastNameChangedDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Merchant) ChangeName(name string) error {
	now := timeutil.Now()
	before30Days = now.Sub(30 * time.Day)
	if m.LastNameChangedDate.After(before30Days) {
		return ErrRecentNameChanged
	}

	m.Name = name
	m.LastNameChangedDate = now
	m.UpdatedAt = now
	return nil
}