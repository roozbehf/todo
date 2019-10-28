package todo

import (
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	models "github.com/roozbehf/todo/gen/models"
	ops "github.com/roozbehf/todo/gen/restapi/operations"
)

// Item holds a todo item
type Item struct {
	ID    string
	Title string
}

var lastID int = 0
var items = make(map[string]Item)

// AddItem adds a new todo item to the list
func AddItem(param ops.AddItemParams) middleware.Responder {
	title := param.ItemTitle
	newID := strconv.Itoa(lastID)
	newItem := Item{ID: newID, Title: title}
	items[newID] = newItem
	lastID++

	resp := models.AddResponse{ID: &newID, Title: &title}
	return ops.NewAddItemAccepted().WithPayload(&resp)
}

// ListItems lists all the todo items
func ListItems(param ops.ListItemsParams) middleware.Responder {
	resp := models.ListResponse{}

	var respItems = make([]*models.ListResponseTodoItemsItems0, len(items))
	var i int64 = 0
	for _, item := range items {
		id := item.ID
		title := item.Title
		respItems[i] = &models.ListResponseTodoItemsItems0{ID: &id, Title: &title}
		i++
	}

	resp.TodoItems = respItems
	return ops.NewListItemsOK().WithPayload(&resp)
}

// RemoveItem removes an item from the list if exists
func RemoveItem(param ops.RemoveItemParams) middleware.Responder {
	key := param.ID

	if item, ok := items[key]; ok {
		id := item.ID
		title := item.Title
		delete(items, key)
		resp := models.RemoveResponse{ID: &id, Title: &title}
		return ops.NewRemoveItemOK().WithPayload(&resp)
	}

	// if item is not found
	errorObj := models.ErrorResponse("item not found")
	return ops.NewRemoveItemDefault(502).WithPayload(errorObj)
}
