package vo

import (
	"fmt"
	"testing"
)

func TestNewMoney(t *testing.T) {
	m := NewMoney(100)
	m.amount = 101
	fmt.Println(m)
}
