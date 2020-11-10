// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/type.proto

package org.mojolang.mojo.lang;

public interface EnumDeclOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.EnumDecl)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   *&#47; position of first character belonging to the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  boolean hasStartPosition();
  /**
   * <pre>
   *&#47; position of first character belonging to the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  org.mojolang.mojo.lang.Position getStartPosition();
  /**
   * <pre>
   *&#47; position of first character belonging to the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position start_position = 1;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getStartPositionOrBuilder();

  /**
   * <pre>
   *&#47; position of first character immediately after the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  boolean hasEndPosition();
  /**
   * <pre>
   *&#47; position of first character immediately after the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  org.mojolang.mojo.lang.Position getEndPosition();
  /**
   * <pre>
   *&#47; position of first character immediately after the Expr
   * </pre>
   *
   * <code>.mojo.lang.Position end_position = 2;</code>
   */
  org.mojolang.mojo.lang.PositionOrBuilder getEndPositionOrBuilder();

  /**
   * <pre>
   *&#47; reserved field no. 3 for string document
   * </pre>
   *
   * <code>.mojo.lang.Document document = 4;</code>
   */
  boolean hasDocument();
  /**
   * <pre>
   *&#47; reserved field no. 3 for string document
   * </pre>
   *
   * <code>.mojo.lang.Document document = 4;</code>
   */
  org.mojolang.mojo.lang.Document getDocument();
  /**
   * <pre>
   *&#47; reserved field no. 3 for string document
   * </pre>
   *
   * <code>.mojo.lang.Document document = 4;</code>
   */
  org.mojolang.mojo.lang.DocumentOrBuilder getDocumentOrBuilder();

  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>string package = 6;</code>
   */
  java.lang.String getPackage();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>string package = 6;</code>
   */
  com.google.protobuf.ByteString
      getPackageBytes();

  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>string name = 10;</code>
   */
  java.lang.String getName();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>string name = 10;</code>
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>repeated .mojo.lang.GenericParameter generic_parameters = 11;</code>
   */
  java.util.List<org.mojolang.mojo.lang.GenericParameter> 
      getGenericParametersList();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>repeated .mojo.lang.GenericParameter generic_parameters = 11;</code>
   */
  org.mojolang.mojo.lang.GenericParameter getGenericParameters(int index);
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>repeated .mojo.lang.GenericParameter generic_parameters = 11;</code>
   */
  int getGenericParametersCount();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>repeated .mojo.lang.GenericParameter generic_parameters = 11;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.GenericParameterOrBuilder> 
      getGenericParametersOrBuilderList();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>repeated .mojo.lang.GenericParameter generic_parameters = 11;</code>
   */
  org.mojolang.mojo.lang.GenericParameterOrBuilder getGenericParametersOrBuilder(
      int index);

  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>repeated .mojo.lang.Attribute attributes = 12;</code>
   */
  java.util.List<org.mojolang.mojo.lang.Attribute> 
      getAttributesList();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>repeated .mojo.lang.Attribute attributes = 12;</code>
   */
  org.mojolang.mojo.lang.Attribute getAttributes(int index);
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>repeated .mojo.lang.Attribute attributes = 12;</code>
   */
  int getAttributesCount();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>repeated .mojo.lang.Attribute attributes = 12;</code>
   */
  java.util.List<? extends org.mojolang.mojo.lang.AttributeOrBuilder> 
      getAttributesOrBuilderList();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>repeated .mojo.lang.Attribute attributes = 12;</code>
   */
  org.mojolang.mojo.lang.AttributeOrBuilder getAttributesOrBuilder(
      int index);

  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>.mojo.lang.NominalType enclosing_type = 13;</code>
   */
  boolean hasEnclosingType();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>.mojo.lang.NominalType enclosing_type = 13;</code>
   */
  org.mojolang.mojo.lang.NominalType getEnclosingType();
  /**
   * <pre>
   *&#47;
   * </pre>
   *
   * <code>.mojo.lang.NominalType enclosing_type = 13;</code>
   */
  org.mojolang.mojo.lang.NominalTypeOrBuilder getEnclosingTypeOrBuilder();

  /**
   * <pre>
   * </pre>
   *
   * <code>.mojo.lang.EnumType type = 15;</code>
   */
  boolean hasType();
  /**
   * <pre>
   * </pre>
   *
   * <code>.mojo.lang.EnumType type = 15;</code>
   */
  org.mojolang.mojo.lang.EnumType getType();
  /**
   * <pre>
   * </pre>
   *
   * <code>.mojo.lang.EnumType type = 15;</code>
   */
  org.mojolang.mojo.lang.EnumTypeOrBuilder getTypeOrBuilder();
}
