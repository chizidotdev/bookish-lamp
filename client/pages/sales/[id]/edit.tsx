import { Button, ButtonGroup, Flex } from '@chakra-ui/react';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { useEffect } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Input, Loading, ProtectedLayout, Text } from '~components';
import { useGetSaleById, useUpdateSale } from '~hooks/sales';
import { SaleBase } from '~lib/types';

export default function EditSale() {
    const { back, query } = useRouter();
    const saleID = query.id as string;
    const sale = useGetSaleById(saleID);
    const updateSale = useUpdateSale();

    const { register, handleSubmit, setValue } = useForm<SaleBase>({
        defaultValues: sale.data,
    });
    useEffect(() => {
        if (sale.isSuccess && sale.data) {
            setValue('quantity_sold', sale.data?.quantity_sold);
            setValue('sale_price', sale.data?.sale_price);
            setValue('customer_name', sale.data?.customer_name);
        }
    }, [sale.status, sale.data, setValue, sale.isSuccess]);

    const onSubmit: SubmitHandler<SaleBase> = (data) => {
        updateSale.mutate({ ...data, saleID });
    };

    if (sale.isLoading) {
        return <Loading />;
    }

    return (
        <ProtectedLayout>
            <Flex justify="space-between" mb={5}>
                <Text variant="h2">Edit Sale</Text>
                <Link href={`/items/${sale.data?.item_id}`}>
                    <Button>View Item</Button>
                </Link>
            </Flex>

            <form onSubmit={handleSubmit(onSubmit)} className="form-control gap-2 max-w-2xl">
                {/* <Input value={.data?.title} readOnly label="Item" /> */}
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
                    <Button isLoading={updateSale.isLoading} type="submit">
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
