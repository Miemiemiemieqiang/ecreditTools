INSERT INTO vaas.variable_mapping(mapping) VALUES ('{
	"name": "max_query_cnt_6h_2wk",
	"returnType": "Integer",
	"description": "变量:近2周在6小时内查询的峰值次数",
	"sourceHandler": [
		{
			"maxQueryCnt6h2wk": []
		}
	],
	"variableType": "numerical"
}'), ('{
	"name": "id_query_cnt_B_3d",
	"returnType": "Integer",
	"description": "变量:近3天身份证B类查询次数",
	"sourceHandler": [
		{
			"idQueryCntB3d": []
		}
	],
	"variableType": "numerical"
}'), ('{
	"name": "id_query_cnt_B_12h",
	"returnType": "Integer",
	"description": "变量:近12小时身份证B类查询次数",
	"sourceHandler": [
		{
			"idQueryCntB12h": []
		}
	],
	"variableType": "numerical"
}'), ('{
	"name": "phone_query_cnt_B_3d",
	"returnType": "Integer",
	"description": "变量:近3天手机号B类查询次数",
	"sourceHandler": [
		{
			"phoneQueryCntB3d": []
		}
	],
	"variableType": "numerical"
}'), ('{
	"name": "phone_query_cnt_B_24h",
	"returnType": "Integer",
	"description": "变量:近24小时手机号B类查询次数",
	"sourceHandler": [
		{
			"phoneQueryCntB24h": []
		}
	],
	"variableType": "numerical"
}');