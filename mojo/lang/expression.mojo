// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/// Expression
///
@discriminator('@type')
@label_format('{}')
type Expression = NullLiteralExpr    @1
                | IntegerLiteralExpr @2
                | FloatLiteralExpr   @3
                | BoolLiteralExpr    @4
                | StringLiteralExpr  @5
                | ObjectLiteralExpr  @6
                | ArrayLiteralExpr   @7
                | MapLiteralExpr     @8
                | RangeLiteralExpr   @9
                | IdentifierExpr     @10
                | NumericSuffixLiteralExpr @11
                | StringPrefixLiteralExpr  @12
                | StringSuffixLiteralExpr  @13
                | StructLiteralExpr  @14
                | ClosureExpr        @15
                | ParenthesizedExpr  @16
                | ImplicitMemberExpr @17
                | WildcardExpr       @18
                | StructConstructionExpr @19
                | TupleExpr @20
                | PrefixUnaryExpr    @30
                | PostfixUnaryExpr   @31
                | FunctionCallExpr   @32
                | ExplicitMemberExpr @33
                | SubscriptExpr @34
                | BinaryExpr @40
                | ConditionalExpr @41
                | TypeCastingExpr @42
                | AssignmentExpr @43
                | ErrorExpr @60


/// Common base for expressions that involve dynamic lookup, which
/// determines at runtime whether a particular method, property, or
/// subscript is available.
type DynamicLookupExpr : Expr

/// UnresolvedMemberExpr - This represents '.foo', an unresolved reference to a
/// member, which is to be resolved with context sensitive type information into
/// bar.foo.  These always have unresolved type.
type UnresolvedMemberExpr

/// An expression node that does not affect the evaluation of its subexpression.
type IdentityExpr : Expr

/// Subscripting expression that applies a keypath to a base.
type KeyPathApplicationExpr : Expr

/// TupleElementExpr - Refer to an element of a tuple,
/// e.g. "(1,field:2).field".
type TupleElementExpr : Expr 

/// ImplicitConversionExpr - An abstract class for expressions which
/// implicitly convert the value of an expression in some way.
//class ImplicitConversionExpr : public Expr

type KeyPathExpr

/// CaptureListExpr - This expression represents the capture list on an explicit
/// closure.  Because the capture list is evaluated outside of the closure, this
/// CaptureList wraps the ClosureExpr.  The dynamic semantics are that evaluates
/// the variable bindings from the capture list, then evaluates the
/// subexpression (the closure itself) and returns the result.
//class CaptureListExpr : public Expr

/// DynamicTypeExpr - "type(of: base)" - Produces a metatype value.
///
/// The metatype value comes from evaluating an expression then retrieving the
/// metatype of the result.
//class DynamicTypeExpr : public Expr 

/// An expression referring to an opaque object of a fixed type.
///
/// Opaque value expressions occur when a particular value within the AST
/// needs to be re-used without being re-evaluated or for a value that is
/// a placeholder. OpaqueValueExpr nodes are introduced by some other AST
/// node (say, a \c DynamicMemberRefExpr) and can only be used within the
/// subexpressions of that AST node.
//class OpaqueValueExpr

/// DotSyntaxCallExpr - Refer to a method of a type, e.g. P.x.  'x'
/// is modeled as a DeclRefExpr or OverloadSetRefExpr on the method.
//class DotSyntaxCallExpr : public SelfApplyExpr

/// ConstructorRefCallExpr - Refer to a constructor for a type P.  The
/// actual reference to function which returns the constructor is modeled
/// as a DeclRefExpr.
//class ConstructorRefCallExpr : public SelfApplyExpr

/// DotSyntaxBaseIgnoredExpr - When a.b resolves to something that does not need
/// the actual value of the base (e.g. when applied to a metatype, module, or
/// the base of a 'static' function) this expression node is created.  The
/// semantics are that its base is evaluated and discarded, then 'b' is
/// evaluated and returned as the result of the expression.
// class DotSyntaxBaseIgnoredExpr : public Expr

/// \brief A pattern production that has been parsed but hasn't been resolved
/// into a complete pattern. Name binding converts these into standalone pattern
/// nodes or raises an error if a pattern production appears in an invalid
/// position.
//class UnresolvedPatternExpr : public Expr


/// Represents the unusual behavior of a . in a \ keypath expression, such as
/// \.[0] and \Foo.?.
//class KeyPathDotExpr : public Expr

