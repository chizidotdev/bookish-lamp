import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { IoMdTrash } from 'react-icons/io';
import { FiEdit } from 'react-icons/fi';
import type { Item } from '~lib/types';
import { useMutation, useQueryClient } from 'react-query';
import { deleteItem } from '~api/item';
import { Button } from '~components/ui';

export type ItemCardProps = {
    item: Item;
};

export function ItemCard({ item }: ItemCardProps) {
    const { push } = useRouter();
    const queryClient = useQueryClient();
    const { mutate } = useMutation(deleteItem, {
        onSuccess: () => {
            push('/items');
            queryClient.invalidateQueries('items');
        },
    });

    const { id, title, buying_price, selling_price, quantity } = item;

    return (
        <div className='card bg-base-200 shadow-md'>
            <div className='card-body flex-row items-center justify-between gap-1 mx-0 py-5'>
                <div
                    onClick={() => push(`/sales?itemID=${id}`)}
                    className='flex-1 flex flex-col cursor-pointer'
                >
                    <div className='indicator'>
                        <span
                            className={`indicator-item indicator-start badge badge-sm ${
                                !quantity ? 'badge-error' : 'badge-secondary'
                            }`}
                        >
                            {quantity}
                        </span>
                        <h2 className='card-title text-lg'>{title}</h2>
                    </div>
                    <div className='flex gap-2 mt-1'>
                        <div className='badge'>₦{buying_price}</div>
                        <div className='badge badge-primary'>₦{selling_price}</div>
                    </div>
                </div>
                <div className='card-actions flex-row gap-4'>
                    <Link href={`/items/${id}/edit`}>
                        <Button>
                            <FiEdit />
                        </Button>
                    </Link>
                    <Button onClick={() => mutate(id)}>
                        <IoMdTrash className='text-error' />
                    </Button>
                </div>
            </div>
        </div>
    );
}

export default ItemCard;
