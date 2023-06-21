import { Button, ButtonGroup, Flex } from '@chakra-ui/react';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Input, Loading, ProtectedLayout, Text } from '~components';
import { useGetItemById } from '~hooks/items';
import { useCreateSale } from '~hooks/sales';
import { SaleBase } from '~lib/types';

export default function NewSale() {
    const { back, query } = useRouter();
    const itemID = query.itemID as string;
    const createSale = useCreateSale();
    const item = useGetItemById(itemID);

    const { register, handleSubmit } = useForm<SaleBase>();
    const onSubmit: SubmitHandler<SaleBase> = (data) => {
        createSale.mutate({ ...data, itemID });
    };

    if (item.isLoading) {
        return <Loading />;
    }

    return (
        <ProtectedLayout>
            <Flex justify="space-between" mb={5}>
                <Text variant="h2">New Sale</Text>
                {item.data && (
                    <Link href={`/items/${itemID}`}>
                        <Button>View Item</Button>
                    </Link>
                )}
            </Flex>

            <form onSubmit={handleSubmit(onSubmit)} className="form-control gap-2 max-w-2xl">
                <Input value={item.data?.title} readOnly label="Item" />
                <div className="flex gap-4">
                    <Input
                        {...register('quantity_sold', {
                            required: true,
                            valueAsNumber: true,
                            min: 1,
                        })}
                        type="number"
                        label="Quantity Sold"
                        placeholder="Quantity Sold"
                    />
                    <Input
                        {...register('sale_price', { required: true, valueAsNumber: true })}
                        type="number"
                        label="Sale Price"
                        placeholder="Sale Price"
                    />
                </div>
                <Input
                    {...register('customer_name')}
                    label="Customer Name"
                    placeholder="Customer Name"
                />
                <Input
                    {...register('sale_date', { required: true, valueAsDate: true })}
                    label="Date Sold"
                    type="date"
                />

                <ButtonGroup mt={4}>
                    <Button isLoading={createSale.isLoading} type="submit">
                        Save
                    </Button>
                    <Button variant="outline" onClick={back}>
                        Cancel
                    </Button>
                </ButtonGroup>
            </form>
        </ProtectedLayout>
    );
}
