import { Route, Routes, Link } from 'react-router-dom';
import { Navbar } from '@copia/ui';
import { Items } from './items';

export function RoutesConfig() {
    return (
        <>
            <Navbar navItems={navItems} />
            <Routes>
                <Route path="/" element={<Items />} />
                <Route
                    path="/page-2"
                    element={
                        <div>
                            <Link to="/">
                                Click here to go back to root page.
                            </Link>
                        </div>
                    }
                />
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
