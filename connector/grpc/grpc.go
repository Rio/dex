// Package grpc implements grpc connector that will call a configured grpc server for login authentication.
package grpc

import (
	"github.com/sirupsen/logrus"
	grpc1 "google.golang.org/grpc"

	"github.com/coreos/dex/connector"
)

type Config struct {
	endpoint string
}

type grpcConnector struct {
	grpcClient ConnectorClient
}

func (c *Config) Open(logger logrus.FieldLogger) (connector.Connector, error) {
	// TODO: implement grpc credentials and other dialing options to be configurable through the config file.
	conn, err := grpc1.Dial(c.endpoint, grpc1.WithInsecure())
	if err != nil {
		logger.Fatalln("Unable to connect to grpc endpoint:", err)
	}

	return grpcConnector{
		grpcClient: NewConnectorClient(conn),
	}, nil
}
