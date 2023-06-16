import React from 'react';
import { useMutation } from 'react-query';
import { newItem } from '~api/item';
import { ItemsLayout, Text } from '~components';
import { useRouter } from 'next/router';
import { SubmitHandler, useForm } from 'react-hook-form';
import { ItemBase } from '~lib/types';
import { Button, Input, FormControl, FormLabel, Box, ButtonGroup, Flex } from '@chakra-ui/react';

export default function NewItem() {
    const { push, back } = useRouter();
    const { mutate, isLoading } = useMutation(newItem, {
        onSuccess: () => {
            push('/items');
        },
    });
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<ItemBase>();

    const onSubmit: SubmitHandler<ItemBase> = (data) => {
        const { title, quantity, buying_price, selling_price } = data;
        mutate({ title, quantity, buying_price, selling_price });
    };

    return (
        <ItemsLayout>
            <Box className='mb-5'>
                <Text variant='h2'>New Item</Text>
            </Box>

            <form onSubmit={handleSubmit(onSubmit)} className='form-control gap-2 max-w-2xl'>
                <FormControl isInvalid={Boolean(errors.title)}>
                    <FormLabel>Title</FormLabel>
                    <Input {...register('title', { required: true })} placeholder='Title' />
                </FormControl>

                <Flex gap='4'>
                    <FormControl isInvalid={Boolean(errors.selling_price)}>
                        <FormLabel>Selling Price</FormLabel>
                        <Input
                            {...register('selling_price', { required: true, valueAsNumber: true })}
                            placeholder='Selling Price'
                            type='number'
                        />
                    </FormControl>
                    <FormControl isInvalid={Boolean(errors.buying_price)}>
                        <FormLabel>Cost Price</FormLabel>
                        <Input
                            {...register('buying_price', { required: true, valueAsNumber: true })}
                            placeholder='Cost Price'
                            type='number'
                        />
                    </FormControl>
                </Flex>

                <FormControl isInvalid={Boolean(errors.quantity)}>
                    <FormLabel>Quantity</FormLabel>
                    <Input
                        {...register('quantity', { required: true, valueAsNumber: true })}
                        placeholder='Quantity'
                        type='number'
                    />
                </FormControl>

                <ButtonGroup mt='5'>
                    <Button isLoading={isLoading} type='submit'>
                        Save
                    </Button>
                    <Button onClick={back} variant='outline'>
                        Cancel
                    </Button>
                </ButtonGroup>
            </form>
        </ItemsLayout>
    );
}
