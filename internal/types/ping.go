package types

type PingResult int

const (
	PingAuthFailure PingResult = iota
	PingServerError
	PingURLBuildFailure
	PingSuccess
)
