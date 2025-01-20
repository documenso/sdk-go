# ErrorNOTFOUND

The error information


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Message`                                                                        | *string*                                                                         | :heavy_check_mark:                                                               | The error message                                                                | Not found                                                                        |
| `Code`                                                                           | *string*                                                                         | :heavy_check_mark:                                                               | The error code                                                                   | NOT_FOUND                                                                        |
| `Issues`                                                                         | [][apierrors.ErrorNOTFOUNDIssues](../../models/apierrors/errornotfoundissues.md) | :heavy_minus_sign:                                                               | An array of issues that were responsible for the error                           | []                                                                               |