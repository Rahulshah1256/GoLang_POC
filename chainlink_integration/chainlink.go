package main

import (
	"fmt"

	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/log"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
)

// ChainlinkIntegration represents the Chainlink integration.
type ChainlinkIntegration struct {
	JobSubscriber    job.JobSubscriber
	LogBroadcaster   log.Broadcaster
	PipelineRunner   pipeline.Runner
	ChainlinkService chainlink.Service
}

// NewChainlinkIntegration creates a new instance of the Chainlink integration.
func NewChainlinkIntegration() (*ChainlinkIntegration, error) {
	// Initialize Chainlink services
	chainlinkService := chainlink.NewChainlinkService()

	// Initialize JobSubscriber
	jobSubscriber, err := job.NewJobSubscriber(chainlinkService)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize JobSubscriber: %v", err)
	}

	// Initialize LogBroadcaster
	logBroadcaster := log.NewBroadcaster(chainlinkService)

	// Initialize PipelineRunner
	pipelineRunner := pipeline.NewRunner()

	return &ChainlinkIntegration{
		JobSubscriber:    jobSubscriber,
		LogBroadcaster:   logBroadcaster,
		PipelineRunner:   pipelineRunner,
		ChainlinkService: chainlinkService,
	}, nil
}
