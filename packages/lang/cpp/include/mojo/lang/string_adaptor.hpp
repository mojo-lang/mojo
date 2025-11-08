#ifndef MOJO_STRING_ADAPTOR_HPP
#define MOJO_STRING_ADAPTOR_HPP

#include <mojo/lang/data_adaptor_of.hpp>

namespace mojo {
namespace lang {

class StringAdaptor : public DataAdaptorOf<String> {
public:
    StringAdaptor() = default;
    virtual ~StringAdaptor() = default;

    virtual int size(AnyConstRef str) const noexcept {
        return  (*str.get<String>()).size();
    }

    virtual bool has() const noexcept {
        return false;
    }

    virtual AnyConstRef at(AnyConstRef str, const String&) const noexcept {

    }

    virtual AnyConstRef at(AnyConstRef str, int) const noexcept {

    }

    virtual ncraft::Status setNull(AnyRef) const noexcept {
    }

    virtual ncraft::Status setBool(AnyRef, bool) const noexcept {
    }

    virtual ncraft::Status set(AnyRef str, AnyConstRef) const noexcept {
    }

    virtual MResult mutableAt(AnyRef str, const String&) const noexcept {
    }

    virtual MResult mutableAt(AnyRef str, int) const noexcept {
    }

    virtual MResult mutateField(AnyRef,
                                const String&,
                                const StructTypeField* = nullptr) const noexcept {
    }

    virtual MResult add(AnyRef str) const noexcept {
    }

    virtual void erase(AnyRef str, const String&) const noexcept {
    }

    virtual void erase(AnyRef, int) const noexcept {
    }

    virtual void clear(AnyRef str) const noexcept {
        //str.clear();
    }
};

}
}

#endif //MOJO_STRING_ADAPTOR_HPP
