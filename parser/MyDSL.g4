grammar MyDSL;

// Parser rules
apiBody: (importDirective | service | type)+ EOF;

importDirective: 'import' STRING_LITERAL;

service: 'service' serviceName '{' serviceBody '}';
serviceName: IDENTIFIER;

serviceBody: (description | method)+;
method: 'method' IDENTIFIER '{' methodBody '}';

methodBody: (description | payload | response)+;

response: 'response' ( ( '{' fieldsBody '}' ) | EQUAL_SIGN IDENTIFIER );

payload: 'payload' ( ( '{' fieldsBody '}' ) | EQUAL_SIGN IDENTIFIER );

description: 'description' EQUAL_SIGN STRING_LITERAL;

fieldsBody: field+;

field: 'field' fieldName fieldType ('{' fieldBody '}')?;
fieldName: IDENTIFIER;
fieldType: IDENTIFIER;
fieldBody: (description | validation | 'required' | 'optional')+;

//required: { !optional() }? 'required';
//optional: { !required() }? 'optional';

validation: 'match' EQUAL_SIGN STRING_LITERAL;

type: 'type' IDENTIFIER '{' typeBody '}';

typeBody: field+;

// Lexer rules
EQUAL_SIGN: '=';
IDENTIFIER: [a-zA-Z] [a-zA-Z0-9]*;

COMMENT: '//' ~ [\r\n]* -> channel(HIDDEN);
MULTILINE_COMMENT: '/*' .*? '*/' -> channel(HIDDEN);

STRING_LITERAL:
    '"' ( '\\' . | ~[\\"] )* '"'
    | '`' ( '\\' . | ~[\\`] )* '`';
WS: [ \t\r\n]+ -> channel(HIDDEN);
