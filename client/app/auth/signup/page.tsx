'use client';
import React from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useMutation } from 'react-query';
import { login } from '~api/user';
import { Button, Checkbox, Input, Text, Link } from '~components';

type FormValues = {
    email: string;
    password: string;
    confirmPassword: string;
    terms: boolean;
};

export default function Page() {
    const {mutate, isLoading, isSuccess} = useMutation(login);
    const { register, handleSubmit } = useForm<FormValues>();
    const onSubmit: SubmitHandler<FormValues> = (data) => {
        mutate({ email: data.email, password: data.password });
    };

    return (
        <>
            <div className='text-center mb-5'>
                <Text variant='h1'>Let&apos;s get started</Text>
            </div>

            <form onSubmit={handleSubmit(onSubmit)} className='form-control gap-2'>
                <Input
                    {...register('email')}
                    type='email'
                    label='Email Address'
                    placeholder='Enter your Email'
                />
                <Input
                    {...register('password')}
                    type='password'
                    label='Password'
                    placeholder='Enter Password'
                />
                <Input
                    {...register('confirmPassword')}
                    type='password'
                    label='Confirm Password'
                    placeholder='Confirm Password'
                />
                <Checkbox {...register('terms')}>
                    Agreed to terms of use and privacy statements.
                </Checkbox>

                <div className='mt-5'>
                    <Button type='submit'>Get Started</Button>
                </div>
            </form>

            <div className='mt-5 flex justify-center gap-1'>
                <Text>Already have an account?</Text>
                <Link href='/auth/login'>Log In</Link>
            </div>
        </>
    );
}
