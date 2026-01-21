package scheduler

import (
	"log"
	"time"

	"github.com/SoyebSarkar/Hiberstack/internal/lifecycle"
	"github.com/SoyebSarkar/Hiberstack/internal/state"
)

type Offloader interface {
	Offload(collection string) error
}

type Scheduler struct {
	store        *state.Store
	lifecycleMgr *lifecycle.Manager
	offloadAfter time.Duration
	gracePeriod  time.Duration
	interval     time.Duration
}

func New(
	store *state.Store,
	lifecycleMgr *lifecycle.Manager,
	offloadAfter time.Duration,
) *Scheduler {
	return &Scheduler{
		store:        store,
		lifecycleMgr: lifecycleMgr,
		offloadAfter: offloadAfter,
		gracePeriod:  30 * time.Second,
		interval:     10 * time.Minute,
	}
}

func (s *Scheduler) Start() {
	ticker := time.NewTicker(s.interval)

	go func() {
		for range ticker.C {
			s.runOnce()
		}
	}()
}

func (s *Scheduler) runOnce() {
	collections := s.store.ListHotOlderThan(s.offloadAfter)

	for _, c := range collections {
		log.Println("scheduler marking draining:", c)
		s.store.Set(c, state.Draining)
		go s.drainAndOffload(c)
	}
}

func (s *Scheduler) drainAndOffload(collection string) {
	time.Sleep(s.gracePeriod)

	// State might have changed
	if s.store.Get(collection) != state.Draining {
		return
	}

	// Activity resumed → cancel offload
	if s.store.WasRecentlyAccessed(collection, s.offloadAfter) {
		log.Println("activity resumed, reverting to HOT:", collection)
		s.store.Set(collection, state.Hot)
		return
	}

	log.Println("scheduler offloading after drain:", collection)
	if err := s.lifecycleMgr.Offload(collection); err != nil {
		log.Println("offload failed:", collection, err)
		// fallback: revert state
		s.store.Set(collection, state.Hot)
	}
}
