/// hello world service
///
/// will do a simple echo
interface HelloWorld {
    /// say hello to the name
    @http.get("/hello-world/v1/echos/{name}")
    get_echo(name: String @1) //< echo name
           -> Echo
}
