package installer

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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

func TestPrintSummaryIncludesFlClashCommand(t *testing.T) {
	t.Parallel()

	var stdout bytes.Buffer
	inst := New(&stdout, &stdout)
	inst.printSummary(&Result{
		Domain:          "xx.example.com",
		NodeName:        DefaultNodeName,
		PublicIP:        "1.2.3.4",
		Port:            443,
		UUID:            "uuid",
		PublicKey:       "pub",
		ShortID:         "short",
		XrayConfigPath:  xrayConfigPath,
		ProxyConfigPath: proxyConfigPath,
	})

	if !strings.Contains(stdout.String(), "sudo cat "+proxyConfigPath) {
		t.Fatalf("summary missing flclash fetch command: %q", stdout.String())
	}
}

func TestPrintDryRunSummaryIncludesFlClashCommandHint(t *testing.T) {
	t.Parallel()

	var stdout bytes.Buffer
	inst := New(&stdout, &stdout)
	inst.printDryRunSummary(&Result{
		Domain:          "xx.example.com",
		NodeName:        DefaultNodeName,
		PublicIP:        "1.2.3.4",
		XrayConfigPath:  xrayConfigPath,
		ProxyConfigPath: proxyConfigPath,
	})

	if !strings.Contains(stdout.String(), "sudo cat "+proxyConfigPath) {
		t.Fatalf("dry-run summary missing flclash fetch command hint: %q", stdout.String())
	}
}

func TestWriteFileAtomicUsesRequestedMode(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	target := filepath.Join(dir, "config.json")
	if err := writeFileAtomic(target, []byte(`{"ok":true}`), xrayConfigMode); err != nil {
		t.Fatalf("writeFileAtomic returned error: %v", err)
	}

	info, err := os.Stat(target)
	if err != nil {
		t.Fatalf("stat target: %v", err)
	}
	if got := info.Mode().Perm(); got != xrayConfigMode {
		t.Fatalf("file mode = %o, want %o", got, xrayConfigMode)
	}
}
