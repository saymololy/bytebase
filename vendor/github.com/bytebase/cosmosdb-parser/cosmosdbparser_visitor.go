// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CosmosDBParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CosmosDBParser.
type CosmosDBParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CosmosDBParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#select_clause.
	VisitSelect_clause(ctx *Select_clauseContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#select_specification.
	VisitSelect_specification(ctx *Select_specificationContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#from_specification.
	VisitFrom_specification(ctx *From_specificationContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#from_source.
	VisitFrom_source(ctx *From_sourceContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#container_expression.
	VisitContainer_expression(ctx *Container_expressionContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#container_name.
	VisitContainer_name(ctx *Container_nameContext) interface{}
}
