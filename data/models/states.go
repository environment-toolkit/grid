package models

type SpecState string

const (
	SpecStateNew        SpecState = "new"
	SpecStateCreating   SpecState = "creating"
	SpecStateCreated    SpecState = "created"
	SpecStateUpdating   SpecState = "updating"
	SpecStateUpdated    SpecState = "updated"
	SpecStateFailed     SpecState = "failed"
	SpecStateDeleting   SpecState = "deleting"
	SpecStateDeleted    SpecState = "deleted"
	SpecStateCancelling SpecState = "cancelling"
)

func (a SpecState) Enum() []interface{} {
	return []interface{}{
		SpecStateNew,
		SpecStateCreating,
		SpecStateCreated,
		SpecStateUpdating,
		SpecStateUpdated,
		SpecStateFailed,
		SpecStateDeleting,
		SpecStateDeleted,
	}
}

type StateState string

const (
	StateStateNew      StateState = "new"
	StateStateUpdating StateState = "updating"
	StateStateUpdated  StateState = "updated"
	StateStateDeleting StateState = "deleting"
	StateStateDeleted  StateState = "deleted"
	StateStateFailed   StateState = "failed"
)

func (a StateState) Enum() []interface{} {
	return []interface{}{
		StateStateNew,
		StateStateUpdating,
		StateStateUpdated,
		StateStateDeleting,
		StateStateDeleted,
		StateStateFailed,
	}
}

type EnvironmentState string

const (
	EnvironmentStateNew        EnvironmentState = "new"
	EnvironmentStateCreating   EnvironmentState = "creating"
	EnvironmentStateCreated    EnvironmentState = "created"
	EnvironmentStateUpdating   EnvironmentState = "updating"
	EnvironmentStateUpdated    EnvironmentState = "updated"
	EnvironmentStateFailed     EnvironmentState = "failed"
	EnvironmentStateDeleting   EnvironmentState = "deleting"
	EnvironmentStateDeleted    EnvironmentState = "deleted"
	EnvironmentStateCancelling EnvironmentState = "cancelling"
)

func (a EnvironmentState) Enum() []interface{} {
	return []interface{}{
		EnvironmentStateNew,
		EnvironmentStateCreating,
		EnvironmentStateCreated,
		EnvironmentStateUpdating,
		EnvironmentStateUpdated,
		EnvironmentStateFailed,
		EnvironmentStateDeleting,
		EnvironmentStateDeleted,
	}
}

type TFStateState string

const (
	TFStateNew      TFStateState = "new"
	TFStateLocked   TFStateState = "locked"
	TFStateUnlocked TFStateState = "unlocked"
	TFStateDeleted  TFStateState = "deleted"
)

func (a TFStateState) Enum() []interface{} {
	return []interface{}{
		TFStateNew,
		TFStateLocked,
		TFStateUnlocked,
		TFStateDeleted,
	}
}
