package elements

import "github.com/pion/ion/pkg/node/avp/pipeline/samplebuilder"

var (
	baseConfig Configs
)

// Element interface
type Element interface {
	ID() string
	Write(*samplebuilder.Sample) error
	Read() <-chan *samplebuilder.Sample
	Close()
}

// Configs for element
type Configs struct {
	On        bool
	WebmSaver WebmSaverConfig
}

// Init elements
func Init(configs Configs) {
	baseConfig = configs
}
