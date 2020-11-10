package org.mojolang.mojo.parser;

import com.google.protobuf.GeneratedMessageV3;
import org.antlr.v4.runtime.tree.TerminalNode;
import org.mojolang.mojo.lang.Expression;
import org.mojolang.mojo.lang.FloatLiteralExpr;
import org.mojolang.mojo.lang.IntegerLiteralExpr;
import org.mojolang.mojo.lang.StringLiteralExpr;

public class ExpressionVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private Expression.Builder builder_ = Expression.newBuilder();

    @Override
    public GeneratedMessageV3 visitExpression(MojoParser.ExpressionContext ctx) {
        MojoParser.Prefix_expressionContext prefixCtx =  ctx.prefix_expression();
        MojoParser.Binary_expressionsContext binaryCtxes = ctx.binary_expressions();

        if (prefixCtx != null) {
            MojoParser.Postfix_expressionContext postfixCtx = prefixCtx.postfix_expression();

            if (postfixCtx instanceof MojoParser.PrimaryContext) {
                MojoParser.Primary_expressionContext primaryCtx = ((MojoParser.PrimaryContext) postfixCtx).primary_expression();
                visitPrimary_expression(primaryCtx);
            }
        }

        if (binaryCtxes != null) {
        }

        return builder_.build();
    }

    @Override
    public GeneratedMessageV3 visitPrimary_expression(MojoParser.Primary_expressionContext ctx) {
        MojoParser.Literal_expressionContext literalExpressionCtx = ctx.literal_expression();
        if (literalExpressionCtx != null) {
            visitLiteral_expression(literalExpressionCtx);
        }

        return super.visitPrimary_expression(ctx);
    }

    @Override
    public GeneratedMessageV3 visitLiteral_expression(MojoParser.Literal_expressionContext ctx) {
        MojoParser.LiteralContext literalCtx = ctx.literal();
        if (literalCtx != null) {
            MojoParser.Numeric_literalContext numericCtx = literalCtx.numeric_literal();
            if (numericCtx != null) {
                MojoParser.Integer_literalContext integerCtx = numericCtx.integer_literal();
                if (integerCtx != null) {
                    IntegerLiteralExpr.Builder builder = IntegerLiteralExpr.newBuilder();

                    long value = Long.parseLong(integerCtx.Decimal_literal().getText());
                    builder.setValue(value);

                    if (numericCtx.negate_prefix_operator() != null) {
                        builder.setIsNegative(!numericCtx.negate_prefix_operator().isEmpty());
                    }

                    builder_.setIntegerLiteralExpr(builder);
                }

                TerminalNode floatCtx = numericCtx.Floating_point_literal();
                if (floatCtx != null) {
                    FloatLiteralExpr.Builder builder = FloatLiteralExpr.newBuilder();

                    double value = Double.parseDouble(floatCtx.getText());
                    builder.setValue(value);

                    if (numericCtx.negate_prefix_operator() != null) {
                        builder.setIsNegative(!numericCtx.negate_prefix_operator().isEmpty());
                    }

                    builder_.setFloatLiteralExpr(builder);
                }
            }

            MojoParser.String_literalContext stringCtx = literalCtx.string_literal();
            if (stringCtx != null) {
                TerminalNode staticCtx = stringCtx.Static_string_literal();
                if (staticCtx != null) {
                    StringLiteralExpr.Builder builder = StringLiteralExpr.newBuilder();
                    builder.setValue(staticCtx.getText());

                    builder_.setStringLiteralExpr(builder);
                }
            }
        }

        return null;
    }
}
