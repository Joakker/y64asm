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

NUM:
    '$'? [0-9]+
    | '0x' [0-9a-fA-F]+
    ;

ID:
    [a-z]+
    ;

NL:
    '\r'? '\n'
    ;

WS:
    ' ' -> skip;
