import { forwardRef } from 'react'

type Props = React.InputHTMLAttributes<HTMLInputElement> & {
    error?: string
}

const Input = forwardRef<HTMLInputElement, Props>(({ error, ...rest }, ref) => (
    <div className="flex w-full flex-col">
        <input
            ref={ref}
            {...rest}
            className="rounded-md border border-black px-3 py-2"
        />
        {error && <span className="text-sm text-red-500">{error}</span>}
    </div>
))

export default Input
