package currency

const (
	JPY = 0
	CNY = 1
	CHF = 2
	CAD = 3
	MXN = 4
	INR = 5
	BRL = 6
	RUB = 7
	KRW = 8
	IDR = 9
	TRY = 10
	SAR = 11
	SEK = 12
	NGN = 13
	PLN = 14
	ARS = 15
	NOK = 16
	TWD = 17
	IRR = 18
	AED = 19
	COP = 20
	THB = 31
	ZAR = 32
	DKK = 33
	MYR = 34
	SGD = 35
	ILS = 36
	HKD = 37
	EGP = 38
	PHP = 39
	CLP = 40
	PKR = 40
	IQD = 41
	DZD = 42
	KZT = 43
	QAR = 44
	CZK = 45
	PEN = 46
	RON = 47
	VND = 48
	BDT = 49
	HUF = 50
	UAH = 51
	AOA = 52
	MAD = 53
	OMR = 54
	CUC = 55
	BYR = 56
	AZN = 57
	LKR = 58
	SDG = 59
	SYP = 60
	MMK = 61
	DOP = 62
	UZS = 63
	KES = 64
	GTQ = 65
	URY = 66
	HRV = 67
	MOP = 68
	ETB = 69
	CRC = 70
	TZS = 71
	TMT = 72
	TND = 73
	PAB = 74
	LBP = 75
	RSD = 76
	LYD = 77
	GHS = 78
	YER = 79
	BOB = 80
	BHD = 81
	CDF = 82
	PYG = 83
	UGX = 84
	SVC = 85
	TTD = 86
	AFN = 87
	NPR = 88
	HNL = 89
	BIH = 90
	BND = 91
	ISK = 92
	KHR = 93
	GEL = 94
	MZN = 95
	BWP = 96
	PGK = 97
	JMD = 98
	XAF = 99
	NAD = 100
	ALL = 101
	SSP = 102
	MUR = 103
	MNT = 104
	NIO = 105
	LAK = 106
	MKD = 107
	AMD = 108
	MGA = 109
	XPF = 110
	TJS = 111
	HTG = 112
	BSD = 113
	MDL = 114
	RWF = 115
	KGS = 116
	GNF = 117
	SRD = 118
	SLL = 119
	XOF = 120
	MWK = 121
	FJD = 122
	ERN = 123
	SZL = 124
	GYD = 125
	BIF = 126
	KYD = 127
	MVR = 128
	LSL = 129
	LRD = 130
	CVE = 131
	DJF = 132
	SCR = 133
	SOS = 134
	GMD = 135
	KMF = 136
	STD = 137
	XRP = 138
	AUD = 139
	BGN = 140
	BTC = 141
	JOD = 142
	GBP = 143
	ETH = 144
	EUR = 145
	LTC = 146
	NZD = 147
)

