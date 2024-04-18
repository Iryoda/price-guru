import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

import { AuthProvider } from './auth'

type Props = {
    children: React.ReactNode
}

const queryClient = new QueryClient()

const AppProvider = ({ children }: Props) => (
    <QueryClientProvider client={queryClient}>
        <AuthProvider>{children}</AuthProvider>
    </QueryClientProvider>
)

export default AppProvider
