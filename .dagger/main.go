package main

type Foo struct {
	A int
	B int
}

func (m *Foo) Test(n int) int {
	return m.A + m.B + n
}
