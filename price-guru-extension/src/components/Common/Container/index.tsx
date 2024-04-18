import { useNavigate } from '@tanstack/react-router'
import { useEffect } from 'react'

import { useAuth } from '@hooks/auth'

import { ROUTES } from '@constants/routes'

type Props = {
    children: React.ReactNode
    className?: string
}

const Container: React.FC<Props> = ({ children, className }) => {
    const navigate = useNavigate()
    const { user } = useAuth()
    useEffect(() => {
        if (!user) {
            navigate({ to: ROUTES.LOGIN })
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])

    return (
        <div
            className={`min-w-screen bg-gradient max-h-scren relative z-0 h-screen ${className}`}
        >
            <div className={`h-full w-full p-4 ${className}`}>{children}</div>
        </div>
    )
}

export default Container
