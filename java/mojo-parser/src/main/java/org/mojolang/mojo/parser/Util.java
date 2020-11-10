package org.mojolang.mojo.parser;

import org.antlr.v4.runtime.Token;
import org.mojolang.mojo.lang.*;

import java.util.ArrayList;
import java.util.List;

public class Util {

    public static Position getPosition(Token token) {
        Position.Builder builder = Position.newBuilder();
        builder.setLine(token.getLine());
        builder.setColumn(token.getCharPositionInLine());
        return builder.build();
    }

    public static Document visitDocument(MojoParser.DocumentContext ctx) {
        DocumentVisitor visitor = new DocumentVisitor();
        return (Document) visitor.visitDocument(ctx);
    }

    public static List<Attribute> visitAttributes(MojoParser.AttributesContext ctx) {
        List<MojoParser.AttributeContext> attributes = ctx.attribute();
        List<Attribute> attrs = new ArrayList<>();
        for (MojoParser.AttributeContext attribute : attributes) {
            AttributeVisitor visitor = new AttributeVisitor();
            Attribute attr = (Attribute) visitor.visitAttribute(attribute);
            attrs.add(attr);
        }
        return attrs;
    }

    public static List<GenericParameter> visitGenericParameters(MojoParser.Generic_parameter_clauseContext ctx) {
        List<MojoParser.Generic_parameterContext> parameters = ctx.generic_parameter_list().generic_parameter();
        List<GenericParameter> params = new ArrayList<>();
        for (MojoParser.Generic_parameterContext parameter : parameters) {
            GenericParameterVisitor visitor = new GenericParameterVisitor();
            GenericParameter param = (GenericParameter) visitor.visitGeneric_parameter(parameter);
            params.add(param);
        }
        return params;
    }

    public static List<GenericArgument> visitGenericArguments(MojoParser.Generic_argument_clauseContext ctx) {
        List<MojoParser.Generic_argumentContext> argumentCtxes = ctx.generic_argument_list().generic_argument();
        List<GenericArgument> arguments = new ArrayList<>();
        for (MojoParser.Generic_argumentContext argumentCtx : argumentCtxes) {
            GenericArgumentVisitor visitor = new GenericArgumentVisitor();
            GenericArgument argument = (GenericArgument) visitor.visitGeneric_argument(argumentCtx);
            arguments.add(argument);
        }
        return arguments;
    }

    public static List<NominalType> visitTypeInheritances(MojoParser.Type_inheritance_clauseContext ctx) {
        List<MojoParser.Type_identifierContext> identifiers = ctx.type_inheritance_list().type_identifier();
        List<NominalType> types = new ArrayList<>();
        for (MojoParser.Type_identifierContext identifier : identifiers) {
            TypeIdentifierVisitor visitor = new TypeIdentifierVisitor();
            NominalType type = (NominalType) visitor.visitType_identifier(identifier);
            types.add(type);
        }
        return types;
    }

    public static List<Expression> visitExpressionList(MojoParser.Expression_listContext ctx) {
        List<MojoParser.ExpressionContext> expressionCtxes = ctx.expression();
        List<Expression> expressions = new ArrayList<>();
        for (MojoParser.ExpressionContext expressionCtx : expressionCtxes) {
            ExpressionVisitor visitor = new ExpressionVisitor();
            Expression expression = (Expression) visitor.visitExpression(expressionCtx);
            expressions.add(expression);
        }
        return expressions;
    }
}
