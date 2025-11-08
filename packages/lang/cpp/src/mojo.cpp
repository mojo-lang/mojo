#include <mojo/lang/package.mojo.hpp>
#include <mojo/lang/source_file.mojo.hpp>
#include <mojo/lang/statement.mojo.hpp>
#include <mojo/lang/struct_decl.mojo.hpp>
#include <mojo/lang/declaration.mojo.hpp>

#include <mojo/lang/func_decl.mojo.hpp>
#include <mojo/lang/func_type.mojo.hpp>
#include <mojo/lang/interface_decl.mojo.hpp>

using namespace mojo;
using namespace mojo::lang;

int main(int argc, char* argv[]) {

    auto package = std::make_shared<Package>();
    auto sourceFile = std::make_shared<SourceFile>();

    auto structField = new StructField;
    structField->name = "foo";
    structField->type.name = "Bar";

    auto structType = new StructType{};
    structType->name = "Foo";
    structType->fields.push_back(structField);

    auto structDecl = std::make_shared<StructDecl>();
    structDecl->name = "Foo";
    structDecl->document = "test";
    structDecl->type = structType;
    structDecl->package = package;
    structDecl->sourceFile = sourceFile;

    auto funcType = std::make_shared<FuncType>();

    auto funcDecl = std::make_shared<FuncDecl>();
    funcDecl->name = "foo";
    funcDecl->type = funcType;

    auto interfaceType = new InterfaceType;
    interfaceType->methods.push_back(funcDecl);

    auto interfaceDecl = std::make_shared<InterfaceDecl>();
    interfaceDecl->name = "FooService";
    interfaceDecl->type = interfaceType;

    sourceFile->name = "test";
    sourceFile->statements.push_back(Statement{Declaration{structDecl}});
    sourceFile->structDecls.emplace(structDecl->name, structDecl);
    sourceFile->interfaceDecls.emplace(interfaceDecl->name, interfaceDecl);
    package->sourceFiles.push_back(std::move(sourceFile));

    return 0;
}


