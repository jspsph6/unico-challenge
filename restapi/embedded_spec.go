// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "Routes for fair API",
    "title": "Fair API",
    "contact": {
      "email": "jspsph@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.11"
  },
  "paths": {
    "/fair": {
      "get": {
        "description": "Obtains a fair",
        "produces": [
          "application/json"
        ],
        "operationId": "getFair",
        "parameters": [
          {
            "type": "integer",
            "description": "The fair register field",
            "name": "register",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Fair"
            }
          },
          "400": {
            "description": "One or more fields is invalid",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "put": {
        "description": "Updates a fair",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "updateFair",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Fair"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          },
          "400": {
            "description": "One or more fields is invalid",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "post": {
        "description": "Creates new fair",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "createFair",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Fair"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          },
          "400": {
            "description": "One or more fields is invalid",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "delete": {
        "description": "Deletes a fair",
        "produces": [
          "application/json"
        ],
        "operationId": "deleteFair",
        "parameters": [
          {
            "type": "integer",
            "description": "The fair register field",
            "name": "register",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          },
          "400": {
            "description": "One or more fields is invalid",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/fairs": {
      "get": {
        "description": "Obtains a list of fairs",
        "produces": [
          "application/json"
        ],
        "operationId": "getFairs",
        "parameters": [
          {
            "type": "string",
            "description": "The district name",
            "name": "districtName",
            "in": "query"
          },
          {
            "type": "string",
            "description": "The region5 name",
            "name": "region5",
            "in": "query"
          },
          {
            "type": "string",
            "description": "The fair name",
            "name": "fairName",
            "in": "query"
          },
          {
            "type": "string",
            "description": "The neighborhood name",
            "name": "neighborhood",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Fairs"
            }
          },
          "400": {
            "description": "One or more fields is invalid",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "description": "Checks if server is up",
        "produces": [
          "text/plain"
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string",
              "example": "OK"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Fair": {
      "type": "object",
      "required": [
        "externalId",
        "setcens",
        "fairName",
        "register",
        "codeDistrict",
        "codeSubprefecture",
        "subprefecture",
        "areap",
        "districtName",
        "region5",
        "region8",
        "street",
        "neighborhood",
        "latitude",
        "longitude"
      ],
      "properties": {
        "areap": {
          "type": "integer",
          "format": "int64",
          "example": 1000
        },
        "codeDistrict": {
          "type": "integer",
          "format": "int64",
          "example": 10
        },
        "codeSubprefecture": {
          "type": "integer",
          "format": "int64",
          "example": 100
        },
        "districtName": {
          "description": "District name",
          "type": "string",
          "example": "Centro"
        },
        "externalId": {
          "type": "integer",
          "format": "int64",
          "example": 1
        },
        "fairName": {
          "description": "Fair name",
          "type": "string",
          "example": "Feira 123"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "example": 1
        },
        "latitude": {
          "type": "integer",
          "format": "int64",
          "example": -23558733
        },
        "longitude": {
          "type": "integer",
          "format": "int64",
          "example": -46550164
        },
        "neighborhood": {
          "type": "string",
          "example": "Jardim Florido"
        },
        "number": {
          "type": "string",
          "example": "123a"
        },
        "reference": {
          "type": "string",
          "example": "Em frente ao poste"
        },
        "region5": {
          "description": "Region 05",
          "type": "string",
          "example": "Leste"
        },
        "region8": {
          "description": "Region 08",
          "type": "string",
          "example": "Centro 2"
        },
        "register": {
          "description": "Fair register",
          "type": "string",
          "example": "123-1"
        },
        "setcens": {
          "type": "integer",
          "format": "int64",
          "example": 10
        },
        "street": {
          "type": "string",
          "example": "Rua Aibi"
        },
        "subprefecture": {
          "description": "Subprefecture name",
          "type": "string",
          "example": "Centro"
        }
      }
    },
    "Fairs": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Fair"
      }
    },
    "Response": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string",
          "example": "Some error msg"
        },
        "success": {
          "type": "string",
          "example": "Some success msg"
        }
      }
    }
  },
  "x-components": {}
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "Routes for fair API",
    "title": "Fair API",
    "contact": {
      "email": "jspsph@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.11"
  },
  "paths": {
    "/fair": {
      "get": {
        "description": "Obtains a fair",
        "produces": [
          "application/json"
        ],
        "operationId": "getFair",
        "parameters": [
          {
            "type": "integer",
            "description": "The fair register field",
            "name": "register",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Fair"
            }
          },
          "400": {
            "description": "One or more fields is invalid",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "put": {
        "description": "Updates a fair",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "updateFair",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Fair"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          },
          "400": {
            "description": "One or more fields is invalid",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "post": {
        "description": "Creates new fair",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "createFair",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Fair"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          },
          "400": {
            "description": "One or more fields is invalid",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      },
      "delete": {
        "description": "Deletes a fair",
        "produces": [
          "application/json"
        ],
        "operationId": "deleteFair",
        "parameters": [
          {
            "type": "integer",
            "description": "The fair register field",
            "name": "register",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          },
          "400": {
            "description": "One or more fields is invalid",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/fairs": {
      "get": {
        "description": "Obtains a list of fairs",
        "produces": [
          "application/json"
        ],
        "operationId": "getFairs",
        "parameters": [
          {
            "type": "string",
            "description": "The district name",
            "name": "districtName",
            "in": "query"
          },
          {
            "type": "string",
            "description": "The region5 name",
            "name": "region5",
            "in": "query"
          },
          {
            "type": "string",
            "description": "The fair name",
            "name": "fairName",
            "in": "query"
          },
          {
            "type": "string",
            "description": "The neighborhood name",
            "name": "neighborhood",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Fairs"
            }
          },
          "400": {
            "description": "One or more fields is invalid",
            "schema": {
              "$ref": "#/definitions/Response"
            }
          }
        }
      }
    },
    "/health": {
      "get": {
        "description": "Checks if server is up",
        "produces": [
          "text/plain"
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string",
              "example": "OK"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Fair": {
      "type": "object",
      "required": [
        "externalId",
        "setcens",
        "fairName",
        "register",
        "codeDistrict",
        "codeSubprefecture",
        "subprefecture",
        "areap",
        "districtName",
        "region5",
        "region8",
        "street",
        "neighborhood",
        "latitude",
        "longitude"
      ],
      "properties": {
        "areap": {
          "type": "integer",
          "format": "int64",
          "example": 1000
        },
        "codeDistrict": {
          "type": "integer",
          "format": "int64",
          "example": 10
        },
        "codeSubprefecture": {
          "type": "integer",
          "format": "int64",
          "example": 100
        },
        "districtName": {
          "description": "District name",
          "type": "string",
          "example": "Centro"
        },
        "externalId": {
          "type": "integer",
          "format": "int64",
          "example": 1
        },
        "fairName": {
          "description": "Fair name",
          "type": "string",
          "example": "Feira 123"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "example": 1
        },
        "latitude": {
          "type": "integer",
          "format": "int64",
          "example": -23558733
        },
        "longitude": {
          "type": "integer",
          "format": "int64",
          "example": -46550164
        },
        "neighborhood": {
          "type": "string",
          "example": "Jardim Florido"
        },
        "number": {
          "type": "string",
          "example": "123a"
        },
        "reference": {
          "type": "string",
          "example": "Em frente ao poste"
        },
        "region5": {
          "description": "Region 05",
          "type": "string",
          "example": "Leste"
        },
        "region8": {
          "description": "Region 08",
          "type": "string",
          "example": "Centro 2"
        },
        "register": {
          "description": "Fair register",
          "type": "string",
          "example": "123-1"
        },
        "setcens": {
          "type": "integer",
          "format": "int64",
          "example": 10
        },
        "street": {
          "type": "string",
          "example": "Rua Aibi"
        },
        "subprefecture": {
          "description": "Subprefecture name",
          "type": "string",
          "example": "Centro"
        }
      }
    },
    "Fairs": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Fair"
      }
    },
    "Response": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string",
          "example": "Some error msg"
        },
        "success": {
          "type": "string",
          "example": "Some success msg"
        }
      }
    }
  },
  "x-components": {}
}`))
}
