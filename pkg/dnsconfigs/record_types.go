// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type RecordType = string

const (
	RecordTypeA     RecordType = "A"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeNS    RecordType = "NS"
	RecordTypeMX    RecordType = "MX"
	RecordTypeSRV   RecordType = "SRV"
	RecordTypeTXT   RecordType = "TXT"
	RecordTypeCAA   RecordType = "CAA"
)

type RecordTypeDefinition struct {
	Type        RecordType `json:"type"`
	Description string     `json:"description"`
}

func FindAllRecordTypeDefinitions() []*RecordTypeDefinition {
	return []*RecordTypeDefinition{
		{
			Type:        RecordTypeA,
			Description: "将域名指向一个IPV4地址",
		},
		{
			Type:        RecordTypeCNAME,
			Description: "将域名指向另外一个域名",
		},
		{
			Type:        RecordTypeAAAA,
			Description: "将域名指向一个IPV6地址",
		},
		{
			Type:        RecordTypeNS,
			Description: "将子域名指定其他DNS服务器解析",
		},
		{
			Type:        RecordTypeMX,
			Description: "将域名指向邮件服务器地址",
		},
		{
			Type:        RecordTypeSRV,
			Description: "记录提供特定的服务的服务器",
		},
		{
			Type:        RecordTypeTXT,
			Description: "文本长度限制512，通常做SPF记录（反垃圾邮件）",
		},
		{
			Type:        RecordTypeCAA,
			Description: "CA证书颁发机构授权校验",
		},
	}
}
