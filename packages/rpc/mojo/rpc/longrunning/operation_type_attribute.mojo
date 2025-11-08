
/// An attribute representing the message types used by a long-running operation.
///
/// ## ResponseType Required. The message name of the primary return type for this
/// long-running operation.
/// This type will be used to deserialize the LRO's response.
///
/// If the response is in a different package from the rpc, a fully-qualified
/// message name must be used (e.g. `google.protobuf.Struct`).
///
/// Note: Altering this value constitutes a breaking change.
///
/// ## MetadataType Required. The message name of the metadata type for this long-running
/// operation.
///
/// If the response is in a different package from the rpc, a fully-qualified
/// message name must be used (e.g. `google.protobuf.Struct`).
///
/// Note: Altering this value constitutes a breaking change.
attribute operation_type<ResponseType, MetadataType>: Bool
