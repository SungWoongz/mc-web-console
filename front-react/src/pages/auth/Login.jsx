import { LoginForm } from '@/components/login-form'

/**
 * 로그인 페이지 컴포넌트
 * @param {Object} props - 컴포넌트 props
 * @param {Function} props.onLogin - 로그인 성공 시 호출되는 함수
 * @returns {JSX.Element} 로그인 페이지
 */
function Login({ onLogin }) {
  const handleLogin = (formData) => {
    console.log('Login attempt:', formData)
    // TODO: API 연동 및 토큰 관리 구현 필요
    // 현재는 바로 로그인 성공으로 처리
    if (onLogin) {
      onLogin()
    }
  }

  return (
    <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-sm">
        <LoginForm onSubmit={handleLogin} />
      </div>
    </div>
  )
}

export default Login
