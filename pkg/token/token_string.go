// Code generated by "stringer -type=ID -output ./token_string.go"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[T_INCLUDE-57346]
	_ = x[T_INCLUDE_ONCE-57347]
	_ = x[T_EXIT-57348]
	_ = x[T_IF-57349]
	_ = x[T_LNUMBER-57350]
	_ = x[T_DNUMBER-57351]
	_ = x[T_STRING-57352]
	_ = x[T_STRING_VARNAME-57353]
	_ = x[T_VARIABLE-57354]
	_ = x[T_NUM_STRING-57355]
	_ = x[T_INLINE_HTML-57356]
	_ = x[T_CHARACTER-57357]
	_ = x[T_BAD_CHARACTER-57358]
	_ = x[T_ENCAPSED_AND_WHITESPACE-57359]
	_ = x[T_CONSTANT_ENCAPSED_STRING-57360]
	_ = x[T_ECHO-57361]
	_ = x[T_DO-57362]
	_ = x[T_WHILE-57363]
	_ = x[T_ENDWHILE-57364]
	_ = x[T_FOR-57365]
	_ = x[T_ENDFOR-57366]
	_ = x[T_FOREACH-57367]
	_ = x[T_ENDFOREACH-57368]
	_ = x[T_DECLARE-57369]
	_ = x[T_ENDDECLARE-57370]
	_ = x[T_AS-57371]
	_ = x[T_SWITCH-57372]
	_ = x[T_ENDSWITCH-57373]
	_ = x[T_CASE-57374]
	_ = x[T_DEFAULT-57375]
	_ = x[T_BREAK-57376]
	_ = x[T_CONTINUE-57377]
	_ = x[T_GOTO-57378]
	_ = x[T_FUNCTION-57379]
	_ = x[T_FN-57380]
	_ = x[T_CONST-57381]
	_ = x[T_RETURN-57382]
	_ = x[T_TRY-57383]
	_ = x[T_CATCH-57384]
	_ = x[T_FINALLY-57385]
	_ = x[T_THROW-57386]
	_ = x[T_USE-57387]
	_ = x[T_INSTEADOF-57388]
	_ = x[T_GLOBAL-57389]
	_ = x[T_VAR-57390]
	_ = x[T_UNSET-57391]
	_ = x[T_ISSET-57392]
	_ = x[T_EMPTY-57393]
	_ = x[T_HALT_COMPILER-57394]
	_ = x[T_CLASS-57395]
	_ = x[T_TRAIT-57396]
	_ = x[T_INTERFACE-57397]
	_ = x[T_EXTENDS-57398]
	_ = x[T_IMPLEMENTS-57399]
	_ = x[T_OBJECT_OPERATOR-57400]
	_ = x[T_DOUBLE_ARROW-57401]
	_ = x[T_LIST-57402]
	_ = x[T_ARRAY-57403]
	_ = x[T_CALLABLE-57404]
	_ = x[T_CLASS_C-57405]
	_ = x[T_TRAIT_C-57406]
	_ = x[T_METHOD_C-57407]
	_ = x[T_FUNC_C-57408]
	_ = x[T_LINE-57409]
	_ = x[T_FILE-57410]
	_ = x[T_COMMENT-57411]
	_ = x[T_DOC_COMMENT-57412]
	_ = x[T_OPEN_TAG-57413]
	_ = x[T_OPEN_TAG_WITH_ECHO-57414]
	_ = x[T_CLOSE_TAG-57415]
	_ = x[T_WHITESPACE-57416]
	_ = x[T_START_HEREDOC-57417]
	_ = x[T_END_HEREDOC-57418]
	_ = x[T_DOLLAR_OPEN_CURLY_BRACES-57419]
	_ = x[T_CURLY_OPEN-57420]
	_ = x[T_PAAMAYIM_NEKUDOTAYIM-57421]
	_ = x[T_NAMESPACE-57422]
	_ = x[T_NS_C-57423]
	_ = x[T_DIR-57424]
	_ = x[T_NS_SEPARATOR-57425]
	_ = x[T_ELLIPSIS-57426]
	_ = x[T_EVAL-57427]
	_ = x[T_REQUIRE-57428]
	_ = x[T_REQUIRE_ONCE-57429]
	_ = x[T_LOGICAL_OR-57430]
	_ = x[T_LOGICAL_XOR-57431]
	_ = x[T_LOGICAL_AND-57432]
	_ = x[T_INSTANCEOF-57433]
	_ = x[T_NEW-57434]
	_ = x[T_CLONE-57435]
	_ = x[T_ELSEIF-57436]
	_ = x[T_ELSE-57437]
	_ = x[T_ENDIF-57438]
	_ = x[T_PRINT-57439]
	_ = x[T_YIELD-57440]
	_ = x[T_STATIC-57441]
	_ = x[T_ABSTRACT-57442]
	_ = x[T_FINAL-57443]
	_ = x[T_PRIVATE-57444]
	_ = x[T_PROTECTED-57445]
	_ = x[T_PUBLIC-57446]
	_ = x[T_INC-57447]
	_ = x[T_DEC-57448]
	_ = x[T_YIELD_FROM-57449]
	_ = x[T_INT_CAST-57450]
	_ = x[T_DOUBLE_CAST-57451]
	_ = x[T_STRING_CAST-57452]
	_ = x[T_ARRAY_CAST-57453]
	_ = x[T_OBJECT_CAST-57454]
	_ = x[T_BOOL_CAST-57455]
	_ = x[T_UNSET_CAST-57456]
	_ = x[T_COALESCE-57457]
	_ = x[T_SPACESHIP-57458]
	_ = x[T_NOELSE-57459]
	_ = x[T_PLUS_EQUAL-57460]
	_ = x[T_MINUS_EQUAL-57461]
	_ = x[T_MUL_EQUAL-57462]
	_ = x[T_POW_EQUAL-57463]
	_ = x[T_DIV_EQUAL-57464]
	_ = x[T_CONCAT_EQUAL-57465]
	_ = x[T_MOD_EQUAL-57466]
	_ = x[T_AND_EQUAL-57467]
	_ = x[T_OR_EQUAL-57468]
	_ = x[T_XOR_EQUAL-57469]
	_ = x[T_SL_EQUAL-57470]
	_ = x[T_SR_EQUAL-57471]
	_ = x[T_COALESCE_EQUAL-57472]
	_ = x[T_BOOLEAN_OR-57473]
	_ = x[T_BOOLEAN_AND-57474]
	_ = x[T_POW-57475]
	_ = x[T_SL-57476]
	_ = x[T_SR-57477]
	_ = x[T_IS_IDENTICAL-57478]
	_ = x[T_IS_NOT_IDENTICAL-57479]
	_ = x[T_IS_EQUAL-57480]
	_ = x[T_IS_NOT_EQUAL-57481]
	_ = x[T_IS_SMALLER_OR_EQUAL-57482]
	_ = x[T_IS_GREATER_OR_EQUAL-57483]
	_ = x[T_NULLSAFE_OBJECT_OPERATOR-57484]
	_ = x[T_MATCH-57485]
	_ = x[T_ATTRIBUTE-57486]
	_ = x[T_NAME_RELATIVE-57487]
	_ = x[T_NAME_QUALIFIED-57488]
	_ = x[T_NAME_FULLY_QUALIFIED-57489]
	_ = x[T_READONLY-57490]
	_ = x[T_ENUM-57491]
	_ = x[T_AMPERSAND_FOLLOWED_BY_VAR_OR_VARARG-57492]
	_ = x[T_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG-57493]
}

