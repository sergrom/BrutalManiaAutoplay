package mouse

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"os/exec"
	"strconv"

	"github.com/fatih/color"
)

// Move ...
func Move(pointer image.Point) {
	_, err := exec.Command("xdotool", "mousemove", strconv.Itoa(pointer.X), strconv.Itoa(pointer.Y)).Output()
	if err != nil {
		color.Red("Move mouse pointer error:", err.Error())
		return
	}
}

// GetPointerCoordinates ...
func GetPointerCoordinates() (int, int, error) {
	var x, y int

	output, err := exec.Command("xdotool", "getmouselocation").Output()
	if err != nil {
		return x, y, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	scanner.Scan()

	_, err = fmt.Sscanf(scanner.Text(), "x:%d y:%d", &x, &y)

	return x, y, err
}

// Click ...
func Click(pointer image.Point) {
	Move(pointer)
	_, err := exec.Command("xdotool", "click", "1").Output()
	if err != nil {
		color.Red("Move click error:", err.Error())
		return
	}
}
