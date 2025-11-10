import eslintPluginYml from 'eslint-plugin-yml';

export default [
  // Apply the "standard" YAML rules to all .yml/.yaml files
  ...eslintPluginYml.configs['flat/standard'],

  // Add your custom rules
  {
    rules: {
      'yml/indent': ['error', 2],
      'yml/quotes': ['error', { prefer: 'single' }]
    }
  },

  // Global ignores
  {
    ignores: [
      'node_modules/',
      '.github/workflows/**',
      '.golangci.yml' // golangci-lint has its own schema that doesn't match ESLint YAML plugin
    ]
  }
];
