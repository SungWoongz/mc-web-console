import { PageHeader } from '@/components/layout'
import { Card, CardContent, CardHeader, CardTitle, Badge } from '@/components/ui'

/**
 * 대시보드 페이지 컴포넌트
 * @returns {JSX.Element} 대시보드 페이지
 */
function Dashboard() {
  return (
    <div className="container mx-auto px-6 py-8">
      <PageHeader
        title="대시보드"
        description="MC-Web-Console 대시보드입니다"
      />

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">시스템 상태</CardTitle>
            <Badge variant="default">정상</Badge>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">정상</div>
            <p className="text-xs text-muted-foreground">
              모든 서비스가 정상적으로 운영 중입니다
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">활성 사용자</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">0명</div>
            <p className="text-xs text-muted-foreground">
              현재 접속 중인 사용자 수
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">총 워크스페이스</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">0개</div>
            <p className="text-xs text-muted-foreground">
              생성된 워크스페이스 수
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">총 역할</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">0개</div>
            <p className="text-xs text-muted-foreground">
              생성된 역할 수
            </p>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}

export default Dashboard
