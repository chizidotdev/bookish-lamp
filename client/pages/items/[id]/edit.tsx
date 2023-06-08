import React from 'react';
import { useMutation, useQuery } from 'react-query';
import { getItemById, updateItem } from '~api/item';
import { Button, Input, ItemsLayout, Loading, Text } from '~components';
import { useRouter } from 'next/router';
import { SubmitHandler, useForm } from 'react-hook-form';
import { ItemBase } from '~lib/types';

export default function EditItem() {
    const { push, back } = useRouter();
    const { query } = useRouter();
    const { data, isLoading } = useQuery(['item', query.id], () => getItemById(query.id as string));
    const { mutate, isLoading: isSaving } = useMutation(updateItem, {
        onSuccess: () => {
            push('/items');
        },
    });
    const { register, handleSubmit } = useForm<ItemBase>({
        values: {
            title: data?.title || '',
            quantity: data?.quantity || 0,
            buying_price: data?.buying_price || 0,
            selling_price: data?.selling_price || 0,
        },
    });

    const onSubmit: SubmitHandler<ItemBase> = (data) => {
        const { title, quantity, buying_price, selling_price } = data;
        mutate({ title, quantity, buying_price, selling_price, id: query.id as string });
    };

    if (isLoading) {
        return <Loading />;
    }

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
                    <Button loading={isSaving} variant='primary' type='submit'>
                        Save
                    </Button>
                    <Button onClick={back}>Cancel</Button>
                </div>
            </form>
        </ItemsLayout>
    );
}
