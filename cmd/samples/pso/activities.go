package main

import (
	"context"
	"math/rand"
	"time"

	"go.uber.org/cadence/activity"
)

/**
 * Sample activities used by file processing sample workflow.
 */
const (
	initParticleActivityName   = "initParticleActivityName"
	updateParticleActivityName = "updateParticleActivityName"
)

var rng *rand.Rand

// This is registration process where you register all your activity handlers.
func init() {
	// initialize the RNG
	// WARNING: the randomness of activity scheduling with multiple workers makes random number generation truly random and not repeatable in debugging
	// worker.ReplayWorkflowHistoryFromJSONFile should be used to troubleshoot a specific workflow failure.
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func initParticleActivity(ctx context.Context, swarm Swarm) (Particle, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("initParticleActivity started.")

	particle := NewParticle(&swarm, rng)
	particle.UpdateFitness(&swarm)

	return *particle, nil
}

func updateParticleActivity(ctx context.Context, swarm Swarm, particleIdx int) (Particle, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("updateParticleActivity started.")

	particle := swarm.Particles[particleIdx]
	particle.UpdateLocation(&swarm, rng)
	particle.UpdateFitness(&swarm)

	return *particle, nil
}
