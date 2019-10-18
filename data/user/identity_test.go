package user

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/parle-io/backend/data"
	"github.com/segmentio/ksuid"
)

func TestMain(m *testing.M) {
	if err := os.Remove("../db/test.db"); err != nil {
		fmt.Println("error reminving the test db", err)
	}

	if err := data.Open("../db/test.db"); err != nil {
		log.Fatal(err)
	}

	data.DB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucket(bucketIdentify); err != nil {
			fmt.Println("error creating bucket", err)
		}

		if _, err := tx.CreateBucket(bucketVisitor); err != nil {

		}
		return nil
	})

	retval := m.Run()
	data.Close()

	os.Exit(retval)
}

func Test_Identify_Visitor(t *testing.T) {

	id, err := Identify(getToken(), "{}")
	if err != nil {
		t.Fatal(err)
	}

	if id.IsVisitor == false {
		t.Error("IsVisitor is true")
	}

	t.Logf("connection token %s", id.ConnectionToken)

	check, err := getIdentity(id.ConnectionToken)
	if err != nil {
		t.Error(err)
	} else if check.ConnectionToken != id.ConnectionToken {
		t.Errorf("invalid connection token %s, should be %s", check.ConnectionToken, id.ConnectionToken)
	}
}

func Test_Identify_Promote_To_Lead(t *testing.T) {
	visitorIdent, err := Identify(getToken(), "{}")
	if err != nil {
		t.Fatal(err)
	}

	lead := Lead{
		Email: "unit@test.com",
	}
	lead.ConnectionToken = visitorIdent.ConnectionToken

	if _, err := Identify(visitorIdent.ConnectionToken, toJSON(lead)); err != nil {
		t.Fatal(err)
	}

	verifyID, err := getIdentity(lead.ConnectionToken)
	if err != nil {
		t.Error(err)
	} else if verifyID.IsLead == false {
		t.Error("identity not tagged as lead")
	}

	//TODO: Assert if the email was properly saved to the leads bucket
}

func getToken() string {
	return ksuid.New().String()
}

func getIdentity(token string) (Identification, error) {
	var id Identification

	err := data.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketIdentify)

		buf := b.Get([]byte(token))
		if err := data.Decode(buf, &id); err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func toJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
