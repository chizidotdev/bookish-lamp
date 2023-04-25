
import React, { useEffect } from 'react'

export const items = (props : {}) => {
    useEffect(() => {
        getItems();
    }, []);

    const getItems = async () => {
        const response = await fetch('http://localhost:3333/api/v1/items');
        const data = await response.json();
        console.log(data);
    };
    return (
        <div>

        </div>
    )
}
