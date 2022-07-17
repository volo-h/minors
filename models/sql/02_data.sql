INSERT INTO roles (name) VALUES ('Admin');
INSERT INTO roles (name) VALUES ('Editor');
INSERT INTO roles (name) VALUES ('Viewer');

INSERT INTO permissions (name) VALUES ('view_users');
INSERT INTO permissions (name) VALUES ('edit_users');
INSERT INTO permissions (name) VALUES ('view_roles');
INSERT INTO permissions (name) VALUES ('edit_roles');
INSERT INTO permissions (name) VALUES ('view_products');
INSERT INTO permissions (name) VALUES ('edit_products');
INSERT INTO permissions (name) VALUES ('view_orders');
INSERT INTO permissions (name) VALUES ('edit_orders');

# add to Admin role (role_id == 1) All persmissions!!!
INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 1);
INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 2);
INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 3);
INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 4);
INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 5);
INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 6);
INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 7);
INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 8);


INSERT INTO products (title, description, price) VALUES ('product1', 'description 1', 10);
INSERT INTO products (title, description, price) VALUES ('product2', 'description 2', 20);

INSERT INTO orders (first_name, last_name, email) VALUES ('test', 'test', 'test@google.com');
INSERT INTO order_items (order_id, product_title, price, quantity) VALUES (1, 'product 1', 10.0, 2);

INSERT INTO orders (first_name, last_name, email) VALUES ('test', 'test', 'test@google.com');
INSERT INTO order_items (order_id, product_title, price, quantity) VALUES (2, 'product 1', 20.0, 2);