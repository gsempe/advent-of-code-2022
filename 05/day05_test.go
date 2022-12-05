package day05

import (
	"reflect"
	"testing"
)

func Test_doMoveWithCrateMover9001(t *testing.T) {
	type args struct {
		stacks   [][]rune
		quantity int
		from     int
		to       int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]rune
		wantErr bool
	}{
		{
			name: "",
			args: args{
				stacks:   [][]rune{[]rune("NZ"), []rune("DCM"), []rune("P")},
				quantity: 1,
				from:     1,
				to:       0,
			},
			want:    [][]rune{[]rune("DNZ"), []rune("CM"), []rune("P")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := doMoveWithCrateMover9001(tt.args.stacks, tt.args.quantity, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("doMoveWithCrateMover9001() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doMoveWithCrateMover9001() got = %v, want %v", got, tt.want)
			}
		})
	}
}
