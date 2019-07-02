package godl

import "testing"

func TestValidateURL(t *testing.T) {
	testcases := []struct {
		name  string
		input string
	}{
		{"normal", "https://example.com"},
		{"error", "https://example.comm"},
		// TODO: status code other than 200 pattern
		// TODO: doesnt support accept-ranges pattern
		// TODO: content length < 1 pattern
	}

	g := NewGodl()

	for _, tt := range testcases {
		tt := tt
		if err := g.ValidateURL(tt.input); err != nil {
			if tt.name != "error" {
				t.Fatalf("unexpected result")
			}
		}
	}
}
