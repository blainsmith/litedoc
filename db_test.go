package litedoc_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/blainsmith/litedoc"
)

type numbers struct {
	Type   string
	Digits string
}

type doc struct {
	Name    string
	Age     int
	Dead    bool
	Numbers []numbers
}

func TestDocument(t *testing.T) {
	ctx := context.Background()

	db, _ := litedoc.Open("./test.db")
	defer os.Remove("./test.db")

	d1 := doc{Name: "Blain Smith", Age: 40, Dead: false, Numbers: []numbers{{Type: "home", Digits: "9784305790"}, {Type: "mobile", Digits: "9784305790"}}}

	err := db.Collection("test").Document("my-doc").Create(ctx, &d1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(d1)

	err = db.Collection("test").Document("my-doc-1").Create(ctx, &d1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(d1)

	var d2 doc
	err = db.Collection("test").Document("my-doc").Get(ctx, &d2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(d2)

	if d1.Name != d2.Name {
		t.Error(".Name mismatch")
	}
	if d1.Age != d2.Age {
		t.Error(".Age mismatch")
	}
	if d1.Dead != d2.Dead {
		t.Error(".Dead mismatch")
	}

	docs, err := db.Collection("test").Query(ctx, "$.Numbers[0].Type", litedoc.OpEqual, "home")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(len(docs))
	for _, doc := range docs {
		var d map[string]any
		doc.DataTo(&d)
		fmt.Println(doc.ID, d)
	}

	if len(docs) < 5 {
		t.Error("not enough docs")
	}
}
