// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type KeyAlgorithmType = string

const (
	KeyAlgorithmTypeHmacSHA1   KeyAlgorithmType = "hmac-sha1."
	KeyAlgorithmTypeHmacSHA224 KeyAlgorithmType = "hmac-sha224."
	KeyAlgorithmTypeHmacSHA256 KeyAlgorithmType = "hmac-sha256."
	KeyAlgorithmTypeHmacSHA384 KeyAlgorithmType = "hmac-sha384."
	KeyAlgorithmTypeHmacSHA512 KeyAlgorithmType = "hmac-sha512."
)

type KeyAlgorithmDefinition struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func FindAllKeyAlgorithmTypes() []*KeyAlgorithmDefinition {
	return []*KeyAlgorithmDefinition{
		{
			Name: "HmacSHA1",
			Code: KeyAlgorithmTypeHmacSHA1,
		},
		{
			Name: "HmacSHA224",
			Code: KeyAlgorithmTypeHmacSHA224,
		},
		{
			Name: "HmacSHA256",
			Code: KeyAlgorithmTypeHmacSHA256,
		},
		{
			Name: "HmacSHA384",
			Code: KeyAlgorithmTypeHmacSHA384,
		},
		{
			Name: "HmacSHA512",
			Code: KeyAlgorithmTypeHmacSHA512,
		},
	}
}

func FindKeyAlgorithmTypeName(algoType KeyAlgorithmType) string {
	for _, def := range FindAllKeyAlgorithmTypes() {
		if def.Code == algoType {
			return def.Name
		}
	}
	return ""
}

type NSKeySecretType = string

const (
	NSKeySecretTypeClear  NSKeySecretType = "clear"
	NSKeySecretTypeBase64 NSKeySecretType = "base64"
)

func FindKeySecretTypeName(secretType NSKeySecretType) string {
	switch secretType {
	case NSKeySecretTypeClear:
		return "明文"
	case NSKeySecretTypeBase64:
		return "BASE64"
	}
	return ""
}
