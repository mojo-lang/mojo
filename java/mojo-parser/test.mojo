type Message {
    a:Int8
    b:String
}

interface Service {
 convert(from :From) -> Result
}
