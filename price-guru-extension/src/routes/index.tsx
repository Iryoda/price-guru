import Home from '@pages/app/home'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
    component: () => <Home />,
})
