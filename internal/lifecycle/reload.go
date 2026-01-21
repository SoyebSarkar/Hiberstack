package lifecycle

import (
	"os"
	"path/filepath"

	"github.com/SoyebSarkar/Hiberstack/internal/state"
)

func (m *Manager) Reload(collection string) error {
	st := m.stateStore.Get(collection)
	if st != state.Cold {
		return nil
	}
	m.stateStore.Set(collection, state.Loading)

	baseDir := filepath.Join(m.snapshotDir, collection)

	schema, err := os.ReadFile(filepath.Join(baseDir, "schema.json"))
	if err != nil {
		return err
	}

	if err := m.ts.CreateCollection(schema); err != nil {
		return err
	}

	file, err := os.Open(filepath.Join(baseDir, "documents.jsonl"))
	if err != nil {
		return err
	}
	defer file.Close()

	if err := m.ts.ImportDocuments(collection, file); err != nil {
		return err
	}

	m.stateStore.Set(collection, state.Hot)
	return nil
}
