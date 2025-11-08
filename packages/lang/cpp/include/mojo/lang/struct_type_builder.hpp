#ifndef MOJO_LANG_STRUCT_TYPE_BUILDER_HPP
#define MOJO_LANG_STRUCT_TYPE_BUILDER_HPP

#include <mojo/lang/data_type_preprocessors.hpp>
#include <mojo/lang/struct_field.hpp>
#include <mojo/lang/struct_field_of.hpp>
#include <mojo/lang/struct_type.hpp>
#include <mojo/lang/struct_type_of.hpp>
#include <mojo/lang/struct_virtual_field.hpp>

namespace mojo {
namespace lang {

template <typename T>
struct StructTypeBuilder {
    StructTypeOf<T>* type;

    void name(String&& name) {
        type->name = ::std::move(name);
    }

    void fullName(String&& fullName) {
        // type->fullName = ::std::move(fullName);
    }

    template <typename M>
    StructFieldOf<T, M>& field(M T::*field, String&& name, int number) {
        auto p = new StructFieldOf<T, M>{field, ::std::move(name), number};
        return addField(p);
    }

    template <typename M>
    StructFieldOf<T, M>& field(M T::*field,
                               String&& name,
                               String&& fullName,
                               int number) {
        auto p = new StructFieldOf<T, M>{field, ::std::move(name), number};
        return addField(p);
    }

    template <typename M>
    StructVirtualField<T, M>& field(M (T::*getter)(),
                                    void (T::*setter)(const M&),
                                    String&& name,
                                    int number) {
        auto p = new StructVirtualField<T, M>{getter, setter, ::std::move(name), number};
        return addField(p);
    }

    StructTypeOf<T>* operator()() {
        return type;
    }

private:
    template <typename F>
    inline F& addField(F* f) {
        type->fields.push_back(f);
        type->fields_.push_back(::std::unique_ptr<StructField<T>>(f));
        return *f;
    }
};
}  // namespace lang
}  // namespace mojo

#define MOJO_STRUCT_BUILDER_NAME(types) \
    BOOST_PP_CAT(StructTypeBuilderFor_, MOJO_TYPES_CAT(types))

#define MOJO_BUILD_STRUCT_TYPE_HELPER(types)                                         \
    struct MOJO_STRUCT_BUILDER_NAME(types)                                           \
        : ::mojo::lang::StructTypeBuilder<::MOJO_TYPES_COMPOSE(types)> {             \
        void build();                                                                \
        void buildWithDefaults() {                                                   \
            name(MOJO_TYPES_LAST_NAME(types));                                       \
            fullName(MOJO_TYPES_FULL_NAME(types));                                   \
            build();                                                                 \
        }                                                                            \
    };                                                                               \
    namespace ncraft {                                                               \
    namespace data {                                                                 \
    inline const ::mojo::lang::StructTypeOf<::MOJO_TYPES_COMPOSE(types)>* buildType( \
        const TypeIdentifier<::MOJO_TYPES_COMPOSE(types)>* dummy) {                  \
        MOJO_STRUCT_BUILDER_NAME(types) builder;                                     \
        builder.type = new ::mojo::lang::StructTypeOf<::MOJO_TYPES_COMPOSE(types)>;  \
        builder.buildWithDefaults();                                                 \
        return builder();                                                            \
    }                                                                                \
    }                                                                                \
    }                                                                                \
    inline void MOJO_STRUCT_BUILDER_NAME(types)::build()

#define MOJO_BUILD_STRUCT_TYPE(...) MOJO_BUILD_STRUCT_TYPE_HELPER(MOJO_TYPES(__VA_ARGS__))

MOJO_BUILD_STRUCT_TYPE(mojo, lang, DataType) {
    field(&mojo::lang::DataType::name, "name", 1);
    // field(&mojo::lang::DataType::fullName, "full_name", "fullName", 2);
    // field(&mojo::lang::DataType::package, "package", 3);
}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, StructType) {
}

//MOJO_BUILD_STRUCT_TYPE(mojo, lang, StructType, Field) {
//}

MOJO_BUILD_STRUCT_TYPE(mojo, lang, Attribute) {
    field(&mojo::lang::Attribute::name, "name", 1);
    field(&mojo::lang::Attribute::expression, "expression", 3);
    field(&mojo::lang::Attribute::package, "package", 4);
}

#endif  // MOJO_LANG_STRUCT_TYPE_BUILDER_HPP
