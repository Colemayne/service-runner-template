package commands

import (
	"fmt"
	"os"
	"syscall"

	"../common"
	serviceHelpers "../helpers/service"
	"github.com/sirupsen/logrus"

	service "github.com/ayufan/golang-kardianos-service"
	"github.com/urfave/cli"
)

// RunCommand is responsible for the logic associated with the service
type RunCommand struct {
	ServiceName      string `short:"service" description:"Use different names for different services"`
	Config           string `short:"config" description:"Specify custom configuration"`
	WorkingDirectory string `short:"working-directory" description:"Specify custom working directory"`
	User             string `short:"user" description:"Which specific user will execute"`
	Syslog           bool   `short:"syslog" description:"Log to system service logger" env:"LOG_SYSLOG"`

	runSignal chan os.Signal

	reloadSignal chan os.Signal

	stopSignals chan os.Signal

	stopSignal os.Signal

	runFinished chan bool

	currentWorkers int
}

// Start is the method implementing the `golang-kardianos-service` `Interface` interface
// responsible for a non-blocking initialization of the process. When it exits, the main control flow is passed to runWait()
func (rc *RunCommand) Start(_ service.Service) error {
	rc.runSignal = make(chan os.Signal, 1)
	rc.reloadSignal = make(chan os.Signal, 1)
	rc.runFinished = make(chan bool, 1)
	rc.stopSignals = make(chan os.Signal)

	if len(rc.WorkingDirectory) > 0 {
		err := os.Chdir(rc.WorkingDirectory)
		if err != nil {
			return err
		}
	}

	// Load configuration here

	// Start shouldn't block. The actual work is async.
	go rc.run()

	return nil

}

// run is the main method of RunCommand. It's started by services support through `Start` method and is responsible for
// responsible for initializing all goroutines and handling concurrent, multi-runner execution of jobs.
func (rc *RunCommand) run() {
	fmt.Println("Running!")

	close(rc.runFinished)
}

func (rc *RunCommand) Stop(_ service.Service) error {
	err := rc.handleGracefulShutdown()
	if err == nil {
		return nil
	}

	// Logic for forceful shutdown

	return err
}

func (rc *RunCommand) handleGracefulShutdown() error {
	// Wait for a SIGQUIT
	for rc.stopSignal == syscall.SIGQUIT {

		// Wait for other signals to finish run
		select {
		case rc.stopSignal = <-rc.stopSignals:
		// We received a new signal

		case <-rc.runFinished:
			// Everything finished we can exit now
			return nil
		}
	}

	return fmt.Errorf("received: %v", rc.stopSignal)
}

func (rc *RunCommand) Execute(_ *cli.Context) error {
	svcConfig := &service.Config{
		Name:        rc.ServiceName,
		DisplayName: rc.ServiceName,
		Description: "A service",
		Arguments:   []string{"run"},
		Option: service.KeyValue{
			"RunWait": rc.runWait,
		},
	}
	svc, err := serviceHelpers.New(rc, svcConfig)
	if err != nil {
		logrus.WithError(err).Fatalln("Service creation failed")
	}

	err = svc.Run()
	if err != nil {
		logrus.WithError(err).Fatal("Service run failed")
	}

	return nil
}

func (rc *RunCommand) runWait() {
	rc.stopSignal = <-rc.stopSignals
}

func init() {

	common.RegisterCommand2("run", "run multi runner service", &RunCommand{
		ServiceName: defaultServiceName,
	})
}
