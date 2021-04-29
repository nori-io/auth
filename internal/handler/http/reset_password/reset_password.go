package reset_password

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/errors"

	"github.com/nori-io/common/v4/pkg/domain/logger"
	error2 "github.com/nori-plugins/authentication/internal/domain/helper/error"
	"github.com/nori-plugins/authentication/internal/domain/service"
	"github.com/nori-plugins/authentication/internal/handler/http/response"
)

type ResetPasswordHandler struct {
	resetPasswordService service.ResetPasswordService
	logger               logger.FieldLogger
	errorHelper          error2.ErrorHelper
}

type Params struct {
	ResetPasswordService service.ResetPasswordService
	Logger               logger.FieldLogger
	ErrorHelper          error2.ErrorHelper
}

func New(params Params) *ResetPasswordHandler {
	return &ResetPasswordHandler{
		resetPasswordService: params.ResetPasswordService,
		logger:               params.Logger,
		errorHelper:          params.ErrorHelper,
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

	response.JSON(w, r, RestorePasswordResponse{
		Success: true,
		Message: "",
	})
}

func (h *ResetPasswordHandler) SetNewPasswordByRestorePasswordEmailToken(w http.ResponseWriter, r *http.Request) {
	data, err := NewSetNewPasswordByRestorePasswordEmailToken(r)
	if err != nil {
		h.logger.Error("%s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.resetPasswordService.SetNewPasswordByRestorePasswordEmailToken(r.Context(), service.SetNewPasswordByRestorePasswordEmailTokenData{
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

	response.JSON(w, r, SetPasswordResponse{
		Success: true,
		Message: "password was changed",
	})
}
