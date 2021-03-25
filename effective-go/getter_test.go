package effective_go

import "testing"

func TestString(t *testing.T) {
	who := &Who{}
	who.SetOwner("hello")
	t.Log(who)
}
