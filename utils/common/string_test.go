package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils_IntToString(t *testing.T) {
	resp := IntToString(4)
	expectResp := "4"
	assert.Equal(t, expectResp, resp)
}