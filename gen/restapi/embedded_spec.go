// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "ToDo",
    "version": "0.1.0"
  },
  "paths": {
    "/add/{itemTitle}": {
      "post": {
        "description": "Add a new todo item",
        "operationId": "addItem",
        "responses": {
          "202": {
            "description": "Accepted",
            "schema": {
              "$ref": "#/definitions/AddResponse"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "itemTitle",
          "in": "path",
          "required": true
        }
      ]
    },
    "/list": {
      "get": {
        "description": "List all todo items",
        "operationId": "listItems",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ListResponse"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/remove/{id}": {
      "delete": {
        "description": "Delete all todo items",
        "operationId": "removeItem",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/RemoveResponse"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AddResponse": {
      "type": "object",
      "required": [
        "id",
        "title"
      ],
      "properties": {
        "id": {
          "description": "the unique id of the newly added item",
          "type": "string"
        },
        "title": {
          "description": "title of the newly added item",
          "type": "string"
        }
      }
    },
    "ErrorResponse": {
      "type": "string"
    },
    "ListResponse": {
      "type": "object",
      "required": [
        "todoItems"
      ],
      "properties": {
        "todoItems": {
          "type": "array",
          "items": {
            "type": "object",
            "required": [
              "id",
              "title"
            ],
            "properties": {
              "id": {
                "description": "the unique id of the newly added item",
                "type": "string"
              },
              "title": {
                "description": "title of the newly added item",
                "type": "string"
              }
            }
          }
        }
      }
    },
    "RemoveResponse": {
      "type": "object",
      "required": [
        "id",
        "title"
      ],
      "properties": {
        "id": {
          "description": "the unique id of the removed item",
          "type": "string"
        },
        "title": {
          "description": "title of the removed item",
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "ToDo",
    "version": "0.1.0"
  },
  "paths": {
    "/add/{itemTitle}": {
      "post": {
        "description": "Add a new todo item",
        "operationId": "addItem",
        "responses": {
          "202": {
            "description": "Accepted",
            "schema": {
              "$ref": "#/definitions/AddResponse"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "itemTitle",
          "in": "path",
          "required": true
        }
      ]
    },
    "/list": {
      "get": {
        "description": "List all todo items",
        "operationId": "listItems",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/ListResponse"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/remove/{id}": {
      "delete": {
        "description": "Delete all todo items",
        "operationId": "removeItem",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/RemoveResponse"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AddResponse": {
      "type": "object",
      "required": [
        "id",
        "title"
      ],
      "properties": {
        "id": {
          "description": "the unique id of the newly added item",
          "type": "string"
        },
        "title": {
          "description": "title of the newly added item",
          "type": "string"
        }
      }
    },
    "ErrorResponse": {
      "type": "string"
    },
    "ListResponse": {
      "type": "object",
      "required": [
        "todoItems"
      ],
      "properties": {
        "todoItems": {
          "type": "array",
          "items": {
            "type": "object",
            "required": [
              "id",
              "title"
            ],
            "properties": {
              "id": {
                "description": "the unique id of the newly added item",
                "type": "string"
              },
              "title": {
                "description": "title of the newly added item",
                "type": "string"
              }
            }
          }
        }
      }
    },
    "RemoveResponse": {
      "type": "object",
      "required": [
        "id",
        "title"
      ],
      "properties": {
        "id": {
          "description": "the unique id of the removed item",
          "type": "string"
        },
        "title": {
          "description": "title of the removed item",
          "type": "string"
        }
      }
    }
  }
}`))
}
