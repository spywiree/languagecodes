package languagecodes

type LanguageCode string

const (
	DETECT_LANGUAGE LanguageCode = "auto"

	AFRIKAANS           LanguageCode = "af"
	ALBANIAN            LanguageCode = "sq"
	AMHARIC             LanguageCode = "am"
	ARABIC              LanguageCode = "ar"
	ARMENIAN            LanguageCode = "hy"
	ASSAMESE            LanguageCode = "as"
	AYMARA              LanguageCode = "ay"
	AZERBAIJANI         LanguageCode = "az"
	BAMBARA             LanguageCode = "bm"
	BASQUE              LanguageCode = "eu"
	BELARUSIAN          LanguageCode = "be"
	BENGALI             LanguageCode = "bn"
	BHOJPURI            LanguageCode = "bho"
	BOSNIAN             LanguageCode = "bs"
	BULGARIAN           LanguageCode = "bg"
	CATALAN             LanguageCode = "ca"
	CEBUANO             LanguageCode = "ceb"
	CHICHEWA            LanguageCode = "ny"
	CHINESE_SIMPLIFIED  LanguageCode = "zh-CN"
	CHINESE_TRADITIONAL LanguageCode = "zh-TW"
	CORSICAN            LanguageCode = "co"
	CROATIAN            LanguageCode = "hr"
	CZECH               LanguageCode = "cs"
	DANISH              LanguageCode = "da"
	DHIVEHI             LanguageCode = "dv"
	DOGRI               LanguageCode = "doi"
	DUTCH               LanguageCode = "nl"
	ENGLISH             LanguageCode = "en"
	ESPERANTO           LanguageCode = "eo"
	ESTONIAN            LanguageCode = "et"
	EWE                 LanguageCode = "ee"
	FILIPINO            LanguageCode = "tl"
	FINNISH             LanguageCode = "fi"
	FRENCH              LanguageCode = "fr"
	FRISIAN             LanguageCode = "fy"
	GALICIAN            LanguageCode = "gl"
	GEORGIAN            LanguageCode = "ka"
	GERMAN              LanguageCode = "de"
	GREEK               LanguageCode = "el"
	GUARANI             LanguageCode = "gn"
	GUJARATI            LanguageCode = "gu"
	HAITIAN_CREOLE      LanguageCode = "ht"
	HAUSA               LanguageCode = "ha"
	HAWAIIAN            LanguageCode = "haw"
	HEBREW              LanguageCode = "iw"
	HINDI               LanguageCode = "hi"
	HMONG               LanguageCode = "hmn"
	HUNGARIAN           LanguageCode = "hu"
	ICELANDIC           LanguageCode = "is"
	IGBO                LanguageCode = "ig"
	ILOCANO             LanguageCode = "ilo"
	INDONESIAN          LanguageCode = "id"
	IRISH               LanguageCode = "ga"
	ITALIAN             LanguageCode = "it"
	JAPANESE            LanguageCode = "ja"
	JAVANESE            LanguageCode = "jw"
	KANNADA             LanguageCode = "kn"
	KAZAKH              LanguageCode = "kk"
	KHMER               LanguageCode = "km"
	KINYARWANDA         LanguageCode = "rw"
	KONKANI             LanguageCode = "gom"
	KOREAN              LanguageCode = "ko"
	KRIO                LanguageCode = "kri"
	KURDISH_KURMANJI    LanguageCode = "ku"
	KURDISH_SORANI      LanguageCode = "ckb"
	KYRGYZ              LanguageCode = "ky"
	LAO                 LanguageCode = "lo"
	LATIN               LanguageCode = "la"
	LATVIAN             LanguageCode = "lv"
	LINGALA             LanguageCode = "ln"
	LITHUANIAN          LanguageCode = "lt"
	LUGANDA             LanguageCode = "lg"
	LUXEMBOURGISH       LanguageCode = "lb"
	MACEDONIAN          LanguageCode = "mk"
	MAITHILI            LanguageCode = "mai"
	MALAGASY            LanguageCode = "mg"
	MALAY               LanguageCode = "ms"
	MALAYALAM           LanguageCode = "ml"
	MALTESE             LanguageCode = "mt"
	MAORI               LanguageCode = "mi"
	MARATHI             LanguageCode = "mr"
	MEITEILON_MANIPURI  LanguageCode = "mni-Mtei"
	MIZO                LanguageCode = "lus"
	MONGOLIAN           LanguageCode = "mn"
	MYANMAR             LanguageCode = "my"
	NEPALI              LanguageCode = "ne"
	NORWEGIAN           LanguageCode = "no"
	ODIA_ORIYA          LanguageCode = "or"
	OROMO               LanguageCode = "om"
	PASHTO              LanguageCode = "ps"
	PERSIAN             LanguageCode = "fa"
	POLISH              LanguageCode = "pl"
	PORTUGUESE          LanguageCode = "pt"
	PUNJABI             LanguageCode = "pa"
	QUECHUA             LanguageCode = "qu"
	ROMANIAN            LanguageCode = "ro"
	RUSSIAN             LanguageCode = "ru"
	SAMOAN              LanguageCode = "sm"
	SANSKRIT            LanguageCode = "sa"
	SCOTS_GAELIC        LanguageCode = "gd"
	SEPEDI              LanguageCode = "nso"
	SERBIAN             LanguageCode = "sr"
	SESOTHO             LanguageCode = "st"
	SHONA               LanguageCode = "sn"
	SINDHI              LanguageCode = "sd"
	SINHALA             LanguageCode = "si"
	SLOVAK              LanguageCode = "sk"
	SLOVENIAN           LanguageCode = "sl"
	SOMALI              LanguageCode = "so"
	SPANISH             LanguageCode = "es"
	SUNDANESE           LanguageCode = "su"
	SWAHILI             LanguageCode = "sw"
	SWEDISH             LanguageCode = "sv"
	TAJIK               LanguageCode = "tg"
	TAMIL               LanguageCode = "ta"
	TATAR               LanguageCode = "tt"
	TELUGU              LanguageCode = "te"
	THAI                LanguageCode = "th"
	TIGRINYA            LanguageCode = "ti"
	TSONGA              LanguageCode = "ts"
	TURKISH             LanguageCode = "tr"
	TURKMEN             LanguageCode = "tk"
	TWI                 LanguageCode = "ak"
	UKRAINIAN           LanguageCode = "uk"
	URDU                LanguageCode = "ur"
	UYGHUR              LanguageCode = "ug"
	UZBEK               LanguageCode = "uz"
	VIETNAMESE          LanguageCode = "vi"
	WELSH               LanguageCode = "cy"
	XHOSA               LanguageCode = "xh"
	YIDDISH             LanguageCode = "yi"
	YORUBA              LanguageCode = "yo"
	ZULU                LanguageCode = "zu"
)

