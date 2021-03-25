package effective_go

import "testing"

func TestData(t *testing.T) {
	s := *new(string)
	n := new(chan string)
	ns := new([]int)
	//nm := new(map[])
	f := *new(float64)
	ed := new(EmptyData)
	ed1 := new(Data)
	edi := new(IData)
	ed2 := new(EmptyData)
	var edd EmptyData
	var eddp *EmptyData
	var i IData
	t.Log(s, f, ed, edd, eddp, i, eddp == nil, new([]int) == nil, edi, ed1, ed2, n, ns)
}
