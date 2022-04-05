// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/halls": {
            "get": {
                "description": "get halls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Halls"
                ],
                "summary": "List halls",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {}
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates hall and returns created object",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Halls"
                ],
                "summary": "Create hall",
                "parameters": [
                    {
                        "description": "The body to create a hall",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {}
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            }
        },
        "/halls/{id}": {
            "get": {
                "description": "Gets hall",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Halls"
                ],
                "summary": "Get hall",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Hall ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Deletes hall",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Halls"
                ],
                "summary": "Delete hall",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Hall ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            }
        },
        "/halls/{id}/sessions": {
            "post": {
                "description": "Creates session and returns created object",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Create session",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Session ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The body to create a session",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {}
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "description": "get movies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "List movie",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {}
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates movie and returns created object",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Create movie",
                "parameters": [
                    {
                        "description": "The body to create a movie",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {}
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            }
        },
        "/movies/{id}": {
            "get": {
                "description": "Gets movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Get movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Deletes movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Delete movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            }
        },
        "/sessions": {
            "get": {
                "description": "get sessions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "List session",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {}
                            }
                        }
                    }
                }
            }
        },
        "/sessions/{id}": {
            "get": {
                "description": "Gets session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Get session",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Session ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Deletes session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sessions"
                ],
                "summary": "Delete session",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Session ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            }
        },
        "/sessions/{id}/tickets": {
            "post": {
                "description": "Creates ticket and returns created object",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickets"
                ],
                "summary": "Create ticket",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ticket ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The body to create a ticket",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {}
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "description": "Signin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Signin",
                "parameters": [
                    {
                        "description": "Email and password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {}
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {}
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Signup",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Signup",
                "parameters": [
                    {
                        "description": "Email and password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {}
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {}
                    }
                }
            }
        },
        "/tickets": {
            "get": {
                "description": "get tickets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickets"
                ],
                "summary": "List ticket",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {}
                            }
                        }
                    }
                }
            }
        },
        "/tickets/{id}": {
            "get": {
                "description": "Gets ticket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickets"
                ],
                "summary": "Get ticket",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ticket ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Deletes ticket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickets"
                ],
                "summary": "Delete ticket",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ticket ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            }
        },
        "/tickets/{id}/download": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickets"
                ],
                "summary": "Download ticket",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ticket ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8085",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Cinetickets API",
	Description:      "This is a sample cinetickets server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
