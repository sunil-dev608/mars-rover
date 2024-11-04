package world

import (
	"io"
	"os"
	"testing"

	"github.com/sunil-dev608/mars-rover/internal/pkg/robot"
)

func Test_world_validateRobotPosition(t *testing.T) {
	type fields struct {
		grid             []int
		robots           []robot.Robot
		commands         []string
		robotsLostStatus []bool
	}
	type args struct {
		robot robot.Robot
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "valid position",
			fields: fields{
				grid: []int{5, 5},
			},
			args: args{
				robot: robot.NewRobot(1, 2, 'N'),
			},
			want: true,
		},
		{
			name: "invalid position",
			fields: fields{
				grid: []int{5, 5},
			},
			args: args{
				robot: robot.NewRobot(-1, 2, 'N'),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &world{
				grid:             tt.fields.grid,
				robots:           tt.fields.robots,
				commands:         tt.fields.commands,
				robotsLostStatus: tt.fields.robotsLostStatus,
			}
			if got := w.validateRobotPosition(tt.args.robot); got != tt.want {
				t.Errorf("world.validateRobotPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func tmpFile(input string) *os.File {

	tmpfile, err := os.CreateTemp("", "tmp-input")
	if err != nil {
		panic(err)
	}

	// Write test data to the file
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		panic(err)

	}

	tmpfile.Close()
	tmpfile, err = os.Open(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	return tmpfile
}

func Test_world_ReadData(t *testing.T) {
	type fields struct {
		grid             []int
		robots           []robot.Robot
		commands         []string
		robotsLostStatus []bool
	}
	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name    string
		input   string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid input",
			input: `4 8
(2, 3, E) LFRFF
(0, 2, N) FFLFRFF
(2, 3, N) FLLFR
(1, 0, S) FFRLF`,
			fields: fields{
				grid:             make([]int, 2),
				robots:           make([]robot.Robot, 0),
				commands:         make([]string, 0),
				robotsLostStatus: make([]bool, 0),
			},
			args: args{
				rd: nil,
			},
			wantErr: false,
		},
		{
			name: "invalid input",
			input: `X Y
			(2, 3, E) LFRFF
			(0, 2, N) FFLFRFF
			(2, 3, N) FLLFR
			(1, 0, S) FFRLF`,
			fields: fields{
				grid:             make([]int, 2),
				robots:           make([]robot.Robot, 0),
				commands:         make([]string, 0),
				robotsLostStatus: make([]bool, 0),
			},
			args: args{
				rd: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &world{
				grid:             tt.fields.grid,
				robots:           tt.fields.robots,
				commands:         tt.fields.commands,
				robotsLostStatus: tt.fields.robotsLostStatus,
			}
			tmpfile := tmpFile(tt.input)
			tt.args.rd = tmpfile
			if err := w.ReadData(tt.args.rd); (err != nil) != tt.wantErr {
				t.Errorf("world.ReadData() error = %v, wantErr %v", err, tt.wantErr)
			}
			defer os.Remove(tmpfile.Name())
		})
	}
}
