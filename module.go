package logistics

import (
	"github.com/zzliekkas/flow-logistics/providers"
	"github.com/zzliekkas/flow/v3"
)

// LogisticsModule implements flow.Module for easy registration into a Flow engine.
type LogisticsModule struct {
	kd100Cfg providers.Kd100Config
}

// LogisticsModuleOption configures the LogisticsModule.
type LogisticsModuleOption func(*LogisticsModule)

// WithKd100 configures Kd100 service for the logistics module.
func WithKd100(cfg providers.Kd100Config) LogisticsModuleOption {
	return func(m *LogisticsModule) {
		m.kd100Cfg = cfg
	}
}

// NewModule creates a new LogisticsModule with the given options.
func NewModule(opts ...LogisticsModuleOption) *LogisticsModule {
	m := &LogisticsModule{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// Name returns the module name.
func (m *LogisticsModule) Name() string {
	return "logistics"
}

// Init registers logistics services into Flow's DI container.
func (m *LogisticsModule) Init(e *flow.Engine) error {
	if m.kd100Cfg.Key != "" {
		svc := providers.NewKd100Service(m.kd100Cfg)
		if err := e.Provide(func() *providers.Kd100Service { return svc }); err != nil {
			return err
		}
	}
	return nil
}
