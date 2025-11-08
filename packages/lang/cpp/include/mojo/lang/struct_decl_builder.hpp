#ifndef LANG_STRUCT_DECL_BUILDER_HPP
#define LANG_STRUCT_DECL_BUILDER_HPP

#include <mojo/core/string.hpp>
#include <mojo/lang/struct_decl.mojo.hpp>

namespace mojo {
namespace lang {

template <typename T>
struct StructDeclBuilder {
    StructDecl* decl;

    void name(String&& name) {
        decl->name = ::std::move(name);
    };

    void fullName(String&& fullName) {
    }

    void buildWithDefaults() {
    }

    StructDecl* operator()() {
        return decl;
    }
};

}  // namespace lang
}  // namespace mojo

#define MOJO_STRUCT_DECL_BUILDER_NAME(types) \
    BOOST_PP_CAT(StructDeclBuilderFor_, MOJO_TYPES_CAT(types))

#define MOJO_BUILD_STRUCT_DECL_HELPER(types)                             \
    struct MOJO_STRUCT_DECL_BUILDER_NAME(types)                          \
        : ::mojo::lang::StructDeclBuilder<::MOJO_TYPES_COMPOSE(types)> { \
        void build();                                                    \
        void buildWithDefaults() {                                       \
            name(MOJO_TYPES_LAST_NAME(types));                           \
            fullName(MOJO_TYPES_FULL_NAME(types));                       \
            build();                                                     \
        }                                                                \
    };                                                                   \
    namespace mojo {                                                     \
    namespace lang {                                                     \
    inline const StructDecl* buildDecl(                                  \
        const TypeIdentifier<::MOJO_TYPES_COMPOSE(types)>* dummy) {      \
        MOJO_STRUCT_DECL_BUILDER_NAME(types) builder;                    \
        builder.decl = new StructDecl;                                   \
        builder.buildWithDefaults();                                     \
        return builder();                                                \
    }                                                                    \
    }                                                                    \
    }                                                                    \
    inline void MOJO_STRUCT_DECL_BUILDER_NAME(types)::build()

#define MOJO_BUILD_STRUCT_DECL(...) MOJO_BUILD_STRUCT_DECL_HELPER(MOJO_TYPES(__VA_ARGS__))

MOJO_BUILD_STRUCT_DECL(mojo, lang, StructDecl) {
}

//MOJO_BUILD_STRUCT_DECL(mojo, lang, Attribute) {
//}

#endif  // LANG_STRUCT_DECL_BUILDER_HPP
