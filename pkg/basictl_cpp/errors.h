#ifndef BASICTL_CPP_ERRORS_H
#define BASICTL_CPP_ERRORS_H

#include <exception>
#include <string_view>
#include <cstring>
#include <string>

namespace basictl {
    enum class tl_error_type {
        STREAM_EOF,
        INCORRECT_SEQUENCE_LENGTH,
        INCORRECT_STRING_PADDING,
        UNEXPECTED_TAG,
        UNKNOWN_SCENARIO,
    };

    template<typename Type>
    class basic_error final : public std::exception {
    public:
        basic_error(Type type, std::string message)
                : type_(type), message_(std::move(message)) {}

        [[nodiscard]] Type type() const noexcept {
            return type_;
        }

        [[nodiscard]] const std::string& message() const & noexcept {
            return message_;
        }

        [[nodiscard]] std::string&& message() && noexcept {
            return std::move(message_);
        }

        [[nodiscard]] const char *what() const noexcept override {
            return message_.c_str();
        }

    private:
        Type type_;
        std::string message_;
    };

    using tl_stream_error = basic_error<tl_error_type>;
    using tl_connector_error = basic_error<std::uint32_t>;
    using tl_error = std::variant<tl_stream_error, tl_connector_error>;
}

#endif //BASICTL_CPP_ERRORS_H
