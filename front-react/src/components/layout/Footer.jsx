/**
 * 푸터 컴포넌트
 * @returns {JSX.Element} 푸터 컴포넌트
 */
function Footer() {
  return (
    <footer className="border-t border-border bg-card">
      <div className="container mx-auto px-6 py-4">
        <div className="flex items-center justify-between">
          <p className="text-sm text-muted-foreground">
            © 2024 MC-Web-Console. All rights reserved.
          </p>
          <div className="flex items-center space-x-4 text-sm text-muted-foreground">
            <span>Version 1.0.0</span>
          </div>
        </div>
      </div>
    </footer>
  )
}

export default Footer
