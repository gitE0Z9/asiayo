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
        "/v1/exchange-rate/{source}/conversion/{target}": {
            "get": {
                "description": "exchange rate conversion",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "exchange rate"
                ],
                "summary": "exchange rate conversion",
                "parameters": [
                    {
                        "enum": [
                            "JPY",
                            "TWD",
                            "USD"
                        ],
                        "type": "string",
                        "description": "source currency",
                        "name": "source",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "JPY",
                            "TWD",
                            "USD"
                        ],
                        "type": "string",
                        "description": "target currency",
                        "name": "target",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"$1,111.05\"",
                        "description": "amount",
                        "name": "amount",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/exchange_rate.ConversionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad parameter\" example({message=\"bad parameter\",amount=\"0\"})",
                        "schema": {
                            "$ref": "#/definitions/exchange_rate.ConversionResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "exchange_rate.ConversionResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "msg": {
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
