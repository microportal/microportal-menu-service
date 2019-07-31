alter table menu drop constraint fk_menu_parent;
alter table menu drop constraint fk_menu_module;
drop table menu;
drop table module;
