
type Cookie {
	name:  String @1 @required
	value: String @2 @required

	path:       String @3   // optional
	domain:     String @4   // optional
	expires:    Timestamp @5 // optional

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	max_age:   Int32 @6
	secure:   Bool @7
	http_only: Bool @8
	same_site: SameSite @9
}

// SameSite allows a server to define a cookie attribute making it impossible for
// the browser to send this cookie along with cross-site requests. The main
// goal is to mitigate the risk of cross-origin information leakage, and provide
// some protection against cross-site request forgery attacks.
//
// See https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00 for details.
enum SameSite {
    unsepecified @0

    lax  @1
    strict @2
    none @3
}
