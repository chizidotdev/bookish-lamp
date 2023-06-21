import {Box, Card, CardBody, Grid, Heading, Image, Stat, StatArrow, StatLabel, StatNumber} from '@chakra-ui/react';
import {useQuery} from 'react-query';
import {getInventory} from '~api/inventory';
import {Loading} from '~components';

export default function Dashboard() {
    const {data, isLoading} = useQuery(['inventory'], {
        queryFn: getInventory,
    });

    if (isLoading) return <Loading/>;
    if (!data) return <Loading/>;

    return (
        <>
            <Heading size="lg" mb="4">
                Dashboard
            </Heading>
            <Grid gridTemplateColumns="repeat(auto-fill, minmax(300px, 1fr))" gap="5">
                <DashboardCard title="Total Items" value={data.total_items}/>
                <DashboardCard title="Low Stock Items" value={data.low_stock_items}/>
                <DashboardStatCard title="Sales Performance" value={data.sales_performance}/>
                {/*<DashboardCard title="Pending Orders" value={data.pending_orders} />*/}
            </Grid>
        </>
    );
}

function DashboardCard({title, value}: { title: string; value: string | number }) {
    return (
        <Card direction="row" variant="filled">
            <CardBody>
                <Stat>
                    <StatLabel>{title}</StatLabel>
                    <StatNumber>{value}</StatNumber>
                </Stat>
            </CardBody>

            <Box marginBlock="auto" mr="5">
                <Image objectFit="cover" w="10" src="/dashboard/chart.svg" alt="Caffe Latte"/>
            </Box>
        </Card>
    );
}

function DashboardStatCard({title, value}: { title: string; value: number }) {
    return (
        <Card direction="row" variant="filled">
            <CardBody>
                <Stat>
                    <StatLabel>{title}</StatLabel>
                    <StatNumber>
                        <StatArrow type={value < 0 ? 'decrease' : 'increase'}/>
                        {value}%
                    </StatNumber>
                    {/*<StatHelpText>*/}
                    {/*    9.05%*/}
                    {/*</StatHelpText>*/}
                </Stat>
            </CardBody>

            <Box marginBlock="auto" mr="5">
                <Image objectFit="cover" w="10" src="/dashboard/chart.svg" alt="Caffe Latte"/>
            </Box>
        </Card>
    );
}
