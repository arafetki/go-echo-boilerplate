package logging

import (
	"io"
	"log/slog"

	"github.com/lmittmann/tint"
)

type Wrapper struct {
	*slog.Logger
	level *slog.LevelVar
}

func (w *Wrapper) SetLevel(level slog.Level) {
	w.level.Set(level)
}

func NewSlogLogger(w io.Writer) *Wrapper {
	levelVar := new(slog.LevelVar)
	levelVar.Set(slog.LevelDebug)
	return &Wrapper{
		Logger: slog.New(tint.NewHandler(w, &tint.Options{Level: levelVar})),
		level:  levelVar,
	}
}
