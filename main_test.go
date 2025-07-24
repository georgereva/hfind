package main

import (
	"os"
	"testing"
)

func Test_getHistoryFilePath(t *testing.T) {
	tests := []struct {
		name    string
		homeDir string
		shell   string
		want    string
		wantErr bool
	}{
		{
			"Get filepath from bash",
			"/home/user",
			"/usr/bin/bash",
			"/home/user/.bash_history",
			false,
		},
		{
			"Get filepath from zsh",
			"/home/user",
			"/usr/bin/zsh",
			"/home/user/.zsh_history",
			false,
		},
		{
			"Unknown shell",
			"/home/user",
			"/usr/bin/test",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			errHome := os.Setenv("HOME", tt.homeDir)
			if errHome != nil {
				return
			}
			errShell := os.Setenv("SHELL", tt.shell)
			if errShell != nil {
				return
			}

			got, err := getHistoryFilePath()
			if (err != nil) != tt.wantErr {
				t.Errorf("getHistoryFilePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getHistoryFilePath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
