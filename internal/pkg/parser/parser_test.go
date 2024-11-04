package parser

import (
	"reflect"
	"testing"
)

func Test_parserImpl_ParseRobotString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		w       *parserImpl
		args    args
		want    *ParsedData
		wantErr bool
	}{
		{
			name: "valid input",
			args: args{
				input: "(2, 3, E) LFRFF",
			},
			want: &ParsedData{
				X:           2,
				Y:           3,
				Orientation: 'E',
				Command:     "LFRFF",
			},
			wantErr: false,
		},
		{
			name: "invalid input",
			args: args{
				input: "(2, 3, SAF",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &parserImpl{}
			got, err := w.ParseRobotString(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parserImpl.ParseRobotString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parserImpl.ParseRobotString() = %v, want %v", got, tt.want)
			}
		})
	}
}
