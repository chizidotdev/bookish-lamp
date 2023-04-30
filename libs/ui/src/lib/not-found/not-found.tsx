import Link from '../link/link';

export interface NotFoundProps {
    children?: React.ReactNode;
}

export function NotFound({ children }: NotFoundProps) {
    return (
        <div className="flex flex-col items-center">
            <img
                src="empty-folder.png"
                className="w-48 mx-auto"
                alt="No items found"
            />
            <div className="text-center my-4">{children}</div>
        </div>
    );
}

export default NotFound;
