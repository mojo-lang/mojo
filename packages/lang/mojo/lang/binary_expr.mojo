
/// BinaryExpr - Infix binary expressions like 'x+y'.  The argument is always
/// an implicit tuple expression of the type expected by the function.
type BinaryExpr : ApplyExpr {
    left_argument: Expression @21
    right_argument: Expression @22
}