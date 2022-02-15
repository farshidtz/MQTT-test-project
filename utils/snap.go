package utils

import (
	"edgex-snap-testing/env"
	"fmt"
	"testing"
)

func SnapInstall(t *testing.T, names ...string) {
	for _, name := range names {
		RunCommand(t, fmt.Sprintf(
			"sudo snap install %s --channel=%s",
			name,
			env.Channel,
		))
	}
}

func SnapRemove(t *testing.T, names ...string) {
	for _, name := range names {
		RunCommand(t, fmt.Sprintf(
			"sudo snap remove --purge %s",
			name,
		))
	}
}

func SnapInstallLocal(t *testing.T, workDir string) {
	// snap install will error and exit if multiple snaps exist
	RunCommand(t, fmt.Sprintf(
		"sudo snap install --dangerous %s/*.snap",
		workDir,
	))
}

func SnapBuild(t *testing.T, workDir string) {
	RunCommand(t, fmt.Sprintf(
		"cd %s && snapcraft",
		workDir,
	))
}

func SnapConnect(t *testing.T, plug, slot string) {
	RunCommand(t, fmt.Sprintf(
		"sudo snap connect %s %s",
		plug, slot,
	))
}
