package route

import (
	"fmt"
	"path"

	"github.com/tomocy/route"
)

func MapRoutes(host, port string) {
	mapRoutes(Web, webRaw, host, port)
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

func mapRoutes(routeMap route.RouteMap, rawMap route.RawMap, host, port string) {
	makeURLAbsolute(rawMap, host, port)
	routeMap.Map(rawMap)
}

func makeURLAbsolute(rmap route.RawMap, host, port string) {
	for key := range rmap {
		rmap[key] = rawAbsoluteURL(host, port, rmap[key])
	}
}

func rawAbsoluteURL(host, port, p string) string {
	return fmt.Sprintf("%s:%s%s", host, port, path.Join("/", p))
}