var currenyList = []interface{}{
	map[string]string{"text": "JPY", "viewText": "", "value": string(JPY)},
	map[string]string{"text": "CNY", "viewText": "", "value": string(CNY)},
	map[string]string{"text": "CHF", "viewText": "", "value": string(CHF)},
	map[string]string{"text": "CAD", "viewText": "", "value": string(CAD)},
	map[string]string{"text": "MXN", "viewText": "", "value": string(MXN)},
	map[string]string{"text": "INR", "viewText": "", "value": string(INR)},
	map[string]string{"text": "BRL", "viewText": "", "value": string(BRL)},
	map[string]string{"text": "RUB", "viewText": "", "value": string(RUB)},
	map[string]string{"text": "KRW", "viewText": "", "value": string(KRW)},
	map[string]string{"text": "IDR", "viewText": "", "value": string(IDR)},
	map[string]string{"text": "TRY", "viewText": "", "value": string(TRY)},
	map[string]string{"text": "SAR", "viewText": "", "value": string(SAR)},
	map[string]string{"text": "SEK", "viewText": "", "value": string(SEK)},
	map[string]string{"text": "NGN", "viewText": "", "value": string(NGN)},
	map[string]string{"text": "PLN", "viewText": "", "value": string(PLN)},
	map[string]string{"text": "ARS", "viewText": "", "value": string(ARS)},
	map[string]string{"text": "NOK", "viewText": "", "value": string(NOK)},
	map[string]string{"text": "TWD", "viewText": "", "value": string(TWD)},
	map[string]string{"text": "IRR", "viewText": "", "value": string(IRR)},
	map[string]string{"text": "AED", "viewText": "", "value": string(AED)},
	map[string]string{"text": "COP", "viewText": "", "value": string(COP)},
	map[string]string{"text": "THB", "viewText": "", "value": string(THB)},
	map[string]string{"text": "ZAR", "viewText": "", "value": string(ZAR)},
	map[string]string{"text": "DKK", "viewText": "", "value": string(DKK)},
	map[string]string{"text": "MYR", "viewText": "", "value": string(MYR)},
	map[string]string{"text": "SGD", "viewText": "", "value": string(SGD)},
	map[string]string{"text": "ILS", "viewText": "", "value": string(ILS)},
	map[string]string{"text": "HKD", "viewText": "", "value": string(HKD)},
	map[string]string{"text": "EGP", "viewText": "", "value": string(EGP)},
	map[string]string{"text": "PHP", "viewText": "", "value": string(PHP)},
	map[string]string{"text": "CLP", "viewText": "", "value": string(CLP)},
	map[string]string{"text": "PKR", "viewText": "", "value": string(PKR)},
	map[string]string{"text": "IQD", "viewText": "", "value": string(IQD)},
	map[string]string{"text": "DZD", "viewText": "", "value": string(DZD)},
	map[string]string{"text": "KZT", "viewText": "", "value": string(KZT)},
	map[string]string{"text": "QAR", "viewText": "", "value": string(QAR)},
	map[string]string{"text": "CZK", "viewText": "", "value": string(CZK)},
	map[string]string{"text": "PEN", "viewText": "", "value": string(PEN)},
	map[string]string{"text": "RON", "viewText": "", "value": string(RON)},
	map[string]string{"text": "VND", "viewText": "", "value": string(VND)},
	map[string]string{"text": "BDT", "viewText": "", "value": string(BDT)},
	map[string]string{"text": "HUF", "viewText": "", "value": string(HUF)},
	map[string]string{"text": "UAH", "viewText": "", "value": string(UAH)},
	map[string]string{"text": "AOA", "viewText": "", "value": string(AOA)},
	map[string]string{"text": "MAD", "viewText": "", "value": string(MAD)},
	map[string]string{"text": "OMR", "viewText": "", "value": string(OMR)},
	map[string]string{"text": "CUC", "viewText": "", "value": string(CUC)},
	map[string]string{"text": "BYR", "viewText": "", "value": string(BYR)},
	map[string]string{"text": "AZN", "viewText": "", "value": string(AZN)},
	map[string]string{"text": "LKR", "viewText": "", "value": string(LKR)},
	map[string]string{"text": "SDG", "viewText": "", "value": string(SDG)},
	map[string]string{"text": "SYP", "viewText": "", "value": string(SYP)},
	map[string]string{"text": "MMK", "viewText": "", "value": string(MMK)},
	map[string]string{"text": "DOP", "viewText": "", "value": string(DOP)},
	map[string]string{"text": "UZS", "viewText": "", "value": string(UZS)},
	map[string]string{"text": "KES", "viewText": "", "value": string(KES)},
	map[string]string{"text": "GTQ", "viewText": "", "value": string(GTQ)},
	map[string]string{"text": "URY", "viewText": "", "value": string(URY)},
	map[string]string{"text": "HRV", "viewText": "", "value": string(HRV)},
	map[string]string{"text": "MOP", "viewText": "", "value": string(MOP)},
	map[string]string{"text": "ETB", "viewText": "", "value": string(ETB)},
	map[string]string{"text": "CRC", "viewText": "", "value": string(CRC)},
	map[string]string{"text": "TZS", "viewText": "", "value": string(TZS)},
	map[string]string{"text": "TMT", "viewText": "", "value": string(TMT)},
	map[string]string{"text": "TND", "viewText": "", "value": string(TND)},
	map[string]string{"text": "PAB", "viewText": "", "value": string(PAB)},
	map[string]string{"text": "LBP", "viewText": "", "value": string(LBP)},
	map[string]string{"text": "RSD", "viewText": "", "value": string(RSD)},
	map[string]string{"text": "LYD", "viewText": "", "value": string(LYD)},
	map[string]string{"text": "GHS", "viewText": "", "value": string(GHS)},
	map[string]string{"text": "YER", "viewText": "", "value": string(YER)},
	map[string]string{"text": "BOB", "viewText": "", "value": string(BOB)},
	map[string]string{"text": "BHD", "viewText": "", "value": string(BHD)},
	map[string]string{"text": "CDF", "viewText": "", "value": string(CDF)},
	map[string]string{"text": "PYG", "viewText": "", "value": string(PYG)},
	map[string]string{"text": "UGX", "viewText": "", "value": string(UGX)},
	map[string]string{"text": "SVC", "viewText": "", "value": string(SVC)},
	map[string]string{"text": "TTD", "viewText": "", "value": string(TTD)},
	map[string]string{"text": "AFN", "viewText": "", "value": string(AFN)},
	map[string]string{"text": "NPR", "viewText": "", "value": string(NPR)},
	map[string]string{"text": "HNL", "viewText": "", "value": string(HNL)},
	map[string]string{"text": "BIH", "viewText": "", "value": string(BIH)},
	map[string]string{"text": "BND", "viewText": "", "value": string(BND)},
	map[string]string{"text": "ISK", "viewText": "", "value": string(ISK)},
	map[string]string{"text": "KHR", "viewText": "", "value": string(KHR)},
	map[string]string{"text": "GEL", "viewText": "", "value": string(GEL)},
	map[string]string{"text": "MZN", "viewText": "", "value": string(MZN)},
	map[string]string{"text": "BWP", "viewText": "", "value": string(BWP)},
	map[string]string{"text": "PGK", "viewText": "", "value": string(PGK)},
	map[string]string{"text": "JMD", "viewText": "", "value": string(JMD)},
	map[string]string{"text": "XAF", "viewText": "", "value": string(XAF)},
	map[string]string{"text": "NAD", "viewText": "", "value": string(NAD)},
	map[string]string{"text": "ALL", "viewText": "", "value": string(ALL)},
	map[string]string{"text": "SSP", "viewText": "", "value": string(SSP)},
	map[string]string{"text": "MUR", "viewText": "", "value": string(MUR)},
	map[string]string{"text": "MNT", "viewText": "", "value": string(MNT)},
	map[string]string{"text": "NIO", "viewText": "", "value": string(NIO)},
	map[string]string{"text": "LAK", "viewText": "", "value": string(LAK)},
	map[string]string{"text": "MKD", "viewText": "", "value": string(MKD)},
	map[string]string{"text": "AMD", "viewText": "", "value": string(AMD)},
	map[string]string{"text": "MGA", "viewText": "", "value": string(MGA)},
	map[string]string{"text": "XPF", "viewText": "", "value": string(XPF)},
	map[string]string{"text": "TJS", "viewText": "", "value": string(TJS)},
	map[string]string{"text": "HTG", "viewText": "", "value": string(HTG)},
	map[string]string{"text": "BSD", "viewText": "", "value": string(BSD)},
	map[string]string{"text": "MDL", "viewText": "", "value": string(MDL)},
	map[string]string{"text": "RWF", "viewText": "", "value": string(RWF)},
	map[string]string{"text": "KGS", "viewText": "", "value": string(KGS)},
	map[string]string{"text": "GNF", "viewText": "", "value": string(GNF)},
	map[string]string{"text": "SRD", "viewText": "", "value": string(SRD)},
	map[string]string{"text": "SLL", "viewText": "", "value": string(SLL)},
	map[string]string{"text": "XOF", "viewText": "", "value": string(XOF)},
	map[string]string{"text": "MWK", "viewText": "", "value": string(MWK)},
	map[string]string{"text": "FJD", "viewText": "", "value": string(FJD)},
	map[string]string{"text": "ERN", "viewText": "", "value": string(ERN)},
	map[string]string{"text": "SZL", "viewText": "", "value": string(SZL)},
	map[string]string{"text": "GYD", "viewText": "", "value": string(GYD)},
	map[string]string{"text": "BIF", "viewText": "", "value": string(BIF)},
	map[string]string{"text": "KYD", "viewText": "", "value": string(KYD)},
	map[string]string{"text": "MVR", "viewText": "", "value": string(MVR)},
	map[string]string{"text": "LSL", "viewText": "", "value": string(LSL)},
	map[string]string{"text": "LRD", "viewText": "", "value": string(LRD)},
	map[string]string{"text": "CVE", "viewText": "", "value": string(CVE)},
	map[string]string{"text": "DJF", "viewText": "", "value": string(DJF)},
	map[string]string{"text": "SCR", "viewText": "", "value": string(SCR)},
	map[string]string{"text": "SOS", "viewText": "", "value": string(SOS)},
	map[string]string{"text": "GMD", "viewText": "", "value": string(GMD)},
	map[string]string{"text": "KMF", "viewText": "", "value": string(KMF)},
	map[string]string{"text": "STD", "viewText": "", "value": string(STD)},
	map[string]string{"text": "XRP", "viewText": "", "value": string(XRP)},
	map[string]string{"text": "AUD", "viewText": "", "value": string(AUD)},
	map[string]string{"text": "BGN", "viewText": "", "value": string(BGN)},
	map[string]string{"text": "BTC", "viewText": "", "value": string(BTC)},
	map[string]string{"text": "JOD", "viewText": "", "value": string(JOD)},
	map[string]string{"text": "GBP", "viewText": "", "value": string(GBP)},
	map[string]string{"text": "ETH", "viewText": "", "value": string(ETH)},
	map[string]string{"text": "EUR", "viewText": "", "value": string(EUR)},
	map[string]string{"text": "LTC", "viewText": "", "value": string(LTC)},
	map[string]string{"text": "NZD", "viewText": "", "value": string(NZD)},
	]
