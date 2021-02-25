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
struct {{Type_name}} : {l% } {
    {%fields}
    {%}
};

{% namespace_end }
{%}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, Struct_decl) {

    {%fields}
    field(&mojo::lang::Struct_decl::type, "type", 1);
    {%}
}

MOJO_BUILD_STRUCT_DECL(mojo, lang, Struct_decl) {

    {%fields}
    field(&mojo::lang::Struct_decl::type, "type", 1);
    {%}
}

#endif {{file_unique_define}}


  Char Expression Statement