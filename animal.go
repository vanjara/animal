package animal

type game struct {
	Running bool
}

func New() game {
	return game{
		Running: true,
	}
}
