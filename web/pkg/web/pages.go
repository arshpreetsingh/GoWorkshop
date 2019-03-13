package web

import (
	"context"
	"fmt"

	pbBounce "github.com/wizelineacademy/GoWorkshop/proto/bounce"
	"github.com/wizelineacademy/GoWorkshop/web/pkg/tpl"

	"github.com/gocraft/web"
)

func (c *Context) home(w web.ResponseWriter, r *web.Request) {
	d := tpl.Data{
		Data: struct {
			Response string
		}{},
	}
	fmt.Println("this is my d", d)
	//d.RenderJson(w, r)
}

func (c *Context) getbounce(w web.ResponseWriter, r *web.Request) {
	// this accont ID I need to be passed through Variabels as well
	var accountID int32
	accountID = 1234

	fmt.Println("this is my account ID as well", accountID)
	resp, err := c.BounceService.GetData(context.Background(), &pbBounce.BounceGetData{
		// instead of Email I have to put account ID here
		AccountId: accountID,
	})
	if err != nil {
		fmt.Println("this is error", err)
	}
	fmt.Println("this is resp", resp)

	d := tpl.Data{
		Data: struct {
			ID   int32
			Resp *pbBounce.BounceGetDataResponse
		}{
			ID:   accountID,
			Resp: resp,
		},
	}
	fmt.Println("this is my d as well!!!!!", d)
	//d.RenderJson(w, r)
}
