package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Monitor AOC", 1000.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Monitor AOC", p.Name)
	assert.Equal(t, 1000.0, p.Price)

}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 1000.0)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Monitor AOC", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Monitor AOC", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}
