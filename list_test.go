package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateList(t *testing.T) {
	l := NewList()
	err := l.Append(1)
	assert.Nil(t, err)
}
