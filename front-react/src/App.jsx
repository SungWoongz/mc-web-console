import { Layout } from '@/components/layout'
import Dashboard from '@/pages/dashboard/Dashboard'
import Login from '@/pages/auth/Login'
import { useState } from 'react'

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false)

  // 인증 상태에 따라 다른 페이지 표시
  if (!isAuthenticated) {
    return <Login onLogin={() => setIsAuthenticated(true)} />
  }

  return (
    <Layout>
      <Dashboard />
    </Layout>
  )
}

export default App
