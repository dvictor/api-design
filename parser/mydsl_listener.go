// Code generated from MyDSL.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // MyDSL

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// MyDSLListener is a complete listener for a parse tree produced by MyDSLParser.
type MyDSLListener interface {
	antlr.ParseTreeListener

	// EnterApiBody is called when entering the apiBody production.
	EnterApiBody(c *ApiBodyContext)

	// EnterImportDirective is called when entering the importDirective production.
	EnterImportDirective(c *ImportDirectiveContext)

	// EnterService is called when entering the service production.
	EnterService(c *ServiceContext)

	// EnterServiceName is called when entering the serviceName production.
	EnterServiceName(c *ServiceNameContext)

	// EnterServiceBody is called when entering the serviceBody production.
	EnterServiceBody(c *ServiceBodyContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterMethodBody is called when entering the methodBody production.
	EnterMethodBody(c *MethodBodyContext)

	// EnterResponse is called when entering the response production.
	EnterResponse(c *ResponseContext)

	// EnterPayload is called when entering the payload production.
	EnterPayload(c *PayloadContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterFieldsBody is called when entering the fieldsBody production.
	EnterFieldsBody(c *FieldsBodyContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterFieldBody is called when entering the fieldBody production.
	EnterFieldBody(c *FieldBodyContext)

	// EnterValidation is called when entering the validation production.
	EnterValidation(c *ValidationContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterTypeBody is called when entering the typeBody production.
	EnterTypeBody(c *TypeBodyContext)

	// ExitApiBody is called when exiting the apiBody production.
	ExitApiBody(c *ApiBodyContext)

	// ExitImportDirective is called when exiting the importDirective production.
	ExitImportDirective(c *ImportDirectiveContext)

	// ExitService is called when exiting the service production.
	ExitService(c *ServiceContext)

	// ExitServiceName is called when exiting the serviceName production.
	ExitServiceName(c *ServiceNameContext)

	// ExitServiceBody is called when exiting the serviceBody production.
	ExitServiceBody(c *ServiceBodyContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitMethodBody is called when exiting the methodBody production.
	ExitMethodBody(c *MethodBodyContext)

	// ExitResponse is called when exiting the response production.
	ExitResponse(c *ResponseContext)

	// ExitPayload is called when exiting the payload production.
	ExitPayload(c *PayloadContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitFieldsBody is called when exiting the fieldsBody production.
	ExitFieldsBody(c *FieldsBodyContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitFieldBody is called when exiting the fieldBody production.
	ExitFieldBody(c *FieldBodyContext)

	// ExitValidation is called when exiting the validation production.
	ExitValidation(c *ValidationContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitTypeBody is called when exiting the typeBody production.
	ExitTypeBody(c *TypeBodyContext)
}
