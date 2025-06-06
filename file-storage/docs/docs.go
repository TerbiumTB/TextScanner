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
        "/download/{id}": {
            "get": {
                "description": "Загружает файл по его уникальному идентификатору в формате multipart/form-data",
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Загрузка"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Уникальный идентификатор файла",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.DownloadResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/record": {
            "get": {
                "description": "Получает метаданные всех файлов",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Инфо"
                ],
                "responses": {
                    "200": {
                        "description": "Список метаданных файлов",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.FileRecord"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/record/{id}": {
            "get": {
                "description": "Получает метаданные файла по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Инфо"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Уникальный идентификатор файла",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Метаданные файла",
                        "schema": {
                            "$ref": "#/definitions/models.FileRecord"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/upload/{filename}": {
            "post": {
                "description": "Загружает файл на сервер и возвращает его уникальный идентификатор",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Загрузка"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Имя файла",
                        "name": "filename",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Файл для загрузки",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.UploadResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.DownloadResponse": {
            "type": "object",
            "properties": {
                "record": {
                    "$ref": "#/definitions/models.FileRecord"
                }
            }
        },
        "handlers.UploadResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "models.FileRecord": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "File Storage API",
	Description:      "API для хранения и управления файлами",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	//LeftDelim:        "{{",
	//RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
