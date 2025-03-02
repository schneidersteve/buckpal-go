package vo

type money struct {
	amount int64
}

func NewMoney(amount int64) *money {
	return &money{amount: amount}
}
