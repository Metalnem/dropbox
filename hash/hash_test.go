package hash

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		size int
		hash string
	}{
		{0, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{2097152, "a092b1804f2a694edacf3f5b9e60dc5bdbd2018fa4d03f3faf2e42014ea87daa"},
		{4194303, "839e2d868a885f260431d186466f0f3c189ee0eee5cc444b54e717740ad7f1c3"},
		{4194304, "c7e946d101855255d919ef0c70718633adf77d3dfb3adeeecf5d0cb4e951be58"},
		{4194305, "14a4d47f23a30177885d9820122f17d2d3a55fe63f7f5c27b95f689e0b2accd6"},
		{6291456, "4c038ad982f42733f8cf5de38af46a51c822e758fd72497bc1187040bb42f5ca"},
		{8388607, "fb9eb4f6e396d3098a09a6557fa5e834e7a199d9629362b36a651c9a7611f35e"},
		{8388608, "03ae066c707c588592d9e27aa2444ca98423e0999024f1ceaa11a153790b37de"},
		{8388609, "0abd0f48fe25004ce9ca2e45df928ba9b1052f20666d7dde21ecdbeb7aa96482"},
	}

	for _, test := range tests {
		size := test.size
		hash := test.hash

		t.Run(fmt.Sprintf("%d-fast", size), func(t *testing.T) {
			t.Parallel()

			h := New()
			h.Write(make([]byte, size))

			got := string(h.Sum(nil))
			want := hash

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})

		t.Run(fmt.Sprintf("%d-slow", size), func(t *testing.T) {
			t.Parallel()

			h := New()
			b := []byte{0}

			for i := 0; i < size; i++ {
				h.Write(b)
			}

			got := string(h.Sum(nil))
			want := hash

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
