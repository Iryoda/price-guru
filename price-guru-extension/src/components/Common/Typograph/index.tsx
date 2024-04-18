/* eslint-disable react-refresh/only-export-components */
type Props = React.FCWC<{ className?: string }>

const H1: Props = ({ children, className }) => (
    <h1 className={`font-inter text-3xl font-bold text-white ${className}`}>
        {children}
    </h1>
)

const H2: Props = ({ children, className }) => (
    <h2 className={`font-inter text-2xl font-bold text-white ${className}`}>
        {children}
    </h2>
)

const H3: Props = ({ children, className }) => (
    <h3 className={`font-inter text-xl font-semibold text-white ${className}`}>
        {children}
    </h3>
)

const H4: Props = ({ children, className }) => (
    <h4 className={`font-inter text-lg font-semibold text-white ${className}`}>
        {children}
    </h4>
)

const Span: Props = ({ children, className }) => (
    <span className={`font-inter text-white ${className}`}>{children}</span>
)

export default { H1, H2, H3, H4, Span }
