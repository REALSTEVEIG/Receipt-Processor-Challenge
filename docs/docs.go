// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/receipts/process": {
            "post": {
                "description": "Submits a receipt for processing and returns a unique receipt ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "receipts"
                ],
                "summary": "Process a receipt",
                "parameters": [
                    {
                        "description": "Receipt to process",
                        "name": "receipt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Receipt"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid receipt",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/receipts/{id}/points": {
            "get": {
                "description": "Retrieves the points awarded for a receipt using its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "receipts"
                ],
                "summary": "Get receipt points",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Receipt ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    },
                    "404": {
                        "description": "Receipt not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Item": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "string"
                },
                "shortDescription": {
                    "type": "string"
                }
            }
        },
        "main.Receipt": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Item"
                    }
                },
                "purchaseDate": {
                    "type": "string"
                },
                "purchaseTime": {
                    "type": "string"
                },
                "retailer": {
                    "type": "string"
                },
                "total": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
