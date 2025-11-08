#ifndef MOJO_LANG_ENUM_TYPE_BUILDER_HPP
#define MOJO_LANG_ENUM_TYPE_BUILDER_HPP

#include <mojo/lang/enum_type.hpp>
#include <mojo/lang/enum_type_of.hpp>

namespace mojo {
namespace lang {

template <typename T>
struct EnumTypeBuilder {
    EnumTypeOf<T>* type = nullptr;

    void name(std::string&& name){};

    void fullName(std::string&& full_name) {
        // msg->name_ = std::move(name);
    }

    ItemOf<T>& item(T const* i, std::string name, int number) {
        // auto p = new StructFieldFor<T, M>{field, std::move(name)};
        // msg->fields_.append(std::unique_ptr<FieldOf<T>>(p));
        // return *p;
    }

    EnumTypeOf<T>* operator()() {
        return type;
    }
};
}
}

#define MOJO_ENUM_BUILDER_NAME(types) \
    BOOST_PP_CAT(EnumTypeBuilderFor_, MOJO_TYPES_CAT(types))

#define MOJO_BUILD_ENUM_TYPE_HELPER(types)                             \
    struct MOJO_ENUM_BUILDER_NAME(types)                               \
        : ::mojo::lang::EnumTypeBuilder<::MOJO_TYPES_COMPOSE(types)> { \
        void build();                                                  \
        void build_with_defaults() {                                   \
            name(MOJO_TYPES_LAST_NAME(types));                         \
            fullName(MOJO_TYPES_FULL_NAME(types));                     \
            build();                                                   \
        }                                                              \
    };                                                                 \
    namespace mojo {                                                   \
    namespace lang {                                                   \
    inline const EnumTypeOf<::MOJO_TYPES_COMPOSE(types)>* buildType(   \
        const TypeIdentifier<::MOJO_TYPES_COMPOSE(types)>* dummy) {    \
        MOJO_ENUM_BUILDER_NAME(types) builder;                         \
        builder.type = new EnumTypeOf<::MOJO_TYPES_COMPOSE(types)>;    \
        builder.build_with_defaults();                                 \
        return builder();                                              \
    }                                                                  \
    }                                                                  \
    }                                                                  \
    inline void MOJO_ENUM_BUILDER_NAME(types)::build()

#define MOJO_BUILD_ENUM_TYPE(...) MOJO_BUILD_ENUM_TYPE_HELPER(MOJO_TYPES(__VA_ARGS__))

#endif  // MOJO_LANG_ENUM_TYPE_BUILDER_HPP