func SupportedLanguages() []string {
	return []string{"afrikaans", "albanian", "amharic", "arabic", "armenian", "assamese", "aymara", "azerbaijani", "bambara", "basque", "belarusian", "bengali", "bhojpuri", "bosnian", "bulgarian", "catalan", "cebuano", "chichewa", "chinese (simplified)", "chinese (traditional)", "corsican", "croatian", "czech", "danish", "dhivehi", "dogri", "dutch", "english", "esperanto", "estonian", "ewe", "filipino", "finnish", "french", "frisian", "galician", "georgian", "german", "greek", "guarani", "gujarati", "haitian creole", "hausa", "hawaiian", "hebrew", "hindi", "hmong", "hungarian", "icelandic", "igbo", "ilocano", "indonesian", "irish", "italian", "japanese", "javanese", "kannada", "kazakh", "khmer", "kinyarwanda", "konkani", "korean", "krio", "kurdish (kurmanji)", "kurdish (sorani)", "kyrgyz", "lao", "latin", "latvian", "lingala", "lithuanian", "luganda", "luxembourgish", "macedonian", "maithili", "malagasy", "malay", "malayalam", "maltese", "maori", "marathi", "meiteilon (manipuri)", "mizo", "mongolian", "myanmar", "nepali", "norwegian", "odia (oriya)", "oromo", "pashto", "persian", "polish", "portuguese", "punjabi", "quechua", "romanian", "russian", "samoan", "sanskrit", "scots gaelic", "sepedi", "serbian", "sesotho", "shona", "sindhi", "sinhala", "slovak", "slovenian", "somali", "spanish", "sundanese", "swahili", "swedish", "tajik", "tamil", "tatar", "telugu", "thai", "tigrinya", "tsonga", "turkish", "turkmen", "twi", "ukrainian", "urdu", "uyghur", "uzbek", "vietnamese", "welsh", "xhosa", "yiddish", "yoruba", "zulu"}
}