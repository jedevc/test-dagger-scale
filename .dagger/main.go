package main

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"dagger/foo/internal/dagger"
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

func (m *Foo) CheckFiles(
	ctx context.Context,
	// +defaultPath="."
	dir *dagger.Directory,
) error {
	entries, err := dir.Entries(ctx)
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		return fmt.Errorf("no files found in the directory")
	}
	if !slices.Contains(entries, "dagger.json") {
		return fmt.Errorf("dagger.json not found in the directory")
	}
	return nil
}
