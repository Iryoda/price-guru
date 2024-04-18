/* eslint-disable @typescript-eslint/ban-types */
/* eslint-disable @typescript-eslint/no-explicit-any */
declare module React {
    interface FCWC<P = {}> {
        (
            props: P & { children?: React.ReactNode },
            context?: any
        ): ReactElement<any, any> | null
    }
}
