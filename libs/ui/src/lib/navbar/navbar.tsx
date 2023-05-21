import { FaStoreAlt } from 'react-icons/fa';
import { NavLink, Outlet } from 'react-router-dom';

export interface NavbarProps {
    navItems?: {
        name: string;
        path: string;
    }[];
}

export const Navbar: React.FC<NavbarProps> = () => {
    return (
        <>
            <nav className="navbar bg-base-100">
                <div className="navbar-start">
                    {/*<HamburgerDropdown navItems={navItems} />*/}
                    <NavLink
                        className="normal-case text-xl flex items-center gap-2"
                        to="/"
                    >
                        <FaStoreAlt className="text-secondary" />
                        Copia
                    </NavLink>
                </div>
                <div className="navbar-end">
                    {/*<div className="hidden lg:flex">
                    <ul className="menu menu-horizontal px-1">
                        {navItems.map(({ name, path }) => (
                            <li key={path}>
                                <NavLink to={path}>{name}</NavLink>
                            </li>
                        ))}
                    </ul>
                </div>*/}
                    <div className="dropdown dropdown-bottom dropdown-end">
                        <label tabIndex={0} className="btn">
                            Add
                        </label>
                        <ul
                            tabIndex={0}
                            className="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52"
                        >
                            <li>
                                <NavLink to="/items/create">New Item</NavLink>
                            </li>
                            <li>
                                <NavLink to="/sales/create">New Sale</NavLink>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
            <Outlet />
        </>
    );
};
