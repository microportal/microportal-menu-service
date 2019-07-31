create table module
(
    id   SMALLSERIAL primary key,
    name VARCHAR(255) not null
);

comment on table module is 'Application modules table';
comment on column module.id is 'Primary key';
comment on column module.name is 'Module name';

create table menu
(
    id          SMALLSERIAL primary key,
    moduleId    SMALLINT     not null,
    name        VARCHAR(100) not null,
    description VARCHAR(1000),
    aggregator  BOOLEAN      not null,
    action      VARCHAR(1000),
    "order"     SMALLINT     not null,
    parentId    SMALLINT
);

comment on table menu is 'Menus for modules';
comment on column menu.id is 'Primary key';
comment on column menu.moduleId is 'FK column for module.id';
comment on column menu.name is 'Menu name';
comment on column menu.description is 'Description for menu item';
comment on column menu.aggregator is 'Flag who indicates if a menu item is a aggregator';
comment on column menu.action is 'Action called for menu item';
comment on column menu."order" is 'Ordering field for menu';
comment on column menu.parentId is 'FK for parent menu item';

alter table menu
    add constraint fk_menu_module foreign key (moduleId) references module (id);
comment on constraint fk_menu_module on menu is 'FK for module table';

alter table menu
    add constraint fk_menu_parent foreign key (parentId) references menu (id);
comment on constraint fk_menu_parent on menu is 'FK for parent menu';
