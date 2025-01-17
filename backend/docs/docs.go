// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "LangBridge Support",
            "url": "http://langbridge.io/support",
            "email": "support@langbridge.io"
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
        "/api/v1/projects/{projectID}/keys": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "keys"
                ],
                "summary": "Create key with translates",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Project ID",
                        "name": "projectID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create key request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/keys.CreateKeyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created key",
                        "schema": {
                            "$ref": "#/definitions/keys.CreateKeyResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internal_api_http_keys.Translate": {
            "type": "object",
            "properties": {
                "language": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "key.Platform": {
            "type": "string",
            "enum": [
                "WEB",
                "IOS",
                "ANDROID",
                "OTHER"
            ],
            "x-enum-varnames": [
                "PlatformWeb",
                "PlatformIOS",
                "PlatformAndroid",
                "PlatformOther"
            ]
        },
        "keys.CreateKeyRequest": {
            "type": "object",
            "properties": {
                "existedTags": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "name": {
                    "type": "string"
                },
                "newTags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "platforms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/key.Platform"
                    }
                },
                "translates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/internal_api_http_keys.Translate"
                    }
                }
            }
        },
        "keys.CreateKeyResponse": {
            "type": "object",
            "properties": {
                "key_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "https://api.langbridge.io",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "LangBridge API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
