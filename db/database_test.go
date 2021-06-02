package db_test

import (
	"bytes"
	"github.com/RosettaFlow/Carrier-Go/db"
	"io/ioutil"
	"os"
	"testing"
)

func newTestLDB() (*db.LDBDatabase, func()) {
	dirname, err := ioutil.TempDir(os.TempDir(), "db_test_")
	if err != nil {
		panic("failed to create test file: " + err.Error())
	}
	db, err := db.NewLDBDatabase(dirname, 0, 0)
	if err != nil {
		panic("failed to create test database: " + err.Error())
	}
	return db, func() {
		db.Close()
		os.RemoveAll(dirname)
	}
}

var test_values = []string{"", "b", "1234", "\x00123\x00",}

func testPutGet(db db.Database, t *testing.T) {
	t.Parallel()

	for _, key := range test_values {
		err := db.Put([]byte(key), nil)
		if err != nil {
			t.Fatalf("put failed: %v", err)
		}
	}

	for _, key := range test_values {
		data, err := db.Get([]byte(key))
		if err != nil {
			t.Fatalf("get failed: %v", err)
		}
		if len(data) != 0 {
			t.Fatalf("get returned wrong result, got %q expected nil", string(data))
		}
	}
	_, err := db.Get([]byte("non-exist-key-test"))
	if err == nil {
		t.Fatalf("expect to return a not found error")
	}

	// put.
	for _, v := range test_values {
		err := db.Put([]byte(v), []byte(v))
		if err != nil {
			t.Fatalf("put failed: %v", err)
		}
	}
	// get
	for _, v := range test_values {
		data, err := db.Get([]byte(v))
		if err != nil {
			t.Fatalf("get failed: %v", err)
		}
		if !bytes.Equal(data, []byte(v)) {
			t.Fatalf("get returned wrong result, got %q expected %q", string(data), v)
		}
	}

	for _, v := range test_values {
		err := db.Put([]byte(v), []byte("?"))
		if err != nil {
			t.Fatalf("put override failed: %v", err)
		}
	}

	for _, v := range test_values {
		data, err := db.Get([]byte(v))
		if err != nil {
			t.Fatalf("get failed: %v", err)
		}
		if !bytes.Equal(data, []byte("?")) {
			t.Fatalf("get returned wrong result, got %q expected ?", string(data))
		}
	}

	for _, v := range test_values {
		orig, err := db.Get([]byte(v))
		if err != nil {
			t.Fatalf("get failed: %v", err)
		}
		orig[0] = byte(0xff)
		data, err := db.Get([]byte(v))
		if err != nil {
			t.Fatalf("get failed: %v", err)
		}
		if !bytes.Equal(data, []byte("?")) {
			t.Fatalf("get returned wrong result, got %q expected ?", string(data))
		}
	}

	for _, v := range test_values {
		err := db.Delete([]byte(v))
		if err != nil {
			t.Fatalf("delete %q failed: %v", v, err)
		}
	}

	for _, v := range test_values {
		_, err := db.Get([]byte(v))
		if err == nil {
			t.Fatalf("got deleted value %q", v)
		}
	}
}

func TestLDB_PutGet(t *testing.T) {
	db, remove := newTestLDB()
	defer remove()
	testPutGet(db, t)
}
