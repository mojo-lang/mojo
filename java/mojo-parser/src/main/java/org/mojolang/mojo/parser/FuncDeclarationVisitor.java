package org.mojolang.mojo.parser;

import com.google.protobuf.GeneratedMessageV3;
import org.mojolang.mojo.lang.FuncDecl;
import org.mojolang.mojo.lang.FuncType;
import org.mojolang.mojo.lang.NominalType;
import org.mojolang.mojo.lang.ValueDecl;

public class FuncDeclarationVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private FuncDecl.Builder builder_ = FuncDecl.newBuilder();
    private FuncType.Builder typeBuilder_= FuncType.newBuilder();

    @Override
    public GeneratedMessageV3 visitFunction_declaration(MojoParser.Function_declarationContext ctx) {
        setFunctionName(ctx.function_name());

        if (ctx.generic_parameter_clause() != null) {
            builder_.addAllGenericParameters(Util.visitGenericParameters(ctx.generic_parameter_clause()));
        }

        visitFunction_signature(ctx.function_signature());

        builder_.setType(typeBuilder_);
        return builder_.build();
    }

    @Override
    public GeneratedMessageV3 visitInterface_method_declaration(MojoParser.Interface_method_declarationContext ctx) {
        setFunctionName(ctx.function_name());

        if (ctx.generic_parameter_clause() != null) {
            builder_.addAllGenericParameters(Util.visitGenericParameters(ctx.generic_parameter_clause()));
        }
        visitFunction_signature(ctx.function_signature());

        builder_.setType(typeBuilder_);
        return builder_.build();
    }

    private void setFunctionName(MojoParser.Function_nameContext ctx) {
        builder_.setName(ctx.declaration_identifier().getText());
        //builder_.setStartPosition(Util.getPosition(ctx.declaration_identifier()));
    }

    @Override
    public GeneratedMessageV3 visitFunction_signature(MojoParser.Function_signatureContext ctx) {
        MojoParser.Function_parameter_listContext parametersCtx = ctx.function_parameter_clause().function_parameter_list();
        if (parametersCtx != null) {
            for (MojoParser.Function_parameterContext parameterCtx : parametersCtx.function_parameter()) {
                ValueDeclarationVisitor visitor = new ValueDeclarationVisitor();
                ValueDecl value = (ValueDecl) visitor.visitFunction_parameter(parameterCtx);

                typeBuilder_.addParameters(value);
            }
        }

        visitFunction_result(ctx.function_result());
        return null;
    }

    @Override
    public GeneratedMessageV3 visitFunction_result(MojoParser.Function_resultContext ctx) {
        if (ctx != null) {
            NominalTypeVisitor visitor = new NominalTypeVisitor();
            NominalType nominalType = (NominalType)visitor.visitTypeAndAttributes(ctx.type_(), ctx.attributes());
            typeBuilder_.setReturnType(nominalType);
        }

        return null;
    }

    @Override
    public GeneratedMessageV3 visitFunction_body(MojoParser.Function_bodyContext ctx) {
        return null;
    }
}
