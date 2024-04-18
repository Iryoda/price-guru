import { api } from '@config/api'

import { Watcher } from '@entities/Watchers'

type CreateWatcherRequest = {
    url: string
    name: string
    node: string
}

export default class WatchersService {
    static async getAll() {
        const { data } = await api.get<Watcher[]>('/secure/watchers/all')
        return data
    }

    static async create(request: CreateWatcherRequest) {
        const { data } = await api.post<Watcher[]>(
            '/secure/watchers/create',
            request
        )
        return data
    }

    static async delete(id: string) {
        const { data } = await api.delete<Watcher[]>(
            '/secure/watchers/delete/' + id
        )
        return data
    }
}
