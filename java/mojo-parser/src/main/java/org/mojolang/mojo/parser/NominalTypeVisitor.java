package org.mojolang.mojo.parser;

import com.google.protobuf.GeneratedMessageV3;
import org.mojolang.mojo.lang.GenericArgument;
import org.mojolang.mojo.lang.NominalType;

public class NominalTypeVisitor extends MojoBaseVisitor<com.google.protobuf.GeneratedMessageV3> {
    private NominalType.Builder builder_ = NominalType.newBuilder();

    @Override
    public GeneratedMessageV3 visitType_annotation(MojoParser.Type_annotationContext ctx) {
        visitType_(ctx.type_());

        if (ctx.attributes() != null) {
            builder_.addAllAttributes(Util.visitAttributes(ctx.attributes()));
        }

        return builder_.build();
    }

    public GeneratedMessageV3 visitTypeAndAttributes(MojoParser.Type_Context typeCtx,
                                                     MojoParser.AttributesContext attributesCtx) {
        visitType_(typeCtx);
        if (attributesCtx != null) {
            builder_.addAllAttributes(Util.visitAttributes(attributesCtx));
        }
        return builder_.build();
    }

    @Override
    public GeneratedMessageV3 visitType_(MojoParser.Type_Context ctx) {
        MojoParser.Array_typeContext arrayCtx = ctx.array_type();
        if (arrayCtx != null) {
            NominalTypeVisitor visitor = new NominalTypeVisitor();
            NominalType nominaltype = (NominalType) visitor.visitTypeAndAttributes(arrayCtx.type_(), arrayCtx.attributes());

            GenericArgument.Builder argumentBuilder = GenericArgument.newBuilder();
            argumentBuilder.addAllAttributes(nominaltype.getAttributesList());
            argumentBuilder.setName(nominaltype.getName());

            builder_.setName("Array");
            builder_.addGenericArguments(argumentBuilder);
        }

        MojoParser.Map_typeContext mapCtx = ctx.map_type();
        if (mapCtx != null) {
        }

        MojoParser.Function_typeContext functionCtx = ctx.function_type();


        MojoParser.Type_identifierContext identifierCtx = ctx.type_identifier();
        if (identifierCtx != null) {
            TypeIdentifierVisitor visitor = new TypeIdentifierVisitor();
            NominalType type = (NominalType) visitor.visitType_identifier(identifierCtx);
            builder_.mergeFrom(type);
        }

        MojoParser.Type_Context typeCtx = ctx.type_();
        if (typeCtx != null && ctx.QUESTION() != null) {
        }

        return super.visitType_(ctx);
    }
}
