# React 마이그레이션 상세 계획서

## 📋 프로젝트 개요
- **현재**: Buffalo 기반 서버사이드 렌더링 + jQuery/Vanilla JS 프론트엔드
- **목표**: React 기반 SPA(Single Page Application)로 전환
- **전략**: 점진적 마이그레이션 (기존 백엔드 API 유지)
- **기간**: 8주 (2024년 12월 ~ 2025년 2월)

## 🚨 개발 워크플로우 규칙 (중요!)

### 개발 진행 원칙
1. **모든 개발 전 토론 필수**: 구현 전에 항상 사용자와 토론
2. **승인 후 진행**: 사용자가 "진행해!" 라고 승인한 후에만 개발 시작
3. **단계별 확인**: 각 단계 완료 후 다음 단계 진행 여부 확인
4. **문서화**: 모든 결정사항을 이 계획서에 기록

### 토론 항목
- [x] 네이밍 규칙 결정 ✅
- [x] ESLint 설정 완료 ✅
- [ ] 컴포넌트 구조 설계
- [ ] 상태 관리 방식
- [ ] API 연동 방식
- [ ] UI/UX 개선 방향
- [ ] 성능 최적화 전략
- [ ] 모듈화 전략 및 재사용성 설계
- [ ] workspace/project 박스 조건부 표시 로직
- [ ] MCI 생성 요청 후 상태 표현 방식
- [ ] shadcn/ui 컴포넌트 활용 전략

### 승인 체크리스트
- [x] 사용자 승인 확인 ✅
- [ ] 구현 범위 명확화
- [ ] 예상 소요 시간 합의
- [ ] 리스크 요소 검토

## 📝 네이밍 규칙 (확정)

### 변수명 패턴
```javascript
// ✅ React 표준 (camelCase)
const appState = { ... }
const userData = { ... }
const isLoading = false
const selectedUser = null

// ❌ 기존 패턴 (PascalCase)
const AppState = { ... }
const UserData = { ... }
```

### 함수명 패턴
```javascript
// ✅ React 표준 (camelCase)
function getUserData() { ... }
function handleUserClick() { ... }
function updateUserInfo() { ... }

// ❌ 기존 패턴
function GetUserData() { ... }
function HandleUserClick() { ... }
```

### 컴포넌트명 패턴
```javascript
// ✅ React 표준 (PascalCase)
function UserManagement() { ... }
function DataTable() { ... }
function LoginForm() { ... }
```

### 파일명 패턴
```javascript
// ✅ React 표준 (camelCase)
UserManagement.jsx
DataTable.jsx
useAuth.js
```

### 상수명 패턴
```javascript
// ✅ React 표준 (UPPER_SNAKE_CASE)
const API_ENDPOINTS = { ... }
const ROLE_TYPES = { ... }
const DEFAULT_CONFIG = { ... }
```

### 훅명 패턴
```javascript
// ✅ React 표준 (use 접두사)
function useAuth() { ... }
function useUserData() { ... }
function useTableState() { ... }
```

### API 함수명 패턴
```javascript
// ✅ 옵션 1: 간단하게
getUsers()
createUser()
updateUser()
deleteUser()
```

### 이벤트 핸들러 패턴
```javascript
// ✅ React 표준 (handle 접두사)
function handleUserClick() { ... }
function handleFormSubmit() { ... }
function handleInputChange() { ... }
```

### 상태 변수명 패턴
```javascript
// ✅ 옵션 1: is 접두사
const isLoading = false
const isVisible = true
const isSelected = false
const isAuthenticated = true
```

## 🛠 기술 스택

### 확정된 기술 스택
```json
{
  "프레임워크": "React 18",
  "상태 관리": "Context API + useReducer",
  "서버 상태": "직접 관리 (axios + useState)",
  "라우팅": "React Router v6",
  "폼 관리": "React Hook Form",
  "UI 라이브러리": "shadcn/ui (Tabler 테마)",
  "스타일링": "Tailwind CSS",
  "렌더링 방식": "CSR",
  "빌드 도구": "Vite",
  "언어": "JavaScript",
  "코드 품질": "ESLint + Prettier",
  "타입 체크": "JSDoc (향후 TypeScript 전환 고려)"
}
```

### 선택 근거
- **React 18**: 최신 기능과 성능 최적화
- **Context API**: 기존 AppState 패턴과 유사, 학습 곡선 낮음
- **React Hook Form**: 폼 검증과 성능 최적화
- **shadcn/ui**: Tabler 디자인과 유사, 컴포넌트 품질 높음
- **Vite**: 빠른 개발 환경과 빌드 최적화

## 🎯 마이그레이션 범위

### 기존 기능 호환성
- **80% 호환 + 20% 개선**
- **핵심 기능**: 100% 동일하게 동작
- **UI/UX**: shadcn/ui 활용하여 개선 허용
- **백엔드 API**: 기존 구조 유지 (변경 없음)

