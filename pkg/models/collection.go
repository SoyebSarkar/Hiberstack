package models

import "time"

type CollectionState string

const (
	HOT     CollectionState = "HOT"
	COLD    CollectionState = "COLD"
	LOADING CollectionState = "LOADING"
	FAILED  CollectionState = "FAILED"
)

type Collection struct {
	Name         string
	State        CollectionState
	LastQueryAt  time.Time
	LastWriteAt  time.Time
	SnapshotPath string
}
