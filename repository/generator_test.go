package repository

import (
	"fmt"
	"testing"
)

func TestSourceGenerator_Generate(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name       string
		g          *SourceGenerator
		args       args
		wantTarget string
		wantErr    bool
	}{
		{
			name:    "ok",
			g:       &SourceGenerator{},
			args:    args{n: 7},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &SourceGenerator{}
			gotTarget, err := g.Generate(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("sourceGenerator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("got = %v\n", gotTarget)
			if len(gotTarget) != tt.args.n {
				t.Errorf("sourceGenerator.Generate() length = %v, want %v", len(gotTarget), tt.args.n)
			}
		})
	}
}
