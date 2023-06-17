import { Button } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useMutation } from 'react-query';
import { signup } from '~api/user';
import { Input, Text, Link, AuthLayout } from '~components';

type FormValues = {
    email: string;
    password: string;
    confirmPassword: string;
};

export default function Page() {
    const { push } = useRouter();
    const { mutate, isLoading } = useMutation(signup, {
        onSuccess: () => {
            push('/auth/login');
        },
    });
    const { register, handleSubmit, watch } = useForm<FormValues>();

    const onSubmit: SubmitHandler<FormValues> = (data) => {
        mutate({ email: data.email, password: data.password });
    };

    return (
        <AuthLayout>
            <div className="text-center mb-5">
                <Text variant="h1">Let&apos;s get started</Text>
            </div>

            <form onSubmit={handleSubmit(onSubmit)} className="form-control gap-2">
                <Input
                    {...register('email', { required: true })}
                    type="email"
                    label="Email Address"
                    placeholder="Enter your Email"
                    autoComplete="off"
                />
                <Input
                    {...register('password', { required: true })}
                    type="password"
                    label="Password"
                    placeholder="Enter Password"
                    autoComplete="off"
                />
                <Input
                    {...register('confirmPassword', { required: true })}
                    pattern={watch('password')}
                    type="password"
                    label="Confirm Password"
                    placeholder="Confirm Password"
                    autoComplete="off"
                />

                <div className="mt-5">
                    <Button w="full" isLoading={isLoading} type="submit">
                        Get Started
                    </Button>
                </div>
            </form>

            <div className="mt-5 flex justify-center gap-1">
                <Text>Already have an account?</Text>
                <Link href="/auth/login">Log In</Link>
            </div>
        </AuthLayout>
    );
}
