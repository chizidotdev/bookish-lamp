import { Box, Card, CardBody, Flex, Grid, Heading, Image, Text } from '@chakra-ui/react';
import { useQuery } from 'react-query';
import { getInventory } from '~api/inventory';
import { Loading } from '~components';
import { Inventory } from '~lib/types';

export default function Dashboard() {
    const { data, isLoading } = useQuery(['inventory'], {
        queryFn: getInventory,
    });

    if (isLoading) return <Loading />;
    if (!data) return <Loading />;

    return (
        <>
            <Heading size="lg" mb="4">
                Dashboard
            </Heading>
            <Grid gridTemplateColumns="repeat(auto-fill, minmax(300px, 1fr))" gap="5">
                <DashboardCard title="Total Items" value={data.total_items} />
                <DashboardCard title="Low Stock Items" value={data.low_stock_items} />
                <DashboardCard title="Sales Performance" value={`${data.sales_performance}%`} />
                <DashboardCard title="Pending Orders" value={data.pending_orders} />
            </Grid>
        </>
    );
}

function DashboardCard({ title, value }: { title: string; value: string | number }) {
    return (
        <Card direction="row" variant="filled">
            <CardBody>
                <Box marginBlock="auto">
                    <Text py="2">{title}</Text>
                    <Heading size="md">{value}</Heading>
                </Box>
            </CardBody>

            <Box marginBlock="auto" mr="5">
                <Image objectFit="cover" w="10" src="/dashboard/chart.svg" alt="Caffe Latte" />
            </Box>
        </Card>
    );
}
