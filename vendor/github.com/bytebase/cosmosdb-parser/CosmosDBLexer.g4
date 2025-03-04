lexer grammar CosmosDBLexer;

options {
	caseInsensitive = true;
}

fragment A: [a];
fragment B: [b];
fragment C: [c];
fragment D: [d];
fragment E: [e];
fragment F: [f];
fragment G: [g];
fragment H: [h];
fragment I: [i];
fragment J: [j];
fragment K: [k];
fragment L: [l];
fragment M: [m];
fragment N: [n];
fragment O: [o];
fragment P: [p];
fragment Q: [q];
fragment R: [r];
fragment S: [s];
fragment T: [t];
fragment U: [u];
fragment V: [v];
fragment W: [w];
fragment X: [x];
fragment Y: [y];
fragment Z: [z];

MULTIPLY_OPERATOR: '*';

SELECT_SYMBOL: 'SELECT';
FROM_SYMBOL: 'FROM';

/* Identifiers */
IDENTIFIER: [a-z] [a-z_0-9]*;

// White space handling
WHITESPACE:
	[ \t\f\r\n] -> channel(HIDDEN); // Ignore whitespaces.