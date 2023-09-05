package logs2

func log() {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"error1": "error1",
		},
	).Info("Something happened")
	log.Warn("You should probably take a look at this.")
}
