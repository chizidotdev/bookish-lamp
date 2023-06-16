import { Button } from '@chakra-ui/react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import React from 'react';
import { CgMenuLeft } from 'react-icons/cg';
import { logout } from '~api/user';
import { Text } from '~components';
import { useUser } from '~store/user-store';

const navbarItems = [
    { title: 'Items', href: '/items' },
    { title: 'Product', href: '/product' },
    { title: 'Company', href: '/company' },
    { title: 'Pricing', href: '/pricing' },
];

export function Navbar() {
    const pathname = usePathname();
    const { user } = useUser();

    const handleLogout = async () => {
        await logout();
    };

    if (pathname === '/' || pathname === '/auth/login' || pathname === '/auth/signup') return null;

    return (
        <nav className='fixed top-0 w-full bg-base-300 z-10'>
            <div className='navbar container mx-auto'>
                <div className='navbar-start'>
                    {/*<div className='dropdown'>
                        <label tabIndex={0} className='btn btn-ghost mr-2 lg:hidden'>
                            <CgMenuLeft size='25' />
                        </label>
                        <ul
                            tabIndex={0}
                            className='menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52'
                        >
                            {navbarItems.map(({ title, href }) => (
                                <li key={href}>
                                    <Link href={href}>{title}</Link>
                                </li>
                            ))}
                        </ul>
                    </div>*/}
                    <Link href='/' className='normal-case text-xl'>
                        Copia
                    </Link>
                </div>

                {/*<div className='navbar-center hidden lg:flex'>
                    <ul className='menu menu-horizontal px-1'>
                        {navbarItems.map(({ title, href }) => (
                            <li key={href}>
                                <Link href={href}>{title}</Link>
                            </li>
                        ))}
                    </ul>
                </div>*/}
                <div className='navbar-end gap-2'>
                    {user ? (
                        <>
                            <Text variant='p'>{user.email}</Text>
                            <Button variant='outline' onClick={handleLogout}>Logout</Button>
                        </>
                    ) : (
                        <>
                            <Link href='/auth/login' className='btn btn-outline'>
                                Login
                            </Link>
                            <Link href='/auth/signup' className='btn'>
                                Get started
                            </Link>
                        </>
                    )}
                </div>
            </div>
        </nav>
    );
}

export default Navbar;
