package mongoutils_test

import (
	"testing"

	"github.com/bopher/mongoutils"
)

func TestPipeline(t *testing.T) {
	var v string
	var err error

	// Add
	v, err = pretty(mongoutils.NewPipe().Add(func(d mongoutils.MongoDoc) mongoutils.MongoDoc {
		return d.Nested("$match", "name", "Jack")
	}).Build())
	if err != nil {
		t.Fatal(err)
	}
	if v != `[[{"Key":"$match","Value":{"name":"Jack"}}]]` {
		t.Log(v)
		t.Fatal("fail Add")
	}

	// Lookup
	v, err = pretty(mongoutils.NewPipe().Lookup("users", "user_id", "_id", "user").Build())
	if err != nil {
		t.Fatal(err)
	}
	if v != `[[{"Key":"$lookup","Value":[{"Key":"from","Value":"users"},{"Key":"localField","Value":"user_id"},{"Key":"foreignField","Value":"_id"},{"Key":"as","Value":"user"}]}]]` {
		t.Log(v)
		t.Fatal("fail Lookup")
	}

	// Unwind
	v, err = pretty(mongoutils.NewPipe().Unwind("services", true).Build())
	if err != nil {
		t.Fatal(err)
	}
	if v != `[[{"Key":"$unwind","Value":[{"Key":"path","Value":"services"},{"Key":"preserveNullAndEmptyArrays","Value":true}]}]]` {
		t.Log(v)
		t.Fatal("fail Unwind")
	}

	// Unwrap
	v, err = pretty(mongoutils.NewPipe().Unwrap("_user", "user").Build())
	if err != nil {
		t.Fatal(err)
	}
	if v != `[[{"Key":"$addFields","Value":[{"Key":"user","Value":{"$first":"_user"}}]}]]` {
		t.Log(v)
		t.Fatal("fail Unwrap")
	}

	// Group
	v, err = pretty(mongoutils.NewPipe().Group(func(d mongoutils.MongoDoc) mongoutils.MongoDoc {
		d.
			Add("_id", "$_id").
			Nested("name", "$first", "$name").
			Nested("total", "$sum", "$invoice")
		return d
	}).Build())
	if err != nil {
		t.Fatal(err)
	}
	if v != `[[{"Key":"$group","Value":[{"Key":"_id","Value":"$_id"},{"Key":"name","Value":{"$first":"$name"}},{"Key":"total","Value":{"$sum":"$invoice"}}]}]]` {
		t.Log(v)
		t.Fatal("fail Group")
	}

	// ReplaceRoot
	v, err = pretty(mongoutils.NewPipe().ReplaceRoot("$my_root").Build())
	if err != nil {
		t.Fatal(err)
	}
	if v != `[[{"Key":"$replaceRoot","Value":[{"Key":"newRoot","Value":"$my_root"}]}]]` {
		t.Log(v)
		t.Fatal("fail ReplaceRoot")
	}

	// MergeRoot
	v, err = pretty(mongoutils.NewPipe().MergeRoot("$my_root", "$$ROOT").Build())
	if err != nil {
		t.Fatal(err)
	}
	if v != `[[{"Key":"$replaceRoot","Value":[{"Key":"newRoot","Value":{"$mergeObjects":["$my_root","$$ROOT"]}}]}]]` {
		t.Log(v)
		t.Fatal("fail MergeRoot")
	}

	// UnProject
	v, err = pretty(mongoutils.NewPipe().UnProject("my_root", "__user").Build())
	if err != nil {
		t.Fatal(err)
	}
	if v != `[[{"Key":"$project","Value":[{"Key":"my_root","Value":0},{"Key":"__user","Value":0}]}]]` {
		t.Log(v)
		t.Fatal("fail UnProject")
	}
}
