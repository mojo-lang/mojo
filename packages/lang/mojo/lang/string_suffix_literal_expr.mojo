

type StringSuffixLiteralExpr: UnaryExpr @field(argument: is<StringLiteralExpr> || is<StringPrefixLiteralExpr>) {
}