### 마이그레이션 대상
- **Tabulator 테이블** → **shadcn/ui Table**
- **Bootstrap 모달** → **shadcn/ui Dialog**
- **jQuery DOM 조작** → **React 상태 관리**
- **Hash 기반 라우팅** → **React Router**
- **ApexCharts** → **Recharts**

### 추가 요구사항
- **모듈화 및 재사용성**: 컴포넌트를 잘게 쪼개어 재사용 가능한 모듈로 개발
- **조건부 UI**: workspace/project 박스는 필요한 화면에만 표시되도록 조건부 렌더링 적용

### TODO 항목 (추후 토론 및 개발)
- [ ] **workspace/project 박스 조건부 표시**: 개발 후에 구현
- [ ] **MCI 생성 요청 후 상태 표현**: 실제 생성된 데이터가 오기 전까지 웹에서 표현 방식 토론 필요
- [ ] **UI/UX 개선**: 개발 진행하면서 점진적 개선

### 기존 코드 패턴 분석
- **공통 패턴 1**: list api 요청 → 테이블에 set → 단건 조회 → 표로 정보 보여주기
- **공통 패턴 2**: 생성 버튼 → 생성폼 → 생성 폼 구성 → spec, image 모달 검색 → 생성 요청 → 리스트
- **복잡한 페이지**: MCI, PMK 관련 HTML들과 JS들
- **중복 패턴**: 대부분의 기능이 유사한 패턴을 따름 (M-CMP 콘솔 특성상)

### 모듈화 우선순위
1. **테이블 컴포넌트** (shadcn/ui Table 활용)
2. **폼 컴포넌트** (shadcn/ui Form + React Hook Form)
3. **그리드 컴포넌트** (shadcn/ui Grid)
4. **버튼 컴포넌트** (shadcn/ui Button)
5. **카드 컴포넌트** (shadcn/ui Card)

## 📁 프로젝트 구조

```
front-react/
├── src/
│   ├── components/
│   │   ├── ui/              # shadcn/ui 컴포넌트
│   │   │   ├── button.jsx
│   │   │   ├── input.jsx
│   │   │   ├── table.jsx
│   │   │   ├── dialog.jsx
│   │   │   └── index.js
│   │   ├── layout/          # 레이아웃 컴포넌트
│   │   │   ├── Sidebar.jsx
│   │   │   ├── Navbar.jsx
│   │   │   ├── PageHeader.jsx
│   │   │   ├── Footer.jsx
│   │   │   └── Layout.jsx
│   │   └── common/          # 공통 컴포넌트
│   │       ├── DataTable.jsx    # Tabulator → shadcn/ui Table
│   │       ├── Modal.jsx        # Bootstrap → shadcn/ui Dialog
│   │       ├── LoadingSpinner.jsx
│   │       ├── Chart.jsx        # ApexCharts → Recharts
│   │       └── index.js
│   ├── pages/               # 페이지 컴포넌트
│   │   ├── auth/
│   │   │   └── Login.jsx
│   │   ├── dashboard/
│   │   │   └── Dashboard.jsx
│   │   ├── settings/
│   │   │   ├── users/
│   │   │   │   └── UserManagement.jsx
│   │   │   └── roles/
│   │   │       └── RoleManagement.jsx
│   │   └── operations/
│   │       ├── workspace/
│   │       │   └── WorkspaceManagement.jsx
│   │       └── manage/
│   │           └── MciManagement.jsx
│   ├── hooks/               # 커스텀 훅
│   │   ├── useAuth.js
│   │   ├── useApi.js
│   │   ├── useTable.js
│   │   ├── useLocalStorage.js
│   │   └── index.js
│   ├── services/            # API 서비스
│   │   ├── api.js          # 기본 API 클라이언트
│   │   ├── authService.js
│   │   ├── userService.js
│   │   ├── roleService.js
│   │   ├── mciService.js
│   │   └── index.js
│   ├── context/             # Context API
│   │   ├── AuthContext.js
│   │   ├── AppContext.js
│   │   ├── TableContext.js
│   │   └── index.js
│   ├── utils/               # 유틸리티 함수
│   │   ├── api.js          # 기존 http.js 변환
│   │   ├── tableUtils.js   # Tabulator → shadcn/ui 변환
│   │   ├── formatters.js   # 데이터 포맷터
│   │   ├── constants.js    # 상수 정의
│   │   └── index.js
│   ├── styles/              # 스타일
│   │   ├── globals.css
│   │   └── components.css
│   ├── App.jsx
│   ├── main.jsx
│   └── index.css
├── public/
├── package.json
├── vite.config.js
├── tailwind.config.js
├── eslint.config.js
├── prettier.config.js
└── index.html
```

## 🔄 마이그레이션 단계

### 1단계: 기반 구축 (1-2주차)
- [x] React 프로젝트 생성 (Vite)
- [x] ESLint + Prettier 설정
- [x] Tailwind CSS 설정
- [x] 네이밍 규칙 결정 ✅
- [x] ESLint 네이밍 규칙 업데이트 ✅
- [x] shadcn/ui 설치 및 설정 ✅
- [x] 프로젝트 구조 생성 ✅
- [x] 기본 레이아웃 구성 ✅
- [ ] API 클라이언트 설정
- [ ] 라우팅 설정
- [ ] Context API 설정

