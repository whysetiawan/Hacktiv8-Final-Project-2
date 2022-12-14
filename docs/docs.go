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
        "/socialmedias": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "get all social medias based on user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpSuccess-array_models_SocialMediaModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "create a social media",
                "parameters": [
                    {
                        "description": "Create Social Media DTO",
                        "name": "SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpsertSocialMediaDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpSuccess-models_SocialMediaModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    }
                }
            }
        },
        "/socialmedias/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "create a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Social Media Based On User",
                        "name": "SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpsertSocialMediaDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpSuccess-models_SocialMediaModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Delete Social Media By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Social Media ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpSuccess-string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "get mutilple users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpSuccess-array_models_UserModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "create a user",
                "parameters": [
                    {
                        "description": "Update User Based On Token",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpsertUserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpSuccess-dto_UpsertUserDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "delete current user based on JWT",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpSuccess-string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "login a user",
                "parameters": [
                    {
                        "description": "Login User DTO",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpSuccess-models_LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "create a user",
                "parameters": [
                    {
                        "description": "Create User DTO",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpsertUserDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpSuccess-dto_UpsertUserDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HttpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.LoginDto": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.UpsertSocialMediaDto": {
            "type": "object",
            "required": [
                "name",
                "social_media_url"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                }
            }
        },
        "dto.UpsertUserDto": {
            "type": "object",
            "required": [
                "age",
                "email",
                "password",
                "username"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.CommentModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.PhotoModel": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CommentModel"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.SocialMediaModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.UserModel": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CommentModel"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PhotoModel"
                    }
                },
                "social_medias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SocialMediaModel"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "utils.HttpError": {
            "type": "object",
            "properties": {
                "message": {},
                "stack_trace": {}
            }
        },
        "utils.HttpSuccess-array_models_SocialMediaModel": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SocialMediaModel"
                    }
                },
                "message": {}
            }
        },
        "utils.HttpSuccess-array_models_UserModel": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UserModel"
                    }
                },
                "message": {}
            }
        },
        "utils.HttpSuccess-dto_UpsertUserDto": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/dto.UpsertUserDto"
                },
                "message": {}
            }
        },
        "utils.HttpSuccess-models_LoginResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.LoginResponse"
                },
                "message": {}
            }
        },
        "utils.HttpSuccess-models_SocialMediaModel": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.SocialMediaModel"
                },
                "message": {}
            }
        },
        "utils.HttpSuccess-string": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "message": {}
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:3030",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
