package effective_go

import (
	"testing"
)

func TestMyIf(t *testing.T) {
	t.Log(myIf(11))
	t.Log(myIf(-11))
}

func TestLoop(t *testing.T) {
	loop(100)
	rangeChannel(mergeCh(wCh(), wCh(), wCh()))
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '+', '=':
		return false
	}

	return true
}
