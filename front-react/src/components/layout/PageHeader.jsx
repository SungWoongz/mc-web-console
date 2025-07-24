/**
 * 페이지 헤더 컴포넌트
 * @param {Object} props - 컴포넌트 props
 * @param {string} props.title - 페이지 제목
 * @param {string} props.description - 페이지 설명
 * @param {React.ReactNode} props.children - 추가 컨텐츠
 * @returns {JSX.Element} 페이지 헤더 컴포넌트
 */
function PageHeader({ title, description, children }) {
  return (
    <div className="mb-8">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">{title}</h1>
          {description && (
            <p className="text-muted-foreground mt-2">{description}</p>
          )}
        </div>
        {children && <div className="flex items-center space-x-2">{children}</div>}
      </div>
    </div>
  )
}

export default PageHeader
