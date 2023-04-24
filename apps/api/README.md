# API DOCS

### Inventory Items

- [GET /items](#get-items)
- [GET /items/:id](#get-itemsid)
- [POST /item](#post-item)
- [PUT /item/:id](#put-itemid)
- [DELETE /item/:id](#delete-itemid)

### Inventory Categories

- [GET /categories](#get-categories)
- [GET /categories/:id](#get-categoriesid)
- [POST /category](#post-category)
- [PUT /category/:id](#put-categoryid)
- [DELETE /category/:id](#delete-categoryid)

### Inventory Sales

- [GET /sales](#get-sales)
- [GET /sales/:id](#get-salesid)
- [POST /sale](#post-sale)
- [PUT /sale/:id](#put-saleid)
- [DELETE /sale/:id](#delete-saleid)

## Inventory Item Object

```ts
{
  id: int,
  name: string,
  category: string,
  quantity: int,
  price: {
      buyingPrice: float,
      sellingPrice: float,
    },
  image: string,
  status: string,
  createdAt: string,
}
```

## Inventory Category Object

```ts
{
  id: int,
  name: string,
  description: string,
  items: [
    {
      id: int,
      name: string,
      category: string,
      quantity: int,
      price: {
          buyingPrice: float,
          sellingPrice: float,
        },
      image: string,
      status: string,
    }
  ]
}
```

## Inventory Sale Object

```ts
{
  id: int,
  total: float,
  quantity: int,
  date: string,
  createdAt: string,
  item: {
    id: int,
    name: string,
    category: string,
    quantity: int,
    price: {
        buyingPrice: float,
        sellingPrice: float,
      },
    image: string,
    status: string,
  },
}
```
