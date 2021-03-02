package authentication

import (
	"net/http"

	"github.com/nori-plugins/authentication/internal/domain/entity"

	"github.com/nori-plugins/authentication/internal/domain/service"
)

type Handler struct {
	AuthenticationService service.AuthenticationService
}

func New(authenticationService service.AuthenticationService) *Handler {
	return &Handler{AuthenticationService: authenticationService}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	data, err := newSignUpData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err := h.AuthenticationService.SignUp(r.Context(), data)
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

func (h *Handler) SigIn(w http.ResponseWriter, r *http.Request) {
	data, err := newSignInData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sess, err := h.AuthenticationService.SignIn(r.Context(), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	JSON(w, r, SignInResponse{
		SessionID: string(sess.SessionKey),
	})
}

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	// todo: extract session ID from context
	sessionIdContext := r.Context().Value("session_id")

	sessionId, _ := sessionIdContext.([]byte)

	if err := h.AuthenticationService.SignOut(r.Context(), &entity.Session{SessionKey: sessionId}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// todo: redirect

	http.Redirect(w, r, "/", 0)
}

func (h *Handler) PutSecret(w http.ResponseWriter, r *http.Request) {
	data, err := newPutSecretData(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	sessionIdContext := r.Context().Value("session_id")
	sessionId, _ := sessionIdContext.([]byte)

	if data.Ssid != sessionIdContext {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	sessionUserId := r.Context().Value("user_id").(uint64)
	login, issuer, err :=
		h.AuthenticationService.PutSecret(r.Context(), &entity.Session{SessionKey: sessionId, UserID: sessionUserId})

	if (login == "") && (issuer == "") {
		http.Error(w, "sign up error", http.StatusInternalServerError)
	}

	JSON(w, r, MfaSecretResponse{
		Login:  login,
		Issuer: issuer,
	})
}
