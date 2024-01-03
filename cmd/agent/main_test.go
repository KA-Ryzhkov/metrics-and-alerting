package main

import (
	"github.com/KA-Ryzhkov/metrics-and-alerting/cmd/agent/metrics"
	"testing"
)

func Test_generateUrl(t *testing.T) {
	type args struct {
		m     metrics.Metric
		gauge bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1. Base test.",
			args: args{
				m: metrics.Metric{
					Name:    "Alloc",
					Gauge:   0.00001,
					Counter: 11,
				},
				gauge: false},
			want: "http://localhost:8080/update/counter/Alloc/11/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateUrl(tt.args.m, tt.args.gauge); got != tt.want {
				t.Errorf("generateUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
