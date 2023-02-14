package orchestrator

import (
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

const (
	ojo_container_name = "ojo"
	ojo_tmrpc_port     = "26657"
	ojo_grpc_port      = "9090"

	price_feeder_container_name = "price-feeder"
	price_feeder_server_port    = "8080"
)

// Orchestrator is responsible for managing docker resources,
// their configuration files, and environment variables.
type Orchestrator struct {
	dockerPool    *dockertest.Pool
	dockerNetwork *dockertest.Network

	ojoResource *dockertest.Resource
	ojoRPC      *rpchttp.HTTP
	ojoChain    *Chain

	priceFeederResource *dockertest.Resource
}

func (o *Orchestrator) InitDockerResources(t *testing.T) error {
	var err error

	t.Log("-> initializing docker network")
	err = o.initNetwork()
	if err != nil {
		return err
	}

	t.Log("-> initializing Ojo validator")
	o.initOjod()

	t.Log("-> initializing price-feeder")

	return nil
}

func (o *Orchestrator) TearDownDockerResources() error {
	return o.dockerPool.Client.RemoveNetwork(o.dockerNetwork.Network.ID)
}

func (o *Orchestrator) initNetwork() error {
	var err error
	o.dockerPool, err = dockertest.NewPool("")
	if err != nil {
		return err
	}

	o.dockerNetwork, err = o.dockerPool.CreateNetwork("e2e_test_network")
	if err != nil {
		return err
	}
	return nil
}

func noRestart(config *docker.HostConfig) {
	config.RestartPolicy = docker.RestartPolicy{
		Name: "no",
	}
}
