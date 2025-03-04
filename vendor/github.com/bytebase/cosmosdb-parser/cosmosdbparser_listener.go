// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CosmosDBParser
import "github.com/antlr4-go/antlr/v4"

// CosmosDBParserListener is a complete listener for a parse tree produced by CosmosDBParser.
type CosmosDBParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterSelect_clause is called when entering the select_clause production.
	EnterSelect_clause(c *Select_clauseContext)

	// EnterSelect_specification is called when entering the select_specification production.
	EnterSelect_specification(c *Select_specificationContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterFrom_specification is called when entering the from_specification production.
	EnterFrom_specification(c *From_specificationContext)

	// EnterFrom_source is called when entering the from_source production.
	EnterFrom_source(c *From_sourceContext)

	// EnterContainer_expression is called when entering the container_expression production.
	EnterContainer_expression(c *Container_expressionContext)

	// EnterContainer_name is called when entering the container_name production.
	EnterContainer_name(c *Container_nameContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitSelect_clause is called when exiting the select_clause production.
	ExitSelect_clause(c *Select_clauseContext)

	// ExitSelect_specification is called when exiting the select_specification production.
	ExitSelect_specification(c *Select_specificationContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitFrom_specification is called when exiting the from_specification production.
	ExitFrom_specification(c *From_specificationContext)

	// ExitFrom_source is called when exiting the from_source production.
	ExitFrom_source(c *From_sourceContext)

	// ExitContainer_expression is called when exiting the container_expression production.
	ExitContainer_expression(c *Container_expressionContext)

	// ExitContainer_name is called when exiting the container_name production.
	ExitContainer_name(c *Container_nameContext)
}
