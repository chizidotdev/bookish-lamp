'use client';
import React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useMutation } from 'react-query';
import { login } from '~api/user';
import { Button, Input, Text, Link } from '~components';

type FormValues = {
    email: string;
    password: string;
};

export default function Page() {
    const { mutate, isLoading, isError, isSuccess, error, data } = useMutation(login);
    const { register, handleSubmit } = useForm<FormValues>();

    const onSubmit: SubmitHandler<FormValues> = (data) => {
        mutate({ email: data.email, password: data.password });
    };

    isSuccess && console.log('success >>> ', { isSuccess, data });
    isError && console.log('error ??? ', { error });

    return (
        <>
            <div className='text-center mb-5'>
                <Text variant='h1'>Login</Text>
            </div>

            <form onSubmit={handleSubmit(onSubmit)} className='form-control gap-2'>
                <Input
                    {...register('email', { required: true })}
                    type='email'
                    label='Email Address'
                    placeholder='Enter your Email'
                />
                <Input
                    {...register('password', { required: true })}
                    type='password'
                    label='Password'
                    placeholder='Enter Password'
                />

                <div className='mt-5'>
                    <Button type='submit'>Login</Button>
                </div>
            </form>

            <div className='mt-5 flex justify-center gap-1'>
                <Text>Don&apos;t have an account?</Text>
                <Link href='/auth/signup'>Sign Up</Link>
            </div>
        </>
    );
}
