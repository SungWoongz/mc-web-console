import { cn } from '@/lib/utils'

/**
 * 로딩 스피너 컴포넌트
 * @param {Object} props - 컴포넌트 props
 * @param {string} props.className - 추가 CSS 클래스
 * @param {string} props.size - 스피너 크기 (sm, md, lg)
 * @param {string} props.text - 로딩 텍스트
 * @returns {JSX.Element} 로딩 스피너 컴포넌트
 */
function LoadingSpinner({ className, size = 'md', text = '로딩 중...' }) {
  const sizeClasses = {
    sm: 'w-4 h-4',
    md: 'w-6 h-6',
    lg: 'w-8 h-8',
  }

  return (
    <div className={cn('flex flex-col items-center justify-center space-y-2', className)}>
      <div
        className={cn(
          'animate-spin rounded-full border-2 border-gray-300 border-t-primary',
          sizeClasses[size],
        )}
      />
      {text && <p className="text-sm text-muted-foreground">{text}</p>}
    </div>
  )
}

export default LoadingSpinner
