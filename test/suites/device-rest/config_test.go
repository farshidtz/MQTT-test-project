package test

import (
	"edgex-snap-testing/test/utils"
	"testing"
)

// Deprecated
func TestEnvConfig(t *testing.T) {
	if utils.FullConfigTest == "off" {
		// make this subtest optional to save testing time
		// similar full config tests have been operated in device-mqtt test suite
		t.Skip()
	} else if utils.FullConfigTest == "on" {
		// start clean
		utils.SnapStop(t, deviceRestSnap)

		t.Run("change service port", func(t *testing.T) {
			t.Cleanup(func() {
				utils.SnapUnset(t, deviceRestSnap, "env")
				utils.SnapStop(t, deviceRestSnap)
			})

			const newPort = "11111"

			// make sure the port is available before using it
			utils.RequirePortAvailable(t, newPort)

			// set env. and validate the new port comes online
			utils.SnapSet(t, deviceRestSnap, "env.service.port", newPort)
			utils.SnapStart(t, deviceRestSnap)

			utils.WaitServiceOnline(t, 60, newPort)

			// unset env. and validate the default port comes online
			utils.SnapUnset(t, deviceRestSnap, "env.service.port")
			utils.SnapRestart(t, deviceRestService)
			utils.WaitServiceOnline(t, 60, defaultServicePort)
		})
	}
}

func TestAppConfig(t *testing.T) {
	// start clean
	utils.SnapStop(t, deviceRestSnap)

	t.Run("set and unset apps.", func(t *testing.T) {
		t.Cleanup(func() {
			// temporary using unset apps and unset config together here to do unset apps' job
			// until this issue been solved: https://github.com/canonical/edgex-snap-hooks/issues/43
			utils.SnapUnset(t, deviceRestSnap, "apps.device-rest.config.service.port")
			utils.SnapUnset(t, deviceRestSnap, "config.service.port")
			utils.SnapStop(t, deviceRestSnap)
		})

		const newPort = "22222"

		// make sure the port is available before using it
		utils.RequirePortAvailable(t, newPort)

		// set apps. and validate the new port comes online
		utils.SnapSet(t, deviceRestSnap, "apps.device-rest.config.service.port", newPort)
		utils.SnapStart(t, deviceRestSnap)

		utils.WaitServiceOnline(t, 60, newPort)

		// unset apps. and validate the default port comes online
		// temporary using unset apps and unset config together here to do unset apps' job
		// until this issue been solved: https://github.com/canonical/edgex-snap-hooks/issues/43
		utils.SnapUnset(t, deviceRestSnap, "apps.device-rest.config.service.port")
		utils.SnapUnset(t, deviceRestSnap, "config.service.port")
		utils.SnapRestart(t, deviceRestService)

		utils.WaitServiceOnline(t, 60, defaultServicePort)
	})
}

func TestGlobalConfig(t *testing.T) {
	// start clean
	utils.SnapStop(t, deviceRestSnap)

	t.Run("set and unset apps.", func(t *testing.T) {
		t.Cleanup(func() {
			utils.SnapUnset(t, deviceRestSnap, "config.service.port")
			utils.SnapStop(t, deviceRestSnap)
		})

		const newPort = "33333"

		// make sure the port is available before using it
		utils.RequirePortAvailable(t, newPort)

		// set config. and validate the new port comes online
		utils.SnapSet(t, deviceRestSnap, "config.service.port", newPort)
		utils.SnapStart(t, deviceRestSnap)

		utils.WaitServiceOnline(t, 60, newPort)

		// unset config. and validate the default port comes online
		utils.SnapUnset(t, deviceRestSnap, "config.service.port")
		utils.SnapRestart(t, deviceRestService)

		utils.WaitServiceOnline(t, 60, defaultServicePort)
	})
}

func TestMixedConfig(t *testing.T) {
	if utils.FullConfigTest == "off" {
		// make this subtest optional to save testing time
		// similar full config tests have been operated in device-mqtt test suite
		t.Skip()
	} else if utils.FullConfigTest == "on" {
		// start clean
		utils.SnapStop(t, deviceRestSnap)

		t.Run("use apps. and config. for different values", func(t *testing.T) {
			t.Cleanup(func() {
				utils.SnapUnset(t, deviceRestSnap, "apps.device-rest.config.service.port")
				utils.SnapUnset(t, deviceRestSnap, "config.service.port")
				utils.SnapStop(t, deviceRestService)
			})

			const newAppPort = "44444"
			const newConfigPort = "55555"

			// make sure the ports are available before using it
			utils.RequirePortAvailable(t, newAppPort)
			utils.RequirePortAvailable(t, newConfigPort)

			// set apps. and config. with different values,
			// and validate that app-specific option has been picked up because it has higher precedence
			utils.SnapSet(t, deviceRestSnap, "apps.device-rest.config.service.port", newAppPort)
			utils.SnapSet(t, deviceRestSnap, "config.service.port", newConfigPort)
			utils.SnapStart(t, deviceRestSnap)

			utils.WaitServiceOnline(t, 60, newAppPort)
		})
	}
}
