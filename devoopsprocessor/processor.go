package devoopsprocessor

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/model/pdata"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

func processTraces(ctx context.Context, td pdata.Traces) (pdata.Traces, error) {
	return td, nil
}
