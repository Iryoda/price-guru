import LoginView from '@pages/auth/login'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/login/')({
    component: () => <LoginView />,
})
