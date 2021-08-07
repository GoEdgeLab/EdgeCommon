// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package dnsconfigs

type Route struct {
	Name       string   `json:"name"`
	Code       string   `json:"code"`
	AliasNames []string `json:"aliasNames"`
}

// AllDefaultISPRoutes 运营商线路
var AllDefaultISPRoutes = []*Route{
	{
		Name:       "电信",
		Code:       "isp:china_telecom",
		AliasNames: []string{"电信"},
	},
	{
		Name:       "联通",
		Code:       "isp:china_unicom",
		AliasNames: []string{"联通"},
	},
	{
		Name:       "移动",
		Code:       "isp:china_mobile",
		AliasNames: []string{"移动"},
	},
	{
		Name:       "鹏博士",
		Code:       "isp:china_drpeng",
		AliasNames: []string{"鹏博士"},
	},
	{
		Name:       "教育网",
		Code:       "isp:china_edu",
		AliasNames: []string{"教育网"},
	},
	{
		Name:       "歌华",
		Code:       "isp:china_gehua",
		AliasNames: []string{"歌华"},
	},
	{
		Name:       "铁通",
		Code:       "isp:china_tietong",
		AliasNames: []string{"铁通"},
	},
	{
		Name:       "阿里巴巴",
		Code:       "isp:alibaba",
		AliasNames: []string{"阿里巴巴"},
	},
	{
		Name:       "亚马逊",
		Code:       "isp:amazon",
		AliasNames: []string{"亚马逊"},
	},
	{
		Name:       "谷歌",
		Code:       "isp:google",
		AliasNames: []string{"谷歌"},
	},
	{
		Name:       "微软",
		Code:       "isp:microsoft",
		AliasNames: []string{"微软"},
	},
	{
		Name:       "腾讯",
		Code:       "isp:tencent",
		AliasNames: []string{"腾讯"},
	},
	{
		Name:       "阿里云",
		Code:       "isp:aliyun",
		AliasNames: []string{"阿里云"},
	},
}

// AllDefaultChinaProvinceRoutes 中国地域线路
var AllDefaultChinaProvinceRoutes = []*Route{
	{
		Name:       "北京市",
		Code:       "china:province:beijing",
		AliasNames: []string{"北京市", "北京"},
	},
	{
		Name:       "天津市",
		Code:       "china:province:tianjin",
		AliasNames: []string{"天津市", "天津"},
	},
	{
		Name:       "河北省",
		Code:       "china:province:heibei",
		AliasNames: []string{"河北省"},
	},
	{
		Name:       "山西省",
		Code:       "china:province:shanxi",
		AliasNames: []string{"山西省"},
	},
	{
		Name:       "内蒙古自治区",
		Code:       "china:province:neimenggu",
		AliasNames: []string{"内蒙古自治区", "内蒙古"},
	},
	{
		Name:       "辽宁省",
		Code:       "china:province:liaoning",
		AliasNames: []string{"辽宁省"},
	},
	{
		Name:       "吉林省",
		Code:       "china:jilin",
		AliasNames: []string{"吉林省"},
	},
	{
		Name:       "黑龙江省",
		Code:       "china:province:heilongjiang",
		AliasNames: []string{"黑龙江省"},
	},
	{
		Name:       "上海市",
		Code:       "china:province:shanghai",
		AliasNames: []string{"上海市", "上海"},
	},
	{
		Name:       "江苏省",
		Code:       "china:province:jiangsu",
		AliasNames: []string{"江苏省"},
	},
	{
		Name:       "浙江省",
		Code:       "china:province:zhejiang",
		AliasNames: []string{"浙江省"},
	},
	{
		Name:       "安徽省",
		Code:       "china:province:anhui",
		AliasNames: []string{"安徽省"},
	},
	{
		Name:       "福建省",
		Code:       "china:province:fujian",
		AliasNames: []string{"福建省"},
	},
	{
		Name:       "江西省",
		Code:       "china:province:jiangxi",
		AliasNames: []string{"江西省"},
	},
	{
		Name:       "山东省",
		Code:       "china:province:shandong",
		AliasNames: []string{"山东省"},
	},
	{
		Name:       "河南省",
		Code:       "china:province:henan",
		AliasNames: []string{"河南省"},
	},
	{
		Name:       "湖北省",
		Code:       "china:province:hubei",
		AliasNames: []string{"湖北省"},
	},
	{
		Name:       "湖南省",
		Code:       "china:province:hunan",
		AliasNames: []string{"湖南省"},
	},
	{
		Name:       "广东省",
		Code:       "china:province:guangdong",
		AliasNames: []string{"广东省"},
	},
	{
		Name:       "广西壮族自治区",
		Code:       "china:province:guangxi",
		AliasNames: []string{"广西壮族自治区", "广西"},
	},
	{
		Name:       "海南省",
		Code:       "china:province:hainan",
		AliasNames: []string{"海南省"},
	},
	{
		Name:       "重庆市",
		Code:       "china:province:chongqing",
		AliasNames: []string{"重庆市", "重庆"},
	},
	{
		Name:       "四川省",
		Code:       "china:province:sichuan",
		AliasNames: []string{"四川省"},
	},
	{
		Name:       "贵州省",
		Code:       "china:province:guizhou",
		AliasNames: []string{"贵州省"},
	},
	{
		Name:       "云南省",
		Code:       "china:province:yunnan",
		AliasNames: []string{"云南省"},
	},
	{
		Name:       "西藏自治区",
		Code:       "china:province:xizang",
		AliasNames: []string{"西藏自治区", "西藏"},
	},
	{
		Name:       "陕西省",
		Code:       "china:province:shaanxi",
		AliasNames: []string{"陕西省"},
	},
	{
		Name:       "甘肃省",
		Code:       "china:province:gansu",
		AliasNames: []string{"甘肃省"},
	},
	{
		Name:       "青海省",
		Code:       "china:province:qinghai",
		AliasNames: []string{"青海省"},
	},
	{
		Name:       "宁夏回族自治区",
		Code:       "china:province:ningxia",
		AliasNames: []string{"宁夏回族自治区", "宁夏"},
	},
	{
		Name:       "新疆维吾尔自治区",
		Code:       "china:province:xinjiang",
		AliasNames: []string{"新疆维吾尔自治区", "新疆"},
	},
	{
		Name:       "香港特别行政区",
		Code:       "china:province:hk",
		AliasNames: []string{"香港特别行政区", "香港"},
	},
	{
		Name:       "澳门特别行政区",
		Code:       "china:province:mo",
		AliasNames: []string{"澳门特别行政区", "澳门"},
	},
	{
		Name:       "台湾省",
		Code:       "china:province:tw",
		AliasNames: []string{"台湾省"},
	},
}

