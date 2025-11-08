#ifndef MOJO_STRUCT_FIELD_FOR_HPP
#define MOJO_STRUCT_FIELD_FOR_HPP

#include <mojo/lang/struct_field.hpp>
#include <ncraft/data/data_visitor.hpp>
#include <ncraft/data/data_mutator.hpp>

namespace mojo {
namespace lang {

template <typename T>
class StructFieldFor : public StructField {
public:
    virtual ~StructFieldFor() = default;

    virtual bool hasValue(const T& msg) const = 0;

    virtual void visit(T&, ncraft::data::DataVisitor& visitor) const = 0;

    virtual ncraft::data::DataMutator mutator(T&) const = 0;
};

}  // namespace lang
}  // namespace mojo

#endif  // MOJO_STRUCT_FIELD_FOR_HPP
