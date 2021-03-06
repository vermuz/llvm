// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"empty",
		"source_filename",
		"=",
		"string_lit",
		"target",
		"datalayout",
		"triple",
		"module",
		"asm",
		"type",
		"opaque",
		"comdat",
		"any",
		"exactmatch",
		"largest",
		"noduplicates",
		"samesize",
		",",
		"externally_initialized",
		"constant",
		"global",
		"declare",
		"define",
		"(",
		")",
		"...",
		"{",
		"}",
		"attributes",
		"!",
		"distinct",
		"global_ident",
		"local_ident",
		"label_ident",
		"attr_group_id",
		"comdat_name",
		"metadata_name",
		"metadata_id",
		"void",
		"int_type",
		"half",
		"float",
		"double",
		"fp128",
		"x86_fp80",
		"ppc_fp128",
		"*",
		"addrspace",
		"<",
		"x",
		">",
		"label",
		"metadata",
		"[",
		"]",
		"int_lit",
		"true",
		"false",
		"float_lit",
		"null",
		"c",
		"zeroinitializer",
		"undef",
		"add",
		"fadd",
		"sub",
		"fsub",
		"mul",
		"fmul",
		"udiv",
		"sdiv",
		"fdiv",
		"urem",
		"srem",
		"frem",
		"shl",
		"lshr",
		"ashr",
		"and",
		"or",
		"xor",
		"extractelement",
		"insertelement",
		"shufflevector",
		"extractvalue",
		"insertvalue",
		"getelementptr",
		"trunc",
		"to",
		"zext",
		"sext",
		"fptrunc",
		"fpext",
		"fptoui",
		"fptosi",
		"uitofp",
		"sitofp",
		"ptrtoint",
		"inttoptr",
		"bitcast",
		"addrspacecast",
		"icmp",
		"fcmp",
		"select",
		"nuw",
		"nsw",
		"arcp",
		"fast",
		"ninf",
		"nnan",
		"nsz",
		"exact",
		"alloca",
		"load",
		"volatile",
		"store",
		"fence",
		"singlethread",
		"acq_rel",
		"acquire",
		"monotonic",
		"release",
		"seq_cst",
		"unordered",
		"cmpxchg",
		"weak",
		"atomicrmw",
		"max",
		"min",
		"nand",
		"umax",
		"umin",
		"xchg",
		"eq",
		"ne",
		"ugt",
		"uge",
		"ult",
		"ule",
		"sgt",
		"sge",
		"slt",
		"sle",
		"oeq",
		"ogt",
		"oge",
		"olt",
		"ole",
		"one",
		"ord",
		"ueq",
		"une",
		"uno",
		"phi",
		"call",
		"tail",
		"musttail",
		"notail",
		"va_arg",
		"landingpad",
		"cleanup",
		"catch",
		"filter",
		"catchpad",
		"within",
		"cleanuppad",
		"none",
		"ret",
		"br",
		"switch",
		"indirectbr",
		"invoke",
		"unwind",
		"resume",
		"catchswitch",
		"caller",
		"catchret",
		"from",
		"cleanupret",
		"unreachable",
		"appending",
		"available_externally",
		"common",
		"internal",
		"linkonce",
		"linkonce_odr",
		"private",
		"weak_odr",
		"extern_weak",
		"external",
		"default",
		"hidden",
		"protected",
		"dllimport",
		"dllexport",
		"thread_local",
		"localdynamic",
		"initialexec",
		"localexec",
		"local_unnamed_addr",
		"unnamed_addr",
		"section",
		"align",
		"gc",
		"prefix",
		"prologue",
		"personality",
		"byval",
		"dereferenceable",
		"dereferenceable_or_null",
		"inalloca",
		"inreg",
		"nest",
		"noalias",
		"nocapture",
		"nonnull",
		"readnone",
		"readonly",
		"returned",
		"signext",
		"sret",
		"swifterror",
		"swiftself",
		"writeonly",
		"zeroext",
		"alignstack",
		"allocsize",
		"alwaysinline",
		"argmemonly",
		"builtin",
		"cold",
		"convergent",
		"inaccessiblemem_or_argmemonly",
		"inaccessiblememonly",
		"inlinehint",
		"jumptable",
		"minsize",
		"naked",
		"nobuiltin",
		"noduplicate",
		"noimplicitfloat",
		"noinline",
		"nonlazybind",
		"norecurse",
		"noredzone",
		"noreturn",
		"nounwind",
		"optnone",
		"optsize",
		"returns_twice",
		"safestack",
		"sanitize_address",
		"sanitize_memory",
		"sanitize_thread",
		"ssp",
		"sspreq",
		"sspstrong",
		"uwtable",
		"amdgpu_cs",
		"amdgpu_gs",
		"amdgpu_kernel",
		"amdgpu_ps",
		"amdgpu_vs",
		"anyregcc",
		"arm_aapcs_vfpcc",
		"arm_aapcscc",
		"arm_apcscc",
		"avr_intrcc",
		"avr_signalcc",
		"cc",
		"ccc",
		"coldcc",
		"cxx_fast_tlscc",
		"fastcc",
		"ghccc",
		"hhvm_ccc",
		"hhvmcc",
		"intel_ocl_bicc",
		"msp430_intrcc",
		"preserve_allcc",
		"preserve_mostcc",
		"ptx_device",
		"ptx_kernel",
		"spir_func",
		"spir_kernel",
		"swiftcc",
		"webkit_jscc",
		"x86_64_sysvcc",
		"x86_64_win64cc",
		"x86_fastcallcc",
		"x86_intrcc",
		"x86_regcallcc",
		"x86_stdcallcc",
		"x86_thiscallcc",
		"x86_vectorcallcc",
		"inbounds",
	},

	idMap: map[string]Type{
		"INVALID":         0,
		"$":               1,
		"empty":           2,
		"source_filename": 3,
		"=":               4,
		"string_lit":      5,
		"target":          6,
		"datalayout":      7,
		"triple":          8,
		"module":          9,
		"asm":             10,
		"type":            11,
		"opaque":          12,
		"comdat":          13,
		"any":             14,
		"exactmatch":      15,
		"largest":         16,
		"noduplicates":    17,
		"samesize":        18,
		",":               19,
		"externally_initialized":        20,
		"constant":                      21,
		"global":                        22,
		"declare":                       23,
		"define":                        24,
		"(":                             25,
		")":                             26,
		"...":                           27,
		"{":                             28,
		"}":                             29,
		"attributes":                    30,
		"!":                             31,
		"distinct":                      32,
		"global_ident":                  33,
		"local_ident":                   34,
		"label_ident":                   35,
		"attr_group_id":                 36,
		"comdat_name":                   37,
		"metadata_name":                 38,
		"metadata_id":                   39,
		"void":                          40,
		"int_type":                      41,
		"half":                          42,
		"float":                         43,
		"double":                        44,
		"fp128":                         45,
		"x86_fp80":                      46,
		"ppc_fp128":                     47,
		"*":                             48,
		"addrspace":                     49,
		"<":                             50,
		"x":                             51,
		">":                             52,
		"label":                         53,
		"metadata":                      54,
		"[":                             55,
		"]":                             56,
		"int_lit":                       57,
		"true":                          58,
		"false":                         59,
		"float_lit":                     60,
		"null":                          61,
		"c":                             62,
		"zeroinitializer":               63,
		"undef":                         64,
		"add":                           65,
		"fadd":                          66,
		"sub":                           67,
		"fsub":                          68,
		"mul":                           69,
		"fmul":                          70,
		"udiv":                          71,
		"sdiv":                          72,
		"fdiv":                          73,
		"urem":                          74,
		"srem":                          75,
		"frem":                          76,
		"shl":                           77,
		"lshr":                          78,
		"ashr":                          79,
		"and":                           80,
		"or":                            81,
		"xor":                           82,
		"extractelement":                83,
		"insertelement":                 84,
		"shufflevector":                 85,
		"extractvalue":                  86,
		"insertvalue":                   87,
		"getelementptr":                 88,
		"trunc":                         89,
		"to":                            90,
		"zext":                          91,
		"sext":                          92,
		"fptrunc":                       93,
		"fpext":                         94,
		"fptoui":                        95,
		"fptosi":                        96,
		"uitofp":                        97,
		"sitofp":                        98,
		"ptrtoint":                      99,
		"inttoptr":                      100,
		"bitcast":                       101,
		"addrspacecast":                 102,
		"icmp":                          103,
		"fcmp":                          104,
		"select":                        105,
		"nuw":                           106,
		"nsw":                           107,
		"arcp":                          108,
		"fast":                          109,
		"ninf":                          110,
		"nnan":                          111,
		"nsz":                           112,
		"exact":                         113,
		"alloca":                        114,
		"load":                          115,
		"volatile":                      116,
		"store":                         117,
		"fence":                         118,
		"singlethread":                  119,
		"acq_rel":                       120,
		"acquire":                       121,
		"monotonic":                     122,
		"release":                       123,
		"seq_cst":                       124,
		"unordered":                     125,
		"cmpxchg":                       126,
		"weak":                          127,
		"atomicrmw":                     128,
		"max":                           129,
		"min":                           130,
		"nand":                          131,
		"umax":                          132,
		"umin":                          133,
		"xchg":                          134,
		"eq":                            135,
		"ne":                            136,
		"ugt":                           137,
		"uge":                           138,
		"ult":                           139,
		"ule":                           140,
		"sgt":                           141,
		"sge":                           142,
		"slt":                           143,
		"sle":                           144,
		"oeq":                           145,
		"ogt":                           146,
		"oge":                           147,
		"olt":                           148,
		"ole":                           149,
		"one":                           150,
		"ord":                           151,
		"ueq":                           152,
		"une":                           153,
		"uno":                           154,
		"phi":                           155,
		"call":                          156,
		"tail":                          157,
		"musttail":                      158,
		"notail":                        159,
		"va_arg":                        160,
		"landingpad":                    161,
		"cleanup":                       162,
		"catch":                         163,
		"filter":                        164,
		"catchpad":                      165,
		"within":                        166,
		"cleanuppad":                    167,
		"none":                          168,
		"ret":                           169,
		"br":                            170,
		"switch":                        171,
		"indirectbr":                    172,
		"invoke":                        173,
		"unwind":                        174,
		"resume":                        175,
		"catchswitch":                   176,
		"caller":                        177,
		"catchret":                      178,
		"from":                          179,
		"cleanupret":                    180,
		"unreachable":                   181,
		"appending":                     182,
		"available_externally":          183,
		"common":                        184,
		"internal":                      185,
		"linkonce":                      186,
		"linkonce_odr":                  187,
		"private":                       188,
		"weak_odr":                      189,
		"extern_weak":                   190,
		"external":                      191,
		"default":                       192,
		"hidden":                        193,
		"protected":                     194,
		"dllimport":                     195,
		"dllexport":                     196,
		"thread_local":                  197,
		"localdynamic":                  198,
		"initialexec":                   199,
		"localexec":                     200,
		"local_unnamed_addr":            201,
		"unnamed_addr":                  202,
		"section":                       203,
		"align":                         204,
		"gc":                            205,
		"prefix":                        206,
		"prologue":                      207,
		"personality":                   208,
		"byval":                         209,
		"dereferenceable":               210,
		"dereferenceable_or_null":       211,
		"inalloca":                      212,
		"inreg":                         213,
		"nest":                          214,
		"noalias":                       215,
		"nocapture":                     216,
		"nonnull":                       217,
		"readnone":                      218,
		"readonly":                      219,
		"returned":                      220,
		"signext":                       221,
		"sret":                          222,
		"swifterror":                    223,
		"swiftself":                     224,
		"writeonly":                     225,
		"zeroext":                       226,
		"alignstack":                    227,
		"allocsize":                     228,
		"alwaysinline":                  229,
		"argmemonly":                    230,
		"builtin":                       231,
		"cold":                          232,
		"convergent":                    233,
		"inaccessiblemem_or_argmemonly": 234,
		"inaccessiblememonly":           235,
		"inlinehint":                    236,
		"jumptable":                     237,
		"minsize":                       238,
		"naked":                         239,
		"nobuiltin":                     240,
		"noduplicate":                   241,
		"noimplicitfloat":               242,
		"noinline":                      243,
		"nonlazybind":                   244,
		"norecurse":                     245,
		"noredzone":                     246,
		"noreturn":                      247,
		"nounwind":                      248,
		"optnone":                       249,
		"optsize":                       250,
		"returns_twice":                 251,
		"safestack":                     252,
		"sanitize_address":              253,
		"sanitize_memory":               254,
		"sanitize_thread":               255,
		"ssp":                           256,
		"sspreq":                        257,
		"sspstrong":                     258,
		"uwtable":                       259,
		"amdgpu_cs":                     260,
		"amdgpu_gs":                     261,
		"amdgpu_kernel":                 262,
		"amdgpu_ps":                     263,
		"amdgpu_vs":                     264,
		"anyregcc":                      265,
		"arm_aapcs_vfpcc":               266,
		"arm_aapcscc":                   267,
		"arm_apcscc":                    268,
		"avr_intrcc":                    269,
		"avr_signalcc":                  270,
		"cc":                            271,
		"ccc":                           272,
		"coldcc":                        273,
		"cxx_fast_tlscc":                274,
		"fastcc":                        275,
		"ghccc":                         276,
		"hhvm_ccc":                      277,
		"hhvmcc":                        278,
		"intel_ocl_bicc":                279,
		"msp430_intrcc":                 280,
		"preserve_allcc":                281,
		"preserve_mostcc":               282,
		"ptx_device":                    283,
		"ptx_kernel":                    284,
		"spir_func":                     285,
		"spir_kernel":                   286,
		"swiftcc":                       287,
		"webkit_jscc":                   288,
		"x86_64_sysvcc":                 289,
		"x86_64_win64cc":                290,
		"x86_fastcallcc":                291,
		"x86_intrcc":                    292,
		"x86_regcallcc":                 293,
		"x86_stdcallcc":                 294,
		"x86_thiscallcc":                295,
		"x86_vectorcallcc":              296,
		"inbounds":                      297,
	},
}
