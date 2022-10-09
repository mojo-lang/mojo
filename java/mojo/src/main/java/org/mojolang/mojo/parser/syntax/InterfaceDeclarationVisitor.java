package org.mojolang.mojo.parser.syntax;

import com.google.protobuf.GeneratedMessageV3;
import org.mojolang.mojo.lang.*;

import java.util.List;

public class InterfaceDeclarationVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private InterfaceDecl.Builder builder_ = InterfaceDecl.newBuilder();
    private InterfaceType.Builder typeBuilder_ = InterfaceType.newBuilder();

    @Override
    public GeneratedMessageV3 visitInterface_declaration(MojoParser.Interface_declarationContext ctx) {
        visitDocument(ctx.document());
        visitAttributes(ctx.attributes());

        visitInterface_name(ctx.interface_name());

        //visitGeneric_parameter_clause(ctx.generic_parameter_clause());
        visitType_inheritance_clause(ctx.type_inheritance_clause());
        visitInterface_body(ctx.interface_body());

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
    public GeneratedMessageV3 visitInterface_name(MojoParser.Interface_nameContext ctx) {
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
    public GeneratedMessageV3 visitInterface_body(MojoParser.Interface_bodyContext ctx) {
        List<MojoParser.Interface_memberContext> members = ctx.interface_member();
        for (MojoParser.Interface_memberContext member : members) {
            MojoParser.Interface_method_declarationContext methodCtx = member.interface_method_declaration();
            if (methodCtx != null) {
                FuncDeclarationVisitor visitor = new FuncDeclarationVisitor();
                FuncDecl func = (FuncDecl) visitor.visitInterface_method_declaration(methodCtx);
                typeBuilder_.addMethods(func);
            }
        }
        return null;
    }
}
