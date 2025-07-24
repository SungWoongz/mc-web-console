import js from '@eslint/js'
import globals from 'globals'
import react from 'eslint-plugin-react'
import reactHooks from 'eslint-plugin-react-hooks'
import reactRefresh from 'eslint-plugin-react-refresh'

export default [
  js.configs.recommended,
  {
    files: ['**/*.{js,jsx}'],
    languageOptions: {
      globals: {
        ...globals.browser,
        ...globals.es2021,
      },
      parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module',
        ecmaFeatures: {
          jsx: true,
        },
      },
    },
    plugins: {
      'react': react,
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
    },
    settings: {
      react: {
        version: 'detect',
      },
    },
    rules: {
      ...react.configs.recommended.rules,
      ...reactHooks.configs.recommended.rules,
      'react-refresh/only-export-components': [
        'warn',
        { allowConstantExport: true },
      ],
      
      // React 17+ JSX 변환 지원
      'react/react-in-jsx-scope': 'off',
      'react/jsx-uses-react': 'off',
      
      // PropTypes 검증 (선택적)
      'react/prop-types': 'off',
      
      // 보안 관련 규칙
      'react/jsx-no-target-blank': 'warn',
      
      // ===== 기본 코드 품질 규칙 =====
      'no-unused-vars': 'warn',
      'no-console': 'warn',
      'prefer-const': 'error',
      'no-var': 'error',
      'no-duplicate-imports': 'error',
      'no-unused-expressions': 'error',
      'prefer-template': 'error',
      
      // ===== 확정된 네이밍 규칙 =====
      
      // 1. 변수명 패턴 (camelCase)
      'camelcase': ['error', { 
        properties: 'always',
        ignoreDestructuring: false,
        ignoreImports: false
      }],
      'id-length': ['warn', { min: 2, max: 30 }],
      'id-match': [
        'error',
        '^[a-zA-Z_$][a-zA-Z0-9_$]*$',
        { properties: true }
      ],
      
      // 2. 함수명 패턴 (camelCase)
      'func-name-matching': 'error',
      
      // 3. 이벤트 핸들러 패턴 (handle 접두사) - 커스텀 규칙으로 처리
      'func-names': 'off',
      
      // 4. 상태 변수명 패턴 (is 접두사) - camelCase 규칙으로 처리됨
      
      // 5. 상수명 패턴 (UPPER_SNAKE_CASE) - camelCase 규칙에서 예외 처리
      
      // 6. 훅명 패턴 (use 접두사)
      'react-hooks/rules-of-hooks': 'error',
      'react-hooks/exhaustive-deps': 'warn',
      
      // 7. 컴포넌트명 패턴 (PascalCase)
      'react/jsx-pascal-case': 'error',
      'react/jsx-filename-extension': [1, { extensions: ['.jsx', '.js'] }],
      
      // 8. API 함수명 패턴 (간단하게)
      'id-denylist': [
        'error',
        'data',
        'err',
        'e',
        'cb',
        'callback'
      ],
      
      // 9. 변수 네이밍 규칙 (기존 AppState → appState)
      'no-underscore-dangle': ['error', { 
        allow: ['_id', '__typename'],
        allowAfterThis: false,
        allowAfterSuper: false
      }],
      
      // ===== 코드 스타일 규칙 =====
      'template-curly-spacing': 'error',
      'object-curly-spacing': ['error', 'always'],
      'array-bracket-spacing': ['error', 'never'],
      'comma-dangle': ['error', 'always-multiline'],
      'semi': ['error', 'never'],
      'quotes': ['error', 'single'],
      'indent': ['error', 2],
      'no-trailing-spaces': 'error',
      'eol-last': 'error',
    },
  },
  {
    ignores: ['dist', 'node_modules'],
  },
]

