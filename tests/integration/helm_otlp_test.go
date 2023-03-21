//go:build onlylatest
// +build onlylatest

package integration

import (
	"testing"

	"github.com/SumoLogic/sumologic-kubernetes-collection/tests/integration/internal"
)

func Test_Helm_OTLP(t *testing.T) {

	expectedMetrics := internal.DefaultExpectedMetrics
	installChecks := []featureCheck{
		CheckSumologicSecret(15),
		CheckOtelcolMetadataLogsInstall,
		CheckOtelcolLogsCollectorInstall,
		CheckOtelcolEventsInstall,
	}

	featInstall := GetInstallFeature(installChecks)

	featLogs := GetLogsFeature()

	featMetrics := GetMetricsFeature(expectedMetrics, Prometheus)

	featTraces := GetTracesFeature()

	featEvents := GetEventsFeature()

	testenv.Test(t, featInstall, featLogs, featMetrics, featEvents, featTraces)
}
