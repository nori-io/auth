package administrator

import (
	errors2 "github.com/nori-plugins/authentication/internal/domain/errors"
	"github.com/nori-plugins/authentication/pkg/enum/session_status"
	"net/http"
	"strconv"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"

	"github.com/go-chi/chi"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"
	"github.com/nori-plugins/authentication/internal/domain/service"
)

type AdminHandler struct {
	sessionService service.SessionService
	userService  service.UserService
	cookieHelper cookie.CookieHelper
	errorHelper  error2.ErrorHelper
	logger       logger.FieldLogger
}

type Params struct {
	SessionService service.SessionService
	UserService  service.UserService
	CookieHelper cookie.CookieHelper
	ErrorHelper  error2.ErrorHelper
	Logger       logger.FieldLogger
}

func New(params Params) *AdminHandler {
	return &AdminHandler{
		sessionService: params.SessionService,
		userService:  params.UserService,
		cookieHelper: params.CookieHelper,
		errorHelper:  params.ErrorHelper,
		logger:       params.Logger,
	}
}

func (h *AdminHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	sessionId, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	session, err:= h.sessionService.GetBySessionKey(r.Context(), service.GetBySessionKeyData{SessionKey: sessionId})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	if session.Status==session_status.Active{
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	offset := r.URL.Query().Get("offset")
	limit := r.URL.Query().Get("limit")
	status := r.URL.Query().Get("status")
	email := r.URL.Query().Get("email")
	phone := r.URL.Query().Get("phone")

	u, err := strconv.ParseUint(status, 10, 8)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userStatus := users_status.UserStatus(u)

	offsetInt, err := strconv.ParseInt(offset, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	users, err := h.userService.GetAll(r.Context(), service.UserFilter{
		EmailPattern: &email,
		PhonePattern: &phone,
		UserStatus:   &userStatus,
		Offset:       int(offsetInt),
		Limit:        int(limitInt),
	})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, convertAll(users))
}

func (h *AdminHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	_, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	id := chi.URLParam(r, "id")
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := h.userService.GetByID(r.Context(), service.GetByIdData{Id: u})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, convert(*user))
}

func (h *AdminHandler) UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	_, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	id := chi.URLParam(r, "id")
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userStatus := chi.URLParam(r, "user_status")
	u, err = strconv.ParseUint(userStatus, 10, 8)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := h.userService.GetByID(r.Context(), service.GetByIdData{Id: u})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := h.userService.UpdateUserStatus(r.Context(), service.UserUpdateStatusData{
		UserID: user.ID,
		Status: users_status.UserStatus(u),
	}); err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, http.StatusOK)
}

func (h *AdminHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
	}

	id := chi.URLParam(r, "id")
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := h.userService.GetByID(r.Context(), service.GetByIdData{Id: u})
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := h.userService.(r.Context(), service.UserUpdateStatusData{
		UserID: user.ID,
		Status: users_status.UserStatus(u),
	}); err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, http.StatusOK)
}
