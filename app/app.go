package app

type App struct {
	OverridePath string
}

func NewApp() *App {
	return &App{
		OverridePath: "",
	}
}
