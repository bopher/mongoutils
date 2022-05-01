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

### FindOption

Generate find option with sorts params.

```go
FindOption(sort interface{}, skip int64, limit int64) *options.FindOptions
```

### AggregateOption

Generate aggregation options.

```go
AggregateOption() *options.AggregateOptions
```

### TxOption

Generate transaction option with majority write and snapshot read.

```go
AggregateOption() *options.AggregateOptions
```

### Array

Generate `primitive.A` from parameters.

```go
Array(args ...interface{}) primitive.A
```

### Map

Generate `primitive.M` from parameters. Parameters count must be even.

```go
// Signature:
Map(args ...interface{}) primitive.M

// Example:
mongoutils.Map("name", "John", "age", 23) // { "name": "John", "age": 23 }
```

### Maps

Generate `[]primitive.M` from parameters. Parameters count must be even.

```go
// Signature:
Maps(args ...interface{}) []primitive.M

// Example:
mongoutils.Maps("name", "John", "age", 23) // [{ "name": "John" }, { "age": 23 }]
```

### Doc

Generate primitive.D from parameters. Parameters count must be even.

```go
// Signature:
Doc(args ...interface{}) primitive.D

// Example:
mongoutils.Doc("name", "John", "age", 23) // { "name": "John", "age": 23 }
```

### Regex

Generate mongo `Regex` doc.

```go
// Signature:
Regex(pattern string, opt string) primitive.Regex

// Example:
mongoutils.Regex("John.*", "i") // { pattern: "John.*", options: "i" }
```

### RegexFor

Generate map with regex parameter.

```go
// Signature
RegexFor(k string, pattern string, opt string) primitive.M

// Example:
mongoutils.RegexFor("name", "John.*", "i") // { "name": { pattern: "John.*", options: "i" } }
```

### In

Generate $in map `{k: {$in: v}}`.

```go
In(k string, v ...interface{}) primitive.M
```

### Set

Generate simple set map  `{$set: v}`.

```go
Set(v interface{}) primitive.M
```

### SetNested

Generate nested set map `{$set: {k: v}}`.

```go
SetNested(k string, v interface{}) primitive.M
```

### Match

Generate nested set map `{$match: v}`.

```go
Match(v interface{}) primitive.M
```

## Base Model

Base model can used as modal parent. This model comes with timestamp and utility functions.

**Note**: Set model bson tag to `inline` for insert timestamps in document root.

```go
// Usage:
import "github.com/bopher/mongoutils"
type Person struct{
 mongoutils.Model  `bson:",inline"`
}

// override methods
func (me Person) IsDeletable() bool{
    return true
}
```

### Available methods

```go
// IsEditable check if document is editable
//
// by default returns true
func (me Model) IsEditable() bool {}

// IsDeletable check if document is deletable
//
// by default returns false
func (me Model) IsDeletable() bool {}

// BeforeInsert function to call before insert
func (me *Model) BeforeInsert() {}

// AfterInsert function to call after insert
func (me Model) AfterInsert(ctx context.Context) {}

// BeforeUpdate function to call before update
func (me *Model) BeforeUpdate() {}

// AfterUpdate function to call after update
func (me Model) AfterUpdate(ctx context.Context) {}

// BeforeDelete function to call before delete
func (me *Model) BeforeDelete() {}

// AfterDelete function to call after delete
func (me Model) AfterDelete(ctx context.Context) {}

// Cleanup document before save
//
// e.g set relation document to nil for ignore saving
func (me *Model) Cleanup() {}
```

### Required Methods

Two `PrepareInsert` and `PrepareUpdate` must called before save model to database.

**Note**: if `true` passed to `PrepareUpdate` method, `updated_at` method not updated.

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

#### DocArray

Add new array element with doc child.

```go
// Signature:
DocArray(k string, cb func(d MongoDoc) MongoDoc) MongoDoc

// Example:
doc.DocArray("$match", func(d mongoutils.MongoDoc) mongoutils.MongoDoc {
    return d.Add("name", "John")
            Add("Family", "Doe")
}) // -> { "$match": [{"name": "John"}, {"Family": "Doe"}] }
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

#### NestedDocArray

Add new nested array element with doc

```go
// Signature:
NestedDocArray(root string, k string, cb func(d MongoDoc) MongoDoc) MongoDoc

// Example:
doc.NestedDocArray("name", "$match", func(d mongoutils.MongoDoc) mongoutils.MongoDoc {
    return d.Add("first", "John")
            Add("last", "Doe")
}) // -> { "name" : {"$match": [{"name": "John"}, {"last": "Doe"}] } }
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

### Match

Add $match stage.

```go
// Signature:
Match(filters interface{}) MongoPipeline

// Example:
pipe.Match(v)
```

### In

Add $in stage.

```go
// Signature:
In(key string, v interface{}) MongoPipeline

// Example:
pipe.In("status", statuses)
```

### Limit

Add $limit stage (ignore negative and zero value).

```go
// Signature:
Limit(limit int64) MongoPipeline

// Example:
pipe.Limit(100)
```

### Skip

Add $skip stage (ignore negative and zero value).

```go
// Signature:
Skip(skip int64) MongoPipeline

// Example:
pipe.Skip(25)
```

### Sort

Add $sort stage (ignore nil value).

```go
// Signature:
Sort(sorts interface{}) MongoPipeline

// Example:
pipe.Sort(primitive.M{"username": 1})
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

#### Unwrap

Get first item of array and insert to doc using $addFields stage. When using lookup result returns as array, use me helper to unwrap lookup result as field.

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

### LoadRelation

Load related document using `$lookup` and `$addField` (Lookup and Unwrap method mix).

```go
// Signature:
LoadRelation(from string, local string, foreign string, as string) MongoPipeline

// Example:
pipe.LoadRelation("users", "user_id", "_id", "user")
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
