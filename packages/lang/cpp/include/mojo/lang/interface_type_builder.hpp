#ifndef MOJO_LANG_SERVICE_TYPE_BUILDER_HPP
#define MOJO_LANG_SERVICE_TYPE_BUILDER_HPP

#include <mojo/lang/interface_type.hpp>
#include <mojo/lang/interface_type_of.hpp>

namespace mojo {
namespace lang {

template <typename T>
struct InterfaceTypeBuilder {
    InterfaceTypeOf<T>* type;

    void name(std::string&& name, std::string&& full_name) {
        // msg->name_ = std::move(name);
    }

    template <typename Req, typename Resp>
    InterfaceMethodOf<T, Req, Resp>& method(Resp T::*func(Req), std::string name, int number) {
        auto p = new MethodFor<T, Req, Resp>{func, std::move(name)};
        // msg->fields_.append(std::unique_ptr<FieldOf<T>>(p));
        return *p;
    }

    InterfaceTypeOf<T>* operator()() {
        return type;
    }
};

}
}

#define MOJO_BUILD_SERVICE_TYPE(TYPE)                                                     \
    struct InterfaceTypeBuilderFor##TYPE : ::ncraft::proto::InterfaceTypeBuilder<TYPE> { \
        void build();                                                                \
        void build_with_defaults() {                                                 \
            name(#TYPE);                                                             \
            build();                                                                 \
        }                                                                            \
    };                                                                               \
    inline const ::ncraft::proto::InterfaceType<TYPE>* buildType(                      \
        const ::ncraft::core::TypeIdentifier<TYPE>* dummy) {                         \
        InterfaceTypeBuilderFor##TYPE builder;                                         \
        builder.type = new ::ncraft::proto::InterfaceTypeOf<TYPE>;                    \
        builder.build_with_defaults();                                               \
        return builder();                                                            \
    }                                                                                \
    inline void InterfaceTypeBuilderFor##TYPE::build()

#endif //MOJO_LANG_SERVICE_TYPE_BUILDER_HPP
