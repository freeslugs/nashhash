package gm

import (
	"testing"
)

func TestOperatorInit(t *testing.T) {
	var gop GameOperator
	gop.Init("0x1")

	assertEqual(t, gop.ContractAddress(), "0x1", "Wrong contract address")

}
