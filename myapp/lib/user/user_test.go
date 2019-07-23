package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserWithHOC(t *testing.T){
	user := &User{
		FName: "New",
		LName: "Knocking",
	}

	assert := assert.New(t)
	assert.Equal("new knocking", user.GetName(withLower()), "Account owner is new knocking")
	// assert.Equal("new ...", user.GetName(withAbvr(4)), "Account owner is new ...")
	
}
