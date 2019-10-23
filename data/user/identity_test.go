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

		if _, err := tx.CreateBucket(bucketLead); err != nil {

		}

		if _, err := tx.CreateBucket(bucketUser); err != nil {

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

	leadIdent, err := Identify(visitorIdent.ConnectionToken, toJSON(lead))
	if err != nil {
		t.Fatal(err)
	}

	verifyID, err := getIdentity(lead.ConnectionToken)
	if err != nil {
		t.Error(err)
	} else if verifyID.IsLead == false {
		t.Error("identity not tagged as lead")
	}

	// is the visitor removed from the visitors bucket
	v, err := getVisitor(visitorIdent.ReferenceID)
	if err != nil {
		t.Fatal(err)
	} else if v != nil {
		t.Errorf("we found the visitor %d but should be removed", visitorIdent.ReferenceID)
	}

	l, err := getLead(leadIdent.ReferenceID)
	if err != nil {
		t.Fatal(err)
	} else if l == nil {
		t.Errorf("we could not found the lead %d in the leads bucket", leadIdent.ReferenceID)
	} else if l.Email != lead.Email {
		t.Errorf("expected email to be %s got %s", lead.Email, l.Email)
	}
}

func Test_Identify_Promote_To_User(t *testing.T) {
	visitorIdent, err := Identify(getToken(), "{}")
	if err != nil {
		t.Fatal(err)
	}

	user := User{
		YourID: getToken(),
	}
	user.ConnectionToken = visitorIdent.ConnectionToken
	user.Email = "bob@test.com"

	userIdent, err := Identify(visitorIdent.ConnectionToken, toJSON(user))
	if err != nil {
		t.Fatal(err)
	}

	verifyID, err := getIdentity(user.ConnectionToken)
	if err != nil {
		t.Error(err)
	} else if verifyID.IsUser == false {
		t.Fatal("identity not tagged as user")
	}

	v, err := getVisitor(visitorIdent.ReferenceID)
	if err != nil {
		t.Fatal(err)
	} else if v != nil {
		t.Errorf("we found the visitor %d but should be removed", visitorIdent.ReferenceID)
	}

	u, err := getUser(userIdent.ReferenceID)
	if err != nil {
		t.Fatal(err)
	} else if u == nil {
		t.Errorf("we could not found the user %d in the users bucket", userIdent.ReferenceID)
	} else if u.YourID != user.YourID {
		t.Errorf("expected YourID to be %s got %s", user.YourID, u.YourID)
	}
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
