package recallengine

// Logger interface API for log.Logger
type Logger interface {
	Printf(string, ...interface{})
}
