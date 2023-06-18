import {
    Button,
    Table,
    TableCaption,
    TableContainer,
    Tbody,
    Td,
    Th,
    Thead,
    Tr,
} from '@chakra-ui/react';
import Link from 'next/link';
import { useRouter } from 'next/router';
import { IoIosAdd } from 'react-icons/io';
import { Loading } from '~components';
import { useGetSales } from '~hooks/sales';
import { formatDate } from '~lib/utils';

export default function Sales() {
    const { query, push } = useRouter();
    const itemID = query.itemID;
    const sales = useGetSales(itemID as string);

    if (sales.isLoading) {
        return <Loading />;
    }

    return (
        <>
            <Link href={{ pathname: '/sales/new', query: { itemID } }}>
                <Button mb={5} rightIcon={<IoIosAdd size={20} />}>
                    Add
                </Button>
            </Link>

            <TableContainer>
                <Table variant="simple">
                    <TableCaption>Sales Records for Inventory</TableCaption>
                    <Thead>
                        <Tr>
                            <Th>Date</Th>
                            <Th isNumeric>Quantity</Th>
                            <Th isNumeric>Price(₦)</Th>
                        </Tr>
                    </Thead>
                    <Tbody>
                        {sales.data?.map((sale) => (
                            <Tr key={sale.id} onClick={() => push(`/sales/${sale.id}/edit`)}>
                                <Td>{formatDate(sale.sale_date)}</Td>
                                <Td isNumeric>{sale.quantity_sold}</Td>
                                <Td isNumeric>{sale.sale_price}</Td>
                            </Tr>
                        ))}
                    </Tbody>
                </Table>
            </TableContainer>
        </>
    );
}
