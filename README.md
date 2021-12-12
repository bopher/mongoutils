# MongoUtils

Mongodb helper functions, document and pipeline builder.

## Helpers

### ParseObjectID

Parse object id from string.

```go
ParseObjectID(id string) *primitive.ObjectID
```

### IsValidObjectId

Check if object id is valid and not zero.

```go
IsValidObjectId(id *primitive.ObjectID) bool
```

## Doc Builder

Document builder is a helper type for creating mongo document (`primitive.D`) with _chained_ methods.

```go
import "github.com/bopher/mongoutils"
doc := mongoutils.NewDoc()
doc.
    Add("name", "John").
    Add("nick", "John2").
    Array("skills", "javascript", "go", "rust", "mongo")
fmt.Println(doc.Build())
// -> {
//   "name": "John",
//   "nick": "John2",
//   "skills": ["javascript","go","rust","mongo"]
// }
```

### Doc Methods

#### Add

Add new element.

```go
// Signature:
Add(k string, v interface{}) MongoDoc

// Example:
doc.Add("name", "Kim")
```

#### Doc

Add new element with nested doc value.

```go
// Signature:
Doc(k string, cb func(d MongoDoc) MongoDoc) MongoDoc

// Example:
doc.Doc("age", func(d mongoutils.MongoDoc) mongoutils.MongoDoc {
    d.Add("$gt", 20)
    d.Add("$lte", 30)
    return d
}) // -> { "age": { "$gt": 20, "$lte": 30 } }
```

#### Array

Add new element with array value.

```go
// Signature:
Array(k string, v ...interface{}) MongoDoc

// Example:
doc.Array("skills", "javascript", "golang") // -> { "skills": ["javascript", "golang"] }
```

#### Nested

Add new nested element.

```go
// Signature:
Nested(root string, k string, v interface{}) MongoDoc

// Example:
doc.Nested("$set", "name", "Jack") // { "$set": { "name": "Jack" } }
```

#### NestedDoc

Add new nested element with doc value.

```go
// Signature:
NestedDoc(root string, k string, cb func(d MongoDoc) MongoDoc) MongoDoc

// Example:
doc.NestedDoc("$set", "address", func(d mongoutils.MongoDoc) mongoutils.MongoDoc {
    d.
        Add("city", "London").
        Add("street", "12th")
    return d
}) // -> { "$set": { "address": { "city": "London", "street": "12th" } } }
```

#### NestedArray

Add new nested element with array value.

```go
// Signature:
NestedArray(root string, k string, v ...interface{}) MongoDoc

// Example:
doc.NestedArray("skill", "$in", "mongo", "golang") // -> { "skill": { "$in": ["mongo", "golang"] } }
```

#### Regex

Add new element with regex value.

```go
// Signature:
Regex(k string, pattern string, opt string) MongoDoc

// Example:
doc.Regex("full_name", "John.*", "i") // -> { "full_name": { pattern: "John.*", options: "i" } }
```

#### Map

Creates a map from the elements of the Doc.

```go
Map() primitive.M
```

#### Build

Generate mongo doc.

```go
Build() primitive.D
```

## Pipeline Builder

Pipeline builder is a helper type for creating mongo pipeline (`[]primitive.D`) with _chained_ methods.

```go
import "github.com/bopher/mongoutils"
pipe := mongoutils.NewPipe()
pipe.
    Add(func(d mongoutils.MongoDoc) mongoutils.MongoDoc{
        d.Nested("$match", "name", "John")
        return d
    }).
    Group(func(d mongoutils.MongoDoc) mongoutils.MongoDoc{
        d.
            Add("_id", "$_id").
            Nested("name", "$first", "$name")
            Nested("total", "$sum", "$invoice")
        return d
    })
fmt.Println(pipe.Build())
// -> [
//   { "$match": { "name": "John"} },
//   { "$group": {
//       "_id": "$_id"
//       "name": { "$first": "$name" },
//       "total": { "$sum": "$invoice" }
//   }}
// ]
```

### Pipeline Methods

#### Add

Add new Doc.

```go
// Signature:
Add(cb func(d MongoDoc) MongoDoc) MongoPipeline

// Example:
pipe.Add(func(d mongoutils.MongoDoc) mongoutils.MongoDoc{
    d.Nested("$match", "name", "John")
    return d
}) // -> [ {"$match": { "name": "John"}} ]
```

#### Lookup

Add $lookup stage.

```go
// Signature:
Lookup(from string, local string, foreign string, as string) MongoPipeline

// Example:
pipe.Lookup("users", "user_id", "_id", "user")
// -> [
//     {"$lookup": {
//         "from": "users",
//         "localField": "user_id",
//         "foreignField": "_id",
//         "as": "user"
//     }}
// ]
```

#### Unwind

Add $unwind stage.

```go
// Signature:
Unwind(path string, prevNullAndEmpty bool) MongoPipeline

// Example:
pipe.Unwind("services", true)
// -> [
//     {"$unwind": {
//         "path": "services",
//         "preserveNullAndEmptyArrays": true,
//     }}
// ]
```

#### Unwrap

Get first item of array and insert to doc using $addFields stage. When using lookup result returns as array, use this helper to unwrap lookup result as field.

```go
// Signature:
Unwrap(field string, as string) MongoPipeline

// Example:
pipe.
    Lookup("users", "user_id", "_id", "__user").
    Unwrap("$__user", "user")
// -> [
//     { "$lookup": {
//         "from": "users",
//         "localField": "user_id",
//         "foreignField": "_id",
//         "as": "user"
//     }},
//     { "$addFields": { "user" : { "$first": "$__user" } } }
// ]
```

#### Group

Add $group stage.

```go
// Signature:
Group(cb func(d MongoDoc) MongoDoc) MongoPipeline

// Example:
pipe.
    Group(func(d mongoutils.MongoDoc) mongoutils.MongoDoc{
        d.
            Add("_id", "$_id").
            Nested("name", "$first", "$name").
            Nested("total", "$sum", "$invoice")
        return d
    })
// -> [
//   { "$group": {
//       "_id": "$_id"
//       "name": { "$first": "$name" },
//       "total": { "$sum": "$invoice" }
//   }}
// ]
```

#### ReplaceRoot

Add $replaceRoot stage.

```go
// Signature:
ReplaceRoot(v interface{}) MongoPipeline

// Example:
pipe.ReplaceRoot("$my_root")
// ->  [{ "$replaceRoot": {"newRoot": "$my_root" } }]
```

#### MergeRoot

Add $replaceRoot stage with $mergeObjects operator.

```go
// Signature:
MergeRoot(fields ...interface{}) MongoPipeline

// Example:
pipe.MergeRoot("$my_root", "$$ROOT")
// -> [
//     {
//         "$replaceRoot": {
//             "newRoot": { "mergeObjects": ["$my_root", "$$ROOT"] }
//         }
//     }
// ]
```

#### UnProject

Generate $project stage to remove fields from result.

```go
// Signature:
UnProject(fields ...string) MongoPipeline

// Example:
pipe.UnProject("my_root", "__user")
// -> [
//     { "$project": { "my_root": 0, "__user": 0 } }
// ]
```

#### Build

Generate mongo pipeline.

```go
Build() mongo.Pipeline
```
