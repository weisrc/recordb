package recordb_test

import (
	"testing"

	"github.com/weisrc/recordb"
	"github.com/weisrc/recordb/types"
)

func TestDB(t *testing.T) {
	db := recordb.New()
	db.Add(recordb.Record{
		Name:  "google.com.",
		Type:  types.A,
		Value: "hello",
	})
	db.Add(recordb.Record{
		Name:  "google.com.",
		Type:  types.AAAA,
		Value: "hello",
	})
	db.Add(recordb.Record{
		Name:  "google.com.",
		Type:  types.CNAME,
		Value: "hello",
	})
	if got, ok := db.Get(types.A, "google.com.", false); !ok || got.Value != "hello" {
		t.Errorf("something went wrong")
	}
	for v := db.List("google.com."); v != nil; v = v.Next {
		println(types.String(v.Type))
	}
	println("==========================")
	db.Remove(types.CNAME, "google.com.")
	for v := db.List("google.com."); v != nil; v = v.Next {
		println(types.String(v.Type))
	}
}
