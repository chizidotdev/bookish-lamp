import { Alert, Button, ComboBox, Input } from '@copia/ui';
import { useState } from 'react';
import { useForm, SubmitHandler } from 'react-hook-form';
import { useMutation, useQuery } from 'react-query';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { addSale, getItems, type Item, SaleBase } from './api';

export const CreateSale = () => {
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();
    const { data: items } = useQuery({
        queryKey: ['get-items'],
        queryFn: getItems,
        refetchOnWindowFocus: false,
    });
    const [selected, setSelected] = useState<Item | undefined>(
        items?.find((item) => item.ID === searchParams.get('itemId'))
    );

    const { mutate, isLoading } = useMutation({
        mutationFn: addSale,
        onSuccess: () => {
            navigate('/');
        },
    });

    const { register, handleSubmit } = useForm<SaleBase>();
    const onSubmit: SubmitHandler<SaleBase> = (data) =>
        mutate({ ...data, id: selected?.ID ?? '' });

    if (!items) return <div>Loading...</div>;

    return (
        <>
            <h1 className="text-2xl font-bold mb-5">Add Sale</h1>

            <section className="mb-10 flex flex-col gap-2">
                <label>Select an item to add a sale.</label>
                <ComboBox
                    data={items}
                    selected={selected}
                    setSelected={setSelected}
                />

                {selected?.quantity === 0 && (
                    <Alert
                        message={`Warning: ${selected?.name} is out of stock. You need to update this item.`}
                        variant="warning"
                    >
                        <Button
                            variant="ghost"
                            size="small"
                            onClick={() =>
                                navigate(`/items/edit/${selected.ID}`)
                            }
                        >
                            Update Item
                        </Button>
                    </Alert>
                )}
            </section>

            <form
                onSubmit={handleSubmit(onSubmit)}
                className="form-control gap-3"
            >
                <Input
                    value={selected?.name}
                    label="Item Name"
                    type="text"
                    disabled
                />
                <div className="grid grid-cols-2 gap-3">
                    <Input
                        {...register('quantity_sold', {
                            required: true,
                            valueAsNumber: true,
                        })}
                        label="Quantity Sold"
                        placeholder="0"
                        disabled={!selected}
                    />
                    <Input
                        {...register('unit_price', {
                            required: true,
                            valueAsNumber: true,
                        })}
                        defaultValue={selected?.selling_price}
                        label="Price"
                        placeholder="N"
                        type="number"
                        disabled={!selected}
                    />
                </div>

                <Input
                    {...register('sale_date', {
                        required: true,
                        valueAsDate: true,
                    })}
                    label="Date"
                    type="date"
                    disabled={!selected}
                />

                <div className="mt-5">
                    <Button variant="primary" loading={isLoading} type="submit">
                        Submit
                    </Button>
                </div>
            </form>
        </>
    );
};
