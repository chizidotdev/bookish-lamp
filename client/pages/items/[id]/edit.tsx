import React, {useEffect} from 'react';
import {useMutation, useQuery} from 'react-query';
import {getItemById, updateItem} from '~api/item';
import {Input, Loading, ProtectedLayout, Text} from '~components';
import {useRouter} from 'next/router';
import {SubmitHandler, useForm} from 'react-hook-form';
import {ItemBase} from '~lib/types';
import {Button, ButtonGroup} from "@chakra-ui/react";

export default function EditItem() {
    const {push, back} = useRouter();
    const {query} = useRouter();
    const {data, isLoading, isSuccess, } = useQuery(['item', query.id], () => getItemById(query.id as string));
    const {mutate, isLoading: isSaving} = useMutation(updateItem, {
        onSuccess: () => {
            push('/items');
        },
    });
    const {register, handleSubmit, setValue} = useForm<ItemBase>();

    useEffect(() => {
        if (isSuccess && data) {
            setValue('title', data?.title);
            setValue('quantity', data?.quantity);
            setValue('selling_price', data?.selling_price);
            setValue('buying_price', data?.buying_price);
        }
    }, [data, setValue, isSuccess]);

    const onSubmit: SubmitHandler<ItemBase> = (data) => {
        const {title, quantity, buying_price, selling_price} = data;
        mutate({title, quantity, buying_price, selling_price, id: query.id as string});
    };

    if (isLoading) {
        return <Loading/>;
    }

    return (
        <ProtectedLayout>
            <div className='mb-5'>
                <Text variant='h2'>Edit Item</Text>
            </div>

            <form onSubmit={handleSubmit(onSubmit)} className='form-control gap-2 max-w-2xl'>
                <Input
                    {...register('title', {required: true})}
                    label='Title'
                    placeholder='Title'
                />
                <div className='flex gap-4'>
                    <Input
                        {...register('selling_price', {required: true, valueAsNumber: true})}
                        label='Selling Price'
                        placeholder='Selling Price'
                    />
                    <Input
                        {...register('buying_price', {required: true, valueAsNumber: true})}
                        label='Cost Price'
                        placeholder='Cost Price'
                    />
                </div>
                <Input
                    {...register('quantity', {required: true, valueAsNumber: true})}
                    label='Quantity'
                    placeholder='Quantity'
                    type='number'
                />

                <ButtonGroup mt={5}>
                    <Button isLoading={isSaving} type='submit'>
                        Save
                    </Button>
                    <Button variant='outline' onClick={back}>Cancel</Button>
                </ButtonGroup>
            </form>
        </ProtectedLayout>
    );
}
