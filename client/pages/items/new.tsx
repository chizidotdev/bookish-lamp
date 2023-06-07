import Link from 'next/link';
import React from 'react';
import { useMutation } from 'react-query';
import { newItem } from '~api/item';
import { Button, Input, ItemsLayout, Text } from '~components';
import { useRouter } from 'next/router';
import { SubmitHandler, useForm } from 'react-hook-form';
import { ItemBase } from '~lib/types';

export default function NewItem() {
    const { push, back } = useRouter();
    const { mutate, isLoading } = useMutation(newItem, {
        onSuccess: () => {
            push('/items');
        },
    });
    const { register, handleSubmit } = useForm<ItemBase>();

    const onSubmit: SubmitHandler<ItemBase> = (data) => {
        const { title, quantity, buying_price, selling_price } = data;
        mutate({ title, quantity, buying_price, selling_price });
    };

    return (
        <ItemsLayout>
            <div className='mb-5'>
                <Text variant='h2'>New Item</Text>
            </div>

            <form onSubmit={handleSubmit(onSubmit)} className='form-control gap-2 max-w-2xl'>
                <Input
                    {...register('title', { required: true })}
                    label='Title'
                    placeholder='Title'
                />
                <div className='flex gap-4'>
                    <Input
                        {...register('selling_price', { required: true, valueAsNumber: true })}
                        label='Selling Price'
                        placeholder='Selling Price'
                    />
                    <Input
                        {...register('buying_price', { required: true, valueAsNumber: true })}
                        label='Cost Price'
                        placeholder='Cost Price'
                    />
                </div>
                <Input
                    {...register('quantity', { required: true, valueAsNumber: true })}
                    label='Quantity'
                    placeholder='Quantity'
                    type='number'
                />

                <div className='flex mt-5 gap-3'>
                    <Button loading={isLoading} variant='primary' type='submit'>
                        Save
                    </Button>
                    <Button onClick={back}>Cancel</Button>
                </div>
            </form>
        </ItemsLayout>
    );
}
