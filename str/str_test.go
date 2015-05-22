package str

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNormalize(t *testing.T) {
	require.Equal(t, "test", Normalize("  TEST  "))
}

func TestReverse(t *testing.T) {
	require.Equal(t, "", Reverse(""))
	require.Equal(t, "cba", Reverse("abc"))
}
