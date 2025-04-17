parser grammar CosmosDBParser;

options {
	tokenVocab = CosmosDBLexer;
}

root: select EOF;

select: select_clause from_clause;

select_clause: SELECT_SYMBOL select_specification;

select_specification: MULTIPLY_OPERATOR;

from_clause: FROM_SYMBOL from_specification;

from_specification: from_source;

from_source: container_expression;

container_expression: container_name;

container_name: IDENTIFIER;