const _ID_name = "T_INCLUDET_INCLUDE_ONCET_EXITT_IFT_LNUMBERT_DNUMBERT_STRINGT_STRING_VARNAMET_VARIABLET_NUM_STRINGT_INLINE_HTMLT_CHARACTERT_BAD_CHARACTERT_ENCAPSED_AND_WHITESPACET_CONSTANT_ENCAPSED_STRINGT_ECHOT_DOT_WHILET_ENDWHILET_FORT_ENDFORT_FOREACHT_ENDFOREACHT_DECLARET_ENDDECLARET_AST_SWITCHT_ENDSWITCHT_CASET_DEFAULTT_BREAKT_CONTINUET_GOTOT_FUNCTIONT_FNT_CONSTT_RETURNT_TRYT_CATCHT_FINALLYT_THROWT_USET_INSTEADOFT_GLOBALT_VART_UNSETT_ISSETT_EMPTYT_HALT_COMPILERT_CLASST_TRAITT_INTERFACET_EXTENDST_IMPLEMENTST_OBJECT_OPERATORT_DOUBLE_ARROWT_LISTT_ARRAYT_CALLABLET_CLASS_CT_TRAIT_CT_METHOD_CT_FUNC_CT_LINET_FILET_COMMENTT_DOC_COMMENTT_OPEN_TAGT_OPEN_TAG_WITH_ECHOT_CLOSE_TAGT_WHITESPACET_START_HEREDOCT_END_HEREDOCT_DOLLAR_OPEN_CURLY_BRACEST_CURLY_OPENT_PAAMAYIM_NEKUDOTAYIMT_NAMESPACET_NS_CT_DIRT_NS_SEPARATORT_ELLIPSIST_EVALT_REQUIRET_REQUIRE_ONCET_LOGICAL_ORT_LOGICAL_XORT_LOGICAL_ANDT_INSTANCEOFT_NEWT_CLONET_ELSEIFT_ELSET_ENDIFT_PRINTT_YIELDT_STATICT_ABSTRACTT_FINALT_PRIVATET_PROTECTEDT_PUBLICT_INCT_DECT_YIELD_FROMT_INT_CASTT_DOUBLE_CASTT_STRING_CASTT_ARRAY_CASTT_OBJECT_CASTT_BOOL_CASTT_UNSET_CASTT_COALESCET_SPACESHIPT_NOELSET_PLUS_EQUALT_MINUS_EQUALT_MUL_EQUALT_POW_EQUALT_DIV_EQUALT_CONCAT_EQUALT_MOD_EQUALT_AND_EQUALT_OR_EQUALT_XOR_EQUALT_SL_EQUALT_SR_EQUALT_COALESCE_EQUALT_BOOLEAN_ORT_BOOLEAN_ANDT_POWT_SLT_SRT_IS_IDENTICALT_IS_NOT_IDENTICALT_IS_EQUALT_IS_NOT_EQUALT_IS_SMALLER_OR_EQUALT_IS_GREATER_OR_EQUALT_NULLSAFE_OBJECT_OPERATORT_MATCHT_ATTRIBUTET_NAME_RELATIVET_NAME_QUALIFIEDT_NAME_FULLY_QUALIFIEDT_READONLYT_ENUMT_AMPERSAND_FOLLOWED_BY_VAR_OR_VARARGT_AMPERSAND_NOT_FOLLOWED_BY_VAR_OR_VARARG"

