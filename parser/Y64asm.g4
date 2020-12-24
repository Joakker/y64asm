grammar Y64asm;
import Y64asmLex;

// Parser rules
prog:
    line*
    ;

line:
    stmt NL?
    ;

refer:
    NUM
    | '(' REG ')'
    | ID
    ;

stmt:
    label // A label declaration.
    | instr
    | preproc
    | NL
    ;

preproc:
    PREPROC NUM
    ;

label:
    ID ':'
    ;

instr:
    OP      REG     ',' REG
    | IOP   refer   ',' REG
    | IRMOV refer   ',' REG
    | RRMOV REG     ',' REG
    | MRMOV refer   ',' REG
    | RMMOV REG     ',' refer
    | RRMOV REG     ',' REG
    | JMP   refer
    | CALL refer
    | RET
    | HALT
    ;
