# ErrorBADREQUEST

The error information


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Message`                                              | *string*                                               | :heavy_check_mark:                                     | The error message                                      | Invalid input data                                     |
| `Code`                                                 | *string*                                               | :heavy_check_mark:                                     | The error code                                         | BAD_REQUEST                                            |
| `Issues`                                               | [][apierrors.Issues](../../models/apierrors/issues.md) | :heavy_minus_sign:                                     | An array of issues that were responsible for the error | []                                                     |