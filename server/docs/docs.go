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
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "get": {
                "description": "checking auth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "checking auth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Checking user authentication",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success authenticate return id",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {}
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Log in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Log in user",
                "parameters": [
                    {
                        "description": "user object for login",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/store.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success login User return cookie",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "not found",
                        "schema": {}
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {}
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "description": "Log out user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Log out user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Log out user",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "void\" \"success log out"
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {}
                    },
                    "401": {
                        "description": "unauthorized",
                        "schema": {}
                    }
                }
            }
        },
        "/restaurants": {
            "get": {
                "description": "giving array of restaurants",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants"
                ],
                "summary": "giving restaurats",
                "responses": {
                    "200": {
                        "description": "success returning array of restaurants",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/store.Restaurant"
                            }
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {}
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Signing up a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Signing up a user",
                "parameters": [
                    {
                        "description": "User object for signing up",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/store.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success create User return id",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {}
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "store.Restaurant": {
            "type": "object",
            "properties": {
                "Category": {
                    "type": "string"
                },
                "CommentsCount": {
                    "type": "integer"
                },
                "DeliveryPrice": {
                    "type": "integer"
                },
                "DeliveryTime": {
                    "type": "integer"
                },
                "ID": {
                    "type": "integer"
                },
                "Icon": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Rating": {
                    "type": "number"
                }
            }
        },
        "store.User": {
            "type": "object",
            "properties": {
                "Birthday": {
                    "type": "string"
                },
                "Email": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "Password": {
                    "type": "string"
                },
                "PhoneNumber": {
                    "type": "string"
                },
                "Username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://84.23.53.216:8001/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Prinesi-Poday API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
