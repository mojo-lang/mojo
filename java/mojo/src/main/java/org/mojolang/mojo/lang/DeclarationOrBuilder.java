// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/lang/declaration.proto

package org.mojolang.mojo.lang;

public interface DeclarationOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.lang.Declaration)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.mojo.lang.EnumDecl enum_decl = 2;</code>
   */
  boolean hasEnumDecl();
  /**
   * <code>.mojo.lang.EnumDecl enum_decl = 2;</code>
   */
  org.mojolang.mojo.lang.EnumDecl getEnumDecl();
  /**
   * <code>.mojo.lang.EnumDecl enum_decl = 2;</code>
   */
  org.mojolang.mojo.lang.EnumDeclOrBuilder getEnumDeclOrBuilder();

  /**
   * <code>.mojo.lang.StructDecl struct_decl = 3;</code>
   */
  boolean hasStructDecl();
  /**
   * <code>.mojo.lang.StructDecl struct_decl = 3;</code>
   */
  org.mojolang.mojo.lang.StructDecl getStructDecl();
  /**
   * <code>.mojo.lang.StructDecl struct_decl = 3;</code>
   */
  org.mojolang.mojo.lang.StructDeclOrBuilder getStructDeclOrBuilder();

  /**
   * <code>.mojo.lang.InterfaceDecl interface_decl = 4;</code>
   */
  boolean hasInterfaceDecl();
  /**
   * <code>.mojo.lang.InterfaceDecl interface_decl = 4;</code>
   */
  org.mojolang.mojo.lang.InterfaceDecl getInterfaceDecl();
  /**
   * <code>.mojo.lang.InterfaceDecl interface_decl = 4;</code>
   */
  org.mojolang.mojo.lang.InterfaceDeclOrBuilder getInterfaceDeclOrBuilder();

  /**
   * <pre>
   * TypeAliasDecl
   * </pre>
   *
   * <code>.mojo.lang.NominalDecl nominal_decl = 5;</code>
   */
  boolean hasNominalDecl();
  /**
   * <pre>
   * TypeAliasDecl
   * </pre>
   *
   * <code>.mojo.lang.NominalDecl nominal_decl = 5;</code>
   */
  org.mojolang.mojo.lang.NominalDecl getNominalDecl();
  /**
   * <pre>
   * TypeAliasDecl
   * </pre>
   *
   * <code>.mojo.lang.NominalDecl nominal_decl = 5;</code>
   */
  org.mojolang.mojo.lang.NominalDeclOrBuilder getNominalDeclOrBuilder();

  /**
   * <code>.mojo.lang.FuncDecl func_decl = 7;</code>
   */
  boolean hasFuncDecl();
  /**
   * <code>.mojo.lang.FuncDecl func_decl = 7;</code>
   */
  org.mojolang.mojo.lang.FuncDecl getFuncDecl();
  /**
   * <code>.mojo.lang.FuncDecl func_decl = 7;</code>
   */
  org.mojolang.mojo.lang.FuncDeclOrBuilder getFuncDeclOrBuilder();

  /**
   * <code>.mojo.lang.ConstDecl const_decl = 10;</code>
   */
  boolean hasConstDecl();
  /**
   * <code>.mojo.lang.ConstDecl const_decl = 10;</code>
   */
  org.mojolang.mojo.lang.ConstDecl getConstDecl();
  /**
   * <code>.mojo.lang.ConstDecl const_decl = 10;</code>
   */
  org.mojolang.mojo.lang.ConstDeclOrBuilder getConstDeclOrBuilder();

  /**
   * <code>.mojo.lang.ValueDecl value_decl = 11;</code>
   */
  boolean hasValueDecl();
  /**
   * <code>.mojo.lang.ValueDecl value_decl = 11;</code>
   */
  org.mojolang.mojo.lang.ValueDecl getValueDecl();
  /**
   * <code>.mojo.lang.ValueDecl value_decl = 11;</code>
   */
  org.mojolang.mojo.lang.ValueDeclOrBuilder getValueDeclOrBuilder();

  /**
   * <code>.mojo.lang.PackageDecl package_decl = 12;</code>
   */
  boolean hasPackageDecl();
  /**
   * <code>.mojo.lang.PackageDecl package_decl = 12;</code>
   */
  org.mojolang.mojo.lang.PackageDecl getPackageDecl();
  /**
   * <code>.mojo.lang.PackageDecl package_decl = 12;</code>
   */
  org.mojolang.mojo.lang.PackageDeclOrBuilder getPackageDeclOrBuilder();

  /**
   * <code>.mojo.lang.ImportDecl import_decl = 13;</code>
   */
  boolean hasImportDecl();
  /**
   * <code>.mojo.lang.ImportDecl import_decl = 13;</code>
   */
  org.mojolang.mojo.lang.ImportDecl getImportDecl();
  /**
   * <code>.mojo.lang.ImportDecl import_decl = 13;</code>
   */
  org.mojolang.mojo.lang.ImportDeclOrBuilder getImportDeclOrBuilder();

  public org.mojolang.mojo.lang.Declaration.DeclarationCase getDeclarationCase();
}