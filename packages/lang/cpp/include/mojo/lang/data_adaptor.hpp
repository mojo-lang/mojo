#ifndef MOJO_DATA_ADAPTOR_HPP
#define MOJO_DATA_ADAPTOR_HPP

#include <ncraft/core/noncopyable.hpp>
#include <ncraft/core/result.hpp>
#include <mojo/lang/data_adaptor_ptr.hpp>

namespace mojo {
namespace lang {

class StructTypeField;

/**
 *
 */
class DataAdaptor : ::ncraft::NonCopyable {
public:
    using Adaptor = ::ncraft::Result<::std::pair<const DataAdaptor*, AnyRef>>;

public:
    virtual ~DataAdaptor() = default;

    /**
     *
     * @return
     */
    virtual int length(AnyConstRef) const noexcept = 0;

    /**
     *
     * @return
     */
    virtual bool has() const noexcept = 0;

    /**
     *
     * @return
     */
    virtual AnyConstRef at(AnyConstRef, const String&) const noexcept = 0;

    /**
     *
     * @return
     */
    virtual AnyConstRef at(AnyConstRef, int) const noexcept = 0;

    /**
     *
     */
    virtual ncraft::Status setNull(AnyRef) const noexcept = 0;

    /**
     *
     */
    virtual ncraft::Status setBool(AnyRef, bool) const noexcept = 0;

    /**
     *
     * @param data
     */
    virtual ncraft::Status set(AnyRef, AnyConstRef) const noexcept = 0;

    /**
     *
     * @param key
     */
    virtual Adaptor mutableAt(AnyRef, const String&) const noexcept = 0;

    /**
     *
     * @param idx
     */
    virtual Adaptor mutableAt(AnyRef, int) const noexcept = 0;

    /**
     *
     * @param name
     * @param fld
     * @return
     */
    virtual Adaptor mutateField(AnyRef,
                                const String&,
                                const StructTypeField* = nullptr) const noexcept = 0;

    /**
     *
     * @return
     */
    virtual Adaptor add(AnyRef) const noexcept = 0;

    /**
     *
     * @param key
     */
    virtual void erase(AnyRef, const String&) const noexcept = 0;

    /**
     *
     * @param index
     */
    virtual void erase(AnyRef, int) const noexcept = 0;

    /**
     *
     */
    virtual void clear(AnyRef) const noexcept = 0;
};
}
}

#endif  // MOJO_DATA_ADAPTOR_HPP
