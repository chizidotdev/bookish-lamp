import { Button, useToast } from '@chakra-ui/react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useMutation } from 'react-query';
import { login } from '~api/user';
import { AuthLayout, Input, Link, Text } from '~components';

type FormValues = {
    email: string;
    password: string;
};

export default function Page() {
    const toast = useToast();
    const { mutate, isLoading } = useMutation(login, {
        onSuccess: (data) => {
            toast({
                description: 'Login successful!',
                status: 'success',
                duration: 3000,
            });
            window.location.href = '/dashboard';
        },
        onError: (error) => {
            console.log(error)
            toast({
                description: error.response.data.error,
                status: 'error',
                duration: 3000,
            });
        }
    });
    const { register, handleSubmit, formState } = useForm<FormValues>();

    const onSubmit: SubmitHandler<FormValues> = (data) => {
        mutate({ email: data.email, password: data.password });
    };

    return (
        <AuthLayout>
            <div className="text-center mb-5">
                <Text variant="h1">Welcome back</Text>
            </div>

            <form onSubmit={handleSubmit(onSubmit)} className="form-control gap-2">
                <Input
                    {...register('email', { required: true })}
                    type="email"
                    label="Email Address"
                    placeholder="Enter your Email"
                    isInvalid={Boolean(formState.errors.email)}
                />
                <Input
                    {...register('password', { required: true })}
                    type="password"
                    label="Password"
                    placeholder="Enter Password"
                    isInvalid={Boolean(formState.errors.password)}
                />

                <Button isLoading={isLoading} type="submit" mt="4">
                    Login
                </Button>
            </form>

            <div className="mt-5 flex justify-center gap-1">
                <Text>Don&apos;t have an account?</Text>
                <Link href="/auth/signup">Sign Up</Link>
            </div>
        </AuthLayout>
    );
}
