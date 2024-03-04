package main

func callbackExit(cfg *config, args ...string) error {
	// avoiding using os.exit(1)
	return nil
}
