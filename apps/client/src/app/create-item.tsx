import { Button, Input } from '@copia/ui';
import { useForm, SubmitHandler } from 'react-hook-form';
import { useMutation } from 'react-query';
import { useNavigate } from 'react-router-dom';
import { addItem, ItemBase } from './api';

export const CreateItem = () => {
    const navigate = useNavigate();
    const { mutate, isLoading } = useMutation({
        mutationFn: addItem,
        onSuccess: () => {
            navigate('/');
        },
    });
    const {
        register,
        handleSubmit,
    } = useForm<ItemBase>();
    const onSubmit: SubmitHandler<ItemBase> = (data) => mutate(data);

    return (
        <form onSubmit={handleSubmit(onSubmit)} className="form-control gap-3">
            <Input
                {...register('name', { required: true })}
                label="Enter Name"
                placeholder="Enter Name"
            />
            <Input
                {...register('quantity', {
                    required: true,
                    valueAsNumber: true,
                })}
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
                    label="Buying Price"
                    placeholder="N"
                    type="number"
                />
                <Input
                    {...register('selling_price', {
                        required: true,
                        valueAsNumber: true,
                    })}
                    label="Selling Price"
                    placeholder="N"
                    type="number"
                />
            </div>

            <div className="mt-5">
                <Button variant='primary' loading={isLoading} type="submit">
                    Submit
                </Button>
            </div>
        </form>
    );
};
