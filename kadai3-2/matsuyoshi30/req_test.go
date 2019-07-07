package godl

import "testing"

func TestMakeRequest(t *testing.T) {
	testcases := []struct {
		name  string
		input *Godl
		// expectedReq []*http.Request
		expectedErr error
	}{
		{"normal", &Godl{url: "https://example.com", ranges: []Range{{0, 600}, {601, 1270}}}, nil},
		{"error", &Godl{url: "https://example.com", ranges: []Range{{600, 0}, {601, 1270}}}, nil},
	}

	for _, tt := range testcases {
		g := tt.input
		_, err := g.MakeRequest()
		if err != nil {
			if tt.name != "error" {
				t.Fatalf("unexpected error")
			}
		}
	}
}
