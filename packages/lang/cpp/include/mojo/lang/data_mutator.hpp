#ifndef MOJO_DATA_MUTATOR_HPP
#define MOJO_DATA_MUTATOR_HPP

#include <ncraft/core/any.hpp>
#include <ncraft/core/movable.hpp>
#include <ncraft/core/noncopyable.hpp>
#include <mojo/lang/data_adaptor.hpp>

namespace mojo {
namespace lang {

class DataMutator : ::ncraft::NonCopyable, ::ncraft::Movable {
public:
    DataMutator() = default;
    DataMutator(::ncraft::AnyRef value, const DataAdaptor* adaptor);

    DataMutator(DataMutator&& mutator);
    DataMutator& operator=(DataMutator&& mutator);

    /**
     *
     */
    void setNull() noexcept {
        //adaptor_->setNull(value_);
    }

    /**
     *
     */
    void setBool(bool value) noexcept {
        //adaptor_->setBool(value_, value);
    }

    /**
     *
     */
    void setInt8(Int8) noexcept;

    /**
     *
     */
    void setInt16(Int16) noexcept;

    /**
     *
     */
    void setInt32(Int32) noexcept;

    /**
     *
     */
    void setInt64(Int64) noexcept;

    /**
     *
     */
    void setUInt8(UInt8) noexcept;

    /**
     *
     */
    void setUInt16(UInt16) noexcept;

    /**
     *
     */
    void setUInt32(UInt32) noexcept;

    /**
     *
     */
    void setUInt64(UInt64) noexcept;

    /**
     *
     */
    void setFloat(Float) noexcept;

    /**
     *
     */
    void setDouble(Double) noexcept;

    /**
     *
     */
    void setString(const ncraft::StringView&) noexcept;

    /**
     *
     * @param data
     * @param type
     */
    void set(AnyConstRef data) noexcept;

    /**
     *
     * @param key
     */
    DataMutator key(const String& key) noexcept;

    /**
     *
     * @param idx
     */
    DataMutator index(int idx) noexcept;

    /**
     *
     * @param name
     * @param fld
     * @return
     */
    DataMutator field(const String& name,
                      const StructTypeField* fld = nullptr) noexcept;

    /**
     *
     * @return
     */
    DataMutator add() noexcept;

    /**
     *
     * @param key
     */
    DataMutator& erase(const String& key) noexcept;

    /**
     *
     * @param index
     */
    DataMutator& erase(int index) noexcept;

    /**
     *
     */
    DataMutator& clear() noexcept;

private:
    AnyRef value_;
    const DataAdaptor* adaptor_ = nullptr;
};
}
}

#endif  // MOJO_DATA_MUTATOR_HPP
