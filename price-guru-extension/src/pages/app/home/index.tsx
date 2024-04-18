import { useMutation, useQuery } from '@tanstack/react-query'
import { useEffect } from 'react'

import { useAuth } from '@hooks/auth'

import { keys } from '@constants/query-keys'

import WatchersService from '@services/watcher-service'

import { handleSelectPrice } from '@utils/price-script'

import Container from '@components/Common/Container'
import Typograph from '@components/Common/Typograph'
import WatcherDisplay from '@components/Common/WatcherDisplay'

const Home = () => {
    const { user } = useAuth()

    const { data, refetch, isLoading } = useQuery({
        queryKey: [keys.watchers, user?.id],
        queryFn: WatchersService.getAll,
    })

    const { mutate } = useMutation({
        mutationFn: WatchersService.create,
        onSuccess: () => {
            console.log('Success')
            refetch()
        },
    })

    const { mutate: deleteMut } = useMutation({
        mutationFn: WatchersService.delete,
        onSuccess: () => {
            refetch()
        },
    })

    useEffect(() => {
        if (chrome.runtime === undefined) return

        chrome.runtime.onMessage.addListener((request) => {
            if (request.message === 'price-clicked') {
                console.log(request)
                mutate({
                    name: 'Test',
                    url: request.location,
                    node: request.target,
                })
            }
        })
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])

    const handleRegister = async () => {
        const [tab] = await chrome.tabs.query({
            active: true,
            currentWindow: true,
        })

        if (tab) {
            await chrome.scripting.executeScript({
                target: { tabId: tab.id! },
                func: handleSelectPrice,
            })
        }
    }

    return (
        <Container>
            <div className="flex w-full justify-end">
                <Typograph.Span className="text-sm lowercase">
                    {user?.name}
                </Typograph.Span>
            </div>

            <div className="w-full pt-6">
                <div className="flex w-full items-baseline pb-4">
                    <Typograph.H3 className="text-lg uppercase text-gray-200">
                        Watchers
                    </Typograph.H3>

                    <button
                        className="pl-4 text-teal-600"
                        onClick={handleRegister}
                    >
                        + Add
                    </button>
                </div>

                <div className="rounded-lg bg-gray-800 p-2">
                    {!isLoading &&
                        data?.map((watcher) => (
                            <WatcherDisplay
                                key={watcher.id}
                                {...watcher}
                                onDelete={deleteMut}
                            />
                        ))}
                </div>
            </div>
        </Container>
    )
}

export default Home
