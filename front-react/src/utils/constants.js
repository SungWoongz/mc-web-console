// API 엔드포인트 상수
export const API_ENDPOINTS = {
  AUTH: {
    LOGIN: '/auth/login',
    LOGOUT: '/auth/logout',
    REFRESH: '/auth/refresh',
  },
  USERS: {
    LIST: '/users',
    CREATE: '/users',
    UPDATE: '/users/:id',
    DELETE: '/users/:id',
  },
  ROLES: {
    LIST: '/roles',
    CREATE: '/roles',
    UPDATE: '/roles/:id',
    DELETE: '/roles/:id',
  },
  WORKSPACES: {
    LIST: '/workspaces',
    CREATE: '/workspaces',
    UPDATE: '/workspaces/:id',
    DELETE: '/workspaces/:id',
  },
}

// 역할 타입 상수
export const ROLE_TYPES = {
  ADMIN: 'admin',
  USER: 'user',
  GUEST: 'guest',
}

// 기본 설정 상수
export const DEFAULT_CONFIG = {
  PAGE_SIZE: 10,
  MAX_PAGE_SIZE: 100,
  RETRY_COUNT: 3,
  TIMEOUT: 30000,
}

// 상태 상수
export const STATUS = {
  LOADING: 'loading',
  SUCCESS: 'success',
  ERROR: 'error',
  IDLE: 'idle',
}
