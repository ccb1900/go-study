package helloworld

import "testing"

func TestHello(t *testing.T) {
	want := "hello world"
	got := hello(want)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	data := []string{"hello", "string"}

	for _, datum := range data {
		t.Run(datum+"111", func(t *testing.T) {
			want := datum
			got := hello(want)

			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}

}
