import Link from 'next/link';
import React from 'react';
import { getItems } from '~api/item';
import { NextPage } from '~lib/types';
import ItemCard from './item-card';

export default async function Items({ searchParams }: NextPage) {
    const pageID = searchParams['page_id'] ?? '1';
    const items = await getItems(pageID, '5');

    const incrementPage = () => {
        const newPageID = Number(pageID) + 1;
        return `/items?page_id=${newPageID}`;
    };

    const decrementPage = () => {
        if (Number(pageID) === 1) return `/items?page_id=${pageID}`;
        const newPageID = Number(pageID) - 1;
        return `/items?page_id=${newPageID}`;
    };

    return (
        <>
            <h1 className="text-2xl font-bold mb-5">Items</h1>
            <div className="flex flex-col gap-3">
                {items && items.map((item) => <ItemCard key={item.ID} item={item} />)}
            </div>

            <div className="mt-5 flex justify-end btn-group">
                <Link href={decrementPage()} className="btn">
                    «
                </Link>
                <button className="btn btn-disabled">Page {pageID}</button>
                <Link href={incrementPage()} className="btn">
                    »
                </Link>
            </div>
        </>
    );
}
