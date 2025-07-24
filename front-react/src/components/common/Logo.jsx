import { cn } from '@/lib/utils'

/**
 * M-CMP 로고 컴포넌트
 * @param {Object} props - 컴포넌트 props
 * @param {string} props.className - 추가 CSS 클래스
 * @param {number} props.width - 로고 너비
 * @param {number} props.height - 로고 높이
 * @returns {JSX.Element} 로고 컴포넌트
 */
export function Logo({ className, width = 110, height = 32, ...props }) {
  return (
    <div className={cn('flex items-center', className)} {...props}>
      <img
        src="/images/mcmp.svg"
        alt="M-CMP Logo"
        width={width}
        height={height}
        className="h-auto"
      />
    </div>
  )
}
