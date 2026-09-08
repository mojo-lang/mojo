package org.mojolang.mojo.parser.syntax;

import com.google.protobuf.GeneratedMessageV3;
import org.mojolang.mojo.lang.GenericParameter;
import org.mojolang.mojo.lang.NominalType;

public class GenericParameterVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private GenericParameter.Builder builder_ = GenericParameter.newBuilder();

    @Override
    public GeneratedMessageV3 visitGeneric_parameter(MojoParser.Generic_parameterContext ctx) {
        builder_.setName(ctx.type_name().Type_name().getText());

        TypeIdentifierVisitor visitor = new TypeIdentifierVisitor();

        if (ctx.type_identifier() != null) {
            NominalType type = (NominalType) visitor.visitType_identifier(ctx.type_identifier());
            builder_.setConstraint(type);
        }

        return builder_.build();
    }
}
