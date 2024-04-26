package php

import (
	"github.com/paranoidxc/php"
	"testing"
)

func TestMd5(t *testing.T) {
	// Happy path test case
	happyStr := "hello world"
	happyExpected := "5eb63bbbe01eeed093cb22bb8f5acdc3"
	happyResult := php.Md5(happyStr)
	if happyResult != happyExpected {
		t.Errorf("Md5(%s) = %s want %s", happyStr, happyResult, happyExpected)
	}

	// Edge case: empty string
	emptyStr := ""
	emptyExpected := "d41d8cd98f00b204e9800998ecf8427e"
	emptyResult := php.Md5(emptyStr)
	if emptyResult != emptyExpected {
		t.Errorf("Md5(%s) = %s want %s", emptyStr, emptyResult, emptyExpected)
	}

	// Edge case: all space string
	spaceStr := "    "
	spaceExpected := "0cf31b2c283ce3431794586df7b0996d"
	spaceResult := php.Md5(spaceStr)
	if spaceResult != spaceExpected {
		t.Errorf("Md5(%s) = %s want %s", spaceStr, spaceResult, spaceExpected)
	}
}

func TestHash256(t *testing.T) {
	// Happy path test case
	expectedHappy := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	resultHappy := php.Hash256("hello")
	if resultHappy != expectedHappy {
		t.Errorf("Hash256('hello') = %s; want %s", resultHappy, expectedHappy)
	}

	// Edge case: one space string
	expectedEmpty := "36a9e7f1c95b82ffb99743e0c5c4ce95d83c9a430aac59f84ef3cbfab6145068"
	resultEmpty := php.Hash256(" ")
	if resultEmpty != expectedEmpty {
		t.Errorf("Hash256('') = %s; want %s", resultEmpty, expectedEmpty)
	}
}
