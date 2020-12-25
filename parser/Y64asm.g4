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

instr0:
    RET
    | HALT
    | NOP
    ;

instr1:
    JMP   refer
    | CALL refer
    ;

instr2:
    OP      REG     ',' REG
    | IOP   refer   ',' REG
    | IRMOV refer   ',' REG
    | RRMOV REG     ',' REG
    | MRMOV refer   ',' REG
    | RMMOV REG     ',' refer
    | RRMOV REG     ',' REG
    ;

instr:
    instr0
    | instr1
    | instr2
    ;
