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
        "/demo/:id": {
            "get": {
                "description": "查询demo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "demo"
                ],
                "summary": "查询demo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标识",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功返回的结果",
                        "schema": {
                            "$ref": "#/definitions/model.ResultGeneric-model_Demo"
                        }
                    },
                    "500": {
                        "description": "请求失败返回的结果",
                        "schema": {
                            "$ref": "#/definitions/model.ResultGeneric-model_Demo"
                        }
                    }
                }
            }
        },
        "/demo/add": {
            "post": {
                "description": "创建demo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "demo"
                ],
                "summary": "创建demo",
                "parameters": [
                    {
                        "description": "demo数据",
                        "name": "demo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Demo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功返回的结果",
                        "schema": {
                            "$ref": "#/definitions/model.ResultGeneric-int"
                        }
                    },
                    "500": {
                        "description": "请求失败返回的结果",
                        "schema": {
                            "$ref": "#/definitions/model.ResultGeneric-int"
                        }
                    }
                }
            }
        },
        "/demo/getDemo": {
            "get": {
                "description": "查询demo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "demo"
                ],
                "summary": "查询demo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标识",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功返回的结果",
                        "schema": {
                            "$ref": "#/definitions/model.ResultGeneric-model_Demo"
                        }
                    },
                    "500": {
                        "description": "请求失败返回的结果",
                        "schema": {
                            "$ref": "#/definitions/model.ResultGeneric-model_Demo"
                        }
                    }
                }
            }
        },
        "/demo/search/:id": {
            "get": {
                "description": "搜索demo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "demo"
                ],
                "summary": "搜索demo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标识",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "名字",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功返回的结果",
                        "schema": {
                            "$ref": "#/definitions/model.ResultGeneric-array_model_Demo"
                        }
                    },
                    "500": {
                        "description": "请求失败返回的结果",
                        "schema": {
                            "$ref": "#/definitions/model.ResultGeneric-array_model_Demo"
                        }
                    }
                }
            }
        },
        "/demo/search2": {
            "post": {
                "description": "搜索demo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "demo"
                ],
                "summary": "搜索demo",
                "parameters": [
                    {
                        "description": "搜索条件",
                        "name": "pms",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Demo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功返回的结果",
                        "schema": {
                            "$ref": "#/definitions/model.ResultGeneric-array_model_Demo"
                        }
                    },
                    "500": {
                        "description": "请求失败返回的结果",
                        "schema": {
                            "$ref": "#/definitions/model.ResultGeneric-array_model_Demo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Demo": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "just ID",
                    "type": "integer"
                },
                "name": {
                    "description": "名字",
                    "type": "string"
                }
            }
        },
        "model.ResultGeneric-array_model_Demo": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误码 200:成功 其他:失败",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Demo"
                    }
                },
                "message": {
                    "description": "错误消息",
                    "type": "string"
                },
                "time": {
                    "description": "耗时 单位:毫秒",
                    "type": "integer"
                },
                "traceId": {
                    "description": "跟踪ID",
                    "type": "string"
                }
            }
        },
        "model.ResultGeneric-int": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误码 200:成功 其他:失败",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "integer"
                },
                "message": {
                    "description": "错误消息",
                    "type": "string"
                },
                "time": {
                    "description": "耗时 单位:毫秒",
                    "type": "integer"
                },
                "traceId": {
                    "description": "跟踪ID",
                    "type": "string"
                }
            }
        },
        "model.ResultGeneric-model_Demo": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误码 200:成功 其他:失败",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.Demo"
                        }
                    ]
                },
                "message": {
                    "description": "错误消息",
                    "type": "string"
                },
                "time": {
                    "description": "耗时 单位:毫秒",
                    "type": "integer"
                },
                "traceId": {
                    "description": "跟踪ID",
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
