package testutil

import "github.com/stretchr/testify/mock"

func ExpectType[T any]() any {
	return mock.MatchedBy(
		func(T) bool { return true },
	)
}

func PointerOf[T any](v T) *T {
	return &v
}
