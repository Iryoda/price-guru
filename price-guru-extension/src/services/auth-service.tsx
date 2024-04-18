import { api } from '@config/api'

import { User } from '@entities/User'

export type LoginRequest = {
    email: string
    password: string
}

type LoginResponse = {
    user: User
    token: string
}

export class AuthService {
    static async login(request: LoginRequest) {
        const { data } = await api.post<LoginResponse>('/auth/login', request)

        return data
    }
}
