package client

import (
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	"chainmaker.org/chainmaker/common/v2/crypto"
	bcx509 "chainmaker.org/chainmaker/common/v2/crypto/x509"
	"chainmaker.org/chainmaker/common/v2/evmutils"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
)

const (
	OrgId1 = "wx-org1.chainmaker.org"
	OrgId2 = "wx-org2.chainmaker.org"
	OrgId3 = "wx-org3.chainmaker.org"
	OrgId4 = "wx-org4.chainmaker.org"
	OrgId5 = "wx-org5.chainmaker.org"

	UserNameOrg1Admin1 = "cmtestuser1"
	UserNameOrg2Admin1 = "cmtestuser2"
	UserNameOrg3Admin1 = "cmtestuser3"
	UserNameOrg4Admin1 = "cmtestuser4"

	Version = "1.0.0"
)

var users = map[string]*User{
	"cmtestuser1": {
		"./crypto-config/TestCMorg1/user/cmtestuser1/cmtestuser1.tls.key",
		"./crypto-config/TestCMorg1/user/cmtestuser1/cmtestuser1.tls.crt",
		"./crypto-config/TestCMorg1/user/cmtestuser1/cmtestuser1.sign.key",
		"./crypto-config/TestCMorg1/user/cmtestuser1/cmtestuser1.sign.crt",
	},
	"cmtestuser2": {
		"./crypto-config/TestCMorg2/user/cmtestuser2/cmtestuser2.tls.key",
		"./crypto-config/TestCMorg2/user/cmtestuser2/cmtestuser2.tls.crt",
		"./crypto-config/TestCMorg2/user/cmtestuser2/cmtestuser2.sign.key",
		"./crypto-config/TestCMorg2/user/cmtestuser2/cmtestuser2.sign.crt",
	},
	"cmtestuser3": {
		"./crypto-config/TestCMorg3/user/cmtestuser3/cmtestuser3.tls.key",
		"./crypto-config/TestCMorg3/user/cmtestuser3/cmtestuser3.tls.crt",
		"./crypto-config/TestCMorg3/user/cmtestuser3/cmtestuser3.sign.key",
		"./crypto-config/TestCMorg3/user/cmtestuser3/cmtestuser3.sign.crt",
	},
	"cmtestuser4": {
		"./crypto-config/TestCMorg4/user/cmtestuser4/cmtestuser4.tls.key",
		"./crypto-config/TestCMorg4/user/cmtestuser4/cmtestuser4.tls.crt",
		"./crypto-config/TestCMorg4/user/cmtestuser4/cmtestuser4.sign.key",
		"./crypto-config/TestCMorg4/user/cmtestuser4/cmtestuser4.sign.crt",
	},
}

type User struct {
	TlsKeyPath, TlsCrtPath   string
	SignKeyPath, SignCrtPath string
}

func GetUser(username string) (*User, error) {
	u, ok := users[username]
	if !ok {
		return nil, errors.New("user not found")
	}

	return u, nil
}

func MakeAddrAndSkiFromCrtFilePath(crtFilePath string) (addressInt string, addressEvm string, addressSki string, err error) {
	crtBytes, err := os.ReadFile(crtFilePath)
	if err != nil {
		return "", "", "", err
	}

	blockCrt, _ := pem.Decode(crtBytes)
	crt, err := bcx509.ParseCertificate(blockCrt.Bytes)
	if err != nil {
		return "", "", "", err
	}

	ski := hex.EncodeToString(crt.SubjectKeyId)
	addrInt, err := evmutils.MakeAddressFromHex(ski)
	if err != nil {
		return "", "", "", err
	}

	return addrInt.String(), fmt.Sprintf("0x%x", addrInt.AsStringKey()), ski, nil
}

func GetEndorsersWithAuthType(hashType crypto.HashType, authType sdk.AuthType, payload *common.Payload, usernames ...string) ([]*common.EndorsementEntry, error) {
	var endorsers []*common.EndorsementEntry

	for _, name := range usernames {
		var entry *common.EndorsementEntry
		var err error
		switch authType {
		case sdk.PermissionedWithCert:
			u, ok := users[name]
			if !ok {
				return nil, errors.New("user not found")
			}
			entry, err = sdkutils.MakeEndorserWithPath(u.SignKeyPath, u.SignCrtPath, payload)
			if err != nil {
				return nil, err
			}

			// case sdk.PermissionedWithKey:
			// 	u, ok := permissionedPkUsers[name]
			// 	if !ok {
			// 		return nil, errors.New("user not found")
			// 	}
			// 	entry, err = sdkutils.MakePkEndorserWithPath(u.SignKeyPath, hashType, u.OrgId, payload)
			// 	if err != nil {
			// 		return nil, err
			// 	}

			// case sdk.Public:
			// u, ok := pkUsers[name]
			// if !ok {
			// 	return nil, errors.New("user not found")
			// }
			// entry, err = sdkutils.MakePkEndorserWithPath(u.SignKeyPath, hashType, "", payload)
			// if err != nil {
			// 	return nil, err
			// }

		default:
			return nil, errors.New("invalid authType")
		}
		endorsers = append(endorsers, entry)
	}

	return endorsers, nil
}
