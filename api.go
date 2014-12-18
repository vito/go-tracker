package tracker

type State string

const (
	StateUnscheduled = "unscheduled"
	StatePlanned     = "planned"
	StateStarted     = "started"
	StateFinished    = "finished"
	StateDelivered   = "delivered"
	StateAccepted    = "accepted"
	StateRejected    = "rejected"
)
