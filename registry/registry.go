package registry

import (
	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/adapter/presenter"
	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/domain/service"
	"github.com/tomocy/archs/infra/memory"
	infraservice "github.com/tomocy/archs/infra/service"
	"github.com/tomocy/archs/infra/web/handler"
	"github.com/tomocy/archs/usecase"
	"github.com/tomocy/archs/usecase/response"
)

type Registry interface {
	NewHandler() handler.Handler
}

type registry struct {
	userRepository repository.UserRepository
}

func NewRegistry() Registry {
	registry := new(registry)
	registry.initRepositories()
	return registry
}

func (r *registry) initRepositories() {
	r.userRepository = r.newUserRepository()
}

func (r registry) NewHandler() handler.Handler {
	return handler.NewHandler(r.newUserHandler())
}

func (r registry) newUserHandler() handler.UserHandler {
	return handler.NewUserHandler(r.newUserController())
}

func (r registry) newUserController() controller.UserController {
	return controller.NewUserController(r.newUserUsecase())
}

func (r registry) newUserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(r.userRepository, r.newUserResponseWriter(), r.newUserService(), r.newSessionService())
}

func (r registry) newUserRepository() repository.UserRepository {
	return memory.NewUserRepository()
}

func (r registry) newUserResponseWriter() response.UserResponseWriter {
	return presenter.NewUserResponseWriter()
}

func (r registry) newUserService() service.UserService {
	return service.NewUserService(r.userRepository, infraservice.NewHashService())
}

func (r registry) newSessionService() service.SessionService {
	return infraservice.NewSessionService()
}
