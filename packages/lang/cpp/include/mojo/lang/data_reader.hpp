#ifndef MOJO_LANG_DATA_READER_HPP
#define MOJO_LANG_DATA_READER_HPP

#include <memory>
#include <sstream>
#include <vector>

#include <ncraft/core/noncopyable.hpp>
#include <ncraft/core/cloning_ptr.hpp>

namespace mojo {
namespace lang {

struct IReaderEnumerator;
using ReaderEnumeratorPtr = CloningPtr<IReaderEnumerator>;

/*
  A Reader traverses data and inspects it as it passes over it.
  It has the ability to dig down into data structures.
*/
class DataReader {
public:
    virtual ~DataReader() {
    }

    virtual DataType* type() const noexcept = 0;

    virtual Optional<Bool> getBool() const noexcept = 0;

    virtual Optional<Integer> getInteger() const noexcept = 0;

    virtual Optional<Real> getReal() const noexcept = 0;

    virtual Optional<String> getString() const noexcept = 0;

    virtual bool has(const String& key) const noexcept = 0;

    virtual size_t length() const noexcept = 0;

    virtual ReaderPtr at(const String& key) const noexcept = 0;

    virtual ReaderPtr at(size_t idx) const noexcept = 0;

    virtual ReaderEnumeratorPtr enumerator() const noexcept = 0;
};

struct IReaderEnumerator {
    virtual ~IReaderEnumerator() {
    }

    virtual ReaderPtr currentValue() const = 0;

    virtual Maybe<String> currentKey() const = 0;

    virtual bool atEnd() const = 0;

    virtual void next() = 0;

    virtual IReaderEnumerator* clone() const = 0;
};

struct NullReader final : DataReader {
    NullReader() {
    }

    SDataType type() const final {
        return SDataType::Nothing;
    }

    Maybe<Boolean> getBoolean() const final {
        return Nothing;
    }

    Maybe<Integer> getInteger() const final {
        return Nothing;
    }

    Maybe<Real> getReal() const final {
        return Nothing;
    }

    Maybe<String> getString() const final {
        return Nothing;
    }

    bool has(const String& key) const {
        return false;
    }

    std::vector<String> keys() const {
        return {};
    }

    ReaderPtr get(const String& key) const {
        return nullptr;
    }

    size_t length() const {
        return 0;
    }

    ReaderPtr at(size_t idx) const {
        return nullptr;
    }

    ReaderEnumeratorPtr enumerator() const {
        return nullptr;
    }
};

struct ReaderEnumeratorAtEnd : Cloneable<ReaderEnumeratorAtEnd, IReaderEnumerator> {
    ReaderPtr currentValue() const final {
        return nullptr;
    }

    Optional<String> currentKey() const final {
        return kNull;
    }

    bool atEnd() const final {
        return true;
    }

    void next() final {
    }
};

}
}

#endif  // MOJO_LANG_DATA_READER_HPP
