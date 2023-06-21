import { Button } from '@chakra-ui/react';
import Link from 'next/link';
import { IoIosAdd } from 'react-icons/io';
import { ProtectedLayout, Text } from '~components';
import ItemCard from '~components/item-card';
import { useGetItems } from '~hooks/items';

export default function Items() {
    const { data: items } = useGetItems();

    return (
        <ProtectedLayout>
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
        </ProtectedLayout>
    );
}
