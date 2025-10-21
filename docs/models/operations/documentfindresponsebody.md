# DocumentFindResponseBody

Successful response


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Data`                                                                       | [][operations.DocumentFindData](../../models/operations/documentfinddata.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `Count`                                                                      | *float64*                                                                    | :heavy_check_mark:                                                           | The total number of items.                                                   |
| `CurrentPage`                                                                | *float64*                                                                    | :heavy_check_mark:                                                           | The current page number, starts at 1.                                        |
| `PerPage`                                                                    | *float64*                                                                    | :heavy_check_mark:                                                           | The number of items per page.                                                |
| `TotalPages`                                                                 | *float64*                                                                    | :heavy_check_mark:                                                           | The total number of pages.                                                   |