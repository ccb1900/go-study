package effective_go

import "testing"

func TestName(t *testing.T) {
	who := Who{}
	who.SetOwner("hello")
	tables := []int{11, 9}
	for _, v := range tables {
		a, e := who.testError(v)
		if e != nil {
			switch e.(type) {
			case *OwnerError:
				t.Log(e)
				break
			default:
				t.Log("没啥问题")
				break
			}
		}
		t.Log(a, e)
	}

}
