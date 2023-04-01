package controllers

import (
	"diwa/kredit-plus/pkg/helper"
	"diwa/kredit-plus/pkg/models"
	"diwa/kredit-plus/pkg/utilities"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

func (c *ctrl) CreditTransactionList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err, ctx := c.uc.CreditTransactionList(ctx)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error Insert data create transaction -> : %v", err)
		Response(w, ctx, 500, false, INTERNALSERVER, nil, nil)
		return
	}

	Response(w, ctx, 200, true, SUCCESS, data, nil)
}

func (c *ctrl) FindByIdCreditTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := chi.URLParam(r, "creditTransactionId")
	id, _ := strconv.Atoi(query)

	ctx, data, err := c.uc.FindByIdCredit(id, ctx)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error Insert data list transaction -> : %v", err)
		Response(w, ctx, 500, false, INTERNALSERVER, nil, nil)
		return
	}

	if data.Id == 0 {
		ctx = utilities.Logf(ctx, "Not row found -> : %v", err)
		Response(w, ctx, 404, false, NOTFOUND, nil, nil)
		return
	}
	Response(w, ctx, 200, true, SUCCESS, data, nil)
}

func (c *ctrl) FindByIdUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := chi.URLParam(r, "userId")
	id, _ := strconv.Atoi(query)

	ctx, data, err := c.uc.FindByIdUser(id, ctx)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error Insert data list transaction -> : %v", err)
		Response(w, ctx, 500, false, INTERNALSERVER, nil, nil)
		return
	}

	if data.Id == 0 {
		ctx = utilities.Logf(ctx, "Not row found -> : %v", err)
		Response(w, ctx, 404, false, NOTFOUND, nil, nil)
		return
	}
	Response(w, ctx, 200, true, SUCCESS, data, nil)
}
