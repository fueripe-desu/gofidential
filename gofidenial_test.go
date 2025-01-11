package gofidential

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sum(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	x := 10
	y := 20
	expected := 30

	// Act
	result := Sum(x, y)

	// Assert
	assert.Equal(result, expected)
}
