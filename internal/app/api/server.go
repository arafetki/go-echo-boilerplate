package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func (api *API) serveHTTP() error {
	shutdownErrChan := make(chan error)

	go func() {

		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan
		ctx, cancel := context.WithTimeout(context.Background(), api.config.Server.ShutdowPeriod)
		defer cancel()

		shutdownErrChan <- api.echo.Shutdown(ctx)

	}()

	api.logger.Info("🚀 server started", "env", api.config.App.Env, "address", api.config.Server.Addr)
	if err := api.echo.Start(api.config.Server.Addr); err != nil && err != http.ErrServerClosed {
		return err
	}

	err := <-shutdownErrChan
	if err != nil {
		return err
	}
	api.wg.Wait()
	api.logger.Warn("server stopped gracefully")
	return nil
}
