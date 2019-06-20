package executor_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	executor "github.com/gopherdojo/dojo5/kadai3-2/nagaa052/pkg/executor"
)

func TestNew(t *testing.T) {
	type args struct {
		maxWorkers int
		timeout    time.Duration
	}
	tests := []struct {
		name string
		args args
		want *executor.Executor
	}{
		{
			name: "Success Test",
			args: args{
				maxWorkers: 4,
				timeout:    2 * time.Second,
			},
			want: &executor.Executor{
				Timeout: 2 * time.Second,
				Jobs:    make([]*executor.Job, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executor.New(tt.args.maxWorkers, tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockPayload struct{}

func (p *mockPayload) Execute(context.Context) error {
	return nil
}

func TestExecutor_Start(t *testing.T) {
	type fields struct {
		Timeout time.Duration
		Jobs    []*executor.Job
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Success Test",
			fields: fields{
				Timeout: 2 * time.Second,
				Jobs: []*executor.Job{
					&executor.Job{
						&mockPayload{},
					},
					&executor.Job{
						&mockPayload{},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ex := &executor.Executor{tt.fields.Timeout, tt.fields.Jobs}
			if err := ex.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Executor.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
