import { useQuery } from 'react-query';
import { Button, Link, NotFound, Table } from '@copia/ui';
import { getItemSales } from './api';
import { Outlet, useNavigate, useParams } from 'react-router-dom';
import { dateFormat } from '../utils';

export const ItemSales = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const { isLoading, data, isSuccess } = useQuery({
        queryKey: ['get-item-sales'],
        queryFn: () => getItemSales(id ?? ''),
        refetchOnWindowFocus: false,
        onError: () => {
            navigate('/');
        },
    });

    return (
        <div className="my-5">
            <h1 className="text-2xl font-bold mb-5 ml-2">Sales</h1>
            {isLoading && <div>Loading...</div>}

            {isSuccess && data?.length === 0 && (
                <NotFound>
                    No Sales found for this item...
                    <br />
                    <Link to="/sales/create">Add Sale</Link>
                </NotFound>
            )}

            <div>
                {data && data.length > 0 && (
                    <>
                        <Table>
                            <Table.Thead>
                                <Table.Row>
                                    <Table.Th>QTY</Table.Th>
                                    <Table.Th>Price</Table.Th>
                                    <Table.Th>Date</Table.Th>
                                </Table.Row>
                            </Table.Thead>
                            <Table.Tbody>
                                {data.map((sale) => (
                                    <Table.Row key={sale.ID}>
                                        <Table.Td>
                                            {sale.quantity_sold}
                                        </Table.Td>
                                        <Table.Td>N{sale.unit_price}</Table.Td>
                                        <Table.Td>
                                            {dateFormat(sale.sale_date)}{' '}
                                        </Table.Td>
                                    </Table.Row>
                                ))}
                            </Table.Tbody>
                        </Table>

                        <div className="mt-5 flex gap-3">
                            <Button
                                onClick={() =>
                                    navigate(`/sales/create?itemId=${id}`)
                                }
                            >
                                Add Sale
                            </Button>
                        </div>
                    </>
                )}
            </div>

            <Outlet />
        </div>
    );
};
