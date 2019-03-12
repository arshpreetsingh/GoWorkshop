package web

import (
	"context"
	"fmt"

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
	account_id := 1234
	fmt.Println("this is my account ID as well", account_id)
	resp, err := c.BounceService.GetData(context.Background(), &pbBounce.BounceGetData{
		// instead of Email I have to put account ID here
		account_id: account_id,
	})
	//resp, err := c.BounceService.GetData(context.Background(), &pbBounce.GetData{
	//	UserId: id,
	//})
}
