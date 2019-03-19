package handlers

import (
	"log"
	"net/http"
	"time"
	"encoding/json"

	"github.com/go-chi/chi"
	"hackathon/TS/helpers"
	"hackathon/TS/otp"
)

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	helpers.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"success": true})
}

func (h *Handlers) Service(w http.ResponseWriter, r *http.Request) {
	reqDetails := struct {
		PhoneNumber string `json:"phone_number"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&reqDetails)
	if err != nil {
		helpers.RespondwithJSON(w, http.StatusBadRequest, map[string]interface{}{"errors": "Invalid phone number"})
		return
	}

	otpID := 0
	otpID, err = otp.GenerateOTP(reqDetails.PhoneNumber)
	if err != nil {
		helpers.RespondwithJSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	helpers.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"ID": otpID})
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("Request processed in %s\n", time.Since(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(r *chi.Mux) {

	// Public routes
	r.Get("/", h.Logger(h.Home))
	r.Get("/service", h.Logger(h.Service))

	// TODO Take care of version
	// TODO
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
