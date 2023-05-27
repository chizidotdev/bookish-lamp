import Link from 'next/link';
import React from 'react';
import { CgMenuLeft } from 'react-icons/cg';

const navbarItems = [
    { title: 'Product', href: '/product' },
    { title: 'Company', href: '/company' },
    { title: 'Pricing', href: '/pricing' },
];

export default function Navbar() {
    return (
        <div className='bg-base-100 z-10'>
            <div className='navbar container mx-auto'>
                <div className='navbar-start'>
                    <div className='dropdown'>
                        <label tabIndex={0} className='btn btn-ghost lg:hidden'>
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
                    </div>
                    <Link href='/' className='normal-case text-xl'>
                        Copia
                    </Link>
                </div>

                <div className='navbar-center hidden lg:flex'>
                    <ul className='menu menu-horizontal px-1'>
                        {navbarItems.map(({ title, href }) => (
                            <li key={href}>
                                <Link href={href}>{title}</Link>
                            </li>
                        ))}
                    </ul>
                </div>
                <div className='navbar-end gap-2'>
                    <a href='http://localhost:8080/login' className='btn btn-outline'>
                        Login
                    </a>
                    <Link href='signup' className='btn'>
                        Get started
                    </Link>
                </div>
            </div>
        </div>
    );
}
