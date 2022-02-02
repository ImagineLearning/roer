package main

import (
	"os"

	"github.com/ImagineLearning/roer/cmd"
	"github.com/ImagineLearning/roer/spinnaker"
	"github.com/sirupsen/logrus"
)

// version is set via ldflags
var version = "dev"

func main() {
	// TODO rz - Don't really like this bit. Standardize a spinnaker config file.
	// maybe worthwhile splitting out this spinnaker API into a standard lib...
	if os.Getenv("SPINNAKER_API") == "" {
		logrus.Fatal("SPINNAKER_API must be set")
	}

	config := spinnaker.ClientConfig{
		Endpoint:          os.Getenv("SPINNAKER_API"),
		HTTPClientFactory: spinnaker.DefaultHTTPClientFactory,
	}
	if err := cmd.NewRoer(version, config).Run(os.Args); err != nil {
		logrus.Fatal(err.Error())
	}
}
