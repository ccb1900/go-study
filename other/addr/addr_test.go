package addr

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	a := 1
	test(&a)
	test(&a)
}

func TestTranslate(t *testing.T) {
	//input := []string{"# a", "## b", "## c", "### d", "# e"}
	input2 := []string{"# a", "## b", "## c", "### d", "# e"}

	fmt.Println(translate(input2))
}
