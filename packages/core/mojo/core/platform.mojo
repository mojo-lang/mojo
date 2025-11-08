
enum Architecture {
    unspecified @0

    x86 @1
    amd64 @2

    arm @5
    arm64 @6

    wasm @10
}

enum OS {
    unspecified @0

    android @1

    darwin @2

    ios @6 //@alias("iOS")

    linux @10

    windows @20

    simulation @30
}

/// the platform information
///
/// Normalization
///
/// Because not all users are familiar with the way the Go runtime represents
/// platforms, several normalizations have been provided to make this package
/// easier to user.
///
/// The following are performed for architectures:
///
///   Value    Normalized
///   aarch64  arm64
///   armhf    arm
///   armel    arm/v6
///   i386     386
///   x86_64   amd64
///   x86-64   amd64
///
/// We also normalize the operating system `macos` to `darwin`.
///
/// ARM Support
///
/// To qualify ARM architecture, the Variant field is used to qualify the arm
/// version. The most common arm version, v7, is represented without the variant
/// unless it is explicitly provided. This is treated as equivalent to armhf. A
/// previous architecture, armel, will be normalized to arm/v6.
@format("{os}/{architecture}{/variant}{-os_name/os_version}")
type Platform {
    /// the CPU architecture, such as "x86" or "amd64".
    architecture: Architecture @1
    
    /// to qualify ARM architecture, the Variant field is used to qualify the arm version.
    variant: String @2 @format('[a-z0-9]+')

    /// the os type in the platform.
    os: OS @10 @required

    /// a string identifying the operating system, such as "Windows NT",
    /// "Mac OS X", or "Ubutun".
    os_name: String @11 @format('[a-zA-Z0-9 ._]+')
    
    /// A string identifying the version of the operating system, such as
    /// "5.1.2600 Service Pack 2" or "10.4.8 8L2127", or "14.02".
    os_version: String @12 @format('[a-zA-Z0-9 ._]+')
}
