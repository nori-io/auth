package administrator

import (
	"net/http"
	"strconv"

	"github.com/nori-plugins/authentication/pkg/enum/users_status"

	"github.com/go-chi/chi"

	"github.com/nori-plugins/authentication/internal/handler/http/response"

	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-plugins/authentication/internal/domain/helper/cookie"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/internal/service/user"
)

type AdminHandler struct {
	userService  user.UserService
	logger       logger.FieldLogger
	cookieHelper cookie.CookieHelper
	errorHelper  error2.ErrorHelper
}

type Params struct {
	UserService  user.UserService
	Logger       logger.FieldLogger
	CookieHelper cookie.CookieHelper
	ErrorHelper  error2.ErrorHelper
}

func New(params Params) *AdminHandler {
	return &AdminHandler{
		userService:  params.UserService,
		logger:       params.Logger,
		cookieHelper: params.CookieHelper,
		errorHelper:  params.ErrorHelper,
	}
}

func (h *AdminHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	_, err := h.cookieHelper.GetSessionID(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, http.ErrNoCookie.Error(), http.StatusUnauthorized)
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

	response.JSON(w, r, convert(*user))
}
