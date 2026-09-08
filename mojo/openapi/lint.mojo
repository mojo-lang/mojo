/// OpenAPI lint result
type LintResult {
    valid: Bool @1 // is valid
    operations: [LintOperation] @2 // operations
}

type LintOperation {
    path: String @1 //< url path
    method: String @2 // http request method, "GET", "POST", "PUT", "DELETE"...
    valid: Bool @3 // is valid
    description: String @4
    how_to_fix: String @5
    line: Int32 @6
}