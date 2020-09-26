-- Products
CREATE TABLE IF NOT EXISTS products (
    id                  serial PRIMARY KEY,
    code                text default '',
    title               text NOT NULL,
    category            text default '',
    maker               text default '',
    description         text default '',
    ean                 text default '',
    ncm                 text default '',
    warranty_month      integer default 0,
    length_mm           integer NOT NULL,
    width_mm            integer NOT NULL,
    height_mm           integer NOT NULL,
    weight_g            integer NOT NULL,
    is_active           boolean default false,    -- System defined.
    is_available        boolean default false,   -- Dealer defined.
    is_sell             boolean default false,   -- User defined.
    stock_origin        text default '',
    stock_count         integer default 0,
    price_buy           integer default 0,
    price_sale          integer default 0,
    supplier_code       text default '',
    supplier_name       text default '',
    created_at          timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    changed_at          timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_deleted          boolean default false
);
