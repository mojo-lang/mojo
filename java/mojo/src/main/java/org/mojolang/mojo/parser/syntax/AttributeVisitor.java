package org.mojolang.mojo.parser.syntax;

import com.google.protobuf.GeneratedMessageV3;
import org.antlr.v4.runtime.tree.TerminalNode;
import org.mojolang.mojo.lang.Attribute;
import org.mojolang.mojo.lang.Expression;

import java.util.List;

public class AttributeVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private Attribute.Builder builder_ = Attribute.newBuilder();

    @Override
    public GeneratedMessageV3 visitAttribute(MojoParser.AttributeContext ctx) {
        visitAttribute_name(ctx.attribute_name());

        if (ctx.attribute_argument_clause() != null) {
            MojoParser.Expression_listContext context = ctx.attribute_argument_clause().expression_list();
            if (context != null) {
                List<Expression> expressions = Util.visitExpressionList(context);
                builder_.addAllExpressions(expressions);
            }
        }

        return builder_.build();
    }

    @Override
    public GeneratedMessageV3 visitAttribute_name(MojoParser.Attribute_nameContext ctx) {
        TerminalNode identifier = ctx.Identifier();
        if (identifier != null) {
            builder_.setName(identifier.getText());
            builder_.setStartPosition(Util.getPosition(identifier.getSymbol()));
            //builder_.setEndPosition(Util.getPosition(identifier.))
        } else {
            TerminalNode number = ctx.Decimal_literal();
            builder_.setName(number.getText());
            builder_.setStartPosition(Util.getPosition(number.getSymbol()));
        }

        return null;
    }
}
