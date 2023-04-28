import { useQuery, useQueryClient } from 'react-query';
import { Table } from '@copia/ui';

export const Items = () => {
    const { isLoading, error, data } = useQuery({
        queryKey: ['get-items'],
        queryFn: getItems,
    });
    const queryClient = useQueryClient();

    function handleInvalidate() {
        // Invalidate and refetch any query with the key 'get-items'
        queryClient.invalidateQueries({ queryKey: ['get-items'] });
    }

    return (
        <>
            <h1>Items</h1>
            <button onClick={handleInvalidate} className="btn mb-2">
                Refetch Items
            </button>

            {isLoading && <div>Loading...</div>}
            {error && <div>Error: {error}</div>}

            <div className="overflow-x-auto">
                {data && (
                    <Table>
                        <Table.Thead>
                            <Table.Row>
                                <Table.Th>#</Table.Th>
                                <Table.Th>Name</Table.Th>
                                <Table.Th>Quantity</Table.Th>
                                <Table.Th>Buying Price</Table.Th>
                                <Table.Th>Selling Price</Table.Th>
                                <Table.Th>Created At</Table.Th>
                            </Table.Row>
                        </Table.Thead>
                        <Table.Tbody>
                            {data.map((item, index) => (
                                <Table.Row key={item.id}>
                                    <Table.Td>{index}</Table.Td>
                                    <Table.Td>{item.name}</Table.Td>
                                    <Table.Td>{item.quantity}</Table.Td>
                                    <Table.Td>{item.buying_price}</Table.Td>
                                    <Table.Td>{item.selling_price}</Table.Td>
                                    <Table.Td>{new Date(item.CreatedAt).toLocaleString()}</Table.Td>
                                </Table.Row>
                            ))}
                        </Table.Tbody>
                    </Table>
                )}
            </div>
        </>
    );
};

type Item = {
    id: string;
    name: string;
    buying_price: number;
    selling_price: number;
    quantity: number;
    CreatedAt: string;
};

const getItems = async (): Promise<Item[]> => {
    const response = await fetch('http://localhost:3333/api/v1/items');
    const data = await response.json();
    return data;
};
