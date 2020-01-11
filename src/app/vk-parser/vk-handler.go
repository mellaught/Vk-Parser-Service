package parser

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttprouter"

	"github.com/mrKitikat/Vk-Parser-Service/src/app/models"

	"github.com/valyala/fasthttp"
)

// Handler for API method "/getProfiles" get profiles by intersection
func (vk *VkParser) GetProfilesHandler(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {

	req := &models.IntersecReq{}
	err := json.Unmarshal(ctx.PostBody(), &req)
	if err != nil {
		fmt.Println(err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	groups, groupsCount, err := vk.GetUserSubscriptions(req.Id)
	if err != nil {
		fmt.Println(err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	fmt.Printf("User %d have %d groups.\nStart searching intersection...\n", req.Id, groupsCount)
	t := time.Now()
	users, err := vk.GetMembers(groups, req)
	if err != nil {
		fmt.Println(err)
		//ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	t2 := time.Now()
	fmt.Println(t2.Sub(t))

	if len(users) == 0 {
		resp := models.Answer{
			Text:     "The list is empty",
			Responce: users,
		}

		bytes, err := json.Marshal(resp)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		ctx.SetBody(bytes)
		ctx.SetStatusCode(fasthttp.StatusOK)
		return
	}

	resp := models.Answer{
		Text:     fmt.Sprintf("We found %d people", len(users)),
		Responce: users,
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetBody(bytes)
	ctx.SetStatusCode(fasthttp.StatusOK)

	return
}
