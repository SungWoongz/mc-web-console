# React 마이그레이션 계획서

## 📋 프로젝트 개요
- **현재**: Buffalo 기반 서버사이드 렌더링 + jQuery/Vanilla JS 프론트엔드
- **목표**: React 기반 SPA(Single Page Application)로 전환
- **전략**: 점진적 마이그레이션 (기존 백엔드 API 유지)

## 🛠 기술 스택

### 확정된 기술 스택
```json
{
  "상태 관리": "Context API",
  "서버 상태": "직접 관리 (axios + useState)",
  "라우팅": "React Router",
  "폼 관리": "직접 관리 (useState)",
  "UI 라이브러리": "shadcn/ui (Tabler 테마)",
  "렌더링 방식": "CSR",
  "빌드 도구": "Vite",
  "언어": "JavaScript"
}
```

### 선택 근거
- **Context API**: 현재 AppState 패턴과 유사, 학습 곡선 낮음
- **직접 관리**: 단순하고 직관적, 기존 axios 패턴 재사용
- **React Router**: 가장 널리 사용, 문서화 잘됨
- **shadcn/ui**: Tabler 디자인과 유사, 컴포넌트 품질 높음
- **CSR**: 실시간 기능에 유리, 점진적 마이그레이션 적합

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

## 📁 프로젝트 구조

```
front-react/
├── src/
│   ├── components/
│   │   ├── ui/              # shadcn/ui 컴포넌트
│   │   ├── layout/          # 레이아웃 컴포넌트
│   │   └── common/          # 공통 컴포넌트
│   ├── pages/               # 페이지 컴포넌트
│   │   ├── auth/            # 인증 관련
│   │   ├── dashboard/       # 대시보드
│   │   ├── settings/        # 설정
│   │   └── configuration/   # 구성 관리
│   ├── hooks/               # 커스텀 훅
│   ├── services/            # API 서비스
│   ├── context/             # Context API
│   ├── utils/               # 유틸리티 함수
│   └── styles/              # 스타일
├── public/
├── package.json
├── vite.config.js
├── tailwind.config.js
└── index.html
```

## 🔄 마이그레이션 단계

### 1단계: 기반 구축 (1주차)
- [ ] React 프로젝트 생성 (Vite)
- [ ] shadcn/ui 설치 및 설정
- [ ] 기본 레이아웃 구성
- [ ] API 클라이언트 설정
- [ ] 라우팅 설정

### 2단계: 로그인 페이지 (2주차)
- [ ] 로그인 폼 구현
- [ ] 인증 Context 설정
- [ ] 기존 API 연동
- [ ] 에러 처리

### 3단계: 대시보드 (3주차)
- [ ] 대시보드 레이아웃
- [ ] 데이터 표시 컴포넌트
- [ ] 차트 컴포넌트 (ApexCharts → Recharts)
- [ ] 실시간 데이터 업데이트

### 4단계: 사용자 관리 (4주차)
- [ ] 사용자 목록 테이블
- [ ] 사용자 생성/수정 모달
- [ ] CRUD 기능 구현
- [ ] 권한 관리

### 5단계: 설정 페이지 (5주차)
- [ ] 복잡한 폼 구현
- [ ] 파일 업로드 (Dropzone → React)
- [ ] 설정 저장/로드

### 6단계: 통합 및 최적화 (6주차)
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
<Table>
  <TableHeader>
    <TableRow>
      <TableHead>Name</TableHead>
      <TableHead>Description</TableHead>
    </TableRow>
  </TableHeader>
  <TableBody>
    {data.map((role) => (
      <TableRow key={role.id}>
        <TableCell>{role.name}</TableCell>
        <TableCell>{role.description}</TableCell>
      </TableRow>
    ))}
  </TableBody>
</Table>
```

### Bootstrap 모달 → shadcn/ui Dialog
```javascript
// 현재 Bootstrap
$("#add-user-modal").modal('show');

// React shadcn/ui Dialog
<Dialog open={isOpen} onOpenChange={onClose}>
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
npm install react-router-dom axios
npm install -D tailwindcss postcss autoprefixer

# shadcn/ui
npm install @radix-ui/react-dialog @radix-ui/react-dropdown-menu
npm install lucide-react class-variance-authority clsx tailwind-merge

# 추가 라이브러리
npm install recharts react-hook-form
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

---

**마지막 업데이트**: 2024년 12월
**버전**: 1.0 