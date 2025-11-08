enum Type {
    unspecified @0
    null    @1
    boolean @2
    integer @3
    number  @4
    string  @5

    array   @10 //< An ordered list of instances, from the JSON "array" production.
    map     @11 //< An unordered list of key value pairs, from the JSON "object" production.
    struct  @12 //< An unordered set of properties mapping a string to an instance, from the JSON "object" production.
}