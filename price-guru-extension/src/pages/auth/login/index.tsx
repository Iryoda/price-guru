import { zodResolver } from '@hookform/resolvers/zod'
import { useNavigate } from '@tanstack/react-router'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

import { useAuth } from '@hooks/auth'

import { ROUTES } from '@constants/routes'

import Button from '@components/Common/Button'
import Container from '@components/Common/Container'
import Input from '@components/Common/Input'
import Typograph from '@components/Common/Typograph'

const schema = z.object({
    email: z.string().email(),
    password: z.string().min(8),
})

type FormValues = z.infer<typeof schema>

const LoginView = () => {
    const navigate = useNavigate()
    const { login } = useAuth()

    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<FormValues>({
        resolver: zodResolver(schema),
    })

    const handleLogin = (data: FormValues) => {
        login(data, {
            onSuccess: () => {
                navigate({ to: ROUTES.HOME })
            },
            onError: () => {
                console.log('Failed to login')
            },
        })
    }

    return (
        <Container className="flex items-center justify-center">
            <div className="w-full max-w-96 rounded-md bg-slate-900 p-8">
                <Typograph.H2 className="text-left">Login</Typograph.H2>

                <form
                    onSubmit={handleSubmit(handleLogin)}
                    className="flex w-full flex-col items-center pt-8"
                >
                    <Input
                        error={errors.email?.message}
                        placeholder="Email"
                        {...register('email')}
                    />

                    <div className="w-full pt-3">
                        <Input
                            error={errors.password?.message}
                            type="password"
                            placeholder="Password"
                            {...register('password')}
                        />
                    </div>

                    <div className="pt-4">
                        <Button>Login</Button>
                    </div>
                </form>
            </div>
        </Container>
    )
}

export default LoginView
