grammar Y64asmLex;

PREPROC:
    '.'
    (
        'quad'
        | 'long'
        | 'pos'
        | 'align'
    )
    ;

OP:
    (
        'add'
        | 'and'
        | 'xor'
        | 'sub'
    ) 'q'?
    ;

IOP:
    'i' OP
    ;

REG:
    '%r'
    (
        [8-9]
        | '1' [0-4]
        | ('s'|'d') 'i'
        | ('s'|'b') 'p'
        | [a-d] 'x'
    )
    ;

MOV:
    'mov' 'q'?
    ;

IRMOV:
    'ir' MOV
    ;

RRMOV:
    'rr' MOV
    ;

MRMOV:
    'mr' MOV
    ;

RMMOV:
    'rm' MOV
    ;

RRMOV:
    'rr' MOV
    ;

JMP:
    'j'
    (
        'ne'
        | 'eq'
        | 'mp'
    )
    ;

RET:
    'ret'
    ;

CALL:
    'call'
    ;

HALT:
    'halt'
    ;

NUM:
    '$'? [0-9]+
    | '0x' [0-9a-fA-F]+
    ;

ID:
    [a-zA-Z]+[0-9]*
    ;

NL:
    '\r'? '\n'
    ;

WS:
    [ \t] -> skip;

LINE_COMMENT:
    '#' ~[\r\n]* -> skip;
