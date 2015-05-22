package math

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddOne(t *testing.T) {
	require.Equal(t, 1, AddOne(0))
	require.Equal(t, 0, AddOne(-1))
}

func TestMultiplyByTwo(t *testing.T) {
	require.Equal(t, 0, MultiplyByTwo(0))
	require.Equal(t, 4, MultiplyByTwo(2))
}
