import { Link } from '../link/link';
import { FaTrash, FaEdit } from 'react-icons/fa';

export type ItemCardProps = {
    item: Item;
};

type Item = {
    id: string;
    name: string;
    buying_price: number;
    selling_price: number;
    quantity: number;
    CreatedAt: string;
};

export function ItemCard({ item }: ItemCardProps) {
    const { name, buying_price, selling_price, quantity, CreatedAt } = item;

    return (
        <div className="card bg-base-200 shadow-md">
            {/*<figure>
                <img
                    src="/images/stock/photo-1635805737707-575885ab0820.jpg"
                    alt="Movie"
                />
            </figure>*/}

            <div className="card-body flex-row items-center justify-between gap-14">
                <div className="flex-1 flex flex-col">
                    <div className="indicator">
                        <span
                            className={`indicator-item indicator-start badge ${
                                !quantity ? 'badge-error' : 'badge-secondary'
                            }`}
                        >
                            {quantity}
                        </span>
                        <h2 className="card-title">{name}</h2>
                    </div>
                    <div className="flex">
                        <div className="badge badge-md">N{buying_price}</div>
                    </div>
                </div>
                {/*<p>If a dog chews shoes whose shoes does he choose?</p>*/}
                <div className="font-bold text-xl">N{selling_price}</div>
                <div className="card-actions gap-1 justify-end">
                    <Link to="/products/edit">
                        <FaEdit className="text-secondary" />
                    </Link>
                    <Link to="/products/delete">
                        <FaTrash className="text-error" />
                    </Link>
                </div>
            </div>
        </div>
    );
}

export default ItemCard;
