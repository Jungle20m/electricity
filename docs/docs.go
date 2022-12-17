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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "get order by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Order"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "get product by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Product"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "string"
                },
                "error_code": {
                    "type": "string"
                },
                "error_message": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "common.SuccessResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "error_code": {
                    "type": "string"
                },
                "error_message": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Order": {
            "type": "object",
            "properties": {
                "brand_name": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "order_at_time": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "update_time": {
                    "type": "string"
                },
                "working_site_id": {
                    "type": "integer"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "code": {
                    "type": "string"
                },
                "commission": {
                    "type": "number"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "order_id": {
                    "type": "integer"
                },
                "point_back": {
                    "type": "integer"
                },
                "point_back_status": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                },
                "working_site_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8000",
	BasePath:         "/example/v1",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
