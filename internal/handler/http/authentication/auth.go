package authentication

import (
	"net/http"

	"github.com/nori-io/authentication/internal/domain/entity"

	"github.com/nori-io/authentication/internal/domain/service"
)

type AuthHandler struct {
	Auth service.AuthenticationService
}

func New(auth service.AuthenticationService) *AuthHandler {
	return &AuthHandler{Auth: auth}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	data, err := newSignUpData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := h.Auth.SignUp(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if user == nil {
		http.Error(w, "sign up error", http.StatusInternalServerError)
	}
	JSON(w, r, SignUpResponse{
		ID:    user.Id,
		Email: user.Email,
	})
}

func (h *AuthHandler) SigIn(w http.ResponseWriter, r *http.Request) {
	data, err := newSignInData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sess, err := h.Auth.SignIn(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	JSON(w, r, SignInResponse{
		SessionID: string(sess.Id),
	})
}

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	// todo: extract session ID from context
	sessionIdContext := r.Context().Value("session_id")

	sessionId, _ := sessionIdContext.([]byte)

	err := h.Auth.SignOut(r.Context(), &entity.Session{Id: sessionId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// todo: redirect
	http.Redirect(w, r, "/", 0)
}