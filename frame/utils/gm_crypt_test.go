package utils

import (
	"encoding/hex"
	"testing"
)

func TestSm2ZD(t *testing.T) {

	pubKey := "048c89b39f68dd7896b90a6b2f5072795c1e9cc6be27dc6c4a671b0e412166a59ac9f4884b4ec0786df6436ea3652c2275c1ec57a7681c9c39fcfd8d4814211e03"
	priKey := "6e3b4f4be11f890a4b02d90761923256f6b7eb0d7f645c0528b0eb528228e271"
	sm2, err := NewSM2(pubKey, priKey, 1)
	if err != nil {
		t.Fatal(err)
	}
	en_bytes, err := sm2.Encrypt([]byte("1678875236.24602"))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("encrypt : %s", hex.EncodeToString(en_bytes))

	destTxt := "042BFE625BA2D19C7244405C7BF68832819A50DF2D31F72939CDCA8D99388259F2A10D7CD8F3A2E5BD7B4072B9B6BF7FBE7E5A8FEF07DB13CB1A2EE66C8799CE7DAD5A95BD5327F6A4C462C17B4B548566FDAB5CFEA033277433FDA3648BD22B9F9215A7F31CFB5694E177ACD02A98D4D7"

	de_bytes, err := hex.DecodeString(destTxt)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("en length:%d , dest length:%d", len(en_bytes), len(de_bytes))

	rst, err := sm2.Decrypt(de_bytes)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("decrpt %s", string(rst))

	// de_bytes, err := sm2.Decrypt(dest_bytes)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// t.Logf("decrypt : %s", string(de_bytes))

	// t.Logf("rst byte leng:%d", len(rst))
	// t.Logf("dest byte leng:%d", len(destLen))
	// t.Logf("rst length:%d", len(enTxt))
	// t.Logf("dest length:%d", len(destEn))

}

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
	sm4, err := NewCBCSM4([]byte("1234567890abcdef"), []byte("1234567890123456"))
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

func TestSm4ZD(t *testing.T) {

	key, err := hex.DecodeString("f83682ef8fac9e391e28f53cfd312cce")
	if err != nil {
		t.Fatal(err)
	}
	sm4, err := NewECBSM4(key)
	if err != nil {
		t.Fatal(err)
	}

	en := "208add776c6102635202c475b165b9f83fdaa27d5100cb7542bcaa99d1501d89e7c6fd8ae03100555297f297827b72fa8ee11a9cacbf8915b99959739281a65440e98f29f7ab629bbb9ab5378b755c23fde1b3e56d16d7d899a1809ef15ea14ff05d0e3aba0874e3a5d5fbb3d21d301701501e0451629be6ffdd2015c766c3a6c4f2c540b1296709e27cf4d465df38b93e3982f91b54d39dca84d16ea5b9c17a43bb79105d22f46124d97facd2af127efbabdd34270b7a9ecdbbd5656f5a193bdce06b613c9b707e107eccab814f3c6b"

	b, err := hex.DecodeString(en)

	if err != nil {
		t.Fatal(err)
	}
	de := sm4.Decrypt(b)

	t.Logf("sm4 decrypt text:%s", string(de))

}
