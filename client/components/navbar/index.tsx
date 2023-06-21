import { usePathname } from 'next/navigation';
import Nav from './navbar';

export function Navbar() {
    const pathname = usePathname();

    if (pathname === '/' || pathname === '/auth/login' || pathname === '/auth/signup') return null;

    return <Nav />;
}

export default Navbar;
