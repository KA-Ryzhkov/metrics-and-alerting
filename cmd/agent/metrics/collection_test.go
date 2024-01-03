package metrics

import (
	"reflect"
	"testing"
)

func TestMetricStart(t *testing.T) {
	type args struct {
		ListNameMetrics []string
	}
	tests := []struct {
		name string
		args args
		want []Metric
	}{
		{
			name: "TEST 1.",
			args: args{ListNameMetrics: []string{"Alloc", "BuckHashSys"}},
			want: []Metric{
				{Name: "Alloc", Gauge: 0.0, Counter: 0},
				{Name: "BuckHashSys", Gauge: 0.0, Counter: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MetricStart(tt.args.ListNameMetrics); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MetricStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
