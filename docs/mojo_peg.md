# mojo的PEG扩展

mojo可以作为PEG解析器的引擎，可以自定义小语言，然后将小语言解析成Term树，然后通过处理，可以变换为MOJO的AST或是用户自定义AST

```
type MatchCase
type MatchExpressionClause : Plus<Seq<Seps, MatchCase>>
```

boost.sprit 是使用操作符重载来实现的