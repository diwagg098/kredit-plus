package controllers

import (
	"diwa/kredit-plus/pkg/helper"
	"diwa/kredit-plus/pkg/models"
	"diwa/kredit-plus/pkg/utilities"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (c *ctrl) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	decode := json.NewDecoder(r.Body)
	tokenStr, err := r.Cookie("token")
	if err != nil {
		Response(w, ctx, 401, false, "Unauthorized", nil, nil)
		return
	}
	id := helper.GetUserId(tokenStr.Value)

	var dataPayload models.CreditTrasaction
	err = decode.Decode(&dataPayload)

	if err != nil {
		ctx = utilities.Logf(ctx, "Error create transaction -> : %v", err)
		Response(w, ctx, http.StatusInternalServerError, false, INTERNALSERVER, nil, nil)
		return
	}

	dataPayload.UserId = id

	// get random uuid for contact id
	contractId := uuid.New().String()
	dataPayload.ContractId = contractId

	ctx, status, msg, data, err := c.uc.CreateTransaction(dataPayload, ctx)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error Insert data create transaction -> : %v", err)
		Response(w, ctx, 500, false, INTERNALSERVER, nil, nil)
		return
	}

	Response(w, ctx, status, true, msg, data, nil)
}
