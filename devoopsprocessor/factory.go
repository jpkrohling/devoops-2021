package devoopsprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const procType = "devoops"

func NewFactory() component.ProcessorFactory {
	return processorhelper.NewFactory(
		procType,
		createDefaultConfig,
		processorhelper.WithTraces(createTracesProcessor))
}

func createDefaultConfig() config.Processor {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(config.NewComponentID(procType)),
	}
}

func createTracesProcessor(_ context.Context, _ component.ProcessorCreateSettings, cfg config.Processor, nextConsumer consumer.Traces) (component.TracesProcessor, error) {
	return processorhelper.NewTracesProcessor(
		cfg,
		nextConsumer,
		processTraces,
		processorhelper.WithCapabilities(processorCapabilities))
}
