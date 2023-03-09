package utils

import (
	"encoding/hex"
	"testing"
)

func TestSm2(t *testing.T) {

	priKey, pubKey, err := GenerateSm2Key()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("priKey : %s", priKey)
	t.Logf("pubKey : %s", pubKey)

	sm2, err := NewSM2(pubKey, priKey, 0)

	if err != nil {
		t.Fatal(err)
	}

	txt := "hello"
	rst, err := sm2.Encrypt([]byte(txt))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("en txt:%s", hex.EncodeToString(rst))

	de, err := sm2.Decrypt(rst)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("de txt:%s", string(de))
	if txt == string(de) {
		t.Log("OK")
	} else {
		t.Fatal("ERR")
	}
	/*
		enstr, err := Sm2Encode("123", pubKey, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("enstr: %s", enstr)
		destr, err := Sm2Decode(enstr, priKey, 0)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("destr: %s", destr)
	*/
}

func TestSm3(t *testing.T) {
	t.Log(Sm3Encode("abcde"))
}

func TestSm4(t *testing.T) {
	sm4, err := NewSM4("1234567890abcdef", "1234567890123456")
	if err != nil {
		t.Fatal(err)
	}

	txt := "hello"

	en := sm4.Encrypt([]byte(txt))

	t.Logf("sm4 encrypt text:%s", hex.EncodeToString(en))

	de := sm4.Decrypt(en)

	t.Logf("sm4 decrypt text:%s", string(de))

	if txt != string(de) {
		t.Fatal("err ")
	} else {
		t.Log("OK")
	}
}
