-- TABLES
CREATE TABLE IF NOT EXISTS ac_roles (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    role VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS ac_categories (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    category VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS ac_actions (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    action VARCHAR(50) NOT NULL
);


-- RELATION TABLE: ROLE / PERMISSIONS
CREATE TABLE IF NOT EXISTS ac_relations (
    role_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    action_id BIGINT NOT NULL,
    PRIMARY KEY (role_id, category_id, action_id),
    FOREIGN KEY (role_id) REFERENCES ac_roles(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES ac_categories(id) ON DELETE CASCADE,
    FOREIGN KEY (action_id) REFERENCES ac_actions(id) ON DELETE CASCADE
);

-- INSERT
INSERT IGNORE INTO ac_roles (role) VALUES
    ('admin'), ('employee'), ('client');

INSERT IGNORE INTO ac_categories (category) VALUES
    ('users'), ('stock');

INSERT IGNORE INTO ac_actions (action) VALUES
    ('view'), ('create'), ('update'), ('delete');

-- -- VARIABLES
-- roles
SET @admin_id = (SELECT id FROM ac_roles WHERE role = 'admin');
SET @employee_id = (SELECT id FROM ac_roles WHERE role = 'employee');
SET @client_id = (SELECT id FROM ac_roles WHERE role = 'client');
-- categories
SET @users_id = (SELECT id FROM ac_categories WHERE category = 'users');
SET @stock_id = (SELECT id FROM ac_categories WHERE category = 'stock');
-- actions
SET @view_id = (SELECT id FROM ac_actions WHERE action = 'view');
SET @create_id = (SELECT id FROM ac_actions WHERE action = 'create');
SET @update_id = (SELECT id FROM ac_actions WHERE action = 'update');
SET @delete_id = (SELECT id FROM ac_actions WHERE action = 'delete');

-- RELATIONS
INSERT IGNORE INTO ac_relations (role_id, category_id, action_id) VALUES
    -- ADMIN PERMISSIONS
    (@admin_id, @users_id, @view_id), -- <= USERS
    (@admin_id, @users_id, @create_id),
    (@admin_id, @users_id, @update_id),
    (@admin_id, @users_id, @delete_id),
    (@admin_id, @stock_id, @view_id), -- <= STOCK
    (@admin_id, @stock_id, @create_id),
    (@admin_id, @stock_id, @update_id),
    (@admin_id, @stock_id, @delete_id),
    -- EMPLOYEE PERMISSIONS
    (@employee_id, @stock_id, @view_id), -- <= STOCK
    (@employee_id, @stock_id, @create_id),
    (@employee_id, @stock_id, @update_id),
    (@employee_id, @stock_id, @delete_id),
    -- CLIENT PERMISSIONS
    (@client_id, @stock_id, @view_id); -- <= STOCK

