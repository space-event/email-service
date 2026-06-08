# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/email.proto](#api_email-proto)
    - [EmailRequest](#email-EmailRequest)
    - [EmailResponse](#email-EmailResponse)
  
    - [EmailService](#email-EmailService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_email-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/email.proto



<a name="email-EmailRequest"></a>

### EmailRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| emailTarget | [string](#string) |  | Target email address (required) Example: &#34;user@example.com&#34; |
| messageText | [string](#string) |  | Email body content (required) Supports both plain text and HTML depending on contentType field |
| subject | [string](#string) |  | Email subject line Should be concise and descriptive |
| contentType | [string](#string) |  | Content type of the email body (required) Allowed values: &#34;text/plain&#34; or &#34;text/html&#34; |






<a name="email-EmailResponse"></a>

### EmailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | Indicates whether the email was sent successfully true = email was queued/sent, false = failed |





 

 

 


<a name="email-EmailService"></a>

### EmailService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Send | [EmailRequest](#email-EmailRequest) | [EmailResponse](#email-EmailResponse) | Send an email to the specified recipient Returns EmailResponse with success status or error |



## Error Codes

| Code | Description | When it occurs |
|------|-------------|----------------|
| InvalidArgument | 3 | Missing required field or invalid email format |
| Internal | 13 | SMTP server error or email sending failed |
| Canceled | 1 | Request was cancelled by client |
| DeadlineExceeded | 4 | Request timeout exceeded |



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