### 2단계: 인증 시스템 (3주차)
- [ ] 로그인 페이지 구현
- [ ] 인증 Context 설정
- [ ] 기존 API 연동
- [ ] 에러 처리
- [ ] 토큰 관리

### 3단계: 테이블 컴포넌트 (4주차)
- [ ] DataTable 컴포넌트 구현
- [ ] Tabulator → shadcn/ui Table 변환
- [ ] 페이지네이션 구현
- [ ] 검색/필터 기능
- [ ] 행 선택 기능

### 4단계: 사용자 관리 (5주차)
- [ ] 사용자 목록 테이블
- [ ] 사용자 생성/수정 모달
- [ ] CRUD 기능 구현
- [ ] 권한 관리

### 5단계: 역할 관리 (6주차)
- [ ] 역할 목록 테이블
- [ ] 역할 생성/수정 폼
- [ ] 권한 할당 기능
- [ ] CSP 역할 매핑

### 6단계: 대시보드 (7주차)
- [ ] 대시보드 레이아웃
- [ ] 데이터 표시 컴포넌트
- [ ] 차트 컴포넌트 (ApexCharts → Recharts)
- [ ] 실시간 데이터 업데이트

### 7단계: 통합 및 최적화 (8주차)
- [ ] 기존 시스템과 통합
- [ ] 성능 최적화
- [ ] 테스트 및 버그 수정
- [ ] 문서화

## 🔧 컴포넌트 변환 가이드

### Tabulator → shadcn/ui Table
```javascript
// 현재 Tabulator
const table = new Tabulator("#roles-table", {
  data: [],
  layout: "fitColumns",
  pagination: true,
  columns: [
    { title: "Name", field: "name" },
    { title: "Description", field: "description" }
  ]
});

// React shadcn/ui Table
const columns = [
  { key: 'name', title: 'Name' },
  { key: 'description', title: 'Description' }
];

<DataTable data={data} columns={columns} />
```

### Bootstrap 모달 → shadcn/ui Dialog
```javascript
// 현재 Bootstrap
$("#add-user-modal").modal('show');

// React shadcn/ui Dialog
const [isOpen, setIsOpen] = useState(false);

<Dialog open={isOpen} onOpenChange={setIsOpen}>
  <DialogContent>
    <DialogHeader>
      <DialogTitle>Add User</DialogTitle>
    </DialogHeader>
    {/* 폼 내용 */}
  </DialogContent>
</Dialog>
```

### AppState → React Context
```javascript
// 현재 AppState
const AppState = {
  users: { list: [], selectedUser: null },
  ui: { viewMode: false, createMode: false }
};

// React Context
const AppContext = createContext();
const [state, dispatch] = useReducer(appReducer, initialState);
```

## 🚀 개발 환경 설정

### 필수 패키지
```bash
# React 프로젝트 생성
npm create vite@latest front-react -- --template react

# 필수 의존성
npm install react-router-dom axios react-hook-form
npm install -D tailwindcss postcss autoprefixer eslint prettier

# shadcn/ui
npm install @radix-ui/react-dialog @radix-ui/react-dropdown-menu
npm install lucide-react class-variance-authority clsx tailwind-merge

# 추가 라이브러리
npm install recharts react-hook-form @hookform/resolvers yup
```

### 개발 서버 설정
```bash
# 백엔드 (기존)
cd api && buffalo dev  # 포트 3000

# React 프론트엔드 (새로운)
cd front-react && npm run dev  # 포트 5173
```

## ⚠️ 주의사항 및 리스크

### 기술적 리스크
- **상태 관리 복잡성**: Context API로 단순화
- **기존 라이브러리 호환성**: Tabulator → shadcn/ui Table
- **API 호환성**: 기존 백엔드 API 유지

### 개발 리스크
- **학습 곡선**: 단순한 컴포넌트부터 시작
- **기존 기능 호환성**: 점진적 마이그레이션으로 최소화

### 성능 고려사항
- **번들 크기**: 코드 스플리팅 적용
- **초기 로딩**: CSR 특성상 초기 로딩 시간 고려
- **실시간 업데이트**: 대시보드 등 실시간 기능 최적화

## 📊 성공 지표

### 기능적 지표
- [ ] 기존 기능 100% 동작
- [ ] 새로운 UI/UX 개선
- [ ] 성능 향상 (로딩 속도, 반응성)

### 개발 지표
- [ ] 코드 재사용성 향상
- [ ] 유지보수성 개선
- [ ] 개발 생산성 향상

## 🔗 참고 자료

- [shadcn/ui 공식 문서](https://ui.shadcn.com/)
- [React Router 문서](https://reactrouter.com/)
- [Vite 문서](https://vitejs.dev/)
- [Tailwind CSS 문서](https://tailwindcss.com/)
- [React Hook Form 문서](https://react-hook-form.com/)

---

**마지막 업데이트**: 2024년 12월
**버전**: 1.0
**담당자**: 개발팀 