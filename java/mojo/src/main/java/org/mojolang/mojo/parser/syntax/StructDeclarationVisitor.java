package org.mojolang.mojo.parser.syntax;

import com.google.protobuf.GeneratedMessageV3;
import org.mojolang.mojo.lang.EnumDecl;
import org.mojolang.mojo.lang.StructDecl;
import org.mojolang.mojo.lang.StructType;
import org.mojolang.mojo.lang.ValueDecl;

import java.util.List;

public class StructDeclarationVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private StructDecl.Builder builder_ = StructDecl.newBuilder();
    private StructType.Builder typeBuilder_ = StructType.newBuilder();

    @Override
    public GeneratedMessageV3 visitStruct_declaration(MojoParser.Struct_declarationContext ctx) {
        visitDocument(ctx.document());
        visitAttributes(ctx.attributes());
        visitStruct_name(ctx.struct_name());
        visitGeneric_parameter_clause(ctx.generic_parameter_clause());
        visitType_inheritance_clause(ctx.type_inheritance_clause());
        visitStruct_body(ctx.struct_body());

        builder_.setType(typeBuilder_);
        return builder_.build();
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
    public GeneratedMessageV3 visitStruct_name(MojoParser.Struct_nameContext ctx) {
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
        typeBuilder_.addAllInherits(Util.visitTypeInheritances(ctx));
        }
        return null;
    }

    @Override
    public GeneratedMessageV3 visitStruct_body(MojoParser.Struct_bodyContext ctx) {
        if (ctx == null) {
            return null;
        }

        List<MojoParser.Struct_memberContext> members = ctx.struct_member();
        for (MojoParser.Struct_memberContext member : members) {
            MojoParser.Struct_declarationContext structCtx = member.struct_declaration();
            if (structCtx != null) {
                StructDeclarationVisitor visitor = new StructDeclarationVisitor();
                StructDecl struct = (StructDecl)visitor.visitStruct_declaration(structCtx);
                builder_.addStructTypeDecls(struct);
            }

            MojoParser.Enum_declarationContext enumCtx = member.enum_declaration();
            if (enumCtx != null) {
                EnumDeclarationVisitor visitor = new EnumDeclarationVisitor();
                EnumDecl declaration = (EnumDecl)visitor.visitEnum_declaration(enumCtx);
                builder_.addEnumTypeDecls(declaration);
            }

            MojoParser.Struct_member_declarationContext memberCtx = member.struct_member_declaration();
            if (memberCtx != null) {
                ValueDeclarationVisitor visitor = new ValueDeclarationVisitor();
                ValueDecl value = (ValueDecl)visitor.visitStruct_member_declaration(memberCtx);
                typeBuilder_.addFields(value);
            }
        }
        return null;
    }
}
