package jscw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	t.Parallel()

	w := Version()
	assert.Equal(t, "0.0.0", w)
}
