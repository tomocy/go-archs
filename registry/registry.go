package registry

import (
	"github.com/tomocy/archs/adapter/controller"
	"github.com/tomocy/archs/domain/repository"
	"github.com/tomocy/archs/domain/service"
	"github.com/tomocy/archs/infra/bcrypt"
	"github.com/tomocy/archs/infra/memory"
	"github.com/tomocy/archs/infra/session"
	"github.com/tomocy/archs/infra/web/http/handler"
	"github.com/tomocy/archs/usecase"
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
	return handler.NewHandler(
		r.newAuthenticationHandler(),
		r.newTweetHandler(),
		r.newUserHandler(),
	)
}

func (r registry) newAuthenticationHandler() handler.AuthenticationHandler {
	return handler.NewAuthenticationHandler(r.newAuthenticationController())
}

func (r registry) newTweetHandler() handler.TweetHandler {
	return handler.NewTweetHandler(r.newTweetController())
}

func (r registry) newUserHandler() handler.UserHandler {
	return handler.NewUserHandler(r.newUserController())
}

func (r registry) newAuthenticationController() controller.AuthenticationController {
	return controller.NewAuthenticationController(r.newAuthenticationUsecase())
}

func (r registry) newTweetController() controller.TweetController {
	return controller.NewTweetController(r.newTweetUsecase())
}

func (r registry) newUserController() controller.UserController {
	return controller.NewUserController(r.newUserUsecase())
}

func (r registry) newAuthenticationUsecase() usecase.AuthenticationUsecase {
	return usecase.NewAuthenticationUsecase(
		r.userRepository,
		r.newHashService(),
		r.newSessionService(),
	)
}

func (r registry) newTweetUsecase() usecase.TweetUsecase {
	return usecase.NewTweetUsecase(r.userRepository)
}

func (r registry) newUserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(
		r.userRepository,
		r.newUserService(),
		r.newHashService(),
	)
}

func (r registry) newUserRepository() repository.UserRepository {
	return memory.UserRepository
}

func (r registry) newHashService() service.HashService {
	return bcrypt.HashService
}

func (r registry) newSessionService() service.SessionService {
	return session.SessionService
}

func (r registry) newUserService() service.UserService {
	return service.NewUserService(r.userRepository, r.newHashService())
}
