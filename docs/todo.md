### TODOs
- remove the repeat-while
- support nested block comment
- string literal refactor (escape, unicode, expression)
- review the expressions grammar and parser

  语句内的空白应该不能包括多个回车符，即一个语句之间存在空行

  1. term -> ast

2. ast -> c++ code (struct, enum, data)

3. struct (c++) -> json / ymal

4. ast -> openapi