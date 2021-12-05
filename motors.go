package ev3goutils

import (
	"errors"
	"log"
	"strings"

	"github.com/ev3go/ev3dev"
)

func newError(message string) error {
	return errors.New("\033[31m ERROR: " + message + "\033[0m")
}

type MotorSet struct {
	A ev3dev.TachoMotor // Motor at the A port.
	B ev3dev.TachoMotor // Motor at the B port.
	C ev3dev.TachoMotor // Motor at the C port.
	D ev3dev.TachoMotor // Motor at the D port.

	connectedMotors []string // Slice of connected motors.
}

func (m *MotorSet) connectMotor(port string, size string) {
	if strings.ToLower(size) != "l" && strings.ToLower(size) != "m" {
		log.Fatal(newError("Motor size" + strings.ToLower(size) + "is not a valid motor size. Use l or m"))
	}
	motorType := "lego-ev3-" + strings.ToLower(size) + "-motor"

	switch strings.ToUpper(port) {
	case "A":
		motor, _ := ev3dev.TachoMotorFor("ev3-ports:outA", motorType)
		m.A = *motor
		m.connectedMotors = append(m.connectedMotors, "a")
	case "B":
		motor, _ := ev3dev.TachoMotorFor("ev3-ports:outB", motorType)
		m.B = *motor
		m.connectedMotors = append(m.connectedMotors, "b")
	case "C":
		motor, _ := ev3dev.TachoMotorFor("ev3-ports:outC", motorType)
		m.C = *motor
		m.connectedMotors = append(m.connectedMotors, "c")
	case "D":
		motor, _ := ev3dev.TachoMotorFor("ev3-ports:outD", motorType)
		m.D = *motor
		m.connectedMotors = append(m.connectedMotors, "d")
	default:
		log.Fatal(newError("Port " + strings.ToUpper(port) + " is not a valid motor port name. Use A, B, C or D."))
	}
}

func (m *MotorSet) checkMotor(motor string) {
	in := false
	for _, v := range m.connectedMotors {
		if v == strings.ToLower(motor) {
			in = true
		}
	}

	if !in {
		log.Fatal(newError("Motor \"" + motor + "\" is not a valid motor name or is not connected."))
	}
}

func (m *MotorSet) ConnectLargeMotor(port string) {
	m.connectMotor(port, "l")
}
func (m *MotorSet) ConnectMediumMotor(port string) {
	m.connectMotor(port, "m")
}

func (m *MotorSet) runDegrees(motor *ev3dev.TachoMotor, degrees int, speed int) {
	motor.SetSpeedSetpoint(speed)
	motor.SetPositionSetpoint(degrees).Command("run-to-rel-pos")
}

func (m *MotorSet) ADegrees(degrees int, speed int) {
	m.runDegrees(&m.A, degrees, speed)
}
func (m *MotorSet) BDegrees(degrees int, speed int) {
	m.runDegrees(&m.B, degrees, speed)
}
func (m *MotorSet) CDegrees(degrees int, speed int) {
	m.runDegrees(&m.C, degrees, speed)
}
func (m *MotorSet) DDegrees(degrees int, speed int) {
	m.runDegrees(&m.D, degrees, speed)
}

func NewMotorSet(ports []string) MotorSet {
	m := MotorSet{}
	return m
}
