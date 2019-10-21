P -> LDE
LDE -> LDE DE
LDE -> DE
DE -> DF
DE -> DT
T -> 'integer'
T -> 'char'
T -> 'boolean'
T -> 'string'
T -> IDU
DT -> 'type' IDD '=' 'array' '[' NUM ']' 'of' T
DT -> 'type' IDD '=' 'struct' NB '{' DC '}'
DT -> 'type' IDD '=' T
DC -> DC ';' LI ':' T
DC -> LI ':' T
DF -> 'function' IDD NF '(' LP ')' ':' T MF B
LP -> LP ',' IDD ':' T
LP -> IDD ':' T
B -> '{' LDV LS '}'
LDV -> LDV DV
LDV -> DV
LS -> LS S
LS -> S
DV -> 'var' LI ':' T ';'
LI -> LI ',' IDD
LI -> IDD
S -> M
S -> U
U -> 'if' '(' E ')' S
U -> 'if' '(' E ')' M 'else' U
M -> 'if' '(' E ')' M 'else' M
M -> 'while' '(' E ')' M
M -> 'do' M 'while' '(' E ')' ';'
M -> NB B
M -> LV '=' E ';'
M -> 'break' ';'
M -> 'continue' ';'
E -> E '&&' L
E -> E '||' L
E -> L
L -> L '<' R
L -> L '>' R
L -> L '<=' R
L -> L '>=' R
L -> L '==' R
L -> L '!=' R
L -> R
R -> R '+' Y
R -> R '-' Y
R -> Y
Y -> Y '*' F
Y -> Y '/' F
Y -> F
F -> LV
F -> '++' LV
F -> '--' LV
F -> LV '++'
F -> LV '--'
F -> '(' E ')'
F -> IDU MC '(' LE ')'
F -> '-' F
F -> '!' F
F -> TRUE
F -> FALSE
F -> CHR
F -> STR
F -> NUM
LE -> LE ';' E
LE -> E
LV -> LV '.' IDU
LV -> LV '[' E ']'
LV -> IDU
NB -> ''
NF -> ''
MF -> ''
MT -> ''
ME -> ''
MW -> ''
MC -> ''
IDD -> ID
IDU -> ID
ID -> Id
TRUE -> 'true'
FALSE -> 'false'
CHR -> c
STR -> s
NUM -> n

