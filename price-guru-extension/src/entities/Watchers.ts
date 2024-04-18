export enum WATECHER_STATUS {
    PENDING = 'pending',
    ACTIVE = 'active',
    INACTIVE = 'inactive',
}
export type Watcher = {
    id: string
    name: string
    status: WATECHER_STATUS
}
