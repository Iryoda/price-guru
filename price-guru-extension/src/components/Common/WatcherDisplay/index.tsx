import { Watcher } from '@entities/Watchers'

import Typograph from '../Typograph'

type Props = Watcher & {
    onDelete?: (id: string) => void
    onRename?: () => void
}

const WatcherDisplay: React.FC<Props> = ({ onDelete, onRename, ...props }) => {
    return (
        <div className="grid w-full grid-cols-2 items-center justify-between p-2">
            <Typograph.Span>{props.name}</Typograph.Span>

            <div className="flex w-full items-center justify-between">
                <Typograph.Span className="lowercase text-green-700">
                    {props.status}
                </Typograph.Span>

                <div>
                    <button
                        className="font-inter ml-2 rounded-md px-2 text-yellow-600 hover:bg-gray-900"
                        onClick={() => onRename && onRename()}
                    >
                        edit
                    </button>

                    <button
                        className="font-inter ml-2 rounded-md px-2 text-red-600 hover:bg-gray-900"
                        onClick={() => onDelete && onDelete(props.id)}
                    >
                        x
                    </button>
                </div>
            </div>
        </div>
    )
}

export default WatcherDisplay
