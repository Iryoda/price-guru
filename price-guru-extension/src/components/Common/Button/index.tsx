type Props = React.ButtonHTMLAttributes<HTMLButtonElement> & {
    children: React.ReactNode
}

const Button: React.FC<Props> = ({ children, className, ...rest }) => (
    <button
        className={`rounded-md bg-teal-600 px-8 py-2 text-gray-100 ${className}`}
        {...rest}
    >
        {children}
    </button>
)

export default Button