var _ID_index = [...]uint16{0, 9, 23, 29, 33, 42, 51, 59, 75, 85, 97, 110, 121, 136, 161, 187, 193, 197, 204, 214, 219, 227, 236, 248, 257, 269, 273, 281, 292, 298, 307, 314, 324, 330, 340, 344, 351, 359, 364, 371, 380, 387, 392, 403, 411, 416, 423, 430, 437, 452, 459, 466, 477, 486, 498, 515, 529, 535, 542, 552, 561, 570, 580, 588, 594, 600, 609, 622, 632, 652, 663, 675, 690, 703, 729, 741, 763, 774, 780, 785, 799, 809, 815, 824, 838, 850, 863, 876, 888, 893, 900, 908, 914, 921, 928, 935, 943, 953, 960, 969, 980, 988, 993, 998, 1010, 1020, 1033, 1046, 1058, 1071, 1082, 1094, 1104, 1115, 1123, 1135, 1148, 1159, 1170, 1181, 1195, 1206, 1217, 1227, 1238, 1248, 1258, 1274, 1286, 1299, 1304, 1308, 1312, 1326, 1344, 1354, 1368, 1389, 1410, 1436, 1443, 1454, 1469, 1485, 1507, 1517, 1523, 1560, 1601}

func (i ID) String() string {
	i -= 57346
	if i < 0 || i >= ID(len(_ID_index)-1) {
		return "ID(" + strconv.FormatInt(int64(i+57346), 10) + ")"
	}
	return _ID_name[_ID_index[i]:_ID_index[i+1]]
}
