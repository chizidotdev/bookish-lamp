import { NavLink } from 'react-router-dom';
import { HamburgerDropdown } from './hamburger-dropdown';

export interface NavbarProps {
    navItems: {
        name: string;
        path: string;
    }[];
}

export const Navbar = ({ navItems }: NavbarProps) => {
    return (
        <nav className="navbar bg-base-100 px-0">
            <div className="navbar-start">
                <HamburgerDropdown navItems={navItems} />
                <NavLink className="normal-case text-xl" to="/">
                    Copia
                </NavLink>
            </div>
            <div className="navbar-end">
                <div className="hidden lg:flex">
                    <ul className="menu menu-horizontal px-1">
                        {navItems.map(({ name, path }) => (
                            <li key={path}>
                                <NavLink to={path}>{name}</NavLink>
                            </li>
                        ))}
                    </ul>
                </div>
                <NavLink to="/items/create" className="btn btn-sm mr-1">
                    Add Item
                </NavLink>
            </div>
        </nav>
    );
};
