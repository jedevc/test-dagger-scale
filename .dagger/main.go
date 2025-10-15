package main

import (
	"context"
	"strconv"
	"strings"
)

type Foo struct {
	Initial int
}

func New(
	// +optional
	initial int,
) *Foo {
	return &Foo{
		Initial: initial,
	}
}

func (m *Foo) Test(n int) int {
	n += m.Initial
	return n
}

func (m *Foo) ScaleTest(n int) int {
	return m.Test(n)
}

func (m *Foo) BusyWork(ctx context.Context, n int) (string, error) {
	n += m.Initial
	out, err := dag.Container().
		From("alpine:latest").
		WithExec([]string{"sh", "-c", "for i in $(seq 1 " + strconv.Itoa(n) + "); do sleep 1; echo $i; done"}).
		WithExec([]string{"sh", "-c", "ls -lh /"}).
		Stdout(ctx)
	return strings.TrimSpace(out), err
}

func (m *Foo) ScaleBusyWork(ctx context.Context, n int) (string, error) {
	return m.BusyWork(ctx, n)
}
