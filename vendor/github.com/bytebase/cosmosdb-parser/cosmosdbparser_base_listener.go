// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CosmosDBParser
import "github.com/antlr4-go/antlr/v4"

// BaseCosmosDBParserListener is a complete listener for a parse tree produced by CosmosDBParser.
type BaseCosmosDBParserListener struct{}

var _ CosmosDBParserListener = &BaseCosmosDBParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCosmosDBParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCosmosDBParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCosmosDBParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCosmosDBParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseCosmosDBParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseCosmosDBParserListener) ExitRoot(ctx *RootContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseCosmosDBParserListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseCosmosDBParserListener) ExitSelect(ctx *SelectContext) {}

// EnterSelect_clause is called when production select_clause is entered.
func (s *BaseCosmosDBParserListener) EnterSelect_clause(ctx *Select_clauseContext) {}

// ExitSelect_clause is called when production select_clause is exited.
func (s *BaseCosmosDBParserListener) ExitSelect_clause(ctx *Select_clauseContext) {}

// EnterSelect_specification is called when production select_specification is entered.
func (s *BaseCosmosDBParserListener) EnterSelect_specification(ctx *Select_specificationContext) {}

// ExitSelect_specification is called when production select_specification is exited.
func (s *BaseCosmosDBParserListener) ExitSelect_specification(ctx *Select_specificationContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BaseCosmosDBParserListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BaseCosmosDBParserListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterFrom_specification is called when production from_specification is entered.
func (s *BaseCosmosDBParserListener) EnterFrom_specification(ctx *From_specificationContext) {}

// ExitFrom_specification is called when production from_specification is exited.
func (s *BaseCosmosDBParserListener) ExitFrom_specification(ctx *From_specificationContext) {}

// EnterFrom_source is called when production from_source is entered.
func (s *BaseCosmosDBParserListener) EnterFrom_source(ctx *From_sourceContext) {}

// ExitFrom_source is called when production from_source is exited.
func (s *BaseCosmosDBParserListener) ExitFrom_source(ctx *From_sourceContext) {}

// EnterContainer_expression is called when production container_expression is entered.
func (s *BaseCosmosDBParserListener) EnterContainer_expression(ctx *Container_expressionContext) {}

// ExitContainer_expression is called when production container_expression is exited.
func (s *BaseCosmosDBParserListener) ExitContainer_expression(ctx *Container_expressionContext) {}

// EnterContainer_name is called when production container_name is entered.
func (s *BaseCosmosDBParserListener) EnterContainer_name(ctx *Container_nameContext) {}

// ExitContainer_name is called when production container_name is exited.
func (s *BaseCosmosDBParserListener) ExitContainer_name(ctx *Container_nameContext) {}
