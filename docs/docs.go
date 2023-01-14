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
        "/foobar": {
            "get": {
                "description": "Shown all Foobar in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foobar"
                ],
                "summary": "Showning all Foobar",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Foobar"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Foobar in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foobar"
                ],
                "summary": "Creating Foobar",
                "parameters": [
                    {
                        "description": "Foobar Data",
                        "name": "foobar",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Foobar"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Foobar"
                        }
                    }
                }
            }
        },
        "/foobar/mock": {
            "get": {
                "description": "Shown all Foobar Mocks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foobar"
                ],
                "summary": "Showning all Foobar Mocks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Foobar"
                        }
                    }
                }
            }
        },
        "/foobar/{id}": {
            "get": {
                "description": "Shown Foobar by ID in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foobar"
                ],
                "summary": "Showning Foobar by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Foobar ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Foobar"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Foobar by ID in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foobar"
                ],
                "summary": "Deleting Foobar by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Foobar ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Foobar"
                        }
                    }
                }
            },
            "patch": {
                "description": "Edit Foobar by ID in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foobar"
                ],
                "summary": "Editing Foobar by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Foobar ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Foobar Data",
                        "name": "foobar",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Foobar"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Foobar"
                        }
                    }
                }
            }
        },
        "/foobar/{reg}": {
            "get": {
                "description": "Shown Foobar by Registration in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foobar"
                ],
                "summary": "Showning Foobar by Registration",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Foobar Registration",
                        "name": "reg",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Foobar"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Foobar": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string",
                    "maxLength": 60
                },
                "registration": {
                    "type": "integer",
                    "minimum": 8
                },
                "updatedAt": {
                    "type": "string"
                },
                "what": {
                    "type": "string",
                    "maxLength": 30
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}