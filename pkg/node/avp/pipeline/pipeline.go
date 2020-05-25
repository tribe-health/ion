package pipeline

import (
	"sync"
	"time"

	"github.com/pion/ion/pkg/log"
	"github.com/pion/ion/pkg/node/avp/elements"
	"github.com/pion/ion/pkg/node/avp/pipeline/samplebuilder"
	"github.com/pion/ion/pkg/rtc/transport"
	"github.com/pion/ion/pkg/util"
)

const (
	liveCycle = 6 * time.Second
)

var (
	pipelineConfig Config
)

// Config for pipeline
type Config struct {
	SampleBuilder samplebuilder.Config
}

// Pipeline constructs a processing graph
//
//                                         +--->node
//                                         |
// pub--->pubCh-->sampleBuilder-->nodeCh---+--->node
//                                         |
//                                         +--->node
type Pipeline struct {
	pub           transport.Transport
	elements      map[string]elements.Element
	elementLock   sync.RWMutex
	elementChans  map[string]chan *samplebuilder.Sample
	sampleBuilder *samplebuilder.SampleBuilder
	stop          bool
	liveTime      time.Time
}

// Init pipeline
func Init(config Config) {
	pipelineConfig = config
}

// NewPipeline return a new Pipeline
func NewPipeline(id string) *Pipeline {
	log.Infof("NewPipeline id=%s", id)
	return &Pipeline{
		elements:      make(map[string]elements.Element),
		elementChans:  make(map[string]chan *samplebuilder.Sample),
		liveTime:      time.Now().Add(liveCycle),
		sampleBuilder: samplebuilder.NewSampleBuilder(pipelineConfig.SampleBuilder),
	}
}

func (p *Pipeline) start() {
	go func() {
		defer util.Recover("[Pipeline.start]")
		for {
			if p.stop {
				return
			}

			sample := p.sampleBuilder.Read()

			p.liveTime = time.Now().Add(liveCycle)
			p.elementLock.RLock()
			// Push to client send queues
			for id := range p.elements {
				// Nonblock sending
				select {
				case p.elementChans[id] <- sample:
				default:
					log.Errorf("Element consumer is backed up. Dropping sample")
				}
			}
			p.elementLock.RUnlock()
		}
	}()
}

// AddElement add a element to router
func (p *Pipeline) AddElement(id string, n elements.Element) {
	//fix panic: assignment to entry in nil map
	if p.stop {
		return
	}
	p.elementLock.Lock()
	defer p.elementLock.Unlock()
	p.elements[id] = n
	p.elementChans[id] = make(chan *samplebuilder.Sample, 100)
	log.Infof("Pipeline.AddElement id=%s p=%p", id, n)
}

// GetElement get a node by id
func (p *Pipeline) GetElement(id string) elements.Element {
	p.elementLock.RLock()
	defer p.elementLock.RUnlock()
	return p.elements[id]
}

// DelElement del node by id
func (p *Pipeline) DelElement(id string) {
	log.Infof("Pipeline.DelElement id=%s", id)
	p.elementLock.Lock()
	defer p.elementLock.Unlock()
	if p.elements[id] != nil {
		p.elements[id].Close()
	}
	if p.elementChans[id] != nil {
		close(p.elementChans[id])
	}
	delete(p.elements, id)
	delete(p.elementChans, id)
}

// DelElements del all node
func (p *Pipeline) DelElements() {
	log.Infof("Pipeline.DelElements")
	p.elementLock.RLock()
	keys := make([]string, 0, len(p.elements))
	for k := range p.elements {
		keys = append(keys, k)
	}
	p.elementLock.RUnlock()

	for _, id := range keys {
		p.DelElement(id)
	}
}

// DelPub del pub
func (p *Pipeline) DelPub() {
	log.Infof("Pipeline.DelPub %v", p.pub)
	if p.pub != nil {
		p.pub.Close()
	}
	p.sampleBuilder.Stop()
	p.pub = nil
}

// Close release all
func (p *Pipeline) Close() {
	if p.stop {
		return
	}
	log.Infof("Pipeline.Close")
	p.DelPub()
	p.stop = true
	p.DelElements()
}

// Alive return pipeline status
func (p *Pipeline) Alive() bool {
	return p.liveTime.After(time.Now())
}
