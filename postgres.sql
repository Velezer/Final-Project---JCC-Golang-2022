-- public.categories definition

-- Drop table

-- DROP TABLE public.categories;

CREATE TABLE public.categories (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NOT NULL,
	CONSTRAINT categories_name_key UNIQUE (name),
	CONSTRAINT categories_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_categories_deleted_at ON public.categories USING btree (deleted_at);


-- public.order_statuses definition

-- Drop table

-- DROP TABLE public.order_statuses;

CREATE TABLE public.order_statuses (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NOT NULL,
	CONSTRAINT order_statuses_name_key UNIQUE (name),
	CONSTRAINT order_statuses_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_order_statuses_deleted_at ON public.order_statuses USING btree (deleted_at);


-- public.roles definition

-- Drop table

-- DROP TABLE public.roles;

CREATE TABLE public.roles (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NOT NULL,
	CONSTRAINT roles_name_key UNIQUE (name),
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_roles_deleted_at ON public.roles USING btree (deleted_at);


-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	email text NOT NULL,
	username varchar(255) NOT NULL,
	"password" text NULL,
	address text NOT NULL,
	role_id int8 NOT NULL,
	CONSTRAINT chk_users_address CHECK ((address <> ''::text)),
	CONSTRAINT chk_users_username CHECK (((username)::text <> ''::text)),
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_username_key UNIQUE (username),
	CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE
);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


-- public.carts definition

-- Drop table

-- DROP TABLE public.carts;

CREATE TABLE public.carts (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NULL,
	is_checkout bool NULL DEFAULT false,
	total_price int8 NULL,
	user_id int8 NULL,
	CONSTRAINT carts_pkey PRIMARY KEY (id),
	CONSTRAINT fk_carts_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE
);
CREATE INDEX idx_carts_deleted_at ON public.carts USING btree (deleted_at);


-- public.orders definition

-- Drop table

-- DROP TABLE public.orders;

CREATE TABLE public.orders (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	status_id int8 NULL,
	user_id int8 NULL,
	cart_id int8 NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (id),
	CONSTRAINT fk_orders_cart FOREIGN KEY (cart_id) REFERENCES public.carts(id) ON UPDATE CASCADE,
	CONSTRAINT fk_orders_status FOREIGN KEY (status_id) REFERENCES public.order_statuses(id),
	CONSTRAINT fk_orders_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE
);
CREATE INDEX idx_orders_deleted_at ON public.orders USING btree (deleted_at);


-- public.stores definition

-- Drop table

-- DROP TABLE public.stores;

CREATE TABLE public.stores (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NOT NULL,
	description text NULL,
	address text NULL,
	image_url text NULL,
	user_id int8 NULL,
	CONSTRAINT chk_stores_name CHECK ((name <> ''::text)),
	CONSTRAINT stores_name_key UNIQUE (name),
	CONSTRAINT stores_pkey PRIMARY KEY (id),
	CONSTRAINT fk_stores_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL ON UPDATE CASCADE
);
CREATE INDEX idx_stores_deleted_at ON public.stores USING btree (deleted_at);


-- public.order_merchant definition

-- Drop table

-- DROP TABLE public.order_merchant;

CREATE TABLE public.order_merchant (
	order_id int8 NOT NULL,
	user_id int8 NOT NULL,
	CONSTRAINT order_merchant_pkey PRIMARY KEY (order_id, user_id),
	CONSTRAINT fk_order_merchant_order FOREIGN KEY (order_id) REFERENCES public.orders(id) ON UPDATE CASCADE,
	CONSTRAINT fk_order_merchant_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE
);


-- public.products definition

-- Drop table

-- DROP TABLE public.products;

CREATE TABLE public.products (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" text NULL,
	description text NULL,
	count int8 NULL,
	price int8 NULL,
	image_url text NULL,
	store_id int8 NULL,
	user_id int8 NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id),
	CONSTRAINT fk_products_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL ON UPDATE CASCADE,
	CONSTRAINT fk_stores_products FOREIGN KEY (store_id) REFERENCES public.stores(id) ON UPDATE CASCADE
);
CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


-- public.cart_items definition

-- Drop table

-- DROP TABLE public.cart_items;

CREATE TABLE public.cart_items (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	product_id int8 NULL,
	count int8 NULL,
	cart_id int8 NULL,
	CONSTRAINT cart_items_pkey PRIMARY KEY (id),
	CONSTRAINT fk_cart_items_product FOREIGN KEY (product_id) REFERENCES public.products(id),
	CONSTRAINT fk_carts_cart_items FOREIGN KEY (cart_id) REFERENCES public.carts(id) ON UPDATE CASCADE
);
CREATE INDEX idx_cart_items_deleted_at ON public.cart_items USING btree (deleted_at);
CREATE UNIQUE INDEX product_cart_pair ON public.cart_items USING btree (product_id, cart_id);


-- public.product_category definition

-- Drop table

-- DROP TABLE public.product_category;

CREATE TABLE public.product_category (
	product_id int8 NOT NULL,
	category_id int8 NOT NULL,
	CONSTRAINT product_category_pkey PRIMARY KEY (product_id, category_id),
	CONSTRAINT fk_product_category_category FOREIGN KEY (category_id) REFERENCES public.categories(id) ON UPDATE CASCADE,
	CONSTRAINT fk_product_category_product FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE
);