import Link from 'next/link';
import {useRouter} from 'next/navigation';
import {IoMdTrash} from 'react-icons/io';
import {FiEdit} from 'react-icons/fi';
import type {Item} from '~lib/types';
import {useMutation, useQueryClient} from 'react-query';
import {deleteItem} from '~api/item';
import {Box, Button, ButtonGroup, Card, CardBody, Tag} from "@chakra-ui/react";

export type ItemCardProps = {
    item: Item;
};

export function ItemCard({item}: ItemCardProps) {
    const {push} = useRouter();
    const queryClient = useQueryClient();
    const {mutate} = useMutation(deleteItem, {
        onSuccess: () => {
            push('/items');
            queryClient.invalidateQueries('items');
        },
    });

    const {id, title, buying_price, selling_price, quantity} = item;

    return (
        <Card>
            <CardBody display='flex' flexDir='row' alignItems='center' justifyContent='space-between' gap={1} py={5}>
                <Box
                    onClick={() => push(`/sales?itemID=${id}`)}
                    flex={1}
                    display='flex'
                    flexDir='column'
                    cursor='pointer'
                >
                    <div className='indicator'>
                        <span
                            className={`indicator-item indicator-start badge badge-sm ${
                                quantity <= 5 ? 'badge-error' : 'badge-ghost'
                            }`}
                        >
                            {quantity}
                        </span>
                        <h2 className='card-title text-lg'>{title}</h2>
                    </div>
                    <Tag width='min-content'>
                        ₦{selling_price}
                    </Tag>
                </Box>

                <ButtonGroup>
                    <Link href={`/items/${id}/edit`}>
                        <Button variant='ghost'>
                            <FiEdit/>
                        </Button>
                    </Link>
                    <Button variant='ghost' onClick={() => mutate(id)}>
                        <IoMdTrash className='text-error'/>
                    </Button>
                </ButtonGroup>
            </CardBody>
        </Card>
    );
}

export default ItemCard;
