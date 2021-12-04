package sendkeys

// KBOpt are options for our wrapper
type KBOpt uint8

const (
	// Stubborn will cause our sequences to continue despite errors.
	// Otherwise, we will stop if our error count is over 0.
	Stubborn KBOpt = iota
	// Noisy will cause all errors to be printed to stdout.
	Noisy
	// Random will use random sleeps throughout the typing process.
	// Otherwise, a static 10 milliseconds will be used.
	Random
	// NoDelay will bypass the 2 second delay for linux, mostly for testing.
	NoDelay
)

func (kb *KBWrap) processOptions(opts ...KBOpt) {
	kbo := map[KBOpt]*bool{
		Stubborn: &kb.stubborn,
		Noisy:    &kb.noisy,
		Random:   &kb.random,
		NoDelay:  &kb.nodelay,
	}
	for _, o := range opts {
		if option, ok := kbo[o]; !ok {
			continue
		} else {
			*option = true
		}
	}
}