// AllDefaultWorldRegionRoutes 世界地域线路
// 参考：https://zh.wikipedia.org/wiki/%E5%9C%8B%E5%AE%B6%E5%9C%B0%E5%8D%80%E4%BB%A3%E7%A2%BC
var AllDefaultWorldRegionRoutes = []*Route{
	{
		Name:       "中国",
		Code:       "world:CN",
		AliasNames: []string{"中国"},
	},
	{
		Name:       "蒙古",
		Code:       "world:MN",
		AliasNames: []string{"蒙古"},
	},
	{
		Name:       "朝鲜",
		Code:       "world:KP",
		AliasNames: []string{"朝鲜"},
	},
	{
		Name:       "韩国",
		Code:       "world:KR",
		AliasNames: []string{"韩国"},
	},
	{
		Name:       "日本",
		Code:       "world:JP",
		AliasNames: []string{"日本"},
	},
	{
		Name:       "菲律宾",
		Code:       "world:PH",
		AliasNames: []string{"菲律宾"},
	},
	{
		Name:       "越南",
		Code:       "world:VN",
		AliasNames: []string{"越南"},
	},
	{
		Name:       "老挝",
		Code:       "world:LA",
		AliasNames: []string{"老挝"},
	},
	{
		Name:       "柬埔寨",
		Code:       "world:KH",
		AliasNames: []string{"柬埔寨"},
	},
	{
		Name:       "缅甸",
		Code:       "world:MM",
		AliasNames: []string{"缅甸"},
	},
	{
		Name:       "泰国",
		Code:       "world:TH",
		AliasNames: []string{"泰国"},
	},
	{
		Name:       "马来西亚",
		Code:       "world:MY",
		AliasNames: []string{"马来西亚"},
	},
	{
		Name:       "文莱",
		Code:       "world:BN",
		AliasNames: []string{"文莱"},
	},
	{
		Name:       "新加坡",
		Code:       "world:SG",
		AliasNames: []string{"新加坡"},
	},
	{
		Name:       "印度尼西亚",
		Code:       "world:ID",
		AliasNames: []string{"印度尼西亚"},
	},
	{
		Name:       "东帝汶",
		Code:       "world:TL",
		AliasNames: []string{"东帝汶"},
	},
	{
		Name:       "尼泊尔",
		Code:       "world:NP",
		AliasNames: []string{"尼泊尔"},
	},
	{
		Name:       "不丹",
		Code:       "world:BT",
		AliasNames: []string{"不丹"},
	},
	{
		Name:       "孟加拉国",
		Code:       "world:BD",
		AliasNames: []string{"孟加拉国", "孟加拉"},
	},
	{
		Name:       "印度",
		Code:       "world:IN",
		AliasNames: []string{"印度"},
	},
	{
		Name:       "巴基斯坦",
		Code:       "world:PK",
		AliasNames: []string{"巴基斯坦"},
	},
	{
		Name:       "斯里兰卡",
		Code:       "world:LK",
		AliasNames: []string{"斯里兰卡"},
	},
	{
		Name:       "马尔代夫",
		Code:       "world:MV",
		AliasNames: []string{"马尔代夫"},
	},
	{
		Name:       "哈萨克斯坦",
		Code:       "world:KZ",
		AliasNames: []string{"哈萨克斯坦"},
	},
	{
		Name:       "吉尔吉斯斯坦",
		Code:       "world:KG",
		AliasNames: []string{"吉尔吉斯斯坦"},
	},
	{
		Name:       "塔吉克斯坦",
		Code:       "world:TJ",
		AliasNames: []string{"塔吉克斯坦"},
	},
	{
		Name:       "乌兹别克斯坦",
		Code:       "world:UZ",
		AliasNames: []string{"乌兹别克斯坦"},
	},
	{
		Name:       "土库曼斯坦",
		Code:       "world:TM",
		AliasNames: []string{"土库曼斯坦"},
	},
	{
		Name:       "阿富汗",
		Code:       "world:AF",
		AliasNames: []string{"阿富汗"},
	},
	{
		Name:       "伊拉克",
		Code:       "world:IQ",
		AliasNames: []string{"伊拉克"},
	},
	{
		Name:       "伊朗",
		Code:       "world:IR",
		AliasNames: []string{"伊朗"},
	},
	{
		Name:       "叙利亚",
		Code:       "world:SY",
		AliasNames: []string{"叙利亚"},
	},
	{
		Name:       "约旦",
		Code:       "world:JO",
		AliasNames: []string{"约旦"},
	},
	{
		Name:       "黎巴嫩",
		Code:       "world:LB",
		AliasNames: []string{"黎巴嫩"},
	},
	{
		Name:       "以色列",
		Code:       "world:IL",
		AliasNames: []string{"以色列"},
	},
	{
		Name:       "巴勒斯坦",
		Code:       "world:PS",
		AliasNames: []string{"巴勒斯坦"},
	},
	{
		Name:       "沙特阿拉伯",
		Code:       "world:SA",
		AliasNames: []string{"沙特阿拉伯"},
	},
	{
		Name:       "巴林",
		Code:       "world:BH",
		AliasNames: []string{"巴林"},
	},
	{
		Name:       "卡塔尔",
		Code:       "world:QA",
		AliasNames: []string{"卡塔尔"},
	},
	{
		Name:       "科威特",
		Code:       "world:KW",
		AliasNames: []string{"科威特"},
	},
	{
		Name:       "阿拉伯联合酋长国",
		Code:       "world:UAR",
		AliasNames: []string{"阿拉伯联合酋长国", "阿联酋"},
	},
	{
		Name:       "阿曼",
		Code:       "world:OM",
		AliasNames: []string{"阿曼"},
	},
	{
		Name:       "也门",
		Code:       "world:YE",
		AliasNames: []string{"也门"},
	},
	{
		Name:       "格鲁吉亚",
		Code:       "world:GE",
		AliasNames: []string{"格鲁吉亚"},
	},
	{
		Name:       "亚美尼亚",
		Code:       "world:AM",
		AliasNames: []string{"亚美尼亚"},
	},
	{
		Name:       "阿塞拜疆",
		Code:       "world:AZ",
		AliasNames: []string{"阿塞拜疆"},
	},
	{
		Name:       "土耳其",
		Code:       "world:TR",
		AliasNames: []string{"土耳其"},
	},
	{
		Name:       "塞浦路斯",
		Code:       "world:CY",
		AliasNames: []string{"塞浦路斯"},
	},
	{
		Name:       "芬兰",
		Code:       "world:FI",
		AliasNames: []string{"芬兰"},
	},
	{
		Name:       "瑞典",
		Code:       "world:SE",
		AliasNames: []string{"瑞典"},
	},
	{
		Name:       "挪威",
		Code:       "world:NO",
		AliasNames: []string{"挪威"},
	},
	{
		Name:       "冰岛",
		Code:       "world:IS",
		AliasNames: []string{"冰岛"},
	},
	{
		Name:       "丹麦",
		Code:       "world:DK",
		AliasNames: []string{"丹麦"},
	},
	{
		Name:       "爱沙尼亚",
		Code:       "world:EE",
		AliasNames: []string{"爱沙尼亚"},
	},
	{
		Name:       "拉脱维亚",
		Code:       "world:LV",
		AliasNames: []string{"拉脱维亚"},
	},
	{
		Name:       "立陶宛",
		Code:       "world:LT",
		AliasNames: []string{"立陶宛"},
	},
	{
		Name:       "白俄罗斯",
		Code:       "world:BY",
		AliasNames: []string{"白俄罗斯"},
	},
	{
		Name:       "俄罗斯",
		Code:       "world:RU",
		AliasNames: []string{"俄罗斯"},
	},
	{
		Name:       "乌克兰",
		Code:       "world:UA",
		AliasNames: []string{"乌克兰"},
	},
	{
		Name:       "摩尔多瓦",
		Code:       "world:MD",
		AliasNames: []string{"摩尔多瓦"},
	},
	{
		Name:       "波兰",
		Code:       "world:PL",
		AliasNames: []string{"波兰"},
	},
	{
		Name:       "捷克",
		Code:       "world:CZ",
		AliasNames: []string{"捷克"},
	},
	{
		Name:       "斯洛伐克",
		Code:       "world:SK",
		AliasNames: []string{"斯洛伐克"},
	},
	{
		Name:       "匈牙利",
		Code:       "world:HU",
		AliasNames: []string{"匈牙利"},
	},
	{
		Name:       "德国",
		Code:       "world:DE",
		AliasNames: []string{"德国"},
	},
	{
		Name:       "奥地利",
		Code:       "world:AT",
		AliasNames: []string{"奥地利"},
	},
	{
		Name:       "瑞士",
		Code:       "world:CH",
		AliasNames: []string{"瑞士"},
	},
	{
		Name:       "列支敦士登",
		Code:       "world:LI",
		AliasNames: []string{"列支敦士登"},
	},
	{
		Name:       "英国",
		Code:       "world:GB",
		AliasNames: []string{"英国"},
	},
	{
		Name:       "爱尔兰",
		Code:       "world:IE",
		AliasNames: []string{"爱尔兰"},
	},
	{
		Name:       "荷兰",
		Code:       "world:NL",
		AliasNames: []string{"荷兰"},
	},
	{
		Name:       "比利时",
		Code:       "world:BE",
		AliasNames: []string{"比利时"},
	},
	{
		Name:       "卢森堡",
		Code:       "world:LU",
		AliasNames: []string{"卢森堡"},
	},
	{
		Name:       "法国",
		Code:       "world:FR",
		AliasNames: []string{"法国"},
	},
	{
		Name:       "摩纳哥",
		Code:       "world:MC",
		AliasNames: []string{"摩纳哥"},
	},
	{
		Name:       "罗马尼亚",
		Code:       "world:RO",
		AliasNames: []string{"罗马尼亚"},
	},
	{
		Name:       "保加利亚",
		Code:       "world:BG",
		AliasNames: []string{"保加利亚"},
	},
	{
		Name:       "塞尔维亚",
		Code:       "world:RS",
		AliasNames: []string{"塞尔维亚"},
	},
	{
		Name:       "马其顿",
		Code:       "world:MK",
		AliasNames: []string{"马其顿"},
	},
	{
		Name:       "阿尔巴尼亚",
		Code:       "world:AL",
		AliasNames: []string{"阿尔巴尼亚"},
	},
	{
		Name:       "希腊",
		Code:       "world:GR",
		AliasNames: []string{"希腊"},
	},
	{
		Name:       "斯洛文尼亚",
		Code:       "world:SI",
		AliasNames: []string{"斯洛文尼亚"},
	},
	{
		Name:       "克罗地亚",
		Code:       "world:HR",
		AliasNames: []string{"克罗地亚"},
	},
	{
		Name:       "意大利",
		Code:       "world:IT",
		AliasNames: []string{"意大利"},
	},
	{
		Name:       "梵蒂冈",
		Code:       "world:VA",
		AliasNames: []string{"梵蒂冈"},
	},
	{
		Name:       "圣马力诺",
		Code:       "world:SM",
		AliasNames: []string{"圣马力诺"},
	},
	{
		Name:       "马耳他",
		Code:       "world:MT",
		AliasNames: []string{"马耳他"},
	},
	{
		Name:       "西班牙",
		Code:       "world:ES",
		AliasNames: []string{"西班牙"},
	},
	{
		Name:       "葡萄牙",
		Code:       "world:PT",
		AliasNames: []string{"葡萄牙"},
	},
	{
		Name:       "安道尔共和国",
		Code:       "world:AD",
		AliasNames: []string{"安道尔共和国", "安道尔"},
	},
	{
		Name:       "埃及",
		Code:       "world:EG",
		AliasNames: []string{"埃及"},
	},
	{
		Name:       "利比亚",
		Code:       "world:LY",
		AliasNames: []string{"利比亚"},
	},
	{
		Name:       "苏丹",
		Code:       "world:SD",
		AliasNames: []string{"苏丹"},
	},
	{
		Name:       "突尼斯",
		Code:       "world:TN",
		AliasNames: []string{"突尼斯"},
	},
	{
		Name:       "阿尔及利亚",
		Code:       "world:DZ",
		AliasNames: []string{"阿尔及利亚"},
	},
	{
		Name:       "摩洛哥",
		Code:       "world:MA",
		AliasNames: []string{"摩洛哥"},
	},
	{
		Name:       "埃塞俄比亚",
		Code:       "world:ET",
		AliasNames: []string{"埃塞俄比亚"},
	},
	{
		Name:       "厄立特里亚",
		Code:       "world:ER",
		AliasNames: []string{"厄立特里亚"},
	},
	{
		Name:       "索马里",
		Code:       "world:SO",
		AliasNames: []string{"索马里"},
	},
	{
		Name:       "吉布提",
		Code:       "world:DJ",
		AliasNames: []string{"吉布提"},
	},
	{
		Name:       "肯尼亚",
		Code:       "world:KE",
		AliasNames: []string{"肯尼亚"},
	},
	{
		Name:       "坦桑尼亚",
		Code:       "world:TZ",
		AliasNames: []string{"坦桑尼亚"},
	},
	{
		Name:       "乌干达",
		Code:       "world:UG",
		AliasNames: []string{"乌干达"},
	},
	{
		Name:       "卢旺达",
		Code:       "world:RW",
		AliasNames: []string{"卢旺达"},
	},
	{
		Name:       "布隆迪",
		Code:       "world:BI",
		AliasNames: []string{"布隆迪"},
	},
	{
		Name:       "塞舌尔",
		Code:       "world:SC",
		AliasNames: []string{"塞舌尔"},
	},
	{
		Name:       "圣多美和普林西比",
		Code:       "world:ST",
		AliasNames: []string{"圣多美及普林西比", "圣多美和普林西比"},
	},
	{
		Name:       "塞内加尔",
		Code:       "world:SN",
		AliasNames: []string{"塞内加尔"},
	},
	{
		Name:       "冈比亚",
		Code:       "world:GM",
		AliasNames: []string{"冈比亚"},
	},
	{
		Name:       "马里",
		Code:       "world:ML",
		AliasNames: []string{"马里"},
	},
	{
		Name:       "布基纳法索",
		Code:       "world:BF",
		AliasNames: []string{"布基纳法索"},
	},
	{
		Name:       "几内亚",
		Code:       "world:GN",
		AliasNames: []string{"几内亚"},
	},
	{
		Name:       "几内亚比绍",
		Code:       "world:GW",
		AliasNames: []string{"几内亚比绍"},
	},
	{
		Name:       "佛得角",
		Code:       "world:CV",
		AliasNames: []string{"佛得角"},
	},
	{
		Name:       "塞拉利昂",
		Code:       "world:SL",
		AliasNames: []string{"塞拉利昂"},
	},
	{
		Name:       "利比里亚",
		Code:       "world:LR",
		AliasNames: []string{"利比里亚"},
	},
	{
		Name:       "科特迪瓦",
		Code:       "world:CI",
		AliasNames: []string{"科特迪瓦"},
	},
	{
		Name:       "加纳",
		Code:       "world:GH",
		AliasNames: []string{"加纳"},
	},
	{
		Name:       "多哥",
		Code:       "world:TG",
		AliasNames: []string{"多哥"},
	},
	{
		Name:       "贝宁",
		Code:       "world:BJ",
		AliasNames: []string{"贝宁"},
	},
	{
		Name:       "尼日尔",
		Code:       "world:NE",
		AliasNames: []string{"尼日尔"},
	},
	{
		Name:       "赞比亚",
		Code:       "world:ZM",
		AliasNames: []string{"赞比亚"},
	},
	{
		Name:       "安哥拉",
		Code:       "world:AO",
		AliasNames: []string{"安哥拉"},
	},
	{
		Name:       "津巴布韦",
		Code:       "world:ZW",
		AliasNames: []string{"津巴布韦"},
	},
	{
		Name:       "马拉维",
		Code:       "world:MW",
		AliasNames: []string{"马拉维"},
	},
	{
		Name:       "莫桑比克",
		Code:       "world:MZ",
		AliasNames: []string{"莫桑比克"},
	},
	{
		Name:       "博茨瓦纳",
		Code:       "world:BW",
		AliasNames: []string{"博茨瓦纳"},
	},
	{
		Name:       "纳米比亚",
		Code:       "world:NA",
		AliasNames: []string{"纳米比亚"},
	},
	{
		Name:       "南非",
		Code:       "world:ZA",
		AliasNames: []string{"南非"},
	},
	{
		Name:       "斯威士兰",
		Code:       "world:SZ",
		AliasNames: []string{"斯威士兰"},
	},
	{
		Name:       "莱索托",
		Code:       "world:LS",
		AliasNames: []string{"莱索托"},
	},
	{
		Name:       "马达加斯加",
		Code:       "world:MG",
		AliasNames: []string{"马达加斯加"},
	},
	{
		Name:       "科摩罗",
		Code:       "world:KM",
		AliasNames: []string{"科摩罗"},
	},
	{
		Name:       "毛里求斯",
		Code:       "world:MU",
		AliasNames: []string{"毛里求斯"},
	},
	{
		Name:       "留尼汪",
		Code:       "world:RE",
		AliasNames: []string{"留尼旺", "留尼汪", "留尼汪岛"},
	},
	{
		Name:       "圣赫勒拿",
		Code:       "world:SH",
		AliasNames: []string{"圣赫勒拿"},
	},
	{
		Name:       "澳大利亚",
		Code:       "world:AU",
		AliasNames: []string{"澳大利亚"},
	},
	{
		Name:       "新西兰",
		Code:       "world:NZ",
		AliasNames: []string{"新西兰"},
	},
	{
		Name:       "巴布亚新几内亚",
		Code:       "world:PG",
		AliasNames: []string{"巴布亚新几内亚"},
	},
	{
		Name:       "所罗门群岛",
		Code:       "world:SB",
		AliasNames: []string{"所罗门群岛"},
	},
	{
		Name:       "瓦努阿图共和国",
		Code:       "world:VU",
		AliasNames: []string{"瓦努阿图共和国", "瓦努阿图"},
	},
	{
		Name:       "密克罗尼西亚",
		Code:       "world:FM",
		AliasNames: []string{"密克罗尼西亚"},
	},
	{
		Name:       "马绍尔群岛",
		Code:       "world:MH",
		AliasNames: []string{"马绍尔群岛"},
	},
	{
		Name:       "帕劳",
		Code:       "world:PW",
		AliasNames: []string{"帕劳"},
	},
	{
		Name:       "瑙鲁",
		Code:       "world:NR",
		AliasNames: []string{"瑙鲁"},
	},
	{
		Name:       "基里巴斯",
		Code:       "world:KI",
		AliasNames: []string{"基里巴斯"},
	},
	{
		Name:       "图瓦卢",
		Code:       "world:TV",
		AliasNames: []string{"图瓦卢"},
	},
	{
		Name:       "萨摩亚",
		Code:       "world:WS",
		AliasNames: []string{"萨摩亚"},
	},
	{
		Name:       "斐济",
		Code:       "world:FJ",
		AliasNames: []string{"斐济"},
	},
	{
		Name:       "汤加",
		Code:       "world:TO",
		AliasNames: []string{"汤加"},
	},
	{
		Name:       "库克群岛",
		Code:       "world:CK",
		AliasNames: []string{"库克群岛"},
	},
	{
		Name:       "关岛",
		Code:       "world:GU",
		AliasNames: []string{"关岛"},
	},
	{
		Name:       "新喀里多尼亚",
		Code:       "world:NC",
		AliasNames: []string{"新喀里多尼亚"},
	},
	{
		Name:       "法属波利尼西亚",
		Code:       "world:PF",
		AliasNames: []string{"法属波利尼西亚"},
	},
	{
		Name:       "皮特凯恩岛",
		Code:       "world:PN",
		AliasNames: []string{"皮特凯恩岛"},
	},
	{
		Name:       "瓦利斯与富图纳",
		Code:       "world:WF",
		AliasNames: []string{"瓦利斯与富图纳", "瓦利斯和富图纳"},
	},
	{
		Name:       "纽埃",
		Code:       "world:NU",
		AliasNames: []string{"纽埃"},
	},
	{
		Name:       "托克劳",
		Code:       "world:TK",
		AliasNames: []string{"托克劳", "托克劳群岛"},
	},
	{
		Name:       "美属萨摩亚",
		Code:       "world:AS",
		AliasNames: []string{"美属萨摩亚"},
	},
	{
		Name:       "北马里亚纳",
		Code:       "world:MP",
		AliasNames: []string{"北马里亚纳", "北马里亚纳群岛"},
	},
	{
		Name:       "加拿大",
		Code:       "world:CA",
		AliasNames: []string{"加拿大"},
	},
	{
		Name:       "美国",
		Code:       "world:US",
		AliasNames: []string{"美国"},
	},
	{
		Name:       "墨西哥",
		Code:       "world:MX",
		AliasNames: []string{"墨西哥"},
	},
	{
		Name:       "格陵兰",
		Code:       "world:GL",
		AliasNames: []string{"格陵兰", "格陵兰岛"},
	},
	{
		Name:       "危地马拉",
		Code:       "world:GT",
		AliasNames: []string{"危地马拉"},
	},
	{
		Name:       "伯利兹",
		Code:       "world:BZ",
		AliasNames: []string{"伯利兹"},
	},
	{
		Name:       "萨尔瓦多",
		Code:       "world:SV",
		AliasNames: []string{"萨尔瓦多"},
	},
	{
		Name:       "洪都拉斯",
		Code:       "world:HN",
		AliasNames: []string{"洪都拉斯"},
	},
	{
		Name:       "尼加拉瓜",
		Code:       "world:NI",
		AliasNames: []string{"尼加拉瓜"},
	},
	{
		Name:       "哥斯达黎加",
		Code:       "world:CR",
		AliasNames: []string{"哥斯达黎加"},
	},
	{
		Name:       "巴拿马",
		Code:       "world:PA",
		AliasNames: []string{"巴拿马"},
	},
	{
		Name:       "巴哈马",
		Code:       "world:BS",
		AliasNames: []string{"巴哈马"},
	},
	{
		Name:       "古巴",
		Code:       "world:CU",
		AliasNames: []string{"古巴"},
	},
	{
		Name:       "牙买加",
		Code:       "world:JM",
		AliasNames: []string{"牙买加"},
	},
	{
		Name:       "海地",
		Code:       "world:HT",
		AliasNames: []string{"海地"},
	},
	{
		Name:       "多米尼加共和国",
		Code:       "world:DO",
		AliasNames: []string{"多米尼加共和国", "多米尼加"},
	},
	{
		Name:       "安提瓜和巴布达",
		Code:       "world:AG",
		AliasNames: []string{"安提瓜和巴布达"},
	},
	{
		Name:       "圣基茨和尼维斯",
		Code:       "world:KN",
		AliasNames: []string{"圣基茨和尼维斯"},
	},
	{
		Name:       "多米尼克",
		Code:       "world:DM",
		AliasNames: []string{"多米尼克"},
	},
	{
		Name:       "圣卢西亚",
		Code:       "world:LC",
		AliasNames: []string{"圣卢西亚"},
	},
	{
		Name:       "圣文森特和格林纳丁斯",
		Code:       "world:VC",
		AliasNames: []string{"圣文森特和格林纳丁斯"},
	},
	{
		Name:       "格林纳达",
		Code:       "world:GD",
		AliasNames: []string{"格林纳达"},
	},
	{
		Name:       "巴巴多斯",
		Code:       "world:BB",
		AliasNames: []string{"巴巴多斯"},
	},
	{
		Name:       "特立尼达和多巴哥",
		Code:       "world:TT",
		AliasNames: []string{"特立尼达和多巴哥"},
	},
	{
		Name:       "波多黎各",
		Code:       "world:PR",
		AliasNames: []string{"波多黎各"},
	},
	{
		Name:       "英属维尔京群岛",
		Code:       "world:VG",
		AliasNames: []string{"英属维尔京群岛"},
	},
	{
		Name:       "美属维尔京群岛",
		Code:       "world:VI",
		AliasNames: []string{"美属维尔京群岛"},
	},
	{
		Name:       "安圭拉",
		Code:       "world:AI",
		AliasNames: []string{"安圭拉"},
	},
	{
		Name:       "蒙特塞拉特岛",
		Code:       "world:MS",
		AliasNames: []string{"蒙特塞拉特岛", "蒙塞拉特岛"},
	},
	{
		Name:       "瓜德罗普",
		Code:       "world:GP",
		AliasNames: []string{"瓜德罗普"},
	},
	{
		Name:       "马提尼克",
		Code:       "world:MQ",
		AliasNames: []string{"马提尼克"},
	},
	{
		Name:       "荷属安的列斯",
		Code:       "world:AN",
		AliasNames: []string{"荷属安的列斯", "安的列斯"},
	},
	{
		Name:       "阿鲁巴",
		Code:       "world:AW",
		AliasNames: []string{"阿鲁巴"},
	},
	{
		Name:       "特克斯和凯科斯群岛",
		Code:       "world:TC",
		AliasNames: []string{"特克斯和凯科斯群岛"},
	},
	{
		Name:       "开曼群岛",
		Code:       "world:KY",
		AliasNames: []string{"开曼群岛"},
	},
	{
		Name:       "百慕大",
		Code:       "world:BM",
		AliasNames: []string{"百慕大", "百慕大群岛"},
	},
	{
		Name:       "哥伦比亚",
		Code:       "world:CO",
		AliasNames: []string{"哥伦比亚"},
	},
	{
		Name:       "委内瑞拉",
		Code:       "world:VE",
		AliasNames: []string{"委内瑞拉"},
	},
	{
		Name:       "圭亚那",
		Code:       "world:GY",
		AliasNames: []string{"圭亚那"},
	},
	{
		Name:       "法属圭亚那",
		Code:       "world:GF",
		AliasNames: []string{"法属圭亚那"},
	},
	{
		Name:       "苏里南",
		Code:       "world:SR",
		AliasNames: []string{"苏里南"},
	},
	{
		Name:       "厄瓜多尔",
		Code:       "world:EC",
		AliasNames: []string{"厄瓜多尔"},
	},
	{
		Name:       "秘鲁",
		Code:       "world:PE",
		AliasNames: []string{"秘鲁"},
	},
	{
		Name:       "玻利维亚",
		Code:       "world:BO",
		AliasNames: []string{"玻利维亚"},
	},
	{
		Name:       "巴西",
		Code:       "world:BR",
		AliasNames: []string{"巴西"},
	},
	{
		Name:       "智利",
		Code:       "world:CL",
		AliasNames: []string{"智利"},
	},
	{
		Name:       "阿根廷",
		Code:       "world:AR",
		AliasNames: []string{"阿根廷"},
	},
	{
		Name:       "乌拉圭",
		Code:       "world:UY",
		AliasNames: []string{"乌拉圭"},
	},
	{
		Name:       "巴拉圭",
		Code:       "world:PY",
		AliasNames: []string{"巴拉圭"},
	},
	{
		Name:       "波黑",
		Code:       "world:BA",
		AliasNames: []string{"波黑"},
	},
	{
		Name:       "直布罗陀",
		Code:       "world:GI",
		AliasNames: []string{"直布罗陀"},
	},
	{
		Name:       "泽西岛",
		Code:       "world:JE",
		AliasNames: []string{"泽西岛"},
	},
	{
		Name:       "黑山",
		Code:       "world:ME",
		AliasNames: []string{"黑山"},
	},
	{
		Name:       "英属马恩岛",
		Code:       "world:IM",
		AliasNames: []string{"英属马恩岛", "马恩岛"},
	},
	{
		Name:       "尼日利亚",
		Code:       "world:NG",
		AliasNames: []string{"尼日利亚"},
	},
	{
		Name:       "喀麦隆",
		Code:       "world:CM",
		AliasNames: []string{"喀麦隆"},
	},
	{
		Name:       "加蓬",
		Code:       "world:GA",
		AliasNames: []string{"加蓬"},
	},
	{
		Name:       "乍得",
		Code:       "world:TD",
		AliasNames: []string{"乍得"},
	},
	{
		Name:       "刚果共和国",
		Code:       "world:CG",
		AliasNames: []string{"刚果共和国", "刚果布"},
	},
	{
		Name:       "中非共和国",
		Code:       "world:CF",
		AliasNames: []string{"中非共和国", "中非"},
	},
	{
		Name:       "南苏丹",
		Code:       "world:SS",
		AliasNames: []string{"南苏丹"},
	},
	{
		Name:       "赤道几内亚",
		Code:       "world:GQ",
		AliasNames: []string{"赤道几内亚"},
	},
	{
		Name:       "毛里塔尼亚",
		Code:       "world:MR",
		AliasNames: []string{"毛里塔尼亚"},
	},
	{
		Name:       "刚果民主共和国",
		Code:       "world:CD",
		AliasNames: []string{"刚果民主共和国", "刚果金"},
	},
	{
		Name:       "法罗群岛",
		Code:       "world:FO",
		AliasNames: []string{"法罗群岛"},
	},
	{
		Name:       "根西岛",
		Code:       "world:GG",
		AliasNames: []string{"根西岛"},
	},
	{
		Name:       "圣皮埃尔和密克隆群岛",
		Code:       "world:PM",
		AliasNames: []string{"圣皮埃尔和密克隆群岛"},
	},
	{
		Name:       "法属圣马丁",
		Code:       "world:MF",
		AliasNames: []string{"法属圣马丁"},
	},
	{
		Name:       "奥兰群岛",
		Code:       "world:AX",
		AliasNames: []string{"奥兰群岛"},
	},
	{
		Name:       "库拉索",
		Code:       "world:CW",
		AliasNames: []string{"库拉索"},
	},
	{
		Name:       "圣巴泰勒米岛",
		Code:       "world:BL",
		AliasNames: []string{"圣巴泰勒米岛", "圣巴泰勒米"},
	},
	{
		Name:       "福克兰群岛",
		Code:       "world:FK",
		AliasNames: []string{"福克兰群岛"},
	},
	{
		Name:       "英属印度洋领地",
		Code:       "world:IO",
		AliasNames: []string{"英属印度洋领地"},
	},
	{
		Name:       "诺福克岛",
		Code:       "world:NF",
		AliasNames: []string{"诺福克岛"},
	},
	{
		Name:       "马约特岛",
		Code:       "world:YT",
		AliasNames: []string{"法属马约特岛", "马约特岛", "马约特"},
	},
	{
		Name:       "科索沃",
		Code:       "world:XK",
		AliasNames: []string{"科索沃"},
	},
	{
		Name:       "荷属圣马丁",
		Code:       "world:SX",
		AliasNames: []string{"荷属圣马丁"},
	},
	{
		Name:       "南乔治亚岛和南桑威奇群岛",
		Code:       "world:GS",
		AliasNames: []string{"南乔治亚岛和南桑威奇群岛"},
	},
	{
		Name:       "南极洲",
		Code:       "world:AQ",
		AliasNames: []string{"南极洲"},
	},
	{
		Name:       "圣诞岛",
		Code:       "world:CX",
		AliasNames: []string{"圣诞岛"},
	},
}
