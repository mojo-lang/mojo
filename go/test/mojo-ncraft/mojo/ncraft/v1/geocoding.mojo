/// Geocoding API
///
@http.base_path('/geocoding/v1{/user}')
interface Geocoding {
    /// geocode for coding the address string to lng lat point and structural address type.
    @http.get('/address/{address}{.format}')
    geocode(address: String @1) -> Address

    /// reverse geocode for decoding from the lng lat point to the structural address type.
    @http.get('/locations/{location}')
    @error_codes([{
        value: 405
        brief: 'Invalid ID supplied'
        category: 'geocoding'
    }, {
        value: 404
        brief: 'Not found'
        category: 'geocoding'
    }])
    reverseGeocode (location:geom.LngLat@5,
        nearby:Bool @10,
        nearby_radius:Int @11
    ) -> Address @2
}
