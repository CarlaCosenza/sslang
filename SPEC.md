# language specification

This is a translation of the Simple Script Language specification.


A program P is composed by a List of External Declarations:

```
P = LED
LED = LED ED | ED
```

An external declaration ED is either a Function Declaration or a Type Declaration or a Variable Declaration:

```
ED = FD | TD | VD
```

A type T can be either any of the keywords 'integer', 'char', 'boolean' and 'string' or an ID represented by a previously declared type:


```
T = 'integer' | 'char' | 'boolean' | 'string' | ID
```

A type declaration TD can be either a declaration of vector kind or a struct or a synonym:

```
TD = 'type' ID '=' 'array' '[' NUM ']' 'of' T
   | 'type' ID '=' 'struct' '{' DC '}'
   | 'type' ID '=' T
CD = CD ';' IL ':' 'T' | IL ':' 'T'
```

A block B is determined by a key and contains a List of Variables Declarations List (LVD) followed by a List of Statements (LS) or Commands:

```
B = '{' LVD LS '}'
LDV = LVD VD | VD
LS = LS S | S

```


A variables declaration VD is composed by the 'var' keyword followed by a List of Identifiers separated by ',' and followed by ':', the declared variables type T and by a ';' at the end:

```
VD = 'var' LI ':' T ';'
LI = LI ',' ID | ID
```

A statement (S) can be either a selection, loop, a block, an assignment or a control command of a loop ('break' or 'continue'):


```
S = 'if' '(' E ')' S
  | 'if' '(' E ')' S 'else' S
  | 'while' '(' E ')' S
  | 'do' S 'while' '(' E ')' ';'
  | B
  | VL = E ';'
  | 'break' ';'
  | 'continue' ';'
```

An expression E can be composed by logical operators, relational or arithmetic, and can also allow the aggregated type component's access and assignments.


```
E = E '&&' L | E '||' L | L
L = L '<' R | L '>' R | L '<=' R | L'>='R | L '==' R | L '!=' R | R
R = R '+' Y | R '-' Y | Y
Y = Y '*' F | Y '/' F | F
F = VL | '++' VL | '--' VL | VL '++' | VL '--'
  | '(' E ')'
  | ID '(' LE ')'
  | '-' F | '!' F
  | TRUE | FALSE | CHR | STR | NUM
LE = LE ',' E
   | E
VL = VL '.' ID | VL '[' E ']' | ID
ID = id
TRUE = 'true'
FALSE = 'false'
CHR = c
STR = s
NUM = n
```

lexems:

```
letter = [a-zA-Z]
digit = [0-9]
id = letter(letter | digit)+
n = [0-9][0-9]+
c = '"'any'"' // character between "
s = '"'(any)+'"' // multiple or none characters between "
```

