package day07

import (
	"testing"
)

func Test_directory_addFile(t *testing.T) {
	type fields struct {
		name        string
		size        int
		directories []*directory
	}
	type args struct {
		name string
		size int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "happy path",
			fields: fields{
				name:        "/",
				size:        0,
				directories: nil,
			},
			args: args{
				name: "file",
				size: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &directory{
				name:        tt.fields.name,
				size:        tt.fields.size,
				directories: tt.fields.directories,
			}
			if got := d.addFile(tt.args.name, tt.args.size); got != tt.want {
				t.Errorf("addFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
