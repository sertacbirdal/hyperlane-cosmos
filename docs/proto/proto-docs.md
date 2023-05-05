<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [hyperlane/mailbox/v1/genesis.proto](#hyperlane/mailbox/v1/genesis.proto)
    - [GenesisState](#hyperlane.mailbox.v1.GenesisState)
    - [GenesisTreeEntry](#hyperlane.mailbox.v1.GenesisTreeEntry)
  
- [hyperlane/mailbox/v1/query.proto](#hyperlane/mailbox/v1/query.proto)
    - [QueryCurrentTreeMetadataRequest](#hyperlane.mailbox.v1.QueryCurrentTreeMetadataRequest)
    - [QueryCurrentTreeMetadataResponse](#hyperlane.mailbox.v1.QueryCurrentTreeMetadataResponse)
  
    - [Query](#hyperlane.mailbox.v1.Query)
  
- [hyperlane/mailbox/v1/tx.proto](#hyperlane/mailbox/v1/tx.proto)
    - [MsgSetDefaultIsm](#hyperlane.mailbox.v1.MsgSetDefaultIsm)
    - [MsgSetDefaultIsmResponse](#hyperlane.mailbox.v1.MsgSetDefaultIsmResponse)
  
    - [Msg](#hyperlane.mailbox.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="hyperlane/mailbox/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/mailbox/v1/genesis.proto



<a name="hyperlane.mailbox.v1.GenesisState"></a>

### GenesisState
Hyperlane mailbox's keeper genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tree_entry` | [GenesisTreeEntry](#hyperlane.mailbox.v1.GenesisTreeEntry) | repeated | Each genesis tree entry |






<a name="hyperlane.mailbox.v1.GenesisTreeEntry"></a>

### GenesisTreeEntry
Hyperlane's tree entry


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [uint32](#uint32) |  | index |
| `message` | [bytes](#bytes) |  | message |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="hyperlane/mailbox/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/mailbox/v1/query.proto



<a name="hyperlane.mailbox.v1.QueryCurrentTreeMetadataRequest"></a>

### QueryCurrentTreeMetadataRequest
QueryCurrentTreeMetadataRequest is the request type for the Query/Tree RPC
method.






<a name="hyperlane.mailbox.v1.QueryCurrentTreeMetadataResponse"></a>

### QueryCurrentTreeMetadataResponse
QueryTreeResponse is the response type for the Query/Tree RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `root` | [bytes](#bytes) |  |  |
| `count` | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="hyperlane.mailbox.v1.Query"></a>

### Query
Query service for hyperlane mailbox module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CurrentTreeMetadata` | [QueryCurrentTreeMetadataRequest](#hyperlane.mailbox.v1.QueryCurrentTreeMetadataRequest) | [QueryCurrentTreeMetadataResponse](#hyperlane.mailbox.v1.QueryCurrentTreeMetadataResponse) | Get current tree metadata | GET|/hyperlane/mailbox/v1/tree|

 <!-- end services -->



<a name="hyperlane/mailbox/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## hyperlane/mailbox/v1/tx.proto



<a name="hyperlane.mailbox.v1.MsgSetDefaultIsm"></a>

### MsgSetDefaultIsm
MsgSetDefaultIsm defines the request type for the SetDefaultIsm rpc.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signer` | [string](#string) |  |  |
| `ism_hash` | [string](#string) |  |  |






<a name="hyperlane.mailbox.v1.MsgSetDefaultIsmResponse"></a>

### MsgSetDefaultIsmResponse
MsgSetDefaultIsmResponse defines the Msg/SetDefaultIsm response type





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="hyperlane.mailbox.v1.Msg"></a>

### Msg
Msg defines the hyperlane mailbox Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SetDefaultIsm` | [MsgSetDefaultIsm](#hyperlane.mailbox.v1.MsgSetDefaultIsm) | [MsgSetDefaultIsmResponse](#hyperlane.mailbox.v1.MsgSetDefaultIsmResponse) | SetDefaultIsm defines a rpc handler method for MsgSetDefaultIsm. | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
