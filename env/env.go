package env

type Environment struct {
	Build struct {
		Name        string
		Version     string
		Commit      string
		Date        string
		Tag         string
		Environment string // eg. prod or dev
	}
}

var Env Environment
