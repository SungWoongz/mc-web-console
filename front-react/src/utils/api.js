import axios from 'axios'
import { DEFAULT_CONFIG } from './constants'

// axios 인스턴스 생성
const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:3000',
  timeout: DEFAULT_CONFIG.TIMEOUT,
  headers: {
    'Content-Type': 'application/json',
  },
})

// 요청 인터셉터
apiClient.interceptors.request.use(
  (config) => {
    // TODO: 인증 토큰 추가 로직
    const token = localStorage.getItem('authToken')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

// 응답 인터셉터
apiClient.interceptors.response.use(
  (response) => {
    return response
  },
  async (error) => {
    const originalRequest = error.config

    // 401 에러 처리 (토큰 만료)
    if (error.response?.status === 401 && !originalRequest.retry) {
      originalRequest.retry = true

      try {
        // TODO: 토큰 갱신 로직
        // const refreshToken = localStorage.getItem('refreshToken')
        // const response = await apiClient.post('/auth/refresh', { refreshToken })
        // localStorage.setItem('authToken', response.data.token)

        // 원래 요청 재시도
        return apiClient(originalRequest)
      } catch (refreshError) {
        // 토큰 갱신 실패 시 로그아웃
        localStorage.removeItem('authToken')
        localStorage.removeItem('refreshToken')
        window.location.href = '/login'
        return Promise.reject(refreshError)
      }
    }

    return Promise.reject(error)
  },
)

/**
 * GET 요청 함수
 * @param {string} url - 요청 URL
 * @param {Object} config - axios 설정
 * @returns {Promise} axios 응답
 */
export const apiGet = async (url, config = {}) => {
  const response = await apiClient.get(url, config)
  return response.data
}

/**
 * POST 요청 함수
 * @param {string} url - 요청 URL
 * @param {Object} requestData - 요청 데이터
 * @param {Object} config - axios 설정
 * @returns {Promise} axios 응답
 */
export const apiPost = async (url, requestData = {}, config = {}) => {
  const response = await apiClient.post(url, requestData, config)
  return response.data
}

/**
 * PUT 요청 함수
 * @param {string} url - 요청 URL
 * @param {Object} requestData - 요청 데이터
 * @param {Object} config - axios 설정
 * @returns {Promise} axios 응답
 */
export const apiPut = async (url, requestData = {}, config = {}) => {
  const response = await apiClient.put(url, requestData, config)
  return response.data
}

/**
 * DELETE 요청 함수
 * @param {string} url - 요청 URL
 * @param {Object} config - axios 설정
 * @returns {Promise} axios 응답
 */
export const apiDelete = async (url, config = {}) => {
  const response = await apiClient.delete(url, config)
  return response.data
}

export default apiClient
