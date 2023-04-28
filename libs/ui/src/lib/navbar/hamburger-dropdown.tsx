import { NavLink } from 'react-router-dom';

export interface HamburgerDropdownProps {
    navItems: {
        name: string;
        path: string;
    }[];
}

export const HamburgerDropdown = ({ navItems }: HamburgerDropdownProps) => {
    return (
        <div className="dropdown">
            <label tabIndex={0} className="btn btn-ghost mr-2 lg:hidden">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-5 w-5"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                >
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth="2"
                        d="M4 6h16M4 12h8m-8 6h16"
                    />
                </svg>
            </label>
            <ul
                tabIndex={0}
                className="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52"
            >
                {navItems.map(({ name, path }) => (
                    <li>
                        <NavLink to={path}>{name}</NavLink>
                    </li>
                ))}
            </ul>
        </div>
    );
};
