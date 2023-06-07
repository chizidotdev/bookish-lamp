import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { FaTrash, FaEdit } from 'react-icons/fa';
import type { Item } from '~lib/types';

export type ItemCardProps = {
    item: Item;
};

export function ItemCard({ item }: ItemCardProps) {
    const { push } = useRouter();
    const { ID, title, buying_price, selling_price, quantity } = item;

    return (
        <div className='card bg-base-200 shadow-md'>
            <div className='card-body flex-row items-center justify-between gap-1 mx-0 py-5'>
                <div
                    onClick={() => push(`/items/${ID}/sales`)}
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
                        <div className='badge badge-accent'>₦{selling_price}</div>
                    </div>
                </div>
                <div className='card-actions flex-row gap-5'>
                    <Link href={`/items/edit/${ID}`}>
                        <FaEdit className='text-primary' />
                    </Link>
                    <Link href={`/items/delete/${ID}`}>
                        <FaTrash className='text-error' />
                    </Link>
                </div>
            </div>
        </div>
    );
}

export default ItemCard;
