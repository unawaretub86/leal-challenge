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
        "/commerce/": {
            "post": {
                "description": "Crea un comercio en el sistema utilizando la información proporcionada.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Commerce"
                ],
                "summary": "Crear un nuevo comercio",
                "parameters": [
                    {
                        "description": "Información del comercio a crear",
                        "name": "commerce",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCommerceDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "comercio creado exitosamente",
                        "schema": {
                            "$ref": "#/definitions/domain.Commerce"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/commerce/branch": {
            "post": {
                "description": "Crea una sucursal en el sistema utilizando la información proporcionada.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Branch"
                ],
                "summary": "Crear una nueva sucursal",
                "parameters": [
                    {
                        "description": "Información de la sucursal a crear",
                        "name": "branch",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBranchDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "sucursal creada exitosamente",
                        "schema": {
                            "$ref": "#/definitions/domain.Branch"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/commerce/branch/campaign/:id": {
            "get": {
                "description": "obtener campañas de una sucursal en el sistema.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Campaign"
                ],
                "summary": "Obtener campañas de una sucursal",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Campaign"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/commerce/campaign": {
            "post": {
                "description": "Crea una campaña en el sistema utilizando la información proporcionada.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Campaign"
                ],
                "summary": "Crear una nueva campaña",
                "parameters": [
                    {
                        "description": "Información de la campaña a crear",
                        "name": "campaign",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCampaignDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "campaña creada exitosamente",
                        "schema": {
                            "$ref": "#/definitions/domain.Campaign"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/commerce/campaign/:id": {
            "get": {
                "description": "obtener campañas de un comercio en el sistema.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Campaign"
                ],
                "summary": "Obtener campañas de un comercio",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Campaign"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/": {
            "post": {
                "description": "Crea un usuario en el sistema utilizando la información proporcionada.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Crear un nuevo usuario",
                "parameters": [
                    {
                        "description": "Información del usuario a crear",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Usuario creado exitosamente",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/redeem": {
            "post": {
                "description": "Registrar una redencion de puntos o cashback.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Redeem"
                ],
                "summary": "Registrar una redencion de puntos o cashback",
                "parameters": [
                    {
                        "description": "Información para redimir",
                        "name": "redeem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RedeemReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Compra registrada exitosamente",
                        "schema": {
                            "$ref": "#/definitions/domain.Redeem"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/register-purchase": {
            "post": {
                "description": "Registra una nueva compra en el sistema utilizando la información proporcionada.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Purchase"
                ],
                "summary": "Registrar una nueva compra",
                "parameters": [
                    {
                        "description": "Información de la compra a registrada",
                        "name": "purchase",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterPurchaseReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Compra registrada exitosamente",
                        "schema": {
                            "$ref": "#/definitions/domain.Purchase"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Award": {
            "type": "object",
            "properties": {
                "campaignID": {
                    "type": "integer"
                },
                "commerceID": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "pointsCost": {
                    "type": "integer"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "domain.Branch": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "commerceID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "domain.Campaign": {
            "type": "object",
            "properties": {
                "award": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Award"
                    }
                },
                "bonusFactor": {
                    "type": "number"
                },
                "bonusType": {
                    "type": "string"
                },
                "branchID": {
                    "type": "integer"
                },
                "commerceID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isForAll": {
                    "type": "boolean"
                },
                "minPurchase": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "domain.Commerce": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "domain.Purchase": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "branchID": {
                    "type": "integer"
                },
                "campaignID": {
                    "type": "integer"
                },
                "commerceID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "earnedCashBack": {
                    "type": "integer"
                },
                "earnedPoints": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "redeemPoints": {
                    "type": "boolean"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "domain.Redeem": {
            "type": "object",
            "properties": {
                "awardID": {
                    "type": "integer"
                },
                "branchID": {
                    "type": "integer"
                },
                "campaignID": {
                    "type": "integer"
                },
                "cashBack": {
                    "type": "integer"
                },
                "commerceID": {
                    "type": "integer"
                },
                "isPointsRedeem": {
                    "type": "boolean"
                },
                "purchaseID": {
                    "type": "integer"
                },
                "redeemedAwardID": {
                    "type": "integer"
                },
                "redeemedCashBack": {
                    "type": "number"
                },
                "redeemedPoints": {
                    "type": "integer"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "properties": {
                "cashback": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "points": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "dto.CreateAwardsCampaign": {
            "type": "object",
            "properties": {
                "campaignID": {
                    "type": "integer"
                },
                "commerce_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "points_cost": {
                    "type": "integer"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "dto.CreateBranchDTO": {
            "type": "object",
            "required": [
                "active_date",
                "address",
                "commerce_id",
                "name"
            ],
            "properties": {
                "active_date": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "commerce_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.CreateCampaignDTO": {
            "type": "object",
            "required": [
                "bonus_factor",
                "bonus_type",
                "commerce_id",
                "end_date",
                "min_purchase",
                "name",
                "start_date"
            ],
            "properties": {
                "awards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateAwardsCampaign"
                    }
                },
                "bonus_factor": {
                    "type": "number"
                },
                "bonus_type": {
                    "type": "string",
                    "enum": [
                        "double",
                        "percentage"
                    ]
                },
                "branch_id": {
                    "type": "integer"
                },
                "commerce_id": {
                    "type": "integer"
                },
                "end_date": {
                    "type": "string"
                },
                "is_for_all": {
                    "type": "boolean"
                },
                "min_purchase": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "dto.CreateCommerceDTO": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.CreateUserReq": {
            "type": "object",
            "required": [
                "email",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.RedeemReq": {
            "type": "object",
            "required": [
                "branch_id",
                "campaign_id",
                "commerce_id",
                "purchase_id",
                "user_id"
            ],
            "properties": {
                "award_id": {
                    "type": "integer"
                },
                "branch_id": {
                    "type": "integer"
                },
                "campaign_id": {
                    "type": "integer"
                },
                "cash_back": {
                    "type": "integer"
                },
                "commerce_id": {
                    "type": "integer"
                },
                "is_points_redeem": {
                    "type": "boolean"
                },
                "purchase_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.RegisterPurchaseReq": {
            "type": "object",
            "required": [
                "amount",
                "branch_id",
                "campaign_id",
                "commerce_id",
                "redeem_points",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "branch_id": {
                    "type": "integer"
                },
                "campaign_id": {
                    "type": "integer"
                },
                "commerce_id": {
                    "type": "integer"
                },
                "redeem_points": {
                    "type": "boolean"
                },
                "user_id": {
                    "type": "integer"
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
