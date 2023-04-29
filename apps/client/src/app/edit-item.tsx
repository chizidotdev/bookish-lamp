import { Button, Input } from '@copia/ui';
import { useForm, SubmitHandler } from 'react-hook-form';
import { useMutation, useQuery, useQueryClient } from 'react-query';
import { useNavigate, useParams } from 'react-router-dom';
import { getItemById, ItemBase, updateItem } from './api';

export const EditItem = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const queryClient = useQueryClient();
    const { isLoading: isFetching, data } = useQuery({
        queryKey: ['get-item-by-id'],
        queryFn: () => getItemById(id ?? ''),
        refetchOnWindowFocus: false,
    });

    const { mutate, isLoading } = useMutation({
        mutationFn: updateItem,
        onSuccess: () => {
            navigate('/');
            queryClient.invalidateQueries(['get-items']);
        },
    });
    const { register, handleSubmit } = useForm<ItemBase>();
    const onSubmit: SubmitHandler<ItemBase> = (data) =>
        mutate({ ...data, id: id ?? '' });

    if (isFetching) return <div>Loading...</div>;

    const { name, quantity, buying_price, selling_price } = data ?? {};

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="form-control gap-3">
            <Input
                {...register('name', { required: true })}
                defaultValue={name}
                label="Enter Name"
                placeholder="Enter Name"
            />
            <Input
                {...register('quantity', {
                    required: true,
                    valueAsNumber: true,
                })}
                defaultValue={quantity}
                label="Quantity"
                placeholder="0"
                type="number"
            />

            <div className="grid grid-cols-2 gap-3">
                <Input
                    {...register('buying_price', {
                        required: true,
                        valueAsNumber: true,
                    })}
                    defaultValue={buying_price}
                    label="Buying Price"
                    placeholder="N"
                    type="number"
                />
                <Input
                    {...register('selling_price', {
                        required: true,
                        valueAsNumber: true,
                    })}
                    defaultValue={selling_price}
                    label="Selling Price"
                    placeholder="N"
                    type="number"
                />
            </div>

            <div className="mt-5">
                <Button loading={isLoading} type="submit">
                    Submit
                </Button>
            </div>
        </form>
    );
};
