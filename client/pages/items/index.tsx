import Link from 'next/link';
import React from 'react';
import ItemCard from '~components/item-card';
import { useQuery } from 'react-query';
import { getItems } from '~api/item';
import { ItemsLayout, Text } from '~components';
import { IoIosAdd } from 'react-icons/io';
import { Button } from '@chakra-ui/react';

export default function Items() {
    const { data: items } = useQuery('items', {
        queryFn: getItems,
    });

    return (
        <ItemsLayout>
            <div className="flex mb-5 justify-between items-center">
                <Text variant="h2">Items</Text>
                <Link href="/items/new">
                    <Button rightIcon={<IoIosAdd size={20} />}>Add</Button>
                </Link>
            </div>
            <div className="flex flex-col gap-3">
                {items &&
                    Boolean(items.length) &&
                    items.map((item) => <ItemCard key={item.id} item={item} />)}
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
