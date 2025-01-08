package util

import (
	"reflect"
	"testing"
)

func TestIfThenElse(t *testing.T) {
	type args[T any] struct {
		condition  bool
		trueValue  T
		falseValue T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "True test",
			args: args[int]{
				condition:  false,
				trueValue:  1,
				falseValue: 0,
			},
			want: 0,
		},
		{
			name: "False test",
			args: args[int]{
				condition:  true,
				trueValue:  1,
				falseValue: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfThenElse(tt.args.condition, tt.args.trueValue, tt.args.falseValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfThenElse() = %v, want %v", got, tt.want)
			}
		})
	}
}
