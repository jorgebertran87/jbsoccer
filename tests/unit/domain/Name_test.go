package domain

import (
	"github.com/jorgebertran87/jbsoccer/core/domain"
	"github.com/jorgebertran87/jbsoccer/tests"
	"testing"
)

const VALUE = "jeje"

func TestItReturnsName(t *testing.T) {
	name := domain.Name{Value: VALUE}

	test := tests.Test{T: t}
	test.AssertSame(VALUE, name.Value)
}

