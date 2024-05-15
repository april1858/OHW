package main

import (
	"os"
	"testing"
)

/*
func TestCopy(t *testing.T) {
	// Place your code here.
}
*/

func TestCopy(t *testing.T) {
	f, err := os.CreateTemp("", "temp")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "1", args: args{fromPath: "testdata/input.txt", toPath: f.Name(), offset: 0, limit: 0}, wantErr: false},
		{name: "2", args: args{fromPath: "testdata/input.txt", toPath: f.Name(), offset: 10, limit: 110}, wantErr: false},
		{name: "3", args: args{fromPath: "testdata/input.txt", toPath: f.Name(), offset: 7000, limit: 0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.fromPath, tt.args.toPath, tt.args.offset, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
