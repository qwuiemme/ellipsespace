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
        "/api/add-object-catologue": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add a record of the object to the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MainAPI"
                ],
                "summary": "Add Object Catalogue",
                "parameters": [
                    {
                        "description": "Object info",
                        "name": "Input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/catalogueobject.CatalogueObject"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    }
                }
            }
        },
        "/api/delete-object-catologue": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a record of the object to the database by name.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MainAPI"
                ],
                "summary": "Delete Object Catalogue",
                "parameters": [
                    {
                        "description": "Object info",
                        "name": "Input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/catalogueobject.CatalogueObjectJsonGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    }
                }
            }
        },
        "/api/get-all-object-catologue": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns all object records in the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MainAPI"
                ],
                "summary": "Get All Objects Catalogue",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/catalogueobject.CatalogueObject"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    }
                }
            }
        },
        "/api/get-object-catologue": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Returns an object record or null object with the passed name.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MainAPI"
                ],
                "summary": "Get Object Catalogue",
                "parameters": [
                    {
                        "description": "Object name",
                        "name": "Input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/catalogueobject.CatalogueObjectJsonGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/catalogueobject.CatalogueObject"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    }
                }
            }
        },
        "/api/session/auth": {
            "get": {
                "description": "Checks the entered data for correctness and returns the JWT token if the check is successful.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Authorize in Session",
                "parameters": [
                    {
                        "description": "Session data",
                        "name": "Input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authorization.SessionJsonGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    }
                }
            }
        },
        "/api/session/create": {
            "post": {
                "description": "Writes a new session to the database and returns its Id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Create Session",
                "parameters": [
                    {
                        "description": "Session data",
                        "name": "Input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authorization.Session"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    }
                }
            }
        },
        "/api/session/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete the session data with the specified Id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Delete Session",
                "parameters": [
                    {
                        "description": "Object info",
                        "name": "Input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authorization.SessionJsonDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    }
                }
            }
        },
        "/api/session/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the info from the current JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Info from the current JWT token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authorization.SessionBase"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/api/session/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates the session data with the specified Id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Update Session",
                "parameters": [
                    {
                        "description": "Session data",
                        "name": "Input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authorization.Session"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    }
                }
            }
        },
        "/api/update-object-catologue": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a record of the object to the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MainAPI"
                ],
                "summary": "Update Object Catalogue",
                "parameters": [
                    {
                        "description": "Object info",
                        "name": "Input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/catalogueobject.CatalogueObject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serverstatus.StatusJson"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authorization.Session": {
            "type": "object",
            "properties": {
                "access-level": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "sname": {
                    "type": "string"
                }
            }
        },
        "authorization.SessionBase": {
            "type": "object",
            "properties": {
                "access-level": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "sname": {
                    "type": "string"
                }
            }
        },
        "authorization.SessionJsonDelete": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "authorization.SessionJsonGet": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "sname": {
                    "type": "string"
                }
            }
        },
        "catalogueobject.CatalogueObject": {
            "type": "object",
            "properties": {
                "avg-radius": {
                    "description": "?????????????? ????????????",
                    "type": "number"
                },
                "description": {
                    "description": "?????????????????? ????????????????",
                    "type": "string"
                },
                "equator-radius": {
                    "description": "???????????????????????????? ????????????",
                    "type": "number"
                },
                "g": {
                    "description": "?????????????????? ???????????????????? ??????????????",
                    "type": "number"
                },
                "inclination": {
                    "description": "????????????????????",
                    "type": "number"
                },
                "m": {
                    "description": "??????????",
                    "type": "number"
                },
                "name": {
                    "description": "???????????????? ????????????????",
                    "type": "string"
                },
                "o-date-time": {
                    "description": "???????? ?????????????????????? EllipseSpace",
                    "type": "string"
                },
                "orbital-vel": {
                    "description": "?????????????????????? ????????????????",
                    "type": "number"
                },
                "p": {
                    "description": "?????????????? ??????????????????",
                    "type": "number"
                },
                "photos": {
                    "description": "????????????????????",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "polar-radius": {
                    "description": "???????????????? ????????????",
                    "type": "number"
                },
                "s": {
                    "description": "??????????????",
                    "type": "number"
                },
                "s-conversion-period": {
                    "description": "???????????????????????? ???????????? ??????????????????",
                    "type": "number"
                },
                "satelites": {
                    "description": "????????????????",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "v": {
                    "description": "??????????",
                    "type": "number"
                },
                "v1": {
                    "description": "???????????? ?????????????????????? ????????????????",
                    "type": "number"
                },
                "v2": {
                    "description": "???????????? ?????????????????????? ????????????????",
                    "type": "number"
                },
                "whose-satelite": {
                    "description": "?????? ??????????????",
                    "type": "string"
                }
            }
        },
        "catalogueobject.CatalogueObjectJsonGet": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "serverstatus.StatusJson": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.2.0",
	Host:             "ellipsespace.onrender.com",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "EllipseSpace API",
	Description:      "API for the Encyclopedia of Space project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
