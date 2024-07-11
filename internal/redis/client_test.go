package redis

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestSingleRedis(t *testing.T) {
	tests := []struct {
		name string
		want *RdbClient
	}{
		// TODO: Add test cases.
		{
			name: "TestSingleRedis",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleRedis(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleRedis() = %v, want %v", got, tt.want)
				got.DB.Set(context.Background(), "superman", "test", time.Minute)
			}
		})
	}
}
