package publicdns

import (
	"fmt"
	"testing"
)

func TestResolve(t *testing.T) {
	for i, tt := range []struct {
		name   string
		rrType RRType
		status int
		rd     bool
		ra     bool
	}{
		{"example.com", A, 0, true, true},
		{"example.net", AAAA, 0, true, true},
		{"example.org", ANY, 0, true, true},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			r, err := Resolve(tt.name, tt.rrType)
			if err != nil {
				t.Fatalf("Resolve(%q, %q) returned unexpected error: %v", tt.name, tt.rrType, err)
			}

			if got, want := r.Status, tt.status; got != want {
				t.Fatalf("r.Status = %d, want %d", got, want)
			}

			if got, want := r.RD, tt.rd; got != want {
				t.Fatalf("r.RD = %v, want %v", got, want)
			}

			if got, want := r.RA, tt.ra; got != want {
				t.Fatalf("r.RA = %v, want %v", got, want)
			}
		})
	}
}

func TestRawurl(t *testing.T) {
	for i, tt := range []struct {
		name   string
		rrType RRType
		want   string
	}{
		{"example.com", A, "https://dns.google.com/resolve?name=example.com&type=A"},
		{"example.net", AAAA, "https://dns.google.com/resolve?name=example.net&type=AAAA"},
		{"example.org", ANY, "https://dns.google.com/resolve?name=example.org&type=ANY"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got, want := rawurl(tt.name, tt.rrType), tt.want; got != want {
				t.Fatalf("rawurl(%q, %q) = %q, want %q", tt.name, tt.rrType, got, want)
			}
		})
	}
}
