package installer

import "testing"

func TestNormalizeNodeName(t *testing.T) {
	t.Parallel()

	got, err := normalizeNodeName("")
	if err != nil {
		t.Fatalf("normalizeNodeName returned error: %v", err)
	}
	if got != DefaultNodeName {
		t.Fatalf("normalizeNodeName default = %q, want %q", got, DefaultNodeName)
	}
}

func TestNormalizeNodeNameRejectsNewline(t *testing.T) {
	t.Parallel()

	if _, err := normalizeNodeName("bad\nname"); err == nil {
		t.Fatal("normalizeNodeName should reject newlines")
	}
}

func TestSupportsDistro(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		id   string
		like string
		want bool
	}{
		{name: "ubuntu", id: "ubuntu", want: true},
		{name: "debian", id: "debian", want: true},
		{name: "archlinux id", id: "archlinux", want: true},
		{name: "arch id", id: "arch", want: true},
		{name: "arch like", id: "manjaro", like: "archlinux", want: true},
		{name: "unsupported", id: "alpine", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := supportsDistro(tt.id, tt.like); got != tt.want {
				t.Fatalf("supportsDistro(%q, %q) = %v, want %v", tt.id, tt.like, got, tt.want)
			}
		})
	}
}
