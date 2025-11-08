#ifndef MOJO_LANG_DATA_WRITER_HPP
#define MOJO_LANG_DATA_WRITER_HPP

#include <ncraft/core/noncopyable.hpp>
#include <mojo/lang/data_type.hpp>
#include <mojo/lang/reader.hpp>

namespace mojo {
namespace lang {

struct Spectator;

class DataWriter : ncraft::NonCopyable {
public:
    virtual ~DataWriter() {
    }

    virtual ncraft::Status setNull() noexcept = 0;

    virtual ncraft::Status setBool(Bool b) noexcept = 0;

    virtual ncraft::Status setInteger(Integer n) noexcept = 0;

    virtual ncraft::Status setReal(Real r) noexcept = 0;

    virtual ncraft::Status setString(String str) noexcept = 0;

    virtual ncraft::Status set(AnyConstRef) noexcept = 0;

    virtual AdapterPtr mutableAt(size_t idx) noexcept = 0;

    virtual AdapterPtr mutableAt(const String& key) noexcept = 0;

    virtual AdapterPtr add() noexcept = 0;

    virtual AdapterPtr add(const String& key) noexcept = 0;

    virtual ncraft::Status erase(const String& key) noexcept = 0;

    virtual ncraft::Status erase(size_t index) noexcept = 0;

    virtual void clear() noexcept = 0;
};

}  // namespace lang
}  // namespace mojo

#endif  // MOJO_LANG_DATA_WRITER_HPP
