package controllers

import (
	"diwa/kredit-plus/pkg/helper"
	"diwa/kredit-plus/pkg/utilities"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (c *ctrl) GetDataUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err, ctx := c.uc.GetDataListUser(ctx)
	if err != nil {
		ctx = utilities.Logf(ctx, "Error Get Data User -> : %v", err)
		println(err.Error())
		Response(w, ctx, 500, false, INTERNALSERVER, nil, nil)
		return
	}
	Response(w, ctx, 200, true, SUCCESS, data, nil)
}

func (c *ctrl) FindByIdLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tokenStr, err := r.Cookie("token")
	if err != nil {
		Response(w, ctx, 500, false, INTERNALSERVER, nil, nil)
		return
	}
	id := helper.GetUserId(tokenStr.Value)

	ctx, data, err := c.uc.GetUserById(id, ctx)
	if err != nil {
		Response(w, ctx, 500, false, INTERNALSERVER, nil, nil)
		return
	}

	if data == nil {
		Response(w, ctx, 404, false, NOTFOUND, nil, nil)
		return
	}

	Response(w, ctx, 200, true, SUCCESS, data, nil)
}
func (c *ctrl) FindById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	query := chi.URLParam(r, "userId")
	id, _ := strconv.Atoi(query)

	ctx, data, err := c.uc.GetUserById(id, ctx)
	if err != nil {
		Response(w, ctx, 500, false, INTERNALSERVER, nil, nil)
		return
	}

	if data.User.Id == 0 {
		Response(w, ctx, 404, false, NOTFOUND, nil, nil)
		return
	}

	Response(w, ctx, 200, true, SUCCESS, data, nil)
}
