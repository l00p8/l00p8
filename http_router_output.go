package l00p8

import (
	"github.com/go-chi/chi"
	"github.com/l00p8/xserver"
)

func NewHandlerRouter(router xserver.Router, mw OutputMiddleware) HandlerRouter {
	return &routerWithOutput{router, mw}
}

type HandlerRouter interface {
	Healthers(healthers ...xserver.Healther)

	Get(prefix string, fn Handler)

	Post(prefix string, fn Handler)

	Put(prefix string, fn Handler)

	Delete(prefix string, fn Handler)

	Patch(prefix string, fn Handler)

	Head(prefix string, fn Handler)

	xserver.Muxer
}

type routerWithOutput struct {
	router xserver.Router
	mw     OutputMiddleware
}

func (r *routerWithOutput) Healthers(healthers ...xserver.Healther) {
	r.router.Healthers(healthers...)
}

func (r *routerWithOutput) Get(prefix string, fn Handler) {
	r.router.Get(prefix, r.mw(fn))
}

func (r *routerWithOutput) Post(prefix string, fn Handler) {
	r.router.Post(prefix, r.mw(fn))
}

func (r *routerWithOutput) Put(prefix string, fn Handler) {
	r.router.Put(prefix, r.mw(fn))
}

func (r *routerWithOutput) Patch(prefix string, fn Handler) {
	r.router.Patch(prefix, r.mw(fn))
}

func (r *routerWithOutput) Head(prefix string, fn Handler) {
	r.router.Head(prefix, r.mw(fn))
}

func (r *routerWithOutput) Delete(prefix string, fn Handler) {
	r.router.Delete(prefix, r.mw(fn))
}

func (r *routerWithOutput) Mux() chi.Router {
	return r.router.Mux()
}
