package plugins

import (
	"fmt"
	"testing"
)

func TestCullTwoArraysString(t *testing.T) {
	origin := []string{"a", "b", "c", "d"}
	input := []string{"c", "d", "e"}
	end := CullTwoArraysString(origin, input)

	fmt.Println("end:", end, " done")
}
