package reset_password

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/helper"

	"github.com/nori-plugins/authentication/internal/domain/errors"

	"github.com/nori-io/common/v5/pkg/domain/logger"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/internal/handler/http/response"
)

type ResetPasswordHandler struct {
	resetPasswordService service.ResetPasswordService
	errorHelper          helper.ErrorHelper
	logger               logger.FieldLogger
}

type Params struct {
	ResetPasswordService service.ResetPasswordService
	ErrorHelper          helper.ErrorHelper
	Logger               logger.FieldLogger
}

func New(params Params) *ResetPasswordHandler {
	return &ResetPasswordHandler{
		resetPasswordService: params.ResetPasswordService,
		errorHelper:          params.ErrorHelper,
		logger:               params.Logger,
	}
}

func (h *ResetPasswordHandler) RequestResetPasswordEmail(w http.ResponseWriter, r *http.Request) {
	data, err := newRequestResetPasswordEmailData(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = h.resetPasswordService.RequestResetPasswordEmail(r.Context(), service.RequestResetPasswordEmailData{
		Email: data.Email,
	}); err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, ResetPasswordResponse{
		Success: true,
		Message: "",
	})
}

func (h *ResetPasswordHandler) SetNewPasswordByResetPasswordEmailToken(w http.ResponseWriter, r *http.Request) {
	data, err := NewSetNewPasswordByResetPasswordEmailToken(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.resetPasswordService.SetNewPasswordByResetPasswordEmailToken(r.Context(), service.SetNewPasswordByResetPasswordEmailTokenData{
		Token:    data.Token,
		Password: data.Password,
	})

	if err == errors.TokenNotFound {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response.JSON(w, r, ResetPasswordSetResponse{
		Success: true,
		Message: "password was changed",
	})
}
