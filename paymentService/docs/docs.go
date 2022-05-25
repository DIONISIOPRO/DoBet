// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Dionisio Paulo",
            "url": "meusite.com",
            "email": "dionisiopaulonamuetho@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/deposit": {
            "get": {
                "description": "this route allows you to deposit your money",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "deposit your money",
                "parameters": [
                    {
                        "description": "deposit",
                        "name": "deposit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Deposit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "sucess",
                        "schema": {
                            "$ref": "#/definitions/controller.SucessResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/withdraw": {
            "get": {
                "description": "this route allows you to WithDraw your money",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "WithDraw your money",
                "parameters": [
                    {
                        "description": "WithDraw",
                        "name": "WithDraw",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.WithDraw"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "sucess",
                        "schema": {
                            "$ref": "#/definitions/controller.SucessResponse"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "controller.SucessResponse": {
            "type": "object",
            "properties": {
                "sucess": {
                    "type": "string"
                }
            }
        },
        "domain.Deposit": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "domain.WithDraw": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9003``",
	BasePath:         "api/v1",
	Schemes:          []string{},
	Title:            "Payment API",
	Description:      "This is a payment service for DoBet Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}