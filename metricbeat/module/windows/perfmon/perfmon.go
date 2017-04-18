// +build windows

package perfmon

import (
	"fmt"
	"strings"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/pkg/errors"
)

type CounterConfig struct {
	Alias       string `config:"alias" validate:"required"`
	Query       string `config:"query" validate:"required"`
	FormatInput string `config:"format"`
	Format      Format
}

func init() {
	if err := mb.Registry.AddMetricSet("windows", "perfmon", New); err != nil {
		panic(err)
	}
}

type MetricSet struct {
	mb.BaseMetricSet
	reader *PerfmonReader
}

// New create a new instance of the MetricSet.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	logp.Warn("BETA: The perfmon metricset is beta")

	config := struct {
		CounterConfig []CounterConfig `config:"perfmon.counters" validate:"required"`
	}{}

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	for _, value := range config.CounterConfig {
		var form Format
		switch strings.ToLower(value.FormatInput) {
		case "":
			value.Format = FLOAT
		case "float":
			value.Format = FLOAT
		case "long":
			value.Format = LONG
		default:
			err := fmt.Errorf("format '%s' for counter '%s' are not valid", value.Format, value.Alias)
			return nil, errors.Wrap(err, "initialization failed")
		}

	}

	reader, err := NewPerfmonReader(config.CounterConfig)
	if err != nil {
		return nil, errors.Wrap(err, "initialization failed")
	}

	return &MetricSet{
		BaseMetricSet: base,
		reader:        reader,
	}, nil
}

func (m *MetricSet) Fetch() (common.MapStr, error) {
	data, err := m.reader.Read()
	if err != nil {
		return nil, errors.Wrap(err, "failed reading counters")
	}

	return data, nil
}
