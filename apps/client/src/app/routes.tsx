import { Route, Routes } from 'react-router-dom';
import { Navbar } from '@copia/ui';
import { Items } from './items';
import { CreateItem } from './create-item';

export function RoutesConfig() {
    return (
        <>
            <Navbar navItems={navItems} />
            <Routes>
                <Route path="/" element={<Items />} />
                <Route path="/items/create" element={<CreateItem />} />
            </Routes>
        </>
    );
}

export const navItems = [
    {
        name: 'Products',
        path: '/products',
    },
    {
        name: 'Sales',
        path: '/sales',
    },
];
