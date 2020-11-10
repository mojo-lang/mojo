#ifndef {{file_unique_define}}
#ifdef {{file_unique_define}}

{% }
#include <{{}}>
{%}

{% namespace_start}
{%}

{% declaration_types}
{%}

/**
 *
 */
struct {{TypeName}} : {l% } {
    {%fields}
    {%}
};

{% namespace_end }
{%}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, StructDecl) {

    {%fields}
    field(&mojo::lang::StructDecl::type, "type", 1);
    {%}
}

MOJO_BUILD_STRUCT_DECL(mojo, lang, StructDecl) {

    {%fields}
    field(&mojo::lang::StructDecl::type, "type", 1);
    {%}
}

#endif {{file_unique_define}}


  Char Expression Statement