package main

import (
	"domain/vo"
	"fmt"
	"testing"
)

func TestNewMoney(t *testing.T) {
	m := vo.NewMoney(100)
	// m.amount = 101
	fmt.Println(m)
}
