package route

import (
	"fmt"
	"path"

	"github.com/tomocy/route"
)

func MapRoutes(addr string) {
	mapRoutes(Web, webRaw, addr)
}

var (
	Web = make(route.RouteMap)
)

var (
	webRaw = route.RawMap{
		"user.new":    "/users/new",
		"user.create": "/users",
	}
)

func mapRoutes(routeMap route.RouteMap, rawMap route.RawMap, addr string) {
	makeURLAbsolute(rawMap, addr)
	routeMap.Map(rawMap)
}

func makeURLAbsolute(rmap route.RawMap, addr string) {
	for key := range rmap {
		rmap[key] = rawAbsoluteURL(addr, rmap[key])
	}
}

func rawAbsoluteURL(addr, p string) string {
	return fmt.Sprintf("%s://%s%s", "http", addr, path.Join("/", p))
}
