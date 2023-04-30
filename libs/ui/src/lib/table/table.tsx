import { ComponentProps } from 'react';

/* eslint-disable-next-line */
export interface TableProps {
    children?: React.ReactNode;
}

function TableBody({ children }: TableProps) {
    return <table className="table w-full">{children}</table>;
}

function TableHeadSection({ children }: TableProps) {
    return <thead>{children}</thead>;
}

function TableBodySection({ children }: TableProps) {
    return <tbody>{children}</tbody>;
}

function TableRow({ children, ...props }: ComponentProps<'tr'> & TableProps) {
    return (
        <tr className="hover" {...props}>
            {children}
        </tr>
    );
}

function TableHead({ children, ...props }: ComponentProps<'th'> & TableProps) {
    return <th className="px-5" {...props}>{children}</th>;
}

function TableDataCell({
    children,
    ...props
}: ComponentProps<'td'> & TableProps) {
    return <td className="px-5" {...props}>{children}</td>;
}

export const Table = Object.assign(TableBody, {
    Thead: TableHeadSection,
    Tbody: TableBodySection,
    Row: TableRow,
    Th: TableHead,
    Td: TableDataCell,
});

export default Table;
