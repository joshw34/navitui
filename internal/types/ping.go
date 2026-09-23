package types

type PingResult int

const (
	AuthFailure PingResult = iota
	ServerError
	URLBuildFailure
	Success
)
