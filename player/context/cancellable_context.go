package context

type CancellableContext struct {
	cancelled bool // cancelled is true if the context is cancelled.
}

// Cancelled returns whether the context is cancelled.
func (ctx *CancellableContext) Cancelled() bool {
	return ctx.cancelled
}

// Cancel cancels the context.
func (ctx *CancellableContext) Cancel() {
	ctx.cancelled = true
}

// UnCancel un-cancels the context.
func (ctx *CancellableContext) UnCancel() {
	ctx.cancelled = false
}
