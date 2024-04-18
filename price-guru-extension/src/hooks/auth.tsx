import { useMutation } from '@tanstack/react-query'
import { createContext, useContext, useEffect, useState } from 'react'

import { api } from '@config/api'

import { STORAGE } from '@constants/storage'

import { AuthService, LoginRequest } from '@services/auth-service'

import { User } from '@entities/User'

type AuthContextType = {
    user: User | null
    login: (
        data: LoginRequest,
        callback: {
            onSuccess?: () => void
            onError?: () => void
        }
    ) => void
    logout: () => void
}

const authContext = createContext<AuthContextType>({} as AuthContextType)

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
    children,
}) => {
    const [user, setUser] = useState({} as User)
    const [, setToken] = useState('')

    useEffect(() => {
        try {
            const tk = localStorage.getItem(STORAGE.TOKEN)
            const user = localStorage.getItem(STORAGE.USER)

            setUser(user ? JSON.parse(user) : null)
            setToken(tk ?? '')

            api.defaults.headers.common['Authorization'] = `Bearer ${tk}`
        } catch (error) {
            console.error(error)
            return
        }
    }, [])

    const { mutate: login } = useMutation({
        mutationFn: (data: LoginRequest) => AuthService.login(data),
        onSuccess: (data) => {
            setUser(data.user)
            setToken(data.token)

            localStorage.setItem('@PriceGuru/token', data.token)
            localStorage.setItem('@PriceGuru/user', JSON.stringify(data.user))

            api.defaults.headers.common['Authorization'] =
                `Bearer ${data.token}`
        },
    })

    const logout = () => {
        setToken('')
        setUser({} as User)
    }

    return (
        <authContext.Provider value={{ user, login, logout }}>
            {children}
        </authContext.Provider>
    )
}

// eslint-disable-next-line react-refresh/only-export-components
export const useAuth = () => useContext(authContext)
