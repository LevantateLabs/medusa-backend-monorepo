package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/akhil-is-watching/medusa-backend-monorepo/pkg/models"
	"github.com/nats-io/nats.go"
)

func (h *BaseHandler) HandleAuthCreated(msg *nats.Msg) {
	log.Println("Received message", string(msg.Data))

	var auth models.Auth
	if err := json.Unmarshal(msg.Data, &auth); err != nil {
		log.Println("Error unmarshalling message", err)
		return
	}

	patient := models.Patient{
		ID:           auth.ID,
		AadharNumber: auth.AadharNumber,
		Name:         auth.Name,
		Email:        auth.Email,
		Phone:        auth.Phone,
		Address:      auth.Address,
		Age:          auth.Age,
		Sex:          auth.Sex,
	}

	if _, err := h.patientRepo.CreatePatient(context.Background(), patient); err != nil {
		log.Println("Error creating patient", err)
		return
	}

}
