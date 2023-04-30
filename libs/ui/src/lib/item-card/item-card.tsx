import { Link } from '../link/link';
import { FaTrash, FaEdit } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';

export type ItemCardProps = {
    item: Item;
};

type Item = {
    ID: string;
    name: string;
    buying_price: number;
    selling_price: number;
    quantity: number;
    CreatedAt: string;
};

export function ItemCard({ item }: ItemCardProps) {
    const navigate = useNavigate();
    const { ID, name, buying_price, selling_price, quantity, CreatedAt } = item;

    return (
        <div className="card bg-base-200 shadow-md">
            <div className="card-body flex-row items-center justify-between gap-1 mx-0 py-5">
                <div
                    onClick={() => navigate(`/items/${ID}/sales`)}
                    className="flex-1 flex flex-col cursor-pointer"
                >
                    <div className="indicator">
                        <span
                            className={`indicator-item indicator-start badge badge-sm ${
                                !quantity ? 'badge-error' : 'badge-secondary'
                            }`}
                        >
                            {quantity}
                        </span>
                        <h2 className="card-title text-lg">{name}</h2>
                    </div>
                    <div className="flex gap-2 mt-1">
                        <div className="badge">₦{buying_price}</div>
                        <div className="badge badge-accent">
                            ₦{selling_price}
                        </div>
                    </div>
                </div>
                {/*<p>If a dog chews shoes whose shoes does he choose?</p>*/}
                <div className="card-actions flex-row gap-5">
                    <Link to={`/items/edit/${ID}`}>
                        <FaEdit className="text-secondary" />
                    </Link>
                    <Link to={`/items/delete/${ID}`}>
                        <FaTrash className="text-error" />
                    </Link>
                </div>
            </div>
        </div>
    );
}

export default ItemCard;
