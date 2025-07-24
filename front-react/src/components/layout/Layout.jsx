import { Sidebar, Navbar, Footer } from './index'

/**
 * 메인 레이아웃 컴포넌트
 * @param {Object} props - 컴포넌트 props
 * @param {React.ReactNode} props.children - 자식 컴포넌트들
 * @returns {JSX.Element} 레이아웃 컴포넌트
 */
function Layout({ children }) {
  return (
    <div className="min-h-screen bg-background flex">
      <Sidebar />
      <div className="flex-1 flex flex-col">
        <Navbar />
        <main className="flex-1 p-6">
          {children}
        </main>
        <Footer />
      </div>
    </div>
  )
}

export default Layout
