package app

import (
	"context"

	"github.com/vl-usp/water_bot/internal/closer"
	"github.com/vl-usp/water_bot/internal/config"
	"github.com/vl-usp/water_bot/internal/tgbot"
	"github.com/vl-usp/water_bot/pkg/logger"
)

// App represents an application.
type App struct {
	bot             *tgbot.TGBot
	serviceProvider *serviceProvider
}

// NewApp creates a new application.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run starts the application.
func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.Add(func() error {
			return logger.CloseFile()
		})
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runTGBot(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initTGBot,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	// TODO: read from flag
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initTGBot(ctx context.Context) error {
	bot, err := tgbot.New(
		a.serviceProvider.TGConfig().Token(),
		a.serviceProvider.UserService(ctx),
	)
	if err != nil {
		return err
	}

	a.bot = bot
	return nil
}

func (a *App) runTGBot(ctx context.Context) error {
	a.bot.Run(ctx)
	return nil
}
