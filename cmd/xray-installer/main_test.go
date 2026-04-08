package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseCLI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    []string
		want    cliOptions
		wantErr string
	}{
		{
			name: "install default",
			args: nil,
			want: cliOptions{command: commandInstall},
		},
		{
			name: "install dry-run",
			args: []string{"--dry-run"},
			want: cliOptions{command: commandInstall, dryRun: true},
		},
		{
			name: "uninstall dry-run",
			args: []string{"uninstall", "--dry-run"},
			want: cliOptions{command: commandUninstall, dryRun: true},
		},
		{
			name: "help",
			args: []string{"--help"},
			want: cliOptions{command: commandInstall, help: true},
		},
		{
			name:    "unexpected arg",
			args:    []string{"example.com"},
			wantErr: "unexpected arguments",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCLI(tt.args)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("parseCLI error = %v, want substring %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseCLI returned error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("parseCLI = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPromptInstallRequest(t *testing.T) {
	t.Parallel()

	input := strings.NewReader("xx.example.com\n我的节点\n")
	var output bytes.Buffer

	req, err := promptInstallRequest(input, &output)
	if err != nil {
		t.Fatalf("promptInstallRequest returned error: %v", err)
	}

	if req.Domain != "xx.example.com" {
		t.Fatalf("Domain = %q, want xx.example.com", req.Domain)
	}
	if req.NodeName != "我的节点" {
		t.Fatalf("NodeName = %q, want 我的节点", req.NodeName)
	}
	if !strings.Contains(output.String(), "请输入域名") || !strings.Contains(output.String(), "请输入 FlClash 节点名称") {
		t.Fatalf("prompt output missing expected labels: %q", output.String())
	}
}
