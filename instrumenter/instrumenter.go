package instumenter

type Instrumenter interface {
	Instrument(string) ([]byte, error)
}
