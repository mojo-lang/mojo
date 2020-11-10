package org.mojolang.mojo.parser;

import com.google.protobuf.GeneratedMessageV3;
import org.mojolang.mojo.lang.NominalType;

import java.util.List;

public class TypeIdentifierVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private NominalType.Builder builder_ = null;

    @Override
    public GeneratedMessageV3 visitType_identifier(MojoParser.Type_identifierContext ctx) {

        List<MojoParser.Type_identifier_clauseContext> typeCtxList = ctx.type_identifier_clause();
        for (MojoParser.Type_identifier_clauseContext typeCtx : typeCtxList) {
            NominalType.Builder builder = visitTypeIdentifierClause(typeCtx);

            if (builder_ != null) {
                builder.setEnclosingType(builder_.build());
            }
            builder_ = builder;
        }

        return builder_.build();
    }

    private NominalType.Builder visitTypeIdentifierClause(MojoParser.Type_identifier_clauseContext ctx) {
        NominalType.Builder builder = NominalType.newBuilder();

        builder.setName(ctx.type_name().getText());

        if (ctx.generic_argument_clause() != null) {
            builder.addAllGenericArguments(Util.visitGenericArguments(ctx.generic_argument_clause()));
        }

        return builder;
    }
}
