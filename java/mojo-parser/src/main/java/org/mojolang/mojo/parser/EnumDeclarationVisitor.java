package org.mojolang.mojo.parser;

import com.google.protobuf.GeneratedMessageV3;
import org.mojolang.mojo.lang.*;

import java.util.List;

public class EnumDeclarationVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    EnumDecl.Builder builder_ = EnumDecl.newBuilder();
    EnumType.Builder typeBuilder_ = EnumType.newBuilder();

    @Override
    public GeneratedMessageV3 visitEnum_declaration(MojoParser.Enum_declarationContext ctx) {
        return super.visitEnum_declaration(ctx);
    }

    @Override
    public GeneratedMessageV3 visitDocument(MojoParser.DocumentContext ctx) {
        if (ctx != null) {
            builder_.setDocument(Util.visitDocument(ctx));
        }
        return null;
    }

    @Override
    public GeneratedMessageV3 visitAttributes(MojoParser.AttributesContext ctx) {
        if (ctx != null) {
            builder_.addAllAttributes(Util.visitAttributes(ctx));
        }
        return null;
    }

    @Override
    public GeneratedMessageV3 visitEnum_name(MojoParser.Enum_nameContext ctx) {
        builder_.setName(ctx.type_name().Type_name().getText());
        return null;
    }

    @Override
    public GeneratedMessageV3 visitGeneric_parameter_clause(MojoParser.Generic_parameter_clauseContext ctx) {
        if (ctx != null) {
            builder_.addAllGenericParameters(Util.visitGenericParameters(ctx));
        }
        return null;
    }

    @Override
    public GeneratedMessageV3 visitType_inheritance_clause(MojoParser.Type_inheritance_clauseContext ctx) {
        if (ctx != null) {
            List<NominalType> types = Util.visitTypeInheritances(ctx);

            if (!types.isEmpty()) {
                typeBuilder_.setUnderlyingType(types.get(0));
            }
        }

        return null;
    }

    @Override
    public GeneratedMessageV3 visitEnum_members(MojoParser.Enum_membersContext ctx) {
        List<MojoParser.Enum_memberContext> membersCtxes = ctx.enum_member();
        for (MojoParser.Enum_memberContext memberCtx : membersCtxes) {
            ValueDeclarationVisitor visitor = new ValueDeclarationVisitor();
            ValueDecl value = (ValueDecl) visitor.visitEnum_member(memberCtx);
            typeBuilder_.addEnumerators(value);
        }
        return null;
    }
}
