package social_provider

import (
	"github.com/nori-plugins/authentication/internal/config"
	"github.com/nori-plugins/authentication/internal/domain/helper/security"
	"github.com/nori-plugins/authentication/internal/domain/repository"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type SocialProviderService struct {
	sessionRepository repository.SessionRepository
	userService       service.UserService
	config            config.Config
	securityHelper    security.SecurityHelper
}
