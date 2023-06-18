import React from 'react';
import { useMutation } from 'react-query';
import { newItem } from '~api/item';
import { Input, ProtectedLayout, Text } from '~components';
import { useRouter } from 'next/router';
import { SubmitHandler, useForm } from 'react-hook-form';
import { ItemBase } from '~lib/types';
import { Button, FormControl, FormLabel, Box, ButtonGroup, Flex } from '@chakra-ui/react';

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
        <ProtectedLayout>
            <Box className="mb-5">
                <Text variant="h2">New Item</Text>
            </Box>

            <form onSubmit={handleSubmit(onSubmit)} className="form-control gap-2 max-w-2xl">
                <FormControl isInvalid={Boolean(errors.title)}>
                    <FormLabel>Title</FormLabel>
                    <Input {...register('title', { required: true })} placeholder="Title" />
                </FormControl>

                <Flex gap="4">
                    <Input
                        {...register('selling_price', { required: true, valueAsNumber: true })}
                        label="Selling Price"
                        placeholder="Selling Price"
                        type="number"
                        isInvalid={Boolean(errors.selling_price)}
                    />
                    <Input
                        {...register('buying_price', { required: true, valueAsNumber: true })}
                        label="Cost Price"
                        placeholder="Cost Price"
                        type="number"
                        isInvalid={Boolean(errors.buying_price)}
                    />
                </Flex>

                <Input
                    {...register('quantity', { required: true, valueAsNumber: true })}
                    label="Quantity"
                    placeholder="Quantity"
                    type="number"
                    isInvalid={Boolean(errors.quantity)}
                />

                <ButtonGroup mt="5">
                    <Button isLoading={isLoading} type="submit">
                        Save
                    </Button>
                    <Button onClick={back} variant="outline">
                        Cancel
                    </Button>
                </ButtonGroup>
            </form>
        </ProtectedLayout>
    );
}
