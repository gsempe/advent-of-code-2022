package day06

import (
	"testing"
)

func Test_isStartOfPacket(t *testing.T) {
	type args struct {
		marker []rune
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "is start",
			args: args{
				marker: []rune("abcd"),
				length: 4,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "is not start",
			args: args{
				marker: []rune("abca"),
				length: 4,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isStartOfPacket(tt.args.marker, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("isStartOfPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isStartOfPacket() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberOfCharacterToStartOfPacket(t *testing.T) {
	type args struct {
		input  string
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				input:  "abcd",
				length: 4,
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input:  "abca",
				length: 4,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "",
			args: args{
				input:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
				length: 4,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "",
			args: args{
				input:  "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
				length: 4,
			},
			want:    11,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NumberOfCharacterToStartOfPacket(tt.args.input, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("NumberOfCharacterToStartOfPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NumberOfCharacterToStartOfPacket() got = %v, want %v", got, tt.want)
			}
		})
	}
}
