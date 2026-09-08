// Copyright 2020 Google LLC
// [source](https://github.com/googleapis/googleapis/google/longrunning/operations.proto)
//
// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/// Manages long-running operations with an API service.
///
/// When an API method normally takes long time to complete, it can be designed
/// to return [Operation][google.longrunning.Operation] to the client, and the client can use this
/// interface to receive the real response asynchronously by polling the
/// operation resource, or pass the operation resource to another API (such as
/// Google Cloud Pub/Sub API) to receive the response.  Any API service that
/// returns long-running operations should implement the `Operations` interface
/// so developers can have a consistent client experience.
@entity("Operation")
interface Operations {
    /// Lists operations that match the specified filter in the request. If the
    /// server doesn't support this method, it returns `UNIMPLEMENTED`.
    ///
    /// NOTE: the `name` binding allows API services to override the binding
    /// to use different resource name schemes, such as `users/*/operations`. To
    /// override the binding, API services can add a binding such as
    /// `"/v1/{name=users/*}/operations"` to their service configuration.
    /// For backwards compatibility, the default name includes the operations
    /// collection id, however overriding users must ensure the name binding
    /// is the parent resource, without the operations collection id.
    @http.get("/operation/v1/{{service}}/operations")
    list_operations(parent: String @1  //< The name of the operation's parent resource.
                    filter: String @4) //< The standard list filter.
                 -> [Operation]

    /// Gets the latest state of a long-running operation.  Clients can use this
    /// method to poll the operation result at intervals as recommended by the API
    /// service.
    @http.get("/operation/v1/{{service}}/operations/{name}")
    get_operation(name: String @1) //< The name of the operation resource.
               -> Operation

    /// Deletes a long-running operation. This method indicates that the client is
    /// no longer interested in the operation result. It does not cancel the
    /// operation. If the server doesn't support this method, it returns
    /// `google.rpc.Code.UNIMPLEMENTED`.
    @http.delete("/operation/v1/{{service}}/operations/{name}")
    delete_operation(name: String @1) //< The name of the operation resource to be deleted.

    /// Starts asynchronous cancellation on a long-running operation.  The server
    /// makes a best effort to cancel the operation, but success is not
    /// guaranteed.  If the server doesn't support this method, it returns
    /// `google.rpc.Code.UNIMPLEMENTED`.  Clients can use
    /// [Operations.GetOperation][google.longrunning.Operations.GetOperation] or
    /// other methods to check whether the cancellation succeeded or whether the
    /// operation completed despite cancellation. On successful cancellation,
    /// the operation is not deleted; instead, it becomes an operation with
    /// an [Operation.error][google.longrunning.Operation.error] value with a [google.rpc.Status.code][google.rpc.Status.code] of 1,
    /// corresponding to `Code.CANCELLED`.
    @http.post("/operation/v1/{{service}}/operations/{name}:cancel")
    cancel_operation(name: String @1) //< The name of the operation resource to be cancelled.

    /// Waits until the specified long-running operation is done or reaches at most
    /// a specified timeout, returning the latest state.  If the operation is
    /// already done, the latest state is immediately returned.  If the timeout
    /// specified is greater than the default HTTP/RPC timeout, the HTTP/RPC
    /// timeout is used.  If the server does not support this method, it returns
    /// `google.rpc.Code.UNIMPLEMENTED`.
    /// Note that this method is on a best-effort basis.  It may return the latest
    /// state before the specified timeout (including immediately), meaning even an
    /// immediate response is no guarantee that the operation is done.
    @http.post("/operation/v1/{{service}}/operations/{name}:wait")
    wait_operation(name: String @1       //< The name of the operation resource to wait on.
                   timeout: Duration @2) //< The maximum duration to wait before timing out. If left blank, the wait
                                         //< will be at most the time permitted by the underlying HTTP/RPC protocol.
                                         //< If RPC context deadline is also specified, the shorter one will be used.
                -> Operation
}
