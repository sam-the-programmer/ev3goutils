package ev3goutils

import (
	"errors"
	"log"
	"strings"

	"github.com/ev3go/ev3dev"
)

type MotorSet struct {
	A ev3dev.TachoMotor // Motor at the A port.
	B ev3dev.TachoMotor // Motor at the B port.
	C ev3dev.TachoMotor // Motor at the C port.
	D ev3dev.TachoMotor // Motor at the D port.
}

func (m *MotorSet) connectMotor(port string, size string) {
	if strings.ToLower(size) != "l" && strings.ToLower(size) != "m" {
		log.Fatal(errors.New("Motor size" + strings.ToLower(size) + "is not a valid motor size. Use l or m"))
	}
	motorType := "lego-ev3-" + strings.ToLower(size) + "-motor"

	switch strings.ToUpper(port) {
	case "A":
		motor, _ := ev3dev.TachoMotorFor("ev3-ports:outA", motorType)
		m.A = *motor
	case "B":
		motor, _ := ev3dev.TachoMotorFor("ev3-ports:outB", motorType)
		m.B = *motor
	case "C":
		motor, _ := ev3dev.TachoMotorFor("ev3-ports:outC", motorType)
		m.C = *motor
	case "D":
		motor, _ := ev3dev.TachoMotorFor("ev3-ports:outD", motorType)
		m.D = *motor
	default:
		log.Fatal(errors.New("Port" + strings.ToUpper(port) + "is not a valid motor port name. Use A, B, C or D."))
	}
}

func (m *MotorSet) ConnectLargeMotor(port string) {
	m.connectMotor(port, "l")
}
func (m *MotorSet) ConnectMediumMotor(port string) {
	m.connectMotor(port, "m")
}

func NewMotorSet(ports []string) MotorSet {
	m := MotorSet{}
	return m
}
