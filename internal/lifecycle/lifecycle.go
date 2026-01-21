package lifecycle

import (
	"github.com/SoyebSarkar/Hiberstack/internal/engine/typesense"
	"github.com/SoyebSarkar/Hiberstack/internal/state"
)

type Manager struct {
	ts          *typesense.Client
	snapshotDir string
	stateStore  *state.Store
}

func NewManager(
	ts *typesense.Client,
	snapshotDir string,
	stateStore *state.Store,
) *Manager {
	return &Manager{
		ts:          ts,
		snapshotDir: snapshotDir,
		stateStore:  stateStore,
	}
}
