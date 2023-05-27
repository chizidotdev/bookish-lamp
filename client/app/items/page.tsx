import React from 'react';
import { getItems } from '~api/item';
import ItemCard from './item-card';

export default async function Items() {
    const items = await getItems();

    return (
        <>
            <h1 className='text-2xl font-bold mb-5'>Items</h1>
            <div className='flex flex-col gap-3'>
                {items && items.map((item) => <ItemCard key={item.ID} item={item} />)}
            </div>

            <div className='mt-5 flex justify-end btn-group'>
                <button className='btn'>«</button>
                <button className='btn'>Page 1</button>
                <button className='btn'>»</button>
            </div>
        </>
    );
}
