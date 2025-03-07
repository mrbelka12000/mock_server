package handler

import "log/slog"

type opt func(d *DynamicRouter)

func WithLogger(log *slog.Logger) opt {
	return func(d *DynamicRouter) {
		d.log = log
	}
}
