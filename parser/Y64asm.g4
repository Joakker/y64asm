grammar Y64asm;
import Y64asmLex;

// Parser rules
prog:
    line*
    ;

line:
    stmt NL?
    ;

stmt:
    ID ':'  // A label declaration.
    | '\t' instr
    ;

label:
    ID ':'
    ;

instr:
    OP      REG ',' REG
    | IOP   NUM ',' REG
    ;

// Lexer rules
