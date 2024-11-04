package world

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/sunil-dev608/mars-rover/internal/pkg/parser"
	"github.com/sunil-dev608/mars-rover/internal/pkg/robot"
)

type World interface {
	ReadData(rd io.Reader) error
	MoveRobots()
	PrintWorld()
}

type world struct {
	grid             []int
	robots           []robot.Robot
	commands         []string
	robotsLostStatus []bool
}

func NewWorld() World {
	return &world{
		grid:             make([]int, 2),
		robots:           make([]robot.Robot, 0),
		commands:         make([]string, 0),
		robotsLostStatus: make([]bool, 0),
	}
}

func (w *world) ReadData(rd io.Reader) error {
	reader := bufio.NewReader(rd)

	isFirstLine := false
	parser := parser.NewParser()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Error reading file: %v\n", err)
			return err
		}
		line = line[:len(line)-1]
		if !isFirstLine {
			isFirstLine = true
			tokens := strings.Split(line, " ")

			w.grid[0], err = strconv.Atoi(tokens[0])
			if err != nil {
				return fmt.Errorf("error parsing grid size: %v", err)
			}
			w.grid[1], err = strconv.Atoi(tokens[1])
			if err != nil {
				return fmt.Errorf("error parsing grid size: %v", err)
			}
		} else {
			parsedData, err := parser.ParseRobotString(line)
			if err != nil {
				return fmt.Errorf("error parsing robot data: %v", err)
			}
			w.robots = append(w.robots, robot.NewRobot(parsedData.X, parsedData.Y, parsedData.Orientation))
			w.commands = append(w.commands, parsedData.Command)
		}
	}
	w.robotsLostStatus = make([]bool, len(w.robots))
	return nil
}

func (w *world) validateRobotPosition(robot robot.Robot) bool {
	x, y, _ := robot.GetPosition()
	if x > w.grid[0] || y > w.grid[1] || x < 0 || y < 0 {
		return false
	}
	return true
}

// Move moves the robot
func (w *world) MoveRobots() {

	for i, r := range w.robots {
		command := w.commands[i]
		for _, c := range command {
			tmp := robot.NewRobot(r.GetPosition())
			r.SingleMove(c)
			if !w.validateRobotPosition(r) {
				w.robots[i] = tmp
				w.robotsLostStatus[i] = true
				break
			}
		}
	}

}

func (w *world) PrintWorld() {
	for i, _ := range w.robots {
		fmt.Print(w.robots[i])
		if w.robotsLostStatus[i] {
			fmt.Print(" LOST")
		}
		fmt.Println()
	}
}
