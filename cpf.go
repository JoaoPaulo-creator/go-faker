package cpf

import (
	"math/rand"
	"strconv"
)

type Cpf struct {
	nums string
}

func (c *Cpf) generateFirstDigit() {
	d1 := 0
	k := 10

	for i := 0; i < len(c.nums); i, k = i+1, k-1 {
		num, _ := strconv.Atoi(c.nums[i : i+1])
		d1 += num * k
	}

	if d1%11 < 2 {
		c.nums += "0"
	} else {
		c.nums += strconv.Itoa(11 - (d1 % 11))
	}
}

func (c *Cpf) generateSecondDigit() {
	d2 := 0
	k := 10

	for i := 0; i < len(c.nums)-1; i, k = i+1, k-1 {
		num, _ := strconv.Atoi(c.nums[i+1 : i+2])
		d2 += num * k
	}

	if d2%11 < 2 {
		c.nums += "0"
	} else {
		c.nums += strconv.Itoa(11 - (d2 % 11))
	}
}

// Gera um novo cpf fake
func New() *Cpf {
	c := &Cpf{}

	for range 9 {
		c.nums += strconv.Itoa(rand.Intn(9))
	}

	c.generateFirstDigit()
	c.generateSecondDigit()

	return c
}
