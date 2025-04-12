module.exports = {
	extends: [
		'plugin:react/recommended',
		'plugin:@typescript-eslint/recommended',
		'prettier',
	],
	plugins: ["prettier", "import"],
	rules: {
		"prettier/prettier": "error",
		'import/order': [
			'warn',
			{
				groups: [
					'builtin',
					'external',
					'internal',
					['parent', 'sibling', 'index'], 
					'object',
					'type',
					'unknown',
				],
				pathGroups: [
					{
						pattern: '*.css',
						group: 'unknown',
						position: 'after',
					},
				],
				pathGroupsExcludedImportTypes: ['builtin'],
				'newlines-between': 'always',
				alphabetize: {
					order: 'asc',
					caseInsensitive: true,
				},
			},
		],
	},
};
