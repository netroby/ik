package plugins

import (
	"github.com/moriyoshi/ik"
	"fmt"
	"log"
	"os"
	"time"
)

type StdoutOutput struct {
	factory *StdoutOutputFactory
	logger  *log.Logger
}

func (output *StdoutOutput) Emit(record []ik.FluentRecord) error {
	for _, record := range record {
		fmt.Fprintf(os.Stdout, "%d %s: %s\n", record.Timestamp, record.Tag, record.Data)
	}
	return nil
}

func (output *StdoutOutput) Factory() ik.OutputFactory {
	return output.factory
}

func (output *StdoutOutput) Run() error {
	time.Sleep(1000000000)
	return ik.Continue
}

func (output *StdoutOutput) Shutdown() error {
	return nil
}

func (output *StdoutOutput) Dispose() {
	output.Shutdown()
}

type StdoutOutputFactory struct {
}

func newStdoutOutput(factory *StdoutOutputFactory, logger *log.Logger) (*StdoutOutput, error) {
	return &StdoutOutput{
		factory: factory,
		logger:  logger,
	}, nil
}

func (factory *StdoutOutputFactory) Name() string {
	return "stdout"
}

func (factory *StdoutOutputFactory) New(engine ik.Engine, _ *ik.ConfigElement) (ik.Output, error) {
	return newStdoutOutput(factory, engine.Logger())
}

var _ = AddPlugin(&StdoutOutputFactory{})