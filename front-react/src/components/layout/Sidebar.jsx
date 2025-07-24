import { Card, CardContent } from '@/components/ui'

/**
 * 사이드바 컴포넌트
 * @returns {JSX.Element} 사이드바 컴포넌트
 */
function Sidebar() {
  return (
    <aside className="w-64 bg-card border-r border-border p-4">
      <Card className="border-0 shadow-none">
        <CardContent className="p-0">
          <div className="space-y-4">
            <div className="flex items-center space-x-2">
              <div className="w-8 h-8 bg-primary rounded-lg flex items-center justify-center">
                <span className="text-primary-foreground font-bold text-sm">M</span>
              </div>
              <h2 className="text-lg font-semibold">MC-Web-Console</h2>
            </div>

            <nav className="space-y-2">
              {/* TODO: 네비게이션 메뉴 추가 예정 */}
              <div className="text-sm text-muted-foreground">
                메뉴가 여기에 추가됩니다
              </div>
            </nav>
          </div>
        </CardContent>
      </Card>
    </aside>
  )
}

export default Sidebar
