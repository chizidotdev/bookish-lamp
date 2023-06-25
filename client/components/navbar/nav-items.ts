export type NavItem = {
    label: string;
    subLabel?: string;
} & (
    | {
          children?: Array<NavItem>;
          href?: never;
      }
    | {
          href: string;
          children?: never;
      }
);

export const NAV_ITEMS: Array<NavItem> = [
    {
        label: 'Dashboard',
        href: '/dashboard',
    },
    {
        label: 'Items',
        href: '/items',
    },
    {
        label: 'Sales',
        href: '/sales',
    },
    {
        label: 'Quick Add',
        children: [
            {
                label: 'Inventory Item',
                subLabel: 'Trending Design to inspire you',
                href: '/items/new'
            },
            {
                label: 'Sale Record',
                subLabel: 'Up-and-coming Designers',
                href: '/sales/new',
            },
        ],
    },
];
