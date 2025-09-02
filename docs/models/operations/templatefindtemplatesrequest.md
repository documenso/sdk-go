# TemplateFindTemplatesRequest


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Query`                                                                 | **string*                                                               | :heavy_minus_sign:                                                      | The search query.                                                       |
| `Page`                                                                  | **float64*                                                              | :heavy_minus_sign:                                                      | The pagination page number, starts at 1.                                |
| `PerPage`                                                               | **float64*                                                              | :heavy_minus_sign:                                                      | The number of items per page.                                           |
| `Type`                                                                  | [*operations.QueryParamType](../../models/operations/queryparamtype.md) | :heavy_minus_sign:                                                      | Filter templates by type.                                               |
| `FolderID`                                                              | **string*                                                               | :heavy_minus_sign:                                                      | The ID of the folder to filter templates by.                            |