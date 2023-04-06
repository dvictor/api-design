// Code generated from MyDSL.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // MyDSL

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseMyDSLListener is a complete listener for a parse tree produced by MyDSLParser.
type BaseMyDSLListener struct{}

var _ MyDSLListener = &BaseMyDSLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMyDSLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMyDSLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMyDSLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMyDSLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterApiBody is called when production apiBody is entered.
func (s *BaseMyDSLListener) EnterApiBody(ctx *ApiBodyContext) {}

// ExitApiBody is called when production apiBody is exited.
func (s *BaseMyDSLListener) ExitApiBody(ctx *ApiBodyContext) {}

// EnterImportDirective is called when production importDirective is entered.
func (s *BaseMyDSLListener) EnterImportDirective(ctx *ImportDirectiveContext) {}

// ExitImportDirective is called when production importDirective is exited.
func (s *BaseMyDSLListener) ExitImportDirective(ctx *ImportDirectiveContext) {}

// EnterService is called when production service is entered.
func (s *BaseMyDSLListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BaseMyDSLListener) ExitService(ctx *ServiceContext) {}

// EnterServiceName is called when production serviceName is entered.
func (s *BaseMyDSLListener) EnterServiceName(ctx *ServiceNameContext) {}

// ExitServiceName is called when production serviceName is exited.
func (s *BaseMyDSLListener) ExitServiceName(ctx *ServiceNameContext) {}

// EnterServiceBody is called when production serviceBody is entered.
func (s *BaseMyDSLListener) EnterServiceBody(ctx *ServiceBodyContext) {}

// ExitServiceBody is called when production serviceBody is exited.
func (s *BaseMyDSLListener) ExitServiceBody(ctx *ServiceBodyContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseMyDSLListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseMyDSLListener) ExitMethod(ctx *MethodContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BaseMyDSLListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BaseMyDSLListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterResponse is called when production response is entered.
func (s *BaseMyDSLListener) EnterResponse(ctx *ResponseContext) {}

// ExitResponse is called when production response is exited.
func (s *BaseMyDSLListener) ExitResponse(ctx *ResponseContext) {}

// EnterPayload is called when production payload is entered.
func (s *BaseMyDSLListener) EnterPayload(ctx *PayloadContext) {}

// ExitPayload is called when production payload is exited.
func (s *BaseMyDSLListener) ExitPayload(ctx *PayloadContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseMyDSLListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseMyDSLListener) ExitDescription(ctx *DescriptionContext) {}

// EnterFieldsBody is called when production fieldsBody is entered.
func (s *BaseMyDSLListener) EnterFieldsBody(ctx *FieldsBodyContext) {}

// ExitFieldsBody is called when production fieldsBody is exited.
func (s *BaseMyDSLListener) ExitFieldsBody(ctx *FieldsBodyContext) {}

// EnterField is called when production field is entered.
func (s *BaseMyDSLListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseMyDSLListener) ExitField(ctx *FieldContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseMyDSLListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseMyDSLListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseMyDSLListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseMyDSLListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterFieldBody is called when production fieldBody is entered.
func (s *BaseMyDSLListener) EnterFieldBody(ctx *FieldBodyContext) {}

// ExitFieldBody is called when production fieldBody is exited.
func (s *BaseMyDSLListener) ExitFieldBody(ctx *FieldBodyContext) {}

// EnterRequired is called when production required is entered.
func (s *BaseMyDSLListener) EnterRequired(ctx *RequiredContext) {}

// ExitRequired is called when production required is exited.
func (s *BaseMyDSLListener) ExitRequired(ctx *RequiredContext) {}

// EnterValidation is called when production validation is entered.
func (s *BaseMyDSLListener) EnterValidation(ctx *ValidationContext) {}

// ExitValidation is called when production validation is exited.
func (s *BaseMyDSLListener) ExitValidation(ctx *ValidationContext) {}

// EnterType is called when production type is entered.
func (s *BaseMyDSLListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseMyDSLListener) ExitType(ctx *TypeContext) {}

// EnterTypeBody is called when production typeBody is entered.
func (s *BaseMyDSLListener) EnterTypeBody(ctx *TypeBodyContext) {}

// ExitTypeBody is called when production typeBody is exited.
func (s *BaseMyDSLListener) ExitTypeBody(ctx *TypeBodyContext) {}
