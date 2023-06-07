import Link from 'next/link';
import React from 'react';
import ItemCard from '~components/item-card';
import { useQuery } from 'react-query';
import { getItems } from '~api/item';
import { ItemsLayout } from '~components';

export default function Items() {
    const { data: items } = useQuery('items', {
        queryFn: getItems,
    });

    return (
        <ItemsLayout>
            <h1 className='text-2xl font-bold mb-5'>Items</h1>
            <div className='flex flex-col gap-3'>
                {items &&
                    Boolean(items.length) &&
                    items.map((item) => <ItemCard key={item.ID} item={item} />)}
            </div>

            {/*<div className='mt-5 flex justify-end btn-group'>
                <Link href={decrementPage()} className='btn'>
                    «
                </Link>
                <button className='btn btn-disabled'>Page {pageID}</button>
                <Link href={incrementPage()} className='btn'>
                    »
                </Link>
            </div>*/}
        </ItemsLayout>
    );
}
