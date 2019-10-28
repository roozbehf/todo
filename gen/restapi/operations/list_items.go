// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListItemsHandlerFunc turns a function with the right signature into a list items handler
type ListItemsHandlerFunc func(ListItemsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListItemsHandlerFunc) Handle(params ListItemsParams) middleware.Responder {
	return fn(params)
}

// ListItemsHandler interface for that can handle valid list items params
type ListItemsHandler interface {
	Handle(ListItemsParams) middleware.Responder
}

// NewListItems creates a new http.Handler for the list items operation
func NewListItems(ctx *middleware.Context, handler ListItemsHandler) *ListItems {
	return &ListItems{Context: ctx, Handler: handler}
}

/*ListItems swagger:route GET /list listItems

List all todo items

*/
type ListItems struct {
	Context *middleware.Context
	Handler ListItemsHandler
}

func (o *ListItems) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListItemsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}