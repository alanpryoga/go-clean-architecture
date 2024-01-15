package domain

type Money struct {
	Currency string
	Value    int64
}

func (m Money) isCurrencyMatch(other Money) error {
	if m.Currency != other.Currency {
		return ErrCurrencyMismatch
	}
	return nil
}

func (m Money) Add(other Money) (Money, error) {
	if err := m.isCurrencyMatch(other); err != nil {
		return Money{}, err
	}

	return Money{
		Currency: m.Currency,
		Value:    m.Value + other.Value,
	}, nil
}

func (m Money) Sub(other Money) (Money, error) {
	if err := m.isCurrencyMatch(other); err != nil {
		return Money{}, err
	}

	return Money{
		Currency: m.Currency,
		Value:    m.Value - other.Value,
	}, nil
}
