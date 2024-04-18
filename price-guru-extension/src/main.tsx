import {
    RouterProvider,
    createMemoryHistory,
    createRouter,
} from '@tanstack/react-router'
import React from 'react'
import ReactDOM from 'react-dom/client'

import AppProvider from './hooks'
import './index.css'
import { routeTree } from './routeTree.gen'

const memoryHistory = createMemoryHistory({
    initialEntries: ['/'],
})

const router = createRouter({
    routeTree,
    defaultPreload: 'intent',
    history: chrome.runtime !== undefined ? memoryHistory : undefined,
})

declare module '@tanstack/react-router' {
    interface Register {
        router: typeof router
    }
}

// Render the app
const rootElement = document.getElementById('app')!
if (!rootElement.innerHTML) {
    const root = ReactDOM.createRoot(rootElement)
    root.render(
        <React.StrictMode>
            <AppProvider>
                <RouterProvider router={router} />
            </AppProvider>
        </React.StrictMode>
    )
}
