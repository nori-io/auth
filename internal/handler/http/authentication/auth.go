package authentication

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/entity"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

type AuthenticationHandler struct {
	authenticationService service.AuthenticationService
}

func New(authenticationService service.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{authenticationService: authenticationService}
}

func (h *AuthenticationHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	data, err := newSignUpData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := h.authenticationService.SignUp(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if user == nil {
		http.Error(w, "sign up error", http.StatusInternalServerError)
	}
	JSON(w, r, SignUpResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}

func (h *AuthenticationHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	data, err := newSignInData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sess, err := h.authenticationService.SignIn(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	JSON(w, r, SignInResponse{
		SessionID: string(sess.SessionKey),
	})
}

func (h *AuthenticationHandler) SignInMfa(w http.ResponseWriter, r *http.Request) {
	data, err := newSignInMfaData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sess, err := h.authenticationService.SignInMfa(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	JSON(w, r, SignInResponse{
		SessionID: string(sess.SessionKey),
	})
}

func (h *AuthenticationHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	// todo: extract session ID from context
	sessionIdContext := r.Context().Value("session_id")

	sessionId, _ := sessionIdContext.([]byte)

	if err := h.authenticationService.SignOut(r.Context(), &entity.Session{SessionKey: sessionId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// todo: redirect

	http.Redirect(w, r, "/", 0)
}
