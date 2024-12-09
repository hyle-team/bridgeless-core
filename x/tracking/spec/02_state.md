# State

## Position

Each position is represented in blockchain with unique address.
The Position object contains specific information to describe a position:
``` protobuf
message Position {
string address = 1;
string recipient =2;
int64 amount = 3;
}
```
There are several commands to operate with Position store:
* To add or update existing position use `SetPosition(ctx sdk.Context, positionAddress string, position types.Position)`.
* To remove position use `DeletePosition(ctGetPositionsWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Position, *query.PageResponse, error)x sdk.Context, positionAddress string)`.
* To get position info use `GetPosition(ctx sdk.Context, positionAddress string) (position types.Position, found bool)`
* To get positions info with pagination use `GetPositionsWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Position, *query.PageResponse, error)`.
