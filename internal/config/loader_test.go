package config

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *Webhook
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				vars: map[string]string{
					"PROCESSOR_HOST":        "fake",
					"GITHUB_WEBHOOK_SECRET": "fake",
				},
			},
			want: &Webhook{
				ProcessorHost: "fake",
				GithubSecret:  "fake",
			},
			wantErr: false,
		},
		{
			name: "failure required fields",
			args: args{
				vars: map[string]string{
					"PROCESSOR_HOST": "fake",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		if err := SetEnv(tt.args.vars); err != nil {
			t.Errorf("SetEnv() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
		if err := UnsetEnv(tt.args.vars); err != nil {
			t.Errorf("UnsetEnv() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}
