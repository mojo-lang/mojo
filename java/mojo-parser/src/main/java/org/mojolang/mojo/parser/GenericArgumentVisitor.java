package org.mojolang.mojo.parser;

import com.google.protobuf.GeneratedMessageV3;
import org.mojolang.mojo.lang.GenericArgument;
import org.mojolang.mojo.lang.NominalType;

public class GenericArgumentVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private GenericArgument.Builder builder_ = GenericArgument.newBuilder();

    @Override
    public GeneratedMessageV3 visitGeneric_argument(MojoParser.Generic_argumentContext ctx) {
        NominalTypeVisitor visitor = new NominalTypeVisitor();
        NominalType type = (NominalType)visitor.visitTypeAndAttributes(ctx.type_(), ctx.attributes());

        builder_.setName(type.getName());
        builder_.addAllGenericArguments(type.getGenericArgumentsList());
        builder_.addAllAttributes(type.getAttributesList());

        return builder_.build();
    }
}
