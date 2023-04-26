/* eslint-disable-next-line */
export interface TableProps {
    children: React.ReactNode;
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

function TableRow({ children }: TableProps) {
    return <tr className='hover'>{children}</tr>;
}

function TableHead({ children }: TableProps) {
    return <th>{children}</th>;
}

function TableDataCell({ children }: TableProps) {
    return <td>{children}</td>;
}

export const Table = Object.assign(TableBody, {
    Thead: TableHeadSection,
    Tbody: TableBodySection,
    Row: TableRow,
    Th: TableHead,
    Td: TableDataCell,
});

export default Table;